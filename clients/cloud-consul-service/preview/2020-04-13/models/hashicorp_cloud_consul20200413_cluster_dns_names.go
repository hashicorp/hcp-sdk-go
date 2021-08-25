// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsul20200413ClusterDNSNames DNSNames holds all of the cluster's DNS names.
//
// swagger:model hashicorp.cloud.consul_20200413.Cluster.DNSNames
type HashicorpCloudConsul20200413ClusterDNSNames struct {

	// private is the DNS name pointing to the cluster's private IP addresses.
	Private string `json:"private,omitempty"`

	// public is the DNS name pointing to the cluster's public IP addresses.
	Public string `json:"public,omitempty"`
}

// Validate validates this hashicorp cloud consul 20200413 cluster DNS names
func (m *HashicorpCloudConsul20200413ClusterDNSNames) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20200413ClusterDNSNames) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20200413ClusterDNSNames) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20200413ClusterDNSNames
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
