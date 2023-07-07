package pingone

import (
	"context"
	"fmt"
	"regexp"

	"github.com/pingidentity/pingctl/internal/pingone/client"
	"github.com/pingidentity/pingctl/internal/ux"
)

type ProfilePingOneCmdAuthConfig struct {
	APIHostname   *string
	AuthHostname  *string
	Region        string
	ClientId      string
	ClientSecret  *string
	EnvironmentId string
	GrantType     string
}

var P1ResourceIDRegexp = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
var IsHostname = regexp.MustCompile(`^(?:[\w-]+\.)+(?:[a-zA-Z]{2,}|xn--[a-zA-Z0-9]+)$`)

func (c *ProfilePingOneCmdAuthConfig) ConfigurePingOneProfile(ctx context.Context) error {

	c.promptForConfig()

	fmt.Println("Testing connection...")

	if err := c.testConnection(ctx); err != nil {
		return err
	}

	fmt.Println("Testing successful.  Saving configuration...")

	return nil
}

func (c *ProfilePingOneCmdAuthConfig) testConnection(ctx context.Context) error {

	clientCfg := client.Config{
		ClientID:             c.ClientId,
		ClientSecret:         *c.ClientSecret,
		EnvironmentID:        c.EnvironmentId,
		Region:               c.Region,
		APIHostnameOverride:  c.APIHostname,
		AuthHostnameOverride: c.AuthHostname,
	}

	client, err := clientCfg.APIClient(ctx, "dev")
	if err != nil {
		return fmt.Errorf("Error connecting to PingOne service: %s", err)
	}

	if client == nil {
		return fmt.Errorf("Error connecting to PingOne service.  Please check your configuration and try again.")
	}

	return nil
}

func (c *ProfilePingOneCmdAuthConfig) promptForConfig() {

	c.EnvironmentId = ux.PromptGetString(ux.PromptContentString{
		Label:           "PingOne environment ID that contains the client to authenticate the CLI with:",
		RegexValidation: P1ResourceIDRegexp,
	})

	c.ClientId = ux.PromptGetString(ux.PromptContentString{
		Label:           "PingOne client ID to authenticate the CLI with:",
		RegexValidation: P1ResourceIDRegexp,
	})

	c.GrantType = ux.PromptGetSelect(ux.PromptContentSelect{
		Label: "What is the grant type to use for the client?",
		Options: []string{
			"Client Credentials",
			// "Authorization Code",
		},
	})

	if c.GrantType == "Client Credentials" {
		clientSecret := ux.PromptGetString(ux.PromptContentString{
			Label:     "PingOne client secret to authenticate the CLI with, using the client credentials grant:",
			Sensitive: true,
		})

		c.ClientSecret = &clientSecret
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
			RegexValidation: IsHostname,
		})
		c.AuthHostname = &authHostname

		apiHostname := ux.PromptGetString(ux.PromptContentString{
			Label:           "PingOne API service hostname (e.g. api.pingone.eu):",
			RegexValidation: IsHostname,
		})
		c.APIHostname = &apiHostname
	} else {
		regionLabels := map[string]string{
			"Europe":        "Europe",
			"North America": "NorthAmerica",
			"Asia Pacific":  "AsiaPacific",
			"Canada":        "Canada",
		}

		c.Region = regionLabels[regionType]
	}

}
