// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geography

const (
	// US is the multi-region United States deployment of HCP
	US Geo = "us"
)

func init() {
	Geographies = append(Geographies, US)
}

func NewConfigUS() *Config {
	return &Config{
		APIAddress:     "api.cloud.hashicorp.com:443",
		PortalURL:      "https://portal.cloud.hashicorp.com",
		AuthURL:        "https://auth.idp.hashicorp.com",
		OAuth2ClientID: "4edd6521-6eb9-4d78-9039-7ce8569d667c",
		SCADAAddress:   "scada.hashicorp.cloud:7224",
	}
}
