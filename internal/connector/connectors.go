package connector

import (
	"context"
)

func Connectors() []func() CmdConnector {
	return []func() CmdConnector{
		NewPingOneCmdConnector,
	}
}

type CmdConnector interface {
	CommandName() string
	ProductName() string
	ConfigureConnector(ctx context.Context, version string) error
	TestConnection(ctx context.Context) error
}

// Will implement settings
type CmdConnectorWithProfileSettings interface {
	ConfigureProfileSettings(ctx context.Context) error
	ProfileSettingsIndex() string
	LoadProfileSettings(ctx context.Context) error
}

// Will implement settings migration
type CmdConnectorWithProfileSettingsMigrate interface {
	ProfileSettingsIndex() string
	ConfigureProfileSettings(ctx context.Context) error
	LoadProfileSettings(ctx context.Context) error
}

// Will implement the "generate" and "docs" commands
type CmdConnectorWithTerraformProvider interface {
	GenerateTerraformHCL(ctx context.Context) error
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
