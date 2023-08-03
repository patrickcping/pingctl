package pingone

import (
	"context"

	"github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/pingtf/generate/internal"
)

type PingOneTerraformProviderGenerate struct {
	apiClient      *pingone.Client
	resourcePrefix string
	resourceSuffix string
	resources      []internal.PingTerraformProviderResourceGenerate
}

var (
	_ internal.PingTerraformProviderGenerate = &PingOneTerraformProviderGenerate{}
)

func NewPingOneTerraformProvider(apiClient *pingone.Client) internal.PingTerraformProviderGenerate {
	resources := []internal.PingTerraformProviderResourceGenerate{
		NewEnvironmentResource(apiClient),
	}

	return &PingOneTerraformProviderGenerate{
		apiClient:      apiClient,
		resourcePrefix: "generated_",
		resources:      resources,
	}
}

func (p *PingOneTerraformProviderGenerate) GetLogicalSchemaMap() map[string]internal.LogicalSchemaResourceMap {
	return internal.ConvertResourceObjectSliceToLogicalSchemaMap(p.resources)
}

func (p *PingOneTerraformProviderGenerate) GetProviderSchemaResources() ([]string, error) {
	return []string{"pingone_environment", "pingone_application"}, nil
}

func (p *PingOneTerraformProviderGenerate) GenerateImportBlocks(ctx context.Context) (*string, error) {

	schemaResources, err := p.GetProviderSchemaResources()
	if err != nil {
		return nil, err
	}

	logicalSchemaMap := p.GetLogicalSchemaMap()

	return internal.GenerateImportBlocks(ctx, schemaResources, logicalSchemaMap, p.resourcePrefix, p.resourceSuffix)
}

func (p *PingOneTerraformProviderGenerate) GetResourcePrefix() string {
	return p.resourcePrefix
}

func (p *PingOneTerraformProviderGenerate) GetResourceSuffix() string {
	return p.resourceSuffix
}

func (p *PingOneTerraformProviderGenerate) SetResourcePrefix(resourcePrefix string) {
	p.resourcePrefix = resourcePrefix
}

func (p *PingOneTerraformProviderGenerate) SetResourceSuffix(resourceSuffix string) {
	p.resourceSuffix = resourceSuffix
}
