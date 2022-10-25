package config

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/hcp-sdk-go/auth"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	// defaultAuthURL is the URL of the production auth endpoint.
	defaultAuthURL = "https://auth.idp.hashicorp.com"

	// defaultOAuth2ClientID is the client ID of the production auth application.
	defaultOAuth2ClientID = "21d86262-6f14-4a30-a90f-07e3fde8b23d"

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
	// Parse default URLs
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

		portalURL: portalURL,

		apiAddress:   defaultAPIAddress,
		apiTLSConfig: &tls.Config{},

		scadaAddress:   defaultSCADAAddress,
		scadaTLSConfig: &tls.Config{},
	}

	// Apply all options
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, fmt.Errorf("failed to apply configuration option: %w", err)
		}
	}

	// Set up a token context with the custom auth TLS config
	tokenTransport := cleanhttp.DefaultPooledTransport()
	tokenTransport.TLSClientConfig = config.authTLSConfig
	tokenContext := context.WithValue(
		context.Background(),
		oauth2.HTTPClient,
		&http.Client{Transport: tokenTransport},
	)

	// Set access token via configured client credentials.
	if config.clientCredentialsConfig.ClientID != "" && config.clientCredentialsConfig.ClientSecret != "" {
		// Set token URL based on auth URL
		tokenURL := config.authURL
		tokenURL.Path = tokenPath
		config.clientCredentialsConfig.TokenURL = tokenURL.String()

		// Create token source from the client credentials configuration
		config.tokenSource = config.clientCredentialsConfig.TokenSource(tokenContext)

	} else { // Set access token via browser login.

		// TODO: Right now we fetch a new token on every init of the client. We need to implement a library that will check for existing tokens in a well-known location.
		// If no token is available or if the available token's max age has exceeded,then we get new token via browser login.
		cache, _ := auth.Read()
		//var tok *oauth2.Token
		var tok *oauth2.Token
		if cache.SessionExpiry.Before(time.Now()) {
			var err error
			tok, err = auth.GetTokenFromBrowser(tokenContext, &config.oauth2Config)
			if err != nil {
				return nil, fmt.Errorf("failed to get access token: %w", err)
			}
		} else {
			tok = &oauth2.Token{
				AccessToken:  cache.AccessToken,
				RefreshToken: cache.RefreshToken,
				Expiry:       cache.AccessTokenExpiry,
			}
		}

		config.tokenSource = config.oauth2Config.TokenSource(tokenContext, tok)
	}

	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("the configuration is not valid: %w", err)
	}

	return config, nil
}
