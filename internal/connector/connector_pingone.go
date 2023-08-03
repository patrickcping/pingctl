package connector

import (
	"context"
	"fmt"
	"regexp"

	"github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/ux"
	"github.com/spf13/viper"
)

type PingOneCmdConnector struct {
	apiClient       *pingone.Client
	profileSettings profileSettings
}

type profileSettings struct {
	APIHostname   *string
	AuthHostname  *string
	Region        string
	ClientId      *string
	ClientSecret  *string
	EnvironmentId *string
	GrantType     string
}

var (
	_ CmdConnector                           = &PingOneCmdConnector{}
	_ CmdConnectorWithProfileSettings        = &PingOneCmdConnector{}
	_ CmdConnectorWithProfileSettingsMigrate = &PingOneCmdConnector{}
	// _ CmdConnectorWithTerraformProvider      = &PingOneCmdConnector{}
)

func NewPingOneCmdConnector() CmdConnector {
	return &PingOneCmdConnector{}
}

func (p *PingOneCmdConnector) CommandName() string {
	return "pingone"
}

func (p *PingOneCmdConnector) ProductName() string {
	return "PingOne"
}

func (p *PingOneCmdConnector) ConfigureConnector(ctx context.Context, version string) error {

	userAgent := fmt.Sprintf("pingctl/%s/go", version)

	apiConfig := pingone.Config{
		ClientID:             p.profileSettings.ClientId,
		ClientSecret:         p.profileSettings.ClientSecret,
		APIHostnameOverride:  p.profileSettings.APIHostname,
		AuthHostnameOverride: p.profileSettings.AuthHostname,
		EnvironmentID:        p.profileSettings.EnvironmentId,
		Region:               p.profileSettings.Region,
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

func (c *PingOneCmdConnector) TestConnection(ctx context.Context) error {
	return nil
}

func (p *PingOneCmdConnector) ConfigureProfileSettings(ctx context.Context) error {
	var p1ResourceIDRegexp = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	var isHostname = regexp.MustCompile(`^(?:[\w-]+\.)+(?:[a-zA-Z]{2,}|xn--[a-zA-Z0-9]+)$`)

	environmentId := ux.PromptGetString(ux.PromptContentString{
		Label:           "PingOne environment ID that contains the client to authenticate the CLI with:",
		RegexValidation: p1ResourceIDRegexp,
	})
	p.profileSettings.EnvironmentId = &environmentId

	clientId := ux.PromptGetString(ux.PromptContentString{
		Label:           "PingOne client ID to authenticate the CLI with:",
		RegexValidation: p1ResourceIDRegexp,
	})
	p.profileSettings.ClientId = &clientId

	p.profileSettings.GrantType = ux.PromptGetSelect(ux.PromptContentSelect{
		Label: "What is the grant type to use for the client?",
		Options: []string{
			"Client Credentials",
			// "Authorization Code",
		},
	})

	if p.profileSettings.GrantType == "Client Credentials" {
		clientSecret := ux.PromptGetString(ux.PromptContentString{
			Label:     "PingOne client secret to authenticate the CLI with, using the client credentials grant:",
			Sensitive: true,
		})

		p.profileSettings.ClientSecret = &clientSecret
	} else {
		// TBC Auth code
	}

	customServiceHostnamesLabel := "Enter custom service hostnames"
	regionOptions := []string{
		"North America",
		"Europe",
		"Asia Pacific",
		"Canada",
	}

	regionOptions = append(regionOptions, customServiceHostnamesLabel)

	regionType := ux.PromptGetSelect(ux.PromptContentSelect{
		Label:   "What is the region of the PingOne tenant?  This is used to populate the correct API endpoints.",
		Options: regionOptions,
	})

	if regionType == customServiceHostnamesLabel {
		authHostname := ux.PromptGetString(ux.PromptContentString{
			Label:           "PingOne authentication service hostname (e.g. auth.pingone.eu):",
			RegexValidation: isHostname,
		})
		p.profileSettings.AuthHostname = &authHostname

		apiHostname := ux.PromptGetString(ux.PromptContentString{
			Label:           "PingOne API service hostname (e.g. api.pingone.eu):",
			RegexValidation: isHostname,
		})
		p.profileSettings.APIHostname = &apiHostname
	} else {
		regionLabels := map[string]string{
			"Europe":        "Europe",
			"North America": "NorthAmerica",
			"Asia Pacific":  "AsiaPacific",
			"Canada":        "Canada",
		}

		p.profileSettings.Region = regionLabels[regionType]
	}

	return nil
}

func (p *PingOneCmdConnector) ProfileSettingsIndex() string {
	return "pingone"
}

func (p *PingOneCmdConnector) LoadProfileSettings(ctx context.Context) error {
	apiHostname := viper.GetString("profiles.default.pingone.apiHostname")
	authHostname := viper.GetString("profiles.default.pingone.authHostname")
	clientId := viper.GetString("profiles.default.pingone.adminClientId")
	clientSecret := viper.GetString("profiles.default.pingone.adminClientSecret")
	environmentId := viper.GetString("profiles.default.pingone.adminEnvironmentId")

	p.profileSettings = profileSettings{
		APIHostname:   &apiHostname,
		AuthHostname:  &authHostname,
		Region:        viper.GetString("profiles.default.pingone.region"),
		ClientId:      &clientId,
		ClientSecret:  &clientSecret,
		EnvironmentId: &environmentId,
		GrantType:     viper.GetString("profiles.default.pingone.grantType"),
	}

	return nil
}

// func (p *PingOneCmdConnector) GenerateTerraformHCL(ctx context.Context) error {

// 	generate := pingtf.NewPingOneTerraformProvider(p.apiClient)
// 	fmt.Println(generate.GenerateImportBlocks(ctx))

// 	return nil
// }
