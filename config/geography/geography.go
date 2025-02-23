// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geography

const (
	// Default is the HCP geography that will be used by default if one is not
	// specified
	Default = NorthAmerica
)

var (
	// Geographies contains all supported geographies of which HCP is deployed
	Geographies []string
)

// ConnectionConfig contains connection strings used to configure HCP clients
type ConnectionConfig struct {
	// AuthURL is the URL of the production auth endpoint
	AuthURL string

	// OAuth2ClientID is the client ID of the production auth application
	OAuth2ClientID string

	// OAuth2RedirectURL is the callback URL for OAuth2
	OAuth2RedirectURL string

	// PortalURL is the URL of the production portal
	PortalURL string

	// APIAddress is the address of the production API
	APIAddress string

	// SCADAAddress is the address of the production SCADA endpoint
	SCADAAddress string
}
