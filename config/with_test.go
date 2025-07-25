// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"os"
	"testing"

	requirepkg "github.com/stretchr/testify/require"

	"github.com/hashicorp/hcp-sdk-go/auth"
	"github.com/hashicorp/hcp-sdk-go/config/geography"
	"github.com/hashicorp/hcp-sdk-go/profile"
)

func TestWith_Geography(t *testing.T) {
	t.Run("us", func(t *testing.T) {
		require := requirepkg.New(t)

		// Setup
		config := &hcpConfig{}
		geo := "us"
		expectation := geography.NewConfigUS()

		// Exercise
		require.NoError(apply(config, WithGeography(geo)))

		// Ensure that the expected config values have been set
		require.Equal(expectation.AuthURL, config.authURL.String())
		require.Equal(expectation.PortalURL, config.portalURL.String())
		require.Equal(expectation.APIAddress, config.apiAddress)
		require.Equal(expectation.SCADAAddress, config.scadaAddress)
		require.Equal(expectation.OAuth2ClientID, config.oauth2Config.ClientID)
	})

	t.Run("eu", func(t *testing.T) {
		require := requirepkg.New(t)

		// Setup
		config := &hcpConfig{}
		geo := "eu"
		expectation := geography.NewConfigEU()

		// Exercise
		require.NoError(apply(config, WithGeography(geo)))

		// Ensure that the expected config values have been set
		require.Equal(expectation.AuthURL, config.authURL.String())
		require.Equal(expectation.PortalURL, config.portalURL.String())
		require.Equal(expectation.APIAddress, config.apiAddress)
		require.Equal(expectation.SCADAAddress, config.scadaAddress)
		require.Equal(expectation.OAuth2ClientID, config.oauth2Config.ClientID)
	})

	t.Run("unsupported", func(t *testing.T) {
		require := requirepkg.New(t)

		// Setup
		config := &hcpConfig{}
		geo := "ap"

		// Exercise
		require.Error(apply(config, WithGeography(geo)))

	})
}

func TestWith_ClientCredentials(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, WithClientCredentials("my-client-id", "my-client-secret")))

	// Ensure that the client credentials have been set
	require.Equal("my-client-id", config.clientID)
	require.Equal("my-client-secret", config.clientSecret)
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

func TestWith_Profile(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	profile := &profile.UserProfile{OrganizationID: "org-id-1234", ProjectID: "project-id-1234"}
	require.NoError(apply(config, WithProfile(profile)))

	// Ensure UserProfile fields match configured values
	require.Equal("org-id-1234", config.Profile().OrganizationID)
	require.Equal("project-id-1234", config.Profile().ProjectID)

}

func TestWithout_BrowserLogin(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, WithoutBrowserLogin()))

	// Ensure  browser login is disabled
	require.True(config.noBrowserLogin)
}

func TestWithout_OpenDefaultBrowser(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, WithoutOpenDefaultBrowser()))

	// Ensure browser login doesn't use the default browser
	require.True(config.noDefaultBrowser)
}

func TestWith_CredentialFile(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise with an invalid credential file
	config := &hcpConfig{}
	cf := &auth.CredentialFile{
		ProjectID: "123",
	}
	require.Error(apply(config, WithCredentialFile(cf)))

	// Exercise with a valid credential file
	cf = &auth.CredentialFile{
		ProjectID: "123",
		Scheme:    auth.CredentialFileSchemeServicePrincipal,
		Oauth: &auth.OauthConfig{
			ClientID:     "123",
			ClientSecret: "456",
		},
	}
	require.NoError(apply(config, WithCredentialFile(cf)))

	// Ensure the cred file is set
	require.Equal(cf, config.credentialFile)
}

func TestWith_CredentialFilePath(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		require := requirepkg.New(t)

		// Write the cred file
		cf := &auth.CredentialFile{
			ProjectID: "123",
			Scheme:    auth.CredentialFileSchemeServicePrincipal,
			Oauth: &auth.OauthConfig{
				ClientID:     "123",
				ClientSecret: "456",
			},
		}
		f, err := os.CreateTemp("", "")
		require.NoError(err)
		require.NoError(auth.WriteCredentialFile(f.Name(), cf))

		// Exercise
		config := &hcpConfig{}
		require.NoError(apply(config, WithCredentialFilePath(f.Name())))

		// Ensure the cred file is set
		require.Equal(cf, config.credentialFile)
	})

	t.Run("not found", func(t *testing.T) {
		// Exercise
		config := &hcpConfig{}
		requirepkg.Error(t, apply(config, WithCredentialFilePath(fmt.Sprintf("random-%d", rand.Int()))))
	})
}
