package pingone

import (
	"context"
	"regexp"

	"github.com/pingidentity/pingctl/internal/ux"
)

func (p *pingOneCmdConnector) ConfigureProfileSettings(ctx context.Context) error {
	var p1ResourceIDRegexp = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	var isHostname = regexp.MustCompile(`^(?:[\w-]+\.)+(?:[a-zA-Z]{2,}|xn--[a-zA-Z0-9]+)$`)

	environmentId := ux.PromptGetString(ux.PromptContentString{
		Label:           "PingOne environment ID that contains the client to authenticate the CLI with:",
		RegexValidation: p1ResourceIDRegexp,
	})
	p.connectorSettings.AdminEnvironmentId = &environmentId

	clientId := ux.PromptGetString(ux.PromptContentString{
		Label:           "PingOne client ID to authenticate the CLI with:",
		RegexValidation: p1ResourceIDRegexp,
	})
	p.connectorSettings.ClientId = &clientId

	p.connectorSettings.GrantType = ux.PromptGetSelect(ux.PromptContentSelect{
		Label: "What is the grant type to use for the client?",
		Options: []string{
			"Client Credentials",
			// "Authorization Code",
		},
	})

	if p.connectorSettings.GrantType == "Client Credentials" {
		clientSecret := ux.PromptGetString(ux.PromptContentString{
			Label:     "PingOne client secret to authenticate the CLI with, using the client credentials grant:",
			Sensitive: true,
		})

		p.connectorSettings.ClientSecret = &clientSecret
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
		p.connectorSettings.AuthHostname = &authHostname

		apiHostname := ux.PromptGetString(ux.PromptContentString{
			Label:           "PingOne API service hostname (e.g. api.pingone.eu):",
			RegexValidation: isHostname,
		})
		p.connectorSettings.APIHostname = &apiHostname
	} else {
		regionLabels := map[string]string{
			"Europe":        "Europe",
			"North America": "NorthAmerica",
			"Asia Pacific":  "AsiaPacific",
			"Canada":        "Canada",
		}

		p.connectorSettings.Region = regionLabels[regionType]
	}

	return nil
}
