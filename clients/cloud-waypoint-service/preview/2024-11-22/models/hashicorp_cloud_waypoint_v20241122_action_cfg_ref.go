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

// HashicorpCloudWaypointV20241122ActionCfgRef hashicorp cloud waypoint v20241122 action cfg ref
//
// swagger:model hashicorp.cloud.waypoint.v20241122.ActionCfgRef
type HashicorpCloudWaypointV20241122ActionCfgRef struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Outputs will be on response messages only
	Outputs *HashicorpCloudWaypointV20241122ActionCfgRefOutputs `json:"outputs,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 action cfg ref
func (m *HashicorpCloudWaypointV20241122ActionCfgRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutputs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122ActionCfgRef) validateOutputs(formats strfmt.Registry) error {
	if swag.IsZero(m.Outputs) { // not required
		return nil
	}

	if m.Outputs != nil {
		if err := m.Outputs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outputs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outputs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint v20241122 action cfg ref based on the context it is used
func (m *HashicorpCloudWaypointV20241122ActionCfgRef) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOutputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122ActionCfgRef) contextValidateOutputs(ctx context.Context, formats strfmt.Registry) error {

	if m.Outputs != nil {

		if swag.IsZero(m.Outputs) { // not required
			return nil
		}

		if err := m.Outputs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outputs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outputs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122ActionCfgRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122ActionCfgRef) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122ActionCfgRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
