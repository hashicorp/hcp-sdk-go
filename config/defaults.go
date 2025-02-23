// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"fmt"
	"net/url"

	"github.com/hashicorp/hcp-sdk-go/config/geography"
)

const (
	// AuthEndpointTokenPath is the path used to retrieve the access token
	AuthEndpointTokenPath = "/oauth2/token"

	// AuthEndpointAuthPath is the auth path for authentication endpoint
	AuthEndpointAuthPath = "/oauth2/auth"
)

var (
	// OAuth2Scopes are the default scopes used for the OAuth2 config
	OAuth2Scopes = []string{
		"openid",
		"offline_access",
	}
)

// HCPConfigFromGeography creates a config with defaults configured to interact
// with a specific geography
func HCPConfigFromGeography(config *hcpConfig, geo string) (*hcpConfig, error) {
	geoConfig := &geography.ConnectionConfig{}

	// Get config based on geographical deployment
	switch geo {
	case geography.NorthAmerica:
		geoConfig = geography.NewConnectionConfigNorthAmerica()
	case geography.Europe:
		geoConfig = geography.NewConnectionConfigEurope()
	default:
		return nil, fmt.Errorf("hcp geography invalid. Supported: %v", geography.Geographies)
	}

	// Parse URLs
	authURL, _ := url.Parse(geoConfig.AuthURL)
	portalURL, _ := url.Parse(geoConfig.PortalURL)

	// Override geography-specific parameters
	config.authURL = authURL
	config.portalURL = portalURL
	config.apiAddress = geoConfig.APIAddress
	config.scadaAddress = geoConfig.SCADAAddress

	config.oauth2Config.Scopes = OAuth2Scopes
	config.oauth2Config.ClientID = geoConfig.OAuth2ClientID
	config.oauth2Config.RedirectURL = geoConfig.OAuth2RedirectURL

	config.oauth2Config.Endpoint.AuthURL = geoConfig.AuthURL + AuthEndpointAuthPath
	config.oauth2Config.Endpoint.TokenURL = geoConfig.AuthURL + AuthEndpointTokenPath

	return config, nil
}
