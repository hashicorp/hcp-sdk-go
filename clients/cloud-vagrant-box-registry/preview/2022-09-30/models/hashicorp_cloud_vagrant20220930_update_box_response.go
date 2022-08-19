// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVagrant20220930UpdateBoxResponse hashicorp cloud vagrant 20220930 update box response
//
// swagger:model hashicorp.cloud.vagrant_20220930.UpdateBoxResponse
type HashicorpCloudVagrant20220930UpdateBoxResponse struct {

	// The updated Box.
	Box *HashicorpCloudVagrant20220930Box `json:"box,omitempty"`
}

// Validate validates this hashicorp cloud vagrant 20220930 update box response
func (m *HashicorpCloudVagrant20220930UpdateBoxResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBox(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrant20220930UpdateBoxResponse) validateBox(formats strfmt.Registry) error {

	if swag.IsZero(m.Box) { // not required
		return nil
	}

	if m.Box != nil {
		if err := m.Box.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("box")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930UpdateBoxResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930UpdateBoxResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrant20220930UpdateBoxResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
