// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/hcp-sdk-go/auth"
	"github.com/hashicorp/hcp-sdk-go/auth/workload"
	"github.com/hashicorp/hcp-sdk-go/profile"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	// defaultAuthURL is the URL of the production auth endpoint.
	defaultAuthURL = "https://auth.idp.hashicorp.com"

	// defaultOAuth2ClientID is the client ID of the production auth application.
	defaultOAuth2ClientID = "4edd6521-6eb9-4d78-9039-7ce8569d667c"

	// defaultPortalURL is the URL of the production portal.
	defaultPortalURL = "https://portal.cloud.hashicorp.com"

	// defaultAPIAddress is the address of the production API.
	defaultAPIAddress = "api.cloud.hashicorp.com:443"

	// defaultSCADAAddress is the address of the production SCADA endpoint.
	defaultSCADAAddress = "scada.hashicorp.cloud:7224"

	// The audience is the API identifier configured in the auth provider and
	// must be provided when requesting an access token for the API. The value
	// is the same regardless of environment.
	aud = "https://api.hashicorp.cloud"

	// tokenPath is the path used to retrieve the access token.
	tokenPath string = "/oauth2/token"
)

// NewHCPConfig will return a HCPConfig. The configuration will be constructed
// from default values and the passed options.
//
// Based on the provided options the configuration values can be provided
// directly or will be read from other places (e.g. the environment).
//
// The configuration will default to values for the HCP production environment,
// but can be overwritten for development purposes.
//
// In addition to the default values the configuration requires client
// credentials to be provided through one of the passed options (e.g. by using
// WithCredentials or by using FromEnv and providing the client credentials via
// environment variables).
func NewHCPConfig(opts ...HCPConfigOption) (HCPConfig, error) {
	// Parse default URLsG
	authURL, _ := url.Parse(defaultAuthURL)
	portalURL, _ := url.Parse(defaultPortalURL)

	// Prepare basic config with default values.
	config := &hcpConfig{
		clientCredentialsConfig: clientcredentials.Config{
			EndpointParams: url.Values{"audience": {aud}},
		},

		authURL:       authURL,
		authTLSConfig: &tls.Config{},
		oauth2Config: oauth2.Config{
			ClientID: defaultOAuth2ClientID,
			Endpoint: oauth2.Endpoint{
				AuthURL:  defaultAuthURL + "/oauth2/auth",
				TokenURL: defaultAuthURL + "/oauth2/token",
			},
			RedirectURL: "http://localhost:8443/oidc/callback",
			Scopes:      []string{"openid", "offline_access"},
		},
		session: &auth.UserSession{},
		profile: &profile.UserProfile{},

		portalURL: portalURL,

		apiAddress:   defaultAPIAddress,
		apiTLSConfig: &tls.Config{},

		scadaAddress:   defaultSCADAAddress,
		scadaTLSConfig: &tls.Config{},
	}

	if config.suppressLogging {
		log.SetOutput(io.Discard)
	}

	if config.noBrowserLogin {
		config.session = &auth.UserSession{
			NoBrowserLogin: true,
		}
	}

	// Apply all options
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, fmt.Errorf("failed to apply configuration option: %w", err)
		}
	}

	// Configure the token source
	if err := config.setTokenSource(); err != nil {
		return nil, err
	}

	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("the configuration is not valid: %w", err)
	}

	return config, nil
}

// setTokenSource sets the token source. If the token source has been explictely
// set, it is a no-op. Otherwise the order of precedence is:
//
// 1. Configured client credentials (either explicit or through environment
// variables).
// 2. Via credential file (sourced first via environment variable and then
// default file location).
// 3. Interactive session.
func (c *hcpConfig) setTokenSource() error {
	// token source is already explicitly configured.
	if c.tokenSource != nil {
		return nil
	}

	// Set up a token context with the custom auth TLS config
	tokenTransport := cleanhttp.DefaultPooledTransport()
	tokenTransport.TLSClientConfig = c.authTLSConfig
	ctx := context.WithValue(
		context.Background(),
		oauth2.HTTPClient,
		&http.Client{Transport: tokenTransport},
	)

	// Set access token via configured client credentials.
	if c.clientCredentialsConfig.ClientID != "" && c.clientCredentialsConfig.ClientSecret != "" {
		c.configureClientCredentialTokenSource(ctx, c.clientCredentialsConfig.ClientID, c.clientCredentialsConfig.ClientSecret)
		return nil
	}

	// If we haven't been given an explicit credential file to use, try to load
	// the credential file from the environment or default location.
	if c.credentialFile == nil {
		credFile, err := auth.GetDefaultCredentialFile()
		if err != nil {
			return err
		}
		c.credentialFile = credFile
	}

	// If we found a credential file use it as a credential source
	if c.credentialFile != nil {
		if c.credentialFile.Scheme == auth.CredentialFileSchemeServicePrincipal {
			c.configureClientCredentialTokenSource(ctx, c.credentialFile.Oauth.ClientID, c.credentialFile.Oauth.SecretID)
			return nil
		} else if c.credentialFile.Scheme == auth.CredentialFileSchemeWorkload {
			w, err := workload.New(c.credentialFile.Workload)
			if err != nil {
				return err
			}

			// Set the API info
			w.SetAPI(c)

			// Use the workload provider as the token source
			c.tokenSource = w
		}

		return nil
	}

	// Set access token via browser login or use token from existing session.
	tok, err := c.session.GetToken(ctx, &c.oauth2Config)
	if err != nil {
		return fmt.Errorf("failed to find existing session or set up new: %w", err)
	}

	// Update HCPConfig with most current token values.
	c.tokenSource = c.oauth2Config.TokenSource(ctx, tok)
	return nil
}

// configureClientCredentialTokenSource configures the credential source to use
// the passed client credentials.
func (c *hcpConfig) configureClientCredentialTokenSource(ctx context.Context, clientID, secretID string) {
	// Set token URL based on auth URL.
	tokenURL := c.authURL
	tokenURL.Path = tokenPath
	c.clientCredentialsConfig.TokenURL = tokenURL.String()

	// Create token source from the client credentials configuration.
	c.tokenSource = c.clientCredentialsConfig.TokenSource(ctx)
}
