// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geography

const (
	// NorthAmerica is the HCP deployment in North America
	NorthAmerica = "north_america"
)

func init() {
	Geographies = append(Geographies, NorthAmerica)
}

func NewConfigNorthAmerica() *Config {
	return &Config{
		APIAddress:        "api.cloud.hashicorp.com:443",
		AuthURL:           "https://auth.idp.hashicorp.com",
		OAuth2ClientID:    "4edd6521-6eb9-4d78-9039-7ce8569d667c",
		OAuth2RedirectURL: "http://localhost:8443/oidc/callback",
		PortalURL:         "https://portal.cloud.hashicorp.com",
		SCADAAddress:      "scada.hashicorp.cloud:7224",
	}
}
