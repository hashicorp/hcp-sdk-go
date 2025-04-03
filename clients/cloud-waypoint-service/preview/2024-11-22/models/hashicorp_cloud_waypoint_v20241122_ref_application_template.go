// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointV20241122RefApplicationTemplate hashicorp cloud waypoint v20241122 ref application template
//
// swagger:model hashicorp.cloud.waypoint.v20241122.Ref.ApplicationTemplate
type HashicorpCloudWaypointV20241122RefApplicationTemplate struct {

	// ID of the ApplicationTemplate
	ID string `json:"id,omitempty"`

	// Name of the ApplicationTemplate
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 ref application template
func (m *HashicorpCloudWaypointV20241122RefApplicationTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint v20241122 ref application template based on context it is used
func (m *HashicorpCloudWaypointV20241122RefApplicationTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122RefApplicationTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122RefApplicationTemplate) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122RefApplicationTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
