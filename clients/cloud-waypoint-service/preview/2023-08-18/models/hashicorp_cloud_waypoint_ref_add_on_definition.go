// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointRefAddOnDefinition hashicorp cloud waypoint ref add on definition
//
// swagger:model hashicorp.cloud.waypoint.Ref.AddOnDefinition
type HashicorpCloudWaypointRefAddOnDefinition struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp cloud waypoint ref add on definition
func (m *HashicorpCloudWaypointRefAddOnDefinition) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint ref add on definition based on context it is used
func (m *HashicorpCloudWaypointRefAddOnDefinition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointRefAddOnDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointRefAddOnDefinition) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointRefAddOnDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
