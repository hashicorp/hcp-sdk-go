// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody hashicorp cloud waypoint v20241122 waypoint service send status log body
//
// swagger:model hashicorp.cloud.waypoint.v20241122.WaypointService.SendStatusLogBody
type HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody struct {

	// action config
	ActionConfig *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig `json:"action_config,omitempty"`

	// namespace
	Namespace *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespace `json:"namespace,omitempty"`

	// The status log to send
	StatusLog *HashicorpCloudWaypointV20241122StatusLog `json:"status_log,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service send status log body
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
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

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody) validateActionConfig(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody) validateNamespace(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespace) { // not required
		return nil
	}

	if m.Namespace != nil {
		if err := m.Namespace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody) validateStatusLog(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service send status log body based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
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

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody) contextValidateActionConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

	if m.Namespace != nil {

		if swag.IsZero(m.Namespace) { // not required
			return nil
		}

		if err := m.Namespace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody) contextValidateStatusLog(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig Action config without ID
//
// # The action config ID to send this to
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig
type HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig struct {

	// URL to trigger an action on. Only used in Custom mode.
	ActionURL string `json:"action_url,omitempty"`

	// The time the action config was created in the database
	// This is mainly a convenience field for the UI and might not always be set.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Description of the action
	Description string `json:"description,omitempty"`

	// Give the action config an optional unique (per-namespace) name
	Name string `json:"name,omitempty"`

	// More explicitly configure the kind of HTTP request to be made
	Request *HashicorpCloudWaypointV20241122ActionConfigRequest `json:"request,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service send status log body action config
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("action_config"+"."+"created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig) validateRequest(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service send status log body action config based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyActionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespace The namespace the action to be listed in
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespace
type HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespace struct {

	// id
	ID string `json:"id,omitempty"`

	// location
	Location *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespaceLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service send status log body namespace
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespace) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service send status log body namespace based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespace) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {

		if swag.IsZero(m.Location) { // not required
			return nil
		}

		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespace) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespaceLocation When used via an API request, this is populated and used to populate id.
//
// When used via an API request, this is populated and used to populate id.
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespaceLocation
type HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespaceLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *cloud.HashicorpCloudLocationRegion `json:"region,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service send status log body namespace location
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespaceLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespaceLocation) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service send status log body namespace location based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespaceLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespaceLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {

		if swag.IsZero(m.Region) { // not required
			return nil
		}

		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespaceLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespaceLocation) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceSendStatusLogBodyNamespaceLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
