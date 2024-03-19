// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointTFModuleVariable TFModuleVariable represents a Terraform module variable.
//
// swagger:model hashicorp.cloud.waypoint.TFModuleVariable
type HashicorpCloudWaypointTFModuleVariable struct {

	// description is the description of the variable.
	Description string `json:"description,omitempty"`

	// name is the name of the variable.
	Name string `json:"name,omitempty"`

	// options is the list of options for the variable, set in Terraform Cloud.
	// 1. If len(options) == 0, then the variable is forcibly left unset.
	// 2. If len(options) == 1, then the variable was set to be that value by the platform engineer, and cannot be set by the app developer.
	// 3. If len(options) > 1, then the variable may be set to one of the options in the list, which can be chosen by the app developer.
	Options []string `json:"options"`

	// indicates if the variable may be set by the app developer.
	UserEditable bool `json:"user_editable,omitempty"`

	// variable_type is the type of Terraform variable. See the documentation for more info:
	// https://developer.hashicorp.com/terraform/language/expressions/types
	VariableType string `json:"variable_type,omitempty"`
}

// Validate validates this hashicorp cloud waypoint t f module variable
func (m *HashicorpCloudWaypointTFModuleVariable) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint t f module variable based on context it is used
func (m *HashicorpCloudWaypointTFModuleVariable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointTFModuleVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointTFModuleVariable) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointTFModuleVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
