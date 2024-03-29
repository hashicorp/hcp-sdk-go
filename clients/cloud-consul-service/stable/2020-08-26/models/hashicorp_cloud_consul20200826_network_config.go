// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudConsul20200826NetworkConfig NetworkConfig configures the network of the Consul cluster.
//
// swagger:model hashicorp.cloud.consul_20200826.NetworkConfig
type HashicorpCloudConsul20200826NetworkConfig struct {

	// network is the link of the HVN to launch the Consul cluster in.
	// The network cannot be updated once the cluster is created.
	Network *cloud.HashicorpCloudLocationLink `json:"network,omitempty"`

	// private indicates if this cluster's instances should be configured in a
	// non-externally accessible way.
	Private bool `json:"private,omitempty"`
}

// Validate validates this hashicorp cloud consul 20200826 network config
func (m *HashicorpCloudConsul20200826NetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826NetworkConfig) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul 20200826 network config based on the context it is used
func (m *HashicorpCloudConsul20200826NetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826NetworkConfig) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.Network != nil {
		if err := m.Network.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826NetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826NetworkConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20200826NetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
