package connector

import "reflect"

type GenerateHCLOpts struct {
	OutputDirectoryPath string
}

type ConnectorParam struct {
	DataType     reflect.Kind
	DefaultValue any
	Description  string
	IsSensitive  bool
	ProfileKey   string
}
