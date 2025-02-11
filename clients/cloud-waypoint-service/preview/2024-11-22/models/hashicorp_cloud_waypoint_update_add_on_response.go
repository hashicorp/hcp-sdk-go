// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointUpdateAddOnResponse UpdateAddOnResponse is the response containing the just-updated Add-on.
//
// swagger:model hashicorp.cloud.waypoint.UpdateAddOnResponse
type HashicorpCloudWaypointUpdateAddOnResponse struct {

	// The newly updated Add-on.
	AddOn *HashicorpCloudWaypointAddOn `json:"add_on,omitempty"`
}

// Validate validates this hashicorp cloud waypoint update add on response
func (m *HashicorpCloudWaypointUpdateAddOnResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUpdateAddOnResponse) validateAddOn(formats strfmt.Registry) error {
	if swag.IsZero(m.AddOn) { // not required
		return nil
	}

	if m.AddOn != nil {
		if err := m.AddOn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("add_on")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("add_on")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint update add on response based on the context it is used
func (m *HashicorpCloudWaypointUpdateAddOnResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUpdateAddOnResponse) contextValidateAddOn(ctx context.Context, formats strfmt.Registry) error {

	if m.AddOn != nil {

		if swag.IsZero(m.AddOn) { // not required
			return nil
		}

		if err := m.AddOn.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("add_on")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("add_on")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointUpdateAddOnResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointUpdateAddOnResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointUpdateAddOnResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
