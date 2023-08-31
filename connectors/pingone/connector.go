package pingone

import (
	"context"
	"fmt"
	"reflect"

	"github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/connector"
)

type pingOneCmdConnector struct {
	apiClient         *pingone.Client
	connectorSettings connectorSettings
}

type connectorSettings struct {
	APIHostname          *string `profile:"apihostname" pingcli:"api-hostname"`
	AuthHostname         *string `profile:"authhostname" pingcli:"auth-hostname"`
	Region               string  `profile:"region" pingcli:"region"`
	ClientId             *string `profile:"adminclientid" pingcli:"admin-client-id"`
	ClientSecret         *string `profile:"adminclientsecret" pingcli:"admin-client-secret"`
	AdminEnvironmentId   *string `profile:"adminenvironmentid" pingcli:"admin-environment-id"`
	GrantType            string  `profile:"adminclientgranttype" pingcli:"grant-type"`
	DefaultEnvironmentId *string `profile:"defaultenvironmentid" pingcli:"default-environment-id"`
}

var (
	_ connector.CmdConnector = &pingOneCmdConnector{}
)

func NewPingOneCmdConnector() connector.CmdConnector {
	return &pingOneCmdConnector{}
}

func (p *pingOneCmdConnector) Metadata() connector.CmdMetadata {
	return connector.CmdMetadata{
		ProductName:        "PingOne",
		CommandName:        "pingone",
		ProfileConfigIndex: "pingone",
	}
}

func (p *pingOneCmdConnector) ConfigureConnector(ctx context.Context, version string, cfg map[string]interface{}) error {

	userAgent := fmt.Sprintf("pingctl/%s/go", version)

	apiConfig := pingone.Config{
		ClientID:             p.connectorSettings.ClientId,
		ClientSecret:         p.connectorSettings.ClientSecret,
		APIHostnameOverride:  p.connectorSettings.APIHostname,
		AuthHostnameOverride: p.connectorSettings.AuthHostname,
		EnvironmentID:        p.connectorSettings.AdminEnvironmentId,
		Region:               p.connectorSettings.Region,
		UserAgentOverride:    &userAgent,
	}

	var err error
	p.apiClient, err = apiConfig.APIClient(ctx)
	if err != nil {
		return err
	}

	if p.apiClient == nil {
		return fmt.Errorf("Error connecting to PingOne service.  Please check your configuration and try again.")
	}

	return nil
}

func (c *pingOneCmdConnector) TestConnection(ctx context.Context) error {
	return nil
}

func (p *pingOneCmdConnector) ConnectorSettings(ctx context.Context) map[string]connector.ConnectorParam {
	return map[string]connector.ConnectorParam{
		"api-hostname": {
			DataType:    reflect.String,
			Description: "PingOne API hostname",
			IsSensitive: false,
			ProfileKey:  "apihostname",
		},
		"auth-hostname": {
			DataType:    reflect.String,
			Description: "PingOne authentication hostname",
			IsSensitive: false,
			ProfileKey:  "authhostname",
		},
		"region": {
			DataType:    reflect.String,
			Description: "PingOne region",
			IsSensitive: false,
			ProfileKey:  "region",
		},
		"admin-client-id": {
			DataType:    reflect.String,
			Description: "PingOne admin client ID",
			IsSensitive: false,
			ProfileKey:  "adminclientid",
		},
		"admin-client-secret": {
			DataType:    reflect.String,
			Description: "PingOne admin client secret",
			IsSensitive: true,
			ProfileKey:  "adminclientsecret",
		},
		"admin-environment-id": {
			DataType:    reflect.String,
			Description: "PingOne admin environment ID",
			IsSensitive: false,
			ProfileKey:  "adminenvironmentid",
		},
		"admin-client-grant-type": {
			DataType:    reflect.String,
			Description: "PingOne admin client grant type",
			IsSensitive: false,
			ProfileKey:  "adminclientgranttype",
		},
		"default-environment-id": {
			DataType:    reflect.String,
			Description: "PingOne default environment ID",
			IsSensitive: false,
			ProfileKey:  "defaultenvironmentid",
		},
	}
}
