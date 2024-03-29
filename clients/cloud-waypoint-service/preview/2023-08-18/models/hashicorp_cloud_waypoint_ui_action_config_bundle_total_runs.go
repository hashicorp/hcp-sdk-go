// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointUIActionConfigBundleTotalRuns hashicorp cloud waypoint UI action config bundle total runs
//
// swagger:model hashicorp.cloud.waypoint.UI.ActionConfigBundle.TotalRuns
type HashicorpCloudWaypointUIActionConfigBundleTotalRuns struct {

	// Total number of times this config has been run in the last day
	LastDay string `json:"last_day,omitempty"`

	// Total number of times this config has been run in the last month
	LastMonth string `json:"last_month,omitempty"`

	// Total number of times this config has been run
	Runs string `json:"runs,omitempty"`
}

// Validate validates this hashicorp cloud waypoint UI action config bundle total runs
func (m *HashicorpCloudWaypointUIActionConfigBundleTotalRuns) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint UI action config bundle total runs based on context it is used
func (m *HashicorpCloudWaypointUIActionConfigBundleTotalRuns) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointUIActionConfigBundleTotalRuns) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointUIActionConfigBundleTotalRuns) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointUIActionConfigBundleTotalRuns
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
