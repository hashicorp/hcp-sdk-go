// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geography

const (
	// EU is the multi-region Europe deployment of HCP
	EU = "eu"
)

func init() {
	Geographies = append(Geographies, EU)
}

func NewConfigEU() *Config {
	return &Config{
		APIAddress:        "api.cloud.eu.hashicorp.com",
		AuthURL:           "https://auth.idp.eu.hashicorp.com",
		OAuth2ClientID:    "",
		OAuth2RedirectURL: "http://localhost:8443/oidc/callback",
		PortalURL:         "https://portal.cloud.eu.hashicorp.com",
		SCADAAddress:      "",
	}
}
