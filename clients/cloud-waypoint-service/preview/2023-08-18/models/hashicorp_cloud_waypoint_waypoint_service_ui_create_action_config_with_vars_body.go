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

// HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody hashicorp cloud waypoint waypoint service UI create action config with vars body
//
// swagger:model hashicorp.cloud.waypoint.WaypointService.UI_CreateActionConfigWithVarsBody
type HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody struct {

	// The action config to create
	ActionConfig *HashicorpCloudWaypointActionConfig `json:"action_config,omitempty"`

	// namespace
	Namespace *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBodyNamespace `json:"namespace,omitempty"`

	// The variables to create and attach the action config scope to
	Variables []*HashicorpCloudWaypointVariable `json:"variables"`
}

// Validate validates this hashicorp cloud waypoint waypoint service UI create action config with vars body
func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody) validateActionConfig(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody) validateNamespace(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody) validateVariables(formats strfmt.Registry) error {
	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for i := 0; i < len(m.Variables); i++ {
		if swag.IsZero(m.Variables[i]) { // not required
			continue
		}

		if m.Variables[i] != nil {
			if err := m.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint waypoint service UI create action config with vars body based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody) contextValidateActionConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody) contextValidateVariables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Variables); i++ {

		if m.Variables[i] != nil {

			if swag.IsZero(m.Variables[i]) { // not required
				return nil
			}

			if err := m.Variables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBodyNamespace The namespace this action config belongs to
//
// swagger:model HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBodyNamespace
type HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBodyNamespace struct {

	// When used via an API request, this is populated and used to populate id.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud waypoint waypoint service UI create action config with vars body namespace
func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBodyNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBodyNamespace) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint waypoint service UI create action config with vars body namespace based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBodyNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBodyNamespace) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBodyNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBodyNamespace) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceUICreateActionConfigWithVarsBodyNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
