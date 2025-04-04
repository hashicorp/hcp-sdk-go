// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointV20241122GetGitHubWorkflowsResponseActionWorkflow hashicorp cloud waypoint v20241122 get git hub workflows response action workflow
//
// swagger:model hashicorp.cloud.waypoint.v20241122.GetGitHubWorkflowsResponse.ActionWorkflow
type HashicorpCloudWaypointV20241122GetGitHubWorkflowsResponseActionWorkflow struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 get git hub workflows response action workflow
func (m *HashicorpCloudWaypointV20241122GetGitHubWorkflowsResponseActionWorkflow) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint v20241122 get git hub workflows response action workflow based on context it is used
func (m *HashicorpCloudWaypointV20241122GetGitHubWorkflowsResponseActionWorkflow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122GetGitHubWorkflowsResponseActionWorkflow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122GetGitHubWorkflowsResponseActionWorkflow) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122GetGitHubWorkflowsResponseActionWorkflow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
