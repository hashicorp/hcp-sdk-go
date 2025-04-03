// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointV20241122IntegrationGitHubRepo hashicorp cloud waypoint v20241122 integration git hub repo
//
// swagger:model hashicorp.cloud.waypoint.v20241122.Integration.GitHub.Repo
type HashicorpCloudWaypointV20241122IntegrationGitHubRepo struct {

	// full_name is the owner/repo
	FullName string `json:"full_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 integration git hub repo
func (m *HashicorpCloudWaypointV20241122IntegrationGitHubRepo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint v20241122 integration git hub repo based on context it is used
func (m *HashicorpCloudWaypointV20241122IntegrationGitHubRepo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122IntegrationGitHubRepo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122IntegrationGitHubRepo) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122IntegrationGitHubRepo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
