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

// HashicorpCloudVagrantCreateRegistryRequest hashicorp cloud vagrant create registry request
//
// swagger:model hashicorp.cloud.vagrant.CreateRegistryRequest
type HashicorpCloudVagrantCreateRegistryRequest struct {

	// The name of the Registry. This is a user-settable, globally-unique name,
	// and forms the first segment of a Box, ie: the "hashicorp" in
	// "hashicorp/vagrant". No two registries can have the same name, regardless
	// of the Organization they are in, or specific Project within that
	// Organization.
	ID string `json:"id,omitempty"`

	// The Location to create the Registry in.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud vagrant create registry request
func (m *HashicorpCloudVagrantCreateRegistryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrantCreateRegistryRequest) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vagrant create registry request based on the context it is used
func (m *HashicorpCloudVagrantCreateRegistryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrantCreateRegistryRequest) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrantCreateRegistryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrantCreateRegistryRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrantCreateRegistryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
