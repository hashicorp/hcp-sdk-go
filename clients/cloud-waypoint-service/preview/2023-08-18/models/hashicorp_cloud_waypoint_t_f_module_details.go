// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointTFModuleDetails hashicorp cloud waypoint t f module details
//
// swagger:model hashicorp.cloud.waypoint.TFModuleDetails
type HashicorpCloudWaypointTFModuleDetails struct {

	// readme
	Readme string `json:"readme,omitempty"`
}

// Validate validates this hashicorp cloud waypoint t f module details
func (m *HashicorpCloudWaypointTFModuleDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint t f module details based on context it is used
func (m *HashicorpCloudWaypointTFModuleDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointTFModuleDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointTFModuleDetails) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointTFModuleDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}