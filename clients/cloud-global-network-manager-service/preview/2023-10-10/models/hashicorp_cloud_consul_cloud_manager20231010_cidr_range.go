// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulCloudManager20231010CidrRange CidrRange specifies an IP address and prefix in CIDR format.
//
// swagger:model hashicorp.cloud.consul_cloud_manager_20231010.CidrRange
type HashicorpCloudConsulCloudManager20231010CidrRange struct {

	// The IPv4 address range, including prefix in CIDR notation. For example, 172.25.16.0/24
	Address string `json:"address,omitempty"`

	// An optional description of the IP address range.
	Description string `json:"description,omitempty"`
}

// Validate validates this hashicorp cloud consul cloud manager 20231010 cidr range
func (m *HashicorpCloudConsulCloudManager20231010CidrRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud consul cloud manager 20231010 cidr range based on context it is used
func (m *HashicorpCloudConsulCloudManager20231010CidrRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulCloudManager20231010CidrRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulCloudManager20231010CidrRange) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulCloudManager20231010CidrRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
