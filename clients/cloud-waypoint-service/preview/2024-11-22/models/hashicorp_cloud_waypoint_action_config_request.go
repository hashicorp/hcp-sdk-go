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

// HashicorpCloudWaypointActionConfigRequest The kind of HTTP request this config should trigger
//
// swagger:model hashicorp.cloud.waypoint.ActionConfig.Request
type HashicorpCloudWaypointActionConfigRequest struct {

	// The configuration for an agent operation
	Agent *HashicorpCloudWaypointActionConfigFlavorAgent `json:"agent,omitempty"`

	// Specify the exact details of the HTTP request to be made
	// Defaults to Custom, and will default to an HTTP POST if no fields are set
	Custom *HashicorpCloudWaypointActionConfigFlavorCustom `json:"custom,omitempty"`

	// The HTTP request will be configured to be sent to GitHub Actions
	Github *HashicorpCloudWaypointActionConfigFlavorGitHub `json:"github,omitempty"`
}

// Validate validates this hashicorp cloud waypoint action config request
func (m *HashicorpCloudWaypointActionConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGithub(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointActionConfigRequest) validateAgent(formats strfmt.Registry) error {
	if swag.IsZero(m.Agent) { // not required
		return nil
	}

	if m.Agent != nil {
		if err := m.Agent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionConfigRequest) validateCustom(formats strfmt.Registry) error {
	if swag.IsZero(m.Custom) { // not required
		return nil
	}

	if m.Custom != nil {
		if err := m.Custom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionConfigRequest) validateGithub(formats strfmt.Registry) error {
	if swag.IsZero(m.Github) { // not required
		return nil
	}

	if m.Github != nil {
		if err := m.Github.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("github")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("github")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint action config request based on the context it is used
func (m *HashicorpCloudWaypointActionConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGithub(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointActionConfigRequest) contextValidateAgent(ctx context.Context, formats strfmt.Registry) error {

	if m.Agent != nil {

		if swag.IsZero(m.Agent) { // not required
			return nil
		}

		if err := m.Agent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionConfigRequest) contextValidateCustom(ctx context.Context, formats strfmt.Registry) error {

	if m.Custom != nil {

		if swag.IsZero(m.Custom) { // not required
			return nil
		}

		if err := m.Custom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointActionConfigRequest) contextValidateGithub(ctx context.Context, formats strfmt.Registry) error {

	if m.Github != nil {

		if swag.IsZero(m.Github) { // not required
			return nil
		}

		if err := m.Github.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("github")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("github")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointActionConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointActionConfigRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointActionConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
