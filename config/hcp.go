package config

import (
	"crypto/tls"
	"fmt"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// HCPConfig provides configuration values that are useful to interact with HCP.
type HCPConfig interface {
	// TokenSource will return a token that can be used to authenticate to HCP.
	oauth2.TokenSource

	// GetPortalURL will return the URL of the portal.
	//
	// The default portal URL will be the production portal, but the value may
	// be overwritten for development purposes.
	GetPortalURL() *url.URL

	// GetAPIAddress will return the HCP API address (<hostname>[:port]).
	//
	// The default address will be the HCP production API, but it may be
	// overwritten for development purposes.
	GetAPIAddress() string

	// GetAPITLSConfig will return the API TLS configuration.
	//
	// TLS will be enabled by default but may be disabled for development
	// purposes.
	GetAPITLSConfig() *tls.Config

	// GetSCADAAddress will return the SCADA address (<hostname>[:port]).
	//
	// The default address will be the HCP production SCADA endpoint, but it may
	// be overwritten for development purposes.
	GetSCADAAddress() string

	// GetSCADATLSConfig will return the SCADA TLS configuration.
	//
	// TLS will be enabled by default but may be disabled for development
	// purposes.
	GetSCADATLSConfig() *tls.Config
}

type hcpConfig struct {
	// clientCredentialsConfig is the configuration that will be used to create
	// the token source.
	clientCredentialsConfig clientcredentials.Config

	// authURL is the URL that will be used to authenticate.
	authURL *url.URL
	// authTLSConfig is the TLS configuration for the auth endpoint. TLS can not
	// be disabled for the auth endpoint, but the configuration can be set to a
	// custom one for tests or development.
	authTLSConfig *tls.Config

	// tokenSource is a self-refreshing token source that returns an access
	// token that can be used to authenticate against the HCP APIs.
	tokenSource oauth2.TokenSource

	// portalURL is the base URL of the portal.
	portalURL *url.URL

	// apiAddress is the address (<hostname>[:port]) of the HCP API.
	apiAddress string
	// apiTLSConfig is the TLS configuration for the HCP API. It can be set to
	// nil if TLS should be disabled.
	apiTLSConfig *tls.Config

	// scadaAddress is the address (<hostname>[:port]) of the SCADA endpoint.
	scadaAddress string
	// scadaTLSConfig is the TLS configuration for the SCADA endpoint. It can be
	// set to nil if TLS should be disabled.
	scadaTLSConfig *tls.Config
}

func (c hcpConfig) Token() (*oauth2.Token, error) {
	return c.tokenSource.Token()
}

func (c hcpConfig) GetPortalURL() *url.URL {
	// Copy the value in order to not return a pointer to the internal one.
	portalURL := *c.portalURL

	return &portalURL
}

func (c hcpConfig) GetAPIAddress() string {
	return c.apiAddress
}

func (c hcpConfig) GetAPITLSConfig() *tls.Config {
	return c.apiTLSConfig.Clone()
}

func (c hcpConfig) GetSCADAAddress() string {
	return c.scadaAddress
}

func (c hcpConfig) GetSCADATLSConfig() *tls.Config {
	return c.scadaTLSConfig.Clone()
}

func (c hcpConfig) validate() error {
	// Ensure client credentials have been provided
	if c.clientCredentialsConfig.ClientID == "" && c.clientCredentialsConfig.ClientSecret == "" {
		return fmt.Errorf("client credentials need to be provided")
	}

	// Ensure the auth URL is valid
	if c.authURL.Host == "" {
		return fmt.Errorf("the auth URL has to be non-empty")
	}

	if c.authURL.Scheme != "https" {
		return fmt.Errorf("the auth URL has to use https at its scheme")
	}

	// Ensure the portal URL is valid
	if c.portalURL.Host == "" {
		return fmt.Errorf("the portal URL has to be non-empty")
	}

	// Ensure an API address is valid
	if c.apiAddress == "" {
		return fmt.Errorf("the API address has to be non-empty")
	}

	// Ensure an SCADA address is valid
	if c.scadaAddress == "" {
		return fmt.Errorf("the SCADA address has to be non-empty")
	}

	return nil
}

var _ HCPConfig = hcpConfig{}
