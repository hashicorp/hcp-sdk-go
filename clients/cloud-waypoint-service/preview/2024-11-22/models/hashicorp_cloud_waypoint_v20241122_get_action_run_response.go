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

// HashicorpCloudWaypointV20241122GetActionRunResponse hashicorp cloud waypoint v20241122 get action run response
//
// swagger:model hashicorp.cloud.waypoint.v20241122.GetActionRunResponse
type HashicorpCloudWaypointV20241122GetActionRunResponse struct {

	// The run that was executed
	ActionRun *HashicorpCloudWaypointV20241122ActionRun `json:"action_run,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 get action run response
func (m *HashicorpCloudWaypointV20241122GetActionRunResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionRun(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122GetActionRunResponse) validateActionRun(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionRun) { // not required
		return nil
	}

	if m.ActionRun != nil {
		if err := m.ActionRun.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_run")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_run")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint v20241122 get action run response based on the context it is used
func (m *HashicorpCloudWaypointV20241122GetActionRunResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionRun(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122GetActionRunResponse) contextValidateActionRun(ctx context.Context, formats strfmt.Registry) error {

	if m.ActionRun != nil {

		if swag.IsZero(m.ActionRun) { // not required
			return nil
		}

		if err := m.ActionRun.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_run")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_run")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122GetActionRunResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122GetActionRunResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122GetActionRunResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
