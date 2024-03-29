// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointRefApplicationTemplate hashicorp cloud waypoint ref application template
//
// swagger:model hashicorp.cloud.waypoint.Ref.ApplicationTemplate
type HashicorpCloudWaypointRefApplicationTemplate struct {

	// ID of the ApplicationTemplate
	ID string `json:"id,omitempty"`

	// Name of the ApplicationTemplate
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp cloud waypoint ref application template
func (m *HashicorpCloudWaypointRefApplicationTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint ref application template based on context it is used
func (m *HashicorpCloudWaypointRefApplicationTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointRefApplicationTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointRefApplicationTemplate) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointRefApplicationTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
