package sdk

import (
	"context"
	"net/http"
	"time"
)

func RetryWrapper(ctx context.Context, timeout time.Duration, f SDKInterfaceFunc, isRetryable Retryable) (interface{}, *http.Response, error) {
	// todo
	return f()
}
