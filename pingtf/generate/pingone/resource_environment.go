package pingone

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sdk"
	"github.com/pingidentity/pingctl/pingtf/generate/internal"
)

const (
	// The primary key
	ENVIRONMENT_ID = "environment_id"
)

type environmentResource struct {
	resourceName string
	idFormat     string
	apiClient    *pingone.Client
	resources    []internal.PingTerraformProviderResourceGenerate
}

var (
	_ internal.PingTerraformProviderResourceGenerate = &environmentResource{}
)

func NewEnvironmentResource(apiClient *pingone.Client) internal.PingTerraformProviderResourceGenerate {
	resources := []internal.PingTerraformProviderResourceGenerate{
		//NewApplicationResource(apiClient),
	}

	return &environmentResource{
		apiClient:    apiClient,
		idFormat:     ENVIRONMENT_ID,
		resourceName: "pingone_environment",
		resources:    resources,
	}
}

func (p *environmentResource) GetResourceName() string {
	return strings.ToLower(p.resourceName)
}

func (p *environmentResource) GetImportIDFormat() string {
	return p.idFormat
}

func (p *environmentResource) GetFetchResourcesFunc() internal.ParseResourcesFunc {
	return func(ctx context.Context, ids ...string) (*internal.TFImportResources, error) {

		var entityArray *management.EntityArray
		if err := sdk.ParseResponse(
			ctx,

			func() (any, *http.Response, error) {
				return p.apiClient.ManagementAPIClient.EnvironmentsApi.ReadAllEnvironments(ctx).Execute()
			},
			"ReadAllEnvironments",
			nil,
			nil,
			&entityArray,
		); err != nil {
			return nil, err
		}

		resources := make([]internal.TFImportResource, 0)
		for _, environment := range entityArray.Embedded.GetEnvironments() {
			resources = append(resources, internal.TFImportResource{
				ResourceImportID:       environment.GetId(),
				ConfigurationReference: fmt.Sprintf("%s (PingOne Environment; %s)", environment.GetName(), string(environment.GetType())),
				ResourceName:           p.resourceName,
				ResourceID:             environment.GetId(),
			})

			// get application resources
		}

		return &internal.TFImportResources{
			Resources: resources,
		}, nil
	}
}

func (p *environmentResource) GetLogicalSchemaMap() map[string]internal.LogicalSchemaResourceMap {
	return internal.ConvertResourceObjectSliceToLogicalSchemaMap(p.resources)
}
