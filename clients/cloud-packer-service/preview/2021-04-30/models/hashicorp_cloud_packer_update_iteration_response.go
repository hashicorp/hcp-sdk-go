// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPackerUpdateIterationResponse hashicorp cloud packer update iteration response
//
// swagger:model hashicorp.cloud.packer.UpdateIterationResponse
type HashicorpCloudPackerUpdateIterationResponse struct {

	// Information about the updated iteration.
	Iteration *HashicorpCloudPackerIteration `json:"iteration,omitempty"`
}

// Validate validates this hashicorp cloud packer update iteration response
func (m *HashicorpCloudPackerUpdateIterationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIteration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerUpdateIterationResponse) validateIteration(formats strfmt.Registry) error {

	if swag.IsZero(m.Iteration) { // not required
		return nil
	}

	if m.Iteration != nil {
		if err := m.Iteration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iteration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerUpdateIterationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerUpdateIterationResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerUpdateIterationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
