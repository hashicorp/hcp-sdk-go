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

// HashicorpCloudConsul20200826ConsulConfig ConsulConfig exposes user settable configurations for a Consul cluster.
//
// swagger:model hashicorp.cloud.consul_20200826.ConsulConfig
type HashicorpCloudConsul20200826ConsulConfig struct {

	// connect_enabled toggles Consul Connect on the Consul servers.
	ConnectEnabled bool `json:"connect_enabled,omitempty"`

	// datacenter sets the datacenter the Consul servers will be running in. If
	// not set, a default value will be used.
	//
	// https://www.consul.io/docs/agent/options.html#datacenter
	Datacenter string `json:"datacenter,omitempty"`

	// primary contains a link to the primary consul cluster in a federation.
	Primary *cloud.HashicorpCloudLocationLink `json:"primary,omitempty"`
}

// Validate validates this hashicorp cloud consul 20200826 consul config
func (m *HashicorpCloudConsul20200826ConsulConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrimary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826ConsulConfig) validatePrimary(formats strfmt.Registry) error {
	if swag.IsZero(m.Primary) { // not required
		return nil
	}

	if m.Primary != nil {
		if err := m.Primary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul 20200826 consul config based on the context it is used
func (m *HashicorpCloudConsul20200826ConsulConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrimary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826ConsulConfig) contextValidatePrimary(ctx context.Context, formats strfmt.Registry) error {

	if m.Primary != nil {
		if err := m.Primary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826ConsulConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826ConsulConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20200826ConsulConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
