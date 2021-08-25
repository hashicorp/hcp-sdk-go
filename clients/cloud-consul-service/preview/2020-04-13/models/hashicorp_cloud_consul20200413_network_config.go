// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsul20200413NetworkConfig NetworkConfig configures the network of the Consul cluster.
//
// swagger:model hashicorp.cloud.consul_20200413.NetworkConfig
type HashicorpCloudConsul20200413NetworkConfig struct {

	// network_id is the ID of the network to launch the Consul cluster in.
	NetworkID string `json:"network_id,omitempty"`

	// private indicates if this cluster's instances should be configured in a
	// non-externally accessible way.
	Private bool `json:"private,omitempty"`
}

// Validate validates this hashicorp cloud consul 20200413 network config
func (m *HashicorpCloudConsul20200413NetworkConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud consul 20200413 network config based on context it is used
func (m *HashicorpCloudConsul20200413NetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20200413NetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20200413NetworkConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20200413NetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
