// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geography

// Geo is a custom type for defining all supported HCP geographies
type Geo string

const (
	// Default is the HCP geography that will be used by default if one is not
	// specified
	Default = US
)

var (
	// Geographies contains all supported geographies of which HCP is deployed
	Geographies []Geo
)

// Config contains connection strings used to configure HCP clients
type Config struct {
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
