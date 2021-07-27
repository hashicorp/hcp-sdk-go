// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudNetwork20200907GetPeeringResponse GetPeeringResponse is a response type for GetPeering endpoint
//
// swagger:model hashicorp.cloud.network_20200907.GetPeeringResponse
type HashicorpCloudNetwork20200907GetPeeringResponse struct {

	// Peering is the requested peering
	Peering *HashicorpCloudNetwork20200907Peering `json:"peering,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 get peering response
func (m *HashicorpCloudNetwork20200907GetPeeringResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeering(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907GetPeeringResponse) validatePeering(formats strfmt.Registry) error {

	if swag.IsZero(m.Peering) { // not required
		return nil
	}

	if m.Peering != nil {
		if err := m.Peering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peering")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907GetPeeringResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907GetPeeringResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907GetPeeringResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
