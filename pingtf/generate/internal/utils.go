package internal

func ConvertResourceObjectSliceToLogicalSchemaMap(resources []PingTerraformProviderResourceGenerate) map[string]LogicalSchemaResourceMap {

	returnVar := make(map[string]LogicalSchemaResourceMap, 0)

	for _, resource := range resources {
		returnVar[resource.GetResourceName()] = LogicalSchemaResourceMap{
			ImportIDPattern:    resource.GetImportIDFormat(),
			FetchResourcesFunc: resource.GetFetchResourcesFunc(),
		}
	}

	return returnVar
}
