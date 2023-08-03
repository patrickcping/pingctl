package pingone

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/patrickcping/pingone-go-sdk-v2/management"
// 	"github.com/patrickcping/pingone-go-sdk-v2/pingone"
// 	"github.com/pingidentity/pingctl/internal/pingone/sdk"
// 	"github.com/pingidentity/pingctl/pingtf/generate/internal"
// )

// const (
// 	// The primary key
// 	APPLICATION_ID = "application_id"
// )

// type applicationResource struct {
// 	resourceName string
// 	idFormat     string
// 	apiClient    *pingone.Client
// 	resources    []internal.PingTerraformProviderResourceGenerate
// }

// var (
// 	_ internal.PingTerraformProviderResourceGenerate = &applicationResource{}
// )

// func NewApplicationResource(apiClient *pingone.Client) internal.PingTerraformProviderResourceGenerate {
// 	resources := []internal.PingTerraformProviderResourceGenerate{}

// 	return &applicationResource{
// 		apiClient:    apiClient,
// 		idFormat:     fmt.Sprintf("%s/%s", ENVIRONMENT_ID, APPLICATION_ID),
// 		resourceName: "pingone_application",
// 		resources:    resources,
// 	}
// }

// func (p *applicationResource) GetResourceName() string {
// 	return strings.ToLower(p.resourceName)
// }

// func (p *applicationResource) GetImportIDFormat() string {
// 	return p.idFormat
// }

// func (p *applicationResource) GetFetchResourcesFunc() internal.ParseResourcesFunc {
// 	return func(ctx context.Context, ids ...string) (*internal.TFImportResources, error) {

// 		//

// 		var entityArray *management.EntityArray
// 		if err := sdk.ParseResponse(
// 			ctx,

// 			func() (any, *http.Response, error) {
// 				return p.apiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(ctx, environmentID).Execute()
// 			},
// 			"ReadAllApplications",
// 			nil,
// 			nil,
// 			&entityArray,
// 		); err != nil {
// 			return nil, err
// 		}

// 		resources := make([]internal.TFImportResource, 0)
// 		for _, application := range entityArray.Embedded.GetApplications() {

// 			var common management.Application

// 			bytesData, err := json.Marshal(application)
// 			if err != nil {
// 				return nil, err
// 			}

// 			if err := json.Unmarshal(bytesData, &common); err != nil {
// 				return nil, err
// 			}

// 			resources = append(resources, internal.TFImportResource{
// 				ResourceImportID:       common.GetId(),
// 				ConfigurationReference: fmt.Sprintf("%s (PingOne Application; %s)", common.GetName(), string(common.GetType())),
// 				ResourceName:           p.resourceName,
// 				ResourceID:             common.GetId(),
// 			})
// 		}

// 		return &internal.TFImportResources{
// 			Resources: resources,
// 		}, nil
// 	}
// }

// func (p *applicationResource) GetLogicalSchemaMap() map[string]internal.LogicalSchemaResourceMap {
// 	return internal.ConvertResourceObjectSliceToLogicalSchemaMap(p.resources)
// }
