package config

import (
	"crypto/tls"
	"fmt"
	"net/url"
)

// WithClientCredentials credentials is an option that can be used to set
// HCP client credentials on the configuration.
func WithClientCredentials(clientID, clientSecret string) HCPConfigOption {
	return func(config *hcpConfig) error {
		config.clientCredentialsConfig.ClientID = clientID
		config.clientCredentialsConfig.ClientSecret = clientSecret

		return nil
	}
}

// WithAPI credentials is an option that can be used to provide a custom
// configuration for the API endpoint.
//
// If nil is provided for the tlsConfig value, TLS will be disabled.
//
// This should only be necessary for development purposes.
func WithAPI(address string, tlsConfig *tls.Config) HCPConfigOption {
	return func(config *hcpConfig) error {
		config.apiAddress = address
		config.apiTLSConfig = cloneTLSConfig(tlsConfig)

		return nil
	}
}

// WithSCADA credentials is an option that can be used to provide a custom
// configuration for the SCADA endpoint.
//
// If nil is provided for the tlsConfig value, TLS will be disabled.
//
// This should only be necessary for development purposes.
func WithSCADA(address string, tlsConfig *tls.Config) HCPConfigOption {
	return func(config *hcpConfig) error {
		config.scadaAddress = address
		config.scadaTLSConfig = cloneTLSConfig(tlsConfig)

		return nil
	}
}

// WithPortalURL credentials is an option that can be used to provide a custom
// URL for the portal.
//
// This should only be necessary for development purposes.
func WithPortalURL(portalURL string) HCPConfigOption {
	return func(config *hcpConfig) error {
		parsedPortalURL, err := url.Parse(portalURL)
		if err != nil {
			return fmt.Errorf("failed to parse portal URL: %w", err)
		}

		config.portalURL = parsedPortalURL

		return nil
	}
}

// WithAuth credentials is an option that can be used to provide a custom URL
// for the auth endpoint.
//
// An alternative TLS configuration can be provided, if non is provided the
// default TLS configuration will be used. It is not possible to disable TLS for
// the auth endpoint.
//
// This should only be necessary for development purposes.
func WithAuth(authURL string, tlsConfig *tls.Config) HCPConfigOption {
	return func(config *hcpConfig) error {
		parsedAuthURL, err := url.Parse(authURL)
		if err != nil {
			return fmt.Errorf("failed to parse auth URL: %w", err)
		}

		// Ensure a TLS configuration is set, as the auth endpoint should always
		// use TLS.
		if tlsConfig == nil {
			tlsConfig = &tls.Config{}
		}

		config.authURL = parsedAuthURL
		config.authTLSConfig = cloneTLSConfig(tlsConfig)

		return nil
	}
}
