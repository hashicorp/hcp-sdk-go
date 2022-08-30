package config

import (
	"crypto/tls"
	"testing"

	requirepkg "github.com/stretchr/testify/require"
)

func TestWith_ClientCredentials(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, WithClientCredentials("my-client-id", "my-client-secret")))

	// Ensure that the client credentials have been set
	require.Equal("my-client-id", config.clientCredentialsConfig.ClientID)
	require.Equal("my-client-secret", config.clientCredentialsConfig.ClientSecret)
}

func TestWith_API(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, WithAPI("my-api:1234", &tls.Config{})))

	// Ensure that the API configuration have been set
	require.Equal("my-api:1234", config.apiAddress)
	require.NotNil(config.APITLSConfig())
}

func TestWith_SCADA(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, WithSCADA("my-scada:1234", &tls.Config{})))

	// Ensure that the SCADA configuration have been set
	require.Equal("my-scada:1234", config.scadaAddress)
	require.NotNil(config.SCADATLSConfig())
}

func TestWith_PortalURL(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, WithPortalURL("http://my-portal:1234")))

	// Ensure that the portal URL has been set
	require.Equal("http://my-portal:1234", config.portalURL.String())
}

func TestWith_Auth(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, WithAuth("http://my-auth:1234", nil)))

	// Ensure that the portal URL has been set
	require.Equal("http://my-auth:1234", config.authURL.String())

	// Ensure OAuth2 config is updated with custom auth URL
	require.Equal("http://my-auth:1234/oauth2/auth", config.oauth2Config.Endpoint.AuthURL)
	require.Equal("http://my-auth:1234/oauth2/token", config.oauth2Config.Endpoint.TokenURL)

	// Ensure auth TLS is configured
	require.NotNil(config.authTLSConfig)
}

func TestWith_Auth_CustomTLSConfig(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, WithAuth("http://my-auth:1234", &tls.Config{InsecureSkipVerify: true})))

	// Ensure auth TLS has custom configuration
	require.True(config.authTLSConfig.InsecureSkipVerify)
}

func TestWith_OAuth2ClientID(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, WithOAuth2ClientID("1a2b3c4d")))

	// Ensure oauth2 config is configured with custom OAuth2 client ID
	require.Equal("1a2b3c4d", config.oauth2Config.ClientID)
}
