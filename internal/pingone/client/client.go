package client

import (
	"context"
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/pingone"
)

type Client struct {
	API         *pingone.Client
	ForceDelete bool
}

func (c *Config) APIClient(ctx context.Context, version string) (*Client, error) {

	userAgent := fmt.Sprintf("pingctl/%s/go", version)

	if v := c.Validate(); v != nil {
		return nil, v
	}

	config := &pingone.Config{
		ClientID:             c.ClientID,
		ClientSecret:         c.ClientSecret,
		EnvironmentID:        c.EnvironmentID,
		AccessToken:          c.AccessToken,
		Region:               c.Region,
		APIHostnameOverride:  c.APIHostnameOverride,
		AuthHostnameOverride: c.AuthHostnameOverride,
		UserAgentOverride:    &userAgent,
	}

	client, err := config.APIClient(ctx)
	if err != nil {
		return nil, err
	}

	cliClient := &Client{
		API:         client,
		ForceDelete: c.ForceDelete,
	}

	return cliClient, nil
}
