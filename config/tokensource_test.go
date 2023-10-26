package config

import (
	"io/ioutil"
	"testing"

	"github.com/hashicorp/hcp-sdk-go/auth"
	"github.com/hashicorp/hcp-sdk-go/auth/workload"
	requirepkg "github.com/stretchr/testify/require"
)

func TestTokenSource_GetTokenSource_WithClientCredentials(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config, err := NewHCPConfig(WithClientCredentials("client-id", "client-secret"))
	require.NoError(err)

	rawConfig, ok := config.(*hcpConfig)
	require.True(ok)

	tokenSource, tokenSourceType, sourceIdentifier, err := rawConfig.getTokenSource()
	require.NoError(err)

	require.NotNil(tokenSource)
	require.Equal(sourceTypeServicePrincipal, tokenSourceType)
	require.Equal("client-id", sourceIdentifier)
}

func TestTokenSource_GetTokenSource_EnvClientCredentials(t *testing.T) {
	require := requirepkg.New(t)

	t.Setenv(envVarClientID, "client-id")
	t.Setenv(envVarClientSecret, "client-secret")

	// Exercise
	config, err := NewHCPConfig(WithClientCredentials("client-id", "client-secret"))
	require.NoError(err)

	rawConfig, ok := config.(*hcpConfig)
	require.True(ok)

	tokenSource, tokenSourceType, sourceIdentifier, err := rawConfig.getTokenSource()
	require.NoError(err)

	require.NotNil(tokenSource)
	require.Equal(sourceTypeServicePrincipal, tokenSourceType)
	require.Equal("client-id", sourceIdentifier)
}

func TestTokenSource_GetTokenSource_CredentialFile_Workload(t *testing.T) {
	require := requirepkg.New(t)

	// Write the cred file
	cf := &auth.CredentialFile{
		ProjectID: "123",
		Scheme:    auth.CredentialFileSchemeWorkload,
		Workload: &workload.IdentityProviderConfig{
			ProviderResourceName: "iam/test",
			AWS:                  &workload.AWSCredentialSource{},
		},
	}

	f, err := ioutil.TempFile("", "")
	require.NoError(err)
	require.NoError(auth.WriteCredentialFile(f.Name(), cf))

	// Exercise
	t.Setenv(auth.EnvHCPCredFile, f.Name())
	config, err := NewHCPConfig()
	require.NoError(err)

	rawConfig, ok := config.(*hcpConfig)
	require.True(ok)

	tokenSource, tokenSourceType, sourceIdentifier, err := rawConfig.getTokenSource()
	require.NoError(err)

	_, ok = tokenSource.(*workload.Provider)
	require.True(ok)
	require.Equal(sourceTypeWorkload, tokenSourceType)
	require.Equal("iam/test", sourceIdentifier)
}

func TestTokenSource_GetTokenSource_CredentialFile_ServicePrincipal(t *testing.T) {
	require := requirepkg.New(t)

	// Write the cred file
	cf := &auth.CredentialFile{
		ProjectID: "123",
		Scheme:    auth.CredentialFileSchemeServicePrincipal,
		Oauth: &auth.OauthConfig{
			ClientID:     "client-id",
			ClientSecret: "client-secret",
		},
	}

	f, err := ioutil.TempFile("", "")
	require.NoError(err)
	require.NoError(auth.WriteCredentialFile(f.Name(), cf))

	// Exercise
	t.Setenv(auth.EnvHCPCredFile, f.Name())
	config, err := NewHCPConfig()
	require.NoError(err)

	rawConfig, ok := config.(*hcpConfig)
	require.True(ok)

	tokenSource, tokenSourceType, sourceIdentifier, err := rawConfig.getTokenSource()
	require.NoError(err)

	require.NotNil(tokenSource)
	require.Equal(sourceTypeServicePrincipal, tokenSourceType)
	require.Equal("client-id", sourceIdentifier)
}

func TestTokenSource_GetTokenSource_Login(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config, err := NewHCPConfig()
	require.NoError(err)

	rawConfig, ok := config.(*hcpConfig)
	require.True(ok)

	tokenSource, tokenSourceType, _, err := rawConfig.getTokenSource()
	require.NoError(err)

	require.NotNil(tokenSource)
	require.Equal(sourceTypeLogin, tokenSourceType)
}
