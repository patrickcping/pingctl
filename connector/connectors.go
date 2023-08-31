package connector

import (
	"context"
)

type CmdMetadata struct {
	ProductName        string
	CommandName        string
	ProfileConfigIndex string
}

type CmdConnector interface {
	Metadata() CmdMetadata
	ConfigureConnector(ctx context.Context, version string, cfg map[string]interface{}) error
	TestConnection(ctx context.Context) error
	ConnectorSettings(ctx context.Context) map[string]ConnectorParam
}

// Will implement settings
type CmdConnectorWithProfileSettings interface {
	ConfigureProfileSettings(ctx context.Context) error
	ProfileSettingsIndex() string
	LoadProfileSettings(ctx context.Context) error
}

// Will implement settings migration
type CmdConnectorWithProfileSettingsMigrate interface {
	ConfigureProfileSettings(ctx context.Context) error
	LoadProfileSettings(ctx context.Context) error
}

// Will implement the "generate" and "docs" commands
type CmdConnectorWithExport interface {
	Export(ctx context.Context, opts GenerateHCLOpts) error
}

// Will implement the "config" command
type CmdConnectorWithConfig interface {
	ConfigCommands(ctx context.Context) (string, error)
}

// Will implement the "export" command
type CmdConnectorWithCustomExport interface {
	Lint() (string, error)
	GenerateDocs(ctx context.Context) (string, error)
}
