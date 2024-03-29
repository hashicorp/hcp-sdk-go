// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointGetActionConfigResponse hashicorp cloud waypoint get action config response
//
// swagger:model hashicorp.cloud.waypoint.GetActionConfigResponse
type HashicorpCloudWaypointGetActionConfigResponse struct {

	// The current requested action config
	ActionConfig *HashicorpCloudWaypointActionConfig `json:"action_config,omitempty"`

	// The latest run for this config. A nil ActionRun means there are not yet any runs
	LatestRun *HashicorpCloudWaypointActionRun `json:"latest_run,omitempty"`

	// The total number of runs (overall) for this config
	TotalRuns *HashicorpCloudWaypointGetActionConfigResponseTotalRuns `json:"total_runs,omitempty"`
}

// Validate validates this hashicorp cloud waypoint get action config response
func (m *HashicorpCloudWaypointGetActionConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestRun(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalRuns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointGetActionConfigResponse) validateActionConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionConfig) { // not required
		return nil
	}

	if m.ActionConfig != nil {
		if err := m.ActionConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointGetActionConfigResponse) validateLatestRun(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestRun) { // not required
		return nil
	}

	if m.LatestRun != nil {
		if err := m.LatestRun.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latest_run")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latest_run")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointGetActionConfigResponse) validateTotalRuns(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalRuns) { // not required
		return nil
	}

	if m.TotalRuns != nil {
		if err := m.TotalRuns.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total_runs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("total_runs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint get action config response based on the context it is used
func (m *HashicorpCloudWaypointGetActionConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatestRun(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalRuns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointGetActionConfigResponse) contextValidateActionConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ActionConfig != nil {

		if swag.IsZero(m.ActionConfig) { // not required
			return nil
		}

		if err := m.ActionConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointGetActionConfigResponse) contextValidateLatestRun(ctx context.Context, formats strfmt.Registry) error {

	if m.LatestRun != nil {

		if swag.IsZero(m.LatestRun) { // not required
			return nil
		}

		if err := m.LatestRun.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latest_run")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latest_run")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointGetActionConfigResponse) contextValidateTotalRuns(ctx context.Context, formats strfmt.Registry) error {

	if m.TotalRuns != nil {

		if swag.IsZero(m.TotalRuns) { // not required
			return nil
		}

		if err := m.TotalRuns.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total_runs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("total_runs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointGetActionConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointGetActionConfigResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointGetActionConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
