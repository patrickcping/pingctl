package internal

import (
	"context"
)

type PingTerraformProviderGenerate interface {
	GenerateImportBlocks(ctx context.Context) (*string, error)
	GetLogicalSchemaMap() map[string]LogicalSchemaResourceMap
	GetProviderSchemaResources() ([]string, error)
	GetResourcePrefix() string
	GetResourceSuffix() string
	SetResourcePrefix(resourcePrefix string)
	SetResourceSuffix(resourceSuffix string)
}

type PingTerraformProviderResourceGenerate interface {
	GetFetchResourcesFunc() ParseResourcesFunc
	GetImportIDFormat() string
	GetLogicalSchemaMap() map[string]LogicalSchemaResourceMap
	GetResourceName() string
}

type ParseResourcesFunc func(ctx context.Context, ids ...string) (*TFImportResources, error)

type LogicalSchemaResourceMap struct {
	FetchResourcesFunc ParseResourcesFunc
	ImportIDPattern    string
}
