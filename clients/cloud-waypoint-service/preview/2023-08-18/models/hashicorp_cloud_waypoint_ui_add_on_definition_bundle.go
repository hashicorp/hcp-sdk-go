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

// HashicorpCloudWaypointUIAddOnDefinitionBundle hashicorp cloud waypoint UI add on definition bundle
//
// swagger:model hashicorp.cloud.waypoint.UI.AddOnDefinitionBundle
type HashicorpCloudWaypointUIAddOnDefinitionBundle struct {

	// add on definition
	AddOnDefinition *HashicorpCloudWaypointAddOnDefinition `json:"add_on_definition,omitempty"`

	// install count
	InstallCount string `json:"install_count,omitempty"`
}

// Validate validates this hashicorp cloud waypoint UI add on definition bundle
func (m *HashicorpCloudWaypointUIAddOnDefinitionBundle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddOnDefinition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUIAddOnDefinitionBundle) validateAddOnDefinition(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint UI add on definition bundle based on the context it is used
func (m *HashicorpCloudWaypointUIAddOnDefinitionBundle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddOnDefinition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUIAddOnDefinitionBundle) contextValidateAddOnDefinition(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointUIAddOnDefinitionBundle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointUIAddOnDefinitionBundle) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointUIAddOnDefinitionBundle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
