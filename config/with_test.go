// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/hashicorp/hcp-sdk-go/auth"
	"github.com/hashicorp/hcp-sdk-go/profile"
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

func TestWith_Session(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config := &hcpConfig{}
	require.NoError(apply(config, WithSession(&auth.MockSession{})))

	// Ensure Sessions is an empty MockSession object
	require.Equal(&auth.MockSession{}, config.session)
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
			ClientID: "123",
			SecretID: "456",
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
				ClientID: "123",
				SecretID: "456",
			},
		}
		f, err := ioutil.TempFile("", "")
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
