// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudWaypointWaypointServiceRunActionBody hashicorp cloud waypoint waypoint service run action body
//
// swagger:model hashicorp.cloud.waypoint.WaypointService.RunActionBody
type HashicorpCloudWaypointWaypointServiceRunActionBody struct {

	// The action config to run
	ActionRef *HashicorpCloudWaypointActionCfgRef `json:"action_ref,omitempty"`

	// namespace
	Namespace *HashicorpCloudWaypointWaypointServiceRunActionBodyNamespace `json:"namespace,omitempty"`

	// Optional scope to set for running the action
	Scope *HashicorpCloudWaypointActionRunScope `json:"scope,omitempty"`

	// Optional variables to override in the action run
	VariableOverrides []*HashicorpCloudWaypointRunActionRequestVariableOverride `json:"variable_overrides"`
}

// Validate validates this hashicorp cloud waypoint waypoint service run action body
func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariableOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) validateActionRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionRef) { // not required
		return nil
	}

	if m.ActionRef != nil {
		if err := m.ActionRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_ref")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_ref")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) validateNamespace(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) validateVariableOverrides(formats strfmt.Registry) error {
	if swag.IsZero(m.VariableOverrides) { // not required
		return nil
	}

	for i := 0; i < len(m.VariableOverrides); i++ {
		if swag.IsZero(m.VariableOverrides[i]) { // not required
			continue
		}

		if m.VariableOverrides[i] != nil {
			if err := m.VariableOverrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variable_overrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variable_overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint waypoint service run action body based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariableOverrides(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) contextValidateActionRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ActionRef != nil {

		if swag.IsZero(m.ActionRef) { // not required
			return nil
		}

		if err := m.ActionRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_ref")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_ref")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if m.Scope != nil {

		if swag.IsZero(m.Scope) { // not required
			return nil
		}

		if err := m.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) contextValidateVariableOverrides(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VariableOverrides); i++ {

		if m.VariableOverrides[i] != nil {

			if swag.IsZero(m.VariableOverrides[i]) { // not required
				return nil
			}

			if err := m.VariableOverrides[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variable_overrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variable_overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceRunActionBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceRunActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointWaypointServiceRunActionBodyNamespace The namespace the action will run in
//
// swagger:model HashicorpCloudWaypointWaypointServiceRunActionBodyNamespace
type HashicorpCloudWaypointWaypointServiceRunActionBodyNamespace struct {

	// When used via an API request, this is populated and used to populate id.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud waypoint waypoint service run action body namespace
func (m *HashicorpCloudWaypointWaypointServiceRunActionBodyNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceRunActionBodyNamespace) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint waypoint service run action body namespace based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceRunActionBodyNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceRunActionBodyNamespace) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointWaypointServiceRunActionBodyNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceRunActionBodyNamespace) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceRunActionBodyNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
