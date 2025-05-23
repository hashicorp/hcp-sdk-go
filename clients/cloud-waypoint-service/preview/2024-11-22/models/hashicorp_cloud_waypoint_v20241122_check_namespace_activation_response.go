// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointV20241122CheckNamespaceActivationResponse hashicorp cloud waypoint v20241122 check namespace activation response
//
// swagger:model hashicorp.cloud.waypoint.v20241122.CheckNamespaceActivationResponse
type HashicorpCloudWaypointV20241122CheckNamespaceActivationResponse struct {

	// If the namespace is in the grace period for actions
	ActionsGracePeriod bool `json:"actions_grace_period,omitempty"`

	// If the checked namespace is considered active
	Active bool `json:"active,omitempty"`

	// If the namespace has the actions entitlement
	HasActions bool `json:"has_actions,omitempty"`

	// If the checked namespace has ever been activated before
	HasActivated bool `json:"has_activated,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 check namespace activation response
func (m *HashicorpCloudWaypointV20241122CheckNamespaceActivationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint v20241122 check namespace activation response based on context it is used
func (m *HashicorpCloudWaypointV20241122CheckNamespaceActivationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122CheckNamespaceActivationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122CheckNamespaceActivationResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122CheckNamespaceActivationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
