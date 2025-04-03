// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointV20241122GetGitHubReposRequestSessionTarget hashicorp cloud waypoint v20241122 get git hub repos request session target
//
// swagger:model hashicorp.cloud.waypoint.v20241122.GetGitHubReposRequest.SessionTarget
type HashicorpCloudWaypointV20241122GetGitHubReposRequestSessionTarget struct {

	// installation id
	InstallationID string `json:"installation_id,omitempty"`

	// session id
	SessionID string `json:"session_id,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 get git hub repos request session target
func (m *HashicorpCloudWaypointV20241122GetGitHubReposRequestSessionTarget) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint v20241122 get git hub repos request session target based on context it is used
func (m *HashicorpCloudWaypointV20241122GetGitHubReposRequestSessionTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122GetGitHubReposRequestSessionTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122GetGitHubReposRequestSessionTarget) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122GetGitHubReposRequestSessionTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
