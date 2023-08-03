package internal

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

type TFImportResources struct {
	Resources []TFImportResource
}

type TFImportResource struct {
	ResourceImportID       string
	ResourceName           string
	ConfigurationReference string
	ResourceID             string
}

func GenerateImportBlocks(ctx context.Context, schemaResources []string, logicalSchemaMap map[string]LogicalSchemaResourceMap, resourcePrefix, resourceSuffix string) (*string, error) {
	if len(schemaResources) == 0 {
		fmt.Println("Nothing to do!")
		return nil, nil
	}

	for _, schemaResource := range schemaResources {

		resourceMapping, ok := logicalSchemaMap[strings.ToLower(schemaResource)]
		if !ok {
			fmt.Printf("No logical schema mapping found for resource: %s\n", schemaResource)
			continue
		}

		resourceImportBlocks, err := resourceMapping.FetchResourcesFunc(ctx)
		if err != nil {
			return nil, err
		}

		resourceImportBlocks.generate(resourcePrefix, resourceSuffix)
	}

	return nil, nil
}

func (r *TFImportResources) generate(resourcePrefix, resourceSuffix string) (*io.Writer, error) {
	var templateString = `
{{ range . }}
# Resource Reference: {{ .ResourceName }}.` + resourcePrefix + `{{ .ResourceID }}` + resourceSuffix + `
{{ if .ConfigurationReference }}# Configuration Reference: {{ .ConfigurationReference }}{{ end }}
# Import ID: {{ .ResourceImportID }}
import {
  # Composite import ID of the resource
  id = "{{ .ResourceImportID }}"

  # Resource address
  to = {{ .ResourceName }}.` + resourcePrefix + `{{ .ResourceID }}` + resourceSuffix + `
}
{{ end }}
`

	tmpl, err := template.New("import").Parse(templateString)
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(os.Stdout, r.Resources)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
