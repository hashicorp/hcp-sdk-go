package config

import (
	"crypto/tls"
	"os"
	"testing"

	requirepkg "github.com/stretchr/testify/require"
)

func TestFromEnv_NoVars(t *testing.T) {
	require := requirepkg.New(t)

	// Clear the environment
	os.Clearenv()

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, FromEnv()))

	// Ensure the config has not been modified
	require.Empty(config)
}

func TestFromEnv_SimpleValues(t *testing.T) {
	require := requirepkg.New(t)

	// Prepare the environment
	os.Clearenv()
	require.NoError(os.Setenv(envVarAuthURL, "https://my-auth:1234"))

	require.NoError(os.Setenv(envVarClientID, "my-client-id"))
	require.NoError(os.Setenv(envVarClientSecret, "my-client-secret"))

	require.NoError(os.Setenv(envVarPortalURL, "http://my-portal:2345"))

	require.NoError(os.Setenv(envVarAPIAddress, "my-api:3456"))
	require.NoError(os.Setenv(envVarSCADAAddress, "my-scada:4567"))

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, FromEnv()))

	// Ensure the auth URL is set correctly
	require.Equal("https", config.authURL.Scheme)
	require.Equal("my-auth:1234", config.authURL.Host)

	// Ensure the client credentials are set correctly
	require.Equal("my-client-id", config.clientCredentialsConfig.ClientID)
	require.Equal("my-client-secret", config.clientCredentialsConfig.ClientSecret)

	// Ensure the portal URL is set correctly
	require.Equal("http", config.portalURL.Scheme)
	require.Equal("my-portal:2345", config.portalURL.Host)

	// Ensure the API address is set correctly
	require.Equal("my-api:3456", config.apiAddress)

	// Ensure the SCADA address is set correctly
	require.Equal("my-scada:4567", config.scadaAddress)
}

func TestFromEnv_LegacyHostname(t *testing.T) {
	require := requirepkg.New(t)

	// Prepare the environment
	os.Clearenv()
	require.NoError(os.Setenv(envVarAPIHostnameLegacy, "https://my-legacy-api"))

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, FromEnv()))

	// Ensure the API address is set correctly
	require.Equal("my-legacy-api", config.apiAddress)
}

func TestFromEnv_TLSConfig_Plain(t *testing.T) {
	require := requirepkg.New(t)

	// Prepare the environment
	os.Clearenv()
	require.NoError(os.Setenv(envVarAPITLS, tlsSettingPlain))
	require.NoError(os.Setenv(envVarSCADATLS, tlsSettingPlain))

	// Exercise
	config := &hcpConfig{
		apiTLSConfig:   &tls.Config{},
		scadaTLSConfig: &tls.Config{},
	}
	require.NoError(apply(config, FromEnv()))

	// Ensure the TLS configuration has been reset
	require.Nil(config.apiTLSConfig)
	require.Nil(config.scadaTLSConfig)
}

func TestFromEnv_TLSConfig_Insecure(t *testing.T) {
	require := requirepkg.New(t)

	// Prepare the environment
	os.Clearenv()
	require.NoError(os.Setenv(envVarAPITLS, tlsSettingInsecure))
	require.NoError(os.Setenv(envVarSCADATLS, tlsSettingInsecure))

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, FromEnv()))

	// Ensure the TLS configuration is set to insecure
	require.True(config.apiTLSConfig.InsecureSkipVerify)
	require.True(config.scadaTLSConfig.InsecureSkipVerify)
}
