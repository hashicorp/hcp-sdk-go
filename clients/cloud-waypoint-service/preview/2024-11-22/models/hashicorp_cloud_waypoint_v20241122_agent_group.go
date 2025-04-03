// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointV20241122AgentGroup hashicorp cloud waypoint v20241122 agent group
//
// swagger:model hashicorp.cloud.waypoint.v20241122.AgentGroup
type HashicorpCloudWaypointV20241122AgentGroup struct {

	// Description of what the group is for (ie: production in us-west-2)
	Description string `json:"description,omitempty"`

	// Name of the group
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 agent group
func (m *HashicorpCloudWaypointV20241122AgentGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint v20241122 agent group based on context it is used
func (m *HashicorpCloudWaypointV20241122AgentGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122AgentGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122AgentGroup) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122AgentGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
