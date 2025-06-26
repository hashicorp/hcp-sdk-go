// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"crypto/tls"
	"testing"

	"github.com/hashicorp/hcp-sdk-go/config/geography"
	"github.com/hashicorp/hcp-sdk-go/profile"
	requirepkg "github.com/stretchr/testify/require"
)

var (
	usConfig = geography.NewConfigUS()
)

func TestNew_Default(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config, err := NewHCPConfig(WithClientCredentials("my-client-id", "my-client-secret"))
	require.NoError(err)

	// Ensure that the default configuration contains the default values
	require.Equal(usConfig.PortalURL, config.PortalURL().String())
	require.Equal(usConfig.APIAddress, config.APIAddress())
	require.Equal(usConfig.SCADAAddress, config.SCADAAddress())

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
		WithProfile(&profile.UserProfile{OrganizationID: "org-id-123", ProjectID: "proj-id-123"}),
	)

	require.NoError(err)
	// Ensure the values have been set accordingly
	require.Equal("https://my-portal:1234", config.PortalURL().String())
	require.Equal("my-api:2345", config.APIAddress())
	require.Equal("my-scada:3456", config.SCADAAddress())
	require.Equal("org-id-123", config.Profile().OrganizationID)
	require.Equal("proj-id-123", config.Profile().ProjectID)
}

func TestNew_WithGeography_US(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config, err := NewHCPConfig(
		WithGeography("us"),
	)

	require.NoError(err)

	// Ensure the values have been set accordingly
	require.Equal("api.cloud.hashicorp.com:443", config.APIAddress())
	require.Equal("https://portal.cloud.hashicorp.com", config.PortalURL().String())
	require.Equal("scada.hashicorp.cloud:7224", config.SCADAAddress())
}

func TestNew_WithGeography_EU(t *testing.T) {
	require := requirepkg.New(t)

	// Exercise
	config, err := NewHCPConfig(
		WithGeography("eu"),
	)

	require.NoError(err)

	// Ensure the values have been set accordingly
	require.Equal("api.cloud.eu.hashicorp.com", config.APIAddress())
	require.Equal("https://portal.cloud.eu.hashicorp.com", config.PortalURL().String())
	require.Equal("", config.SCADAAddress())
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
		{
			name: "empty project ID with populated org ID",
			options: []HCPConfigOption{
				WithClientCredentials("my-client-id", "my-client-secret"),
				WithProfile(&profile.UserProfile{
					OrganizationID: "abc123",
				}),
			},
			expectedError: "the configuration is not valid: when setting a user profile, both organization ID and project ID must be provided",
		},
		{
			name: "empty org ID with populated project ID",
			options: []HCPConfigOption{
				WithClientCredentials("my-client-id", "my-client-secret"),
				WithProfile(&profile.UserProfile{
					ProjectID: "abc123",
				}),
			},
			expectedError: "the configuration is not valid: when setting a user profile, both organization ID and project ID must be provided",
		},
		{
			name: "unsupported geography config",
			options: []HCPConfigOption{
				WithGeography("ap"),
			},
			expectedError: "failed to apply configuration option: hcp geography invalid. Supported: [eu us]",
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

// This test will open your default browser and ask you to login with HCP.
func TestNew_NoConfigPassed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode (CI).")
	}
	require := requirepkg.New(t)

	// Exercise
	config, err := NewHCPConfig()
	require.NoError(err)

	// Ensure the values have been set accordingly
	token, err := config.Token()
	require.NoError(err)
	require.NotNil(token)

	require.Equal(usConfig.PortalURL, config.PortalURL().String())
	require.Equal(usConfig.APIAddress, config.APIAddress())
	require.Equal(usConfig.SCADAAddress, config.SCADAAddress())
}
