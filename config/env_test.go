// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
	clearEnv()
	defer clearEnv()

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, FromEnv()))

	// Ensure the config has not been modified
	require.Empty(config)
}

func TestFromEnv_SimpleValues(t *testing.T) {
	require := requirepkg.New(t)

	// Clear the environment
	clearEnv()
	defer clearEnv()

	// Prepare the environment
	require.NoError(os.Setenv(envVarAuthURL, "https://my-auth:1234"))

	require.NoError(os.Setenv(envVarClientID, "my-client-id"))
	require.NoError(os.Setenv(envVarClientSecret, "my-client-secret"))

	require.NoError(os.Setenv(envVarPortalURL, "http://my-portal:2345"))

	require.NoError(os.Setenv(envVarOAuth2ClientID, "1a2b3c4d"))

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

	// Ensure the oauth2 config is set correctly
	require.Equal("1a2b3c4d", config.oauth2Config.ClientID)
	require.Equal("https://my-auth:1234/oauth2/auth", config.oauth2Config.Endpoint.AuthURL)
	require.Equal("https://my-auth:1234/oauth2/token", config.oauth2Config.Endpoint.TokenURL)

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

	// Clear the environment
	clearEnv()
	defer clearEnv()

	// Prepare the environment
	require.NoError(os.Setenv(envVarAPIHostnameLegacy, "https://my-legacy-api"))

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, FromEnv()))

	// Ensure the API address is set correctly
	require.Equal("my-legacy-api", config.apiAddress)
}

func TestFromEnv_TLSConfig_Plain(t *testing.T) {
	require := requirepkg.New(t)

	// Clear the environment
	clearEnv()
	defer clearEnv()

	// Prepare the environment
	require.NoError(os.Setenv(envVarAPITLS, tlsSettingDisabled))
	require.NoError(os.Setenv(envVarSCADATLS, tlsSettingDisabled))
	require.NoError(os.Setenv(envVarAuthTLS, tlsSettingDisabled))

	// Exercise
	config := &hcpConfig{
		apiTLSConfig:   &tls.Config{},
		scadaTLSConfig: &tls.Config{},
		authTLSConfig:  &tls.Config{},
	}
	require.NoError(apply(config, FromEnv()))

	// Ensure the TLS configuration has been reset
	require.Nil(config.apiTLSConfig)
	require.Nil(config.scadaTLSConfig)
	require.Nil(config.authTLSConfig)
}

func TestFromEnv_TLSConfig_Insecure(t *testing.T) {
	require := requirepkg.New(t)

	// Clear the environment
	clearEnv()
	defer clearEnv()

	// Prepare the environment
	require.NoError(os.Setenv(envVarAPITLS, tlsSettingInsecure))
	require.NoError(os.Setenv(envVarSCADATLS, tlsSettingInsecure))
	require.NoError(os.Setenv(envVarAuthTLS, tlsSettingInsecure))

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, FromEnv()))

	// Ensure the TLS configuration is set to insecure
	require.True(config.apiTLSConfig.InsecureSkipVerify)
	require.True(config.scadaTLSConfig.InsecureSkipVerify)
	require.True(config.authTLSConfig.InsecureSkipVerify)
}

// clearEnv will unset any environment variables that FromEnv might read.
func clearEnv() {
	os.Unsetenv(envVarAuthURL)
	os.Unsetenv(envVarClientID)
	os.Unsetenv(envVarClientSecret)
	os.Unsetenv(envVarOAuth2ClientID)
	os.Unsetenv(envVarPortalURL)
	os.Unsetenv(envVarAPIAddress)
	os.Unsetenv(envVarAPIHostnameLegacy)
	os.Unsetenv(envVarAPITLS)
	os.Unsetenv(envVarSCADAAddress)
	os.Unsetenv(envVarSCADATLS)
}
