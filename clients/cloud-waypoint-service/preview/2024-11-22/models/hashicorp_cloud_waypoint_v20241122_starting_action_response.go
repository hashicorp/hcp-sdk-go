// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointV20241122StartingActionResponse hashicorp cloud waypoint v20241122 starting action response
//
// swagger:model hashicorp.cloud.waypoint.v20241122.StartingActionResponse
type HashicorpCloudWaypointV20241122StartingActionResponse struct {

	// The external identifier for the action run created
	ActionRunID string `json:"action_run_id,omitempty"`

	// The sequence number used to determine how many times this has been run
	Sequence string `json:"sequence,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 starting action response
func (m *HashicorpCloudWaypointV20241122StartingActionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint v20241122 starting action response based on context it is used
func (m *HashicorpCloudWaypointV20241122StartingActionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122StartingActionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122StartingActionResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122StartingActionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
