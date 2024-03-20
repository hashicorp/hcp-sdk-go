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

// HashicorpCloudWaypointWaypointServiceSendStatusLogBody hashicorp cloud waypoint waypoint service send status log body
//
// swagger:model hashicorp.cloud.waypoint.WaypointService.SendStatusLogBody
type HashicorpCloudWaypointWaypointServiceSendStatusLogBody struct {

	// action config
	ActionConfig *HashicorpCloudWaypointWaypointServiceSendStatusLogBodyActionConfig `json:"action_config,omitempty"`

	// The namespace the action to be listed in
	Namespace interface{} `json:"namespace,omitempty"`

	// The status log to send
	StatusLog *HashicorpCloudWaypointStatusLog `json:"status_log,omitempty"`
}

// Validate validates this hashicorp cloud waypoint waypoint service send status log body
func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusLog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBody) validateActionConfig(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBody) validateStatusLog(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusLog) { // not required
		return nil
	}

	if m.StatusLog != nil {
		if err := m.StatusLog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_log")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint waypoint service send status log body based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBody) contextValidateActionConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBody) contextValidateStatusLog(ctx context.Context, formats strfmt.Registry) error {

	if m.StatusLog != nil {

		if swag.IsZero(m.StatusLog) { // not required
			return nil
		}

		if err := m.StatusLog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_log")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceSendStatusLogBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointWaypointServiceSendStatusLogBodyActionConfig Action config without ID
//
// # The action config ID to send this to
//
// swagger:model HashicorpCloudWaypointWaypointServiceSendStatusLogBodyActionConfig
type HashicorpCloudWaypointWaypointServiceSendStatusLogBodyActionConfig struct {

	// URL to trigger an action on. Only used in Custom mode.
	ActionURL string `json:"action_url,omitempty"`

	// Description of the action
	Description string `json:"description,omitempty"`

	// Give the action config an optional unique (per-namespace) name
	Name string `json:"name,omitempty"`

	// More explicitly configure the kind of HTTP request to be made
	Request *HashicorpCloudWaypointActionConfigRequest `json:"request,omitempty"`
}

// Validate validates this hashicorp cloud waypoint waypoint service send status log body action config
func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBodyActionConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBodyActionConfig) validateRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_config" + "." + "request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_config" + "." + "request")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint waypoint service send status log body action config based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBodyActionConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBodyActionConfig) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {

		if swag.IsZero(m.Request) { // not required
			return nil
		}

		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_config" + "." + "request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_config" + "." + "request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBodyActionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceSendStatusLogBodyActionConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceSendStatusLogBodyActionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}