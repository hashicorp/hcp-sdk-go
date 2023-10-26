// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"crypto/tls"
	"fmt"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/hcp-sdk-go/auth"
	"github.com/hashicorp/hcp-sdk-go/auth/tokencache"
	"github.com/hashicorp/hcp-sdk-go/auth/workload"
	"github.com/hashicorp/hcp-sdk-go/profile"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
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

type sourceType = string

var (
	sourceTypeLogin            = sourceType("login")
	sourceTypeServicePrincipal = sourceType("service-principal")
	sourceTypeWorkload         = sourceType("workload")
)

func (c *hcpConfig) setTokenSource() error {
	// Get the credential cache path
	// TODO: make this configurable
	userHome, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to retrieve user's home directory path: %v", err)
	}

	cacheFile := path.Join(userHome, ".config/hcp/creds-cache.json")

	tokenSource, sourceType, sourceIdentifier, err := c.getTokenSource()
	if err != nil {
		return err
	}

	switch sourceType {
	case sourceTypeLogin:
		c.tokenSource = tokencache.NewLoginTokenSource(
			cacheFile,
			c.forceLogin,
			tokenSource,
			&c.oauth2Config,
		)
	case sourceTypeServicePrincipal:
		c.tokenSource = tokencache.NewServicePrincipalTokenSource(
			cacheFile,
			sourceIdentifier,
			tokenSource,
		)
	case sourceTypeWorkload:
		c.tokenSource = tokencache.NewWorkloadTokenSource(
			cacheFile,
			sourceIdentifier,
			tokenSource,
		)
	}

	return nil
}

// sgetTokenSource gets the token source. The order of precedence is:
//
// 1. Configured client credentials (either explicit or through environment
// variables).
// 2. Via credential file (sourced first via environment variable and then
// default file location).
// 3. Interactive session.
func (c *hcpConfig) getTokenSource() (oauth2.TokenSource, sourceType, string, error) {
	// Set up a token context with the custom auth TLS config
	tokenTransport := cleanhttp.DefaultPooledTransport()
	tokenTransport.TLSClientConfig = c.authTLSConfig
	ctx := context.WithValue(
		context.Background(),
		oauth2.HTTPClient,
		&http.Client{Transport: tokenTransport},
	)

	// Set client credentials token URL based on auth URL.
	tokenURL := c.authURL
	tokenURL.Path = tokenPath

	clientCredentials := clientcredentials.Config{
		EndpointParams: url.Values{"audience": {aud}},
		TokenURL:       tokenURL.String(),
	}

	// Set access token via configured client credentials.
	if c.clientID != "" && c.clientSecret != "" {
		// Create token source for client secrets
		clientCredentials.ClientID = c.clientID
		clientCredentials.ClientSecret = c.clientSecret

		return clientCredentials.TokenSource(ctx), sourceTypeServicePrincipal, clientCredentials.ClientID, nil
	}

	// Use workload provider config if it was provided
	if c.workloadProviderConfig != nil {
		provider, err := workload.New(c.workloadProviderConfig)
		if err != nil {
			return nil, "", "", err
		}
		provider.SetAPI(c)
		return oauth2.ReuseTokenSource(nil, provider), sourceTypeWorkload, c.workloadProviderConfig.ProviderResourceName, nil
	}

	// If we haven't been given an explicit credential file to use, try to load
	// the credential file from the environment or default location.
	if c.credentialFile == nil {
		credFile, err := auth.GetDefaultCredentialFile()
		if err != nil {
			return nil, "", "", err
		}
		c.credentialFile = credFile
	}

	// If we found a credential file use it as a credential source
	if c.credentialFile != nil {
		if c.credentialFile.Scheme == auth.CredentialFileSchemeServicePrincipal {
			// Set credentials on client credentials configuration
			clientCredentials.ClientID = c.credentialFile.Oauth.ClientID
			clientCredentials.ClientSecret = c.credentialFile.Oauth.ClientSecret

			// Create token source from the client credentials configuration.
			return clientCredentials.TokenSource(ctx), sourceTypeServicePrincipal, clientCredentials.ClientID, nil
		} else if c.credentialFile.Scheme == auth.CredentialFileSchemeWorkload {
			w, err := workload.New(c.credentialFile.Workload)
			if err != nil {
				return nil, "", "", err
			}

			// Set the API info
			w.SetAPI(c)

			// Use the workload provider as the token source
			return w, sourceTypeWorkload, c.credentialFile.Workload.ProviderResourceName, nil
		}
	}

	var loginTokenSource oauth2.TokenSource
	if !c.noBrowserLogin {
		loginTokenSource = auth.NewBrowserLogin(&c.oauth2Config)
	}

	return loginTokenSource, sourceTypeLogin, "", nil
}
