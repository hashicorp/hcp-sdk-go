// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geography

const (
	// Europe is the HCP deployment in Europe
	Europe = "europe"
)

func init() {
	Geographies = append(Geographies, Europe)
}

func NewConfigEurope() *Config {
	return &Config{
		APIAddress:        "api.cloud.eu.hashicorp.com",
		AuthURL:           "https://auth.idp.eu.hashicorp.com",
		OAuth2ClientID:    "",
		OAuth2RedirectURL: "http://localhost:8443/oidc/callback",
		PortalURL:         "portal.cloud.eu.hashicorp.com",
		SCADAAddress:      "",
	}
}
