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

// HashicorpCloudWaypointCreateAddOnDefinitionResponse CreateAddOnDefinitionResponse is the response containing the just-created
// Add-on definition
//
// swagger:model hashicorp.cloud.waypoint.CreateAddOnDefinitionResponse
type HashicorpCloudWaypointCreateAddOnDefinitionResponse struct {

	// add on definition
	AddOnDefinition *HashicorpCloudWaypointAddOnDefinition `json:"add_on_definition,omitempty"`
}

// Validate validates this hashicorp cloud waypoint create add on definition response
func (m *HashicorpCloudWaypointCreateAddOnDefinitionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddOnDefinition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointCreateAddOnDefinitionResponse) validateAddOnDefinition(formats strfmt.Registry) error {
	if swag.IsZero(m.AddOnDefinition) { // not required
		return nil
	}

	if m.AddOnDefinition != nil {
		if err := m.AddOnDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("add_on_definition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("add_on_definition")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint create add on definition response based on the context it is used
func (m *HashicorpCloudWaypointCreateAddOnDefinitionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddOnDefinition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointCreateAddOnDefinitionResponse) contextValidateAddOnDefinition(ctx context.Context, formats strfmt.Registry) error {

	if m.AddOnDefinition != nil {

		if swag.IsZero(m.AddOnDefinition) { // not required
			return nil
		}

		if err := m.AddOnDefinition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("add_on_definition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("add_on_definition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointCreateAddOnDefinitionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointCreateAddOnDefinitionResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointCreateAddOnDefinitionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
