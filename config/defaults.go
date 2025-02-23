// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"crypto/tls"
	"fmt"
	"net/url"

	"golang.org/x/oauth2"

	"github.com/hashicorp/hcp-sdk-go/config/geography"
	"github.com/hashicorp/hcp-sdk-go/profile"
)

const (
	// AuthEndpointTokenPath is the path used to retrieve the access token
	AuthEndpointTokenPath = "/oauth2/token"

	// AuthEndpointAuthPath is the auth path for authentication endpoint
	AuthEndpointAuthPath = "/oauth2/auth"

	// APIAudienceID is the API identifier configured in the auth provider and
	// must be provided when requesting an access token for the API. The value
	// is the same regardless of environment
	APIAudienceID = "https://api.hashicorp.cloud"
)

// HCPConfigFromGeography creates a config with defaults configured to interact
// with a specific geography
func HCPConfigFromGeography(geo string) (*hcpConfig, error) {
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

	config := &hcpConfig{
		authURL:       authURL,
		authTLSConfig: &tls.Config{},
		oauth2Config: oauth2.Config{
			ClientID:    geoConfig.OAuth2ClientID,
			RedirectURL: geoConfig.OAuth2RedirectURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:  geoConfig.AuthURL + AuthEndpointAuthPath,
				TokenURL: geoConfig.AuthURL + AuthEndpointTokenPath,
			},
			Scopes: []string{
				"openid",
				"offline_access",
			},
		},
		profile:        &profile.UserProfile{},
		portalURL:      portalURL,
		apiAddress:     geoConfig.APIAddress,
		apiTLSConfig:   &tls.Config{},
		scadaAddress:   geoConfig.SCADAAddress,
		scadaTLSConfig: &tls.Config{},
	}

	return config, nil
}
