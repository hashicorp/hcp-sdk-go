// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudBoundary20211221SessionsResponse SessionsResponse is a response from getting an existing cluster.
//
// swagger:model hashicorp.cloud.boundary_20211221.SessionsResponse
type HashicorpCloudBoundary20211221SessionsResponse struct {

	// Sessions are the retrieved list of Boundary sessions.
	Sessions []*HashicorpCloudBoundary20211221Session `json:"sessions"`
}

// Validate validates this hashicorp cloud boundary 20211221 sessions response
func (m *HashicorpCloudBoundary20211221SessionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSessions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudBoundary20211221SessionsResponse) validateSessions(formats strfmt.Registry) error {

	if swag.IsZero(m.Sessions) { // not required
		return nil
	}

	for i := 0; i < len(m.Sessions); i++ {
		if swag.IsZero(m.Sessions[i]) { // not required
			continue
		}

		if m.Sessions[i] != nil {
			if err := m.Sessions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudBoundary20211221SessionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudBoundary20211221SessionsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudBoundary20211221SessionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
