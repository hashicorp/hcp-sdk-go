// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"crypto/tls"
	"fmt"
	"net/url"

	"github.com/hashicorp/hcp-sdk-go/config/geography"
	"github.com/hashicorp/hcp-sdk-go/profile"
	"golang.org/x/oauth2"
)

const (
	// AuthEndpointTokenPath is the path used to retrieve the access token
	AuthEndpointTokenPath = "/oauth2/token"

	// AuthEndpointAuthPath is the auth path for authentication endpoint
	AuthEndpointAuthPath = "/oauth2/auth"

	// AuthRedirectURL is the callback URL for OAuth2
	AuthRedirectURL = "http://localhost:8443/oidc/callback"
)

var (
	// OAuth2Scopes are the default scopes used for the OAuth2 config
	OAuth2Scopes = []string{
		"openid",
		"offline_access",
	}
)

// configFromGeography creates a config with defaults configured to interact
// with a specific geography
func configFromGeography(config *hcpConfig, geo geography.Geo) (*hcpConfig, error) {
	geoConfig := &geography.Config{}

	// Get config based on geographical deployment
	switch geo {
	case geography.US:
		geoConfig = geography.NewConfigUS()
	case geography.EU:
		geoConfig = geography.NewConfigEU()
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

	// Oauth2 config
	config.oauth2Config.ClientID = geoConfig.OAuth2ClientID
	config.oauth2Config.RedirectURL = AuthRedirectURL
	config.oauth2Config.Endpoint.AuthURL = geoConfig.AuthURL + AuthEndpointAuthPath
	config.oauth2Config.Endpoint.TokenURL = geoConfig.AuthURL + AuthEndpointTokenPath

	return config, nil
}

// newDefaultConfig builds an empty HCP Config
func newDefaultConfig() *hcpConfig {
	hcpConfig := &hcpConfig{}
	tlsConfig := &tls.Config{}

	hcpConfig.authTLSConfig = tlsConfig
	hcpConfig.apiTLSConfig = tlsConfig
	hcpConfig.scadaTLSConfig = tlsConfig
	hcpConfig.profile = &profile.UserProfile{}
	hcpConfig.oauth2Config = newDefaultOAuth2Config()

	return hcpConfig
}

// newDefaultOAuth2Config builds a default OAuth2 config
func newDefaultOAuth2Config() oauth2.Config {
	return oauth2.Config{
		Scopes:   OAuth2Scopes,
		Endpoint: oauth2.Endpoint{},
	}
}
