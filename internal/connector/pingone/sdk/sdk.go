package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
)

type SDKInterfaceFunc func() (any, *http.Response, error)

type CustomError func(model.P1Error) error

type Retryable func(context.Context, *http.Response, *model.P1Error) bool

var (
	DefaultCustomError = func(error model.P1Error) error { return nil }
	DefaultRetryable   = func(ctx context.Context, r *http.Response, p1error *model.P1Error) bool { return false }
)

func ParseResponse(ctx context.Context, f SDKInterfaceFunc, requestID string, customError CustomError, customRetryConditions Retryable, targetObject any) error {
	defaultTimeout := 10
	return ParseResponseWithCustomTimeout(ctx, f, requestID, customError, customRetryConditions, targetObject, time.Duration(defaultTimeout)*time.Minute)
}

func ParseResponseWithCustomTimeout(ctx context.Context, f SDKInterfaceFunc, requestID string, customError CustomError, customRetryConditions Retryable, targetObject any, timeout time.Duration) error {

	if customError == nil {
		customError = DefaultCustomError
	}

	if customRetryConditions == nil {
		customRetryConditions = DefaultRetryable
	}

	resp, r, err := RetryWrapper(
		ctx,
		timeout,
		f,
		customRetryConditions,
	)

	if err != nil || r.StatusCode >= 300 {

		switch t := err.(type) {
		case *model.GenericOpenAPIError:

			if v, ok := t.Model().(model.P1Error); ok && v.GetId() != "" {

				summaryText := fmt.Sprintf("Error when calling `%s`: %v", requestID, v.GetMessage())
				detailText := fmt.Sprintf("PingOne Error Details:\nID: %s\nCode: %s\nMessage: %s", v.GetId(), v.GetCode(), v.GetMessage())

				// diags = customError(v)
				// if diags != nil {
				// 	return diags
				// }

				if details, ok := v.GetDetailsOk(); ok {
					detailsBytes, err := json.Marshal(details)
					if err != nil {
						return fmt.Errorf("Cannot parse details object - There is an internal problem.  Please raise an issue with the project maintainers.")
					}

					detailText = fmt.Sprintf("%s\nDetails object: %+v", detailText, string(detailsBytes[:]))
				}

				return fmt.Errorf("%s - %s", summaryText, detailText)
			}

			return fmt.Errorf("Error when calling `%s`: %v", requestID, t.Error())

		case *url.Error:
			return fmt.Errorf("Error when calling `%s`: %v", requestID, t.Error())

		default:
			return fmt.Errorf("Error when calling `%s`: %v", requestID, fmt.Sprintf("A generic error has occurred.\nError details: %+v", t))
		}

	}

	if targetObject != nil {
		v := reflect.ValueOf(targetObject)
		if v.Kind() != reflect.Ptr {
			return fmt.Errorf("Invalid target object.  Target object must be a pointer.  Please raise an issue with the project maintainers.")
		}
		if !v.Elem().IsValid() {
			return fmt.Errorf("Invalid target object.  Target object is not valid.  Please raise an issue with the project maintainers.")
		}

		if resp != nil {
			v.Elem().Set(reflect.ValueOf(resp))
		}
	}

	return nil

}
