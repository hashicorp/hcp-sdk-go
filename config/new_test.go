package config

import (
	"crypto/tls"
	"testing"

	"github.com/hashicorp/hcp-sdk-go/auth"
	requirepkg "github.com/stretchr/testify/require"
)

func TestNew_Default(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config, err := NewHCPConfig(WithClientCredentials("my-client-id", "my-client-secret"))
	require.NoError(err)

	// Ensure that the default configuration contains the default values
	require.Equal(defaultPortalURL, config.PortalURL().String())
	require.Equal(defaultAPIAddress, config.APIAddress())
	require.Equal(defaultSCADAAddress, config.SCADAAddress())

	// Ensure the default configuration uses secure TLS
	require.NotNil(config.APITLSConfig())
	require.False(config.APITLSConfig().InsecureSkipVerify)
	require.NotNil(config.SCADATLSConfig())
	require.False(config.SCADATLSConfig().InsecureSkipVerify)
}

func TestNew_Options(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config, err := NewHCPConfig(
		WithClientCredentials("my-client-id", "my-client-secret"),
		WithPortalURL("https://my-portal:1234"),
		WithAPI("my-api:2345", &tls.Config{}),
		WithSCADA("my-scada:3456", &tls.Config{}),
		WithProfile("org-id-234", "proj-id-2345", "dev-env"),
	)

	require.NoError(err)
	// Ensure the values have been set accordingly
	require.Equal("https://my-portal:1234", config.PortalURL().String())
	require.Equal("my-api:2345", config.APIAddress())
	require.Equal("my-scada:3456", config.SCADAAddress())
}

func TestNew_MockSession(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config, err := NewHCPConfig(
		WithSession(&auth.MockSession{}),
	)

	require.NoError(err)
	// Ensure the values have been set accordingly
	tok, err := config.Token()
	require.NoError(err)
	require.Equal("SomeAccessToken", tok.AccessToken)
	require.Equal("SomeRefreshToken", tok.RefreshToken)
}

func TestNew_Invalid(t *testing.T) {
	testCases := []struct {
		name          string
		options       []HCPConfigOption
		expectedError string
	}{
		{
			name: "empty portal URL",
			options: []HCPConfigOption{
				WithClientCredentials("my-client-id", "my-client-secret"),
				WithPortalURL(""),
			},
			expectedError: "the configuration is not valid: the portal URL has to be non-empty",
		},
		{
			name: "empty auth URL",
			options: []HCPConfigOption{
				WithClientCredentials("my-client-id", "my-client-secret"),
				WithAuth("", nil),
			},
			expectedError: "the configuration is not valid: the auth URL has to be non-empty",
		},
		{
			name: "auth URL with invalid scheme",
			options: []HCPConfigOption{
				WithClientCredentials("my-client-id", "my-client-secret"),
				WithAuth("http://auth", nil),
			},
			expectedError: "the configuration is not valid: the auth URL has to use https at its scheme",
		},
		{
			name: "empty API address",
			options: []HCPConfigOption{
				WithClientCredentials("my-client-id", "my-client-secret"),
				WithAPI("", nil),
			},
			expectedError: "the configuration is not valid: the API address has to be non-empty",
		},
		{
			name: "empty SCADA address",
			options: []HCPConfigOption{
				WithClientCredentials("my-client-id", "my-client-secret"),
				WithSCADA("", nil),
			},
			expectedError: "the configuration is not valid: the SCADA address has to be non-empty",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			require := requirepkg.New(t)

			_, err := NewHCPConfig(testCase.options...)
			require.EqualError(err, testCase.expectedError)
		})
	}
}

// TODO: uncomment these tests when we've set up a CI test user to login via browser (HCE-606)
// func TestNew_NoConfigPassed(t *testing.T) {
// 	require := requirepkg.New(t)

// 	// Exercise
// 	config, err := NewHCPConfig()
// 	require.NoError(err)

// 	// Ensure the values have been set accordingly
// 	token, err := config.Token()
// 	require.NoError(err)
// 	require.NotNil(token)

// 	require.Equal(defaultPortalURL, config.PortalURL().String())
// 	require.Equal(defaultAPIAddress, config.APIAddress())
// 	require.Equal(defaultSCADAAddress, config.SCADAAddress())
// }
