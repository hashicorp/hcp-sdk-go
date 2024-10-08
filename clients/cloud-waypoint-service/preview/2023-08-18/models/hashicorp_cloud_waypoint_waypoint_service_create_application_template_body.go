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

// HashicorpCloudWaypointWaypointServiceCreateApplicationTemplateBody hashicorp cloud waypoint waypoint service create application template body
//
// swagger:model hashicorp.cloud.waypoint.WaypointService.CreateApplicationTemplateBody
type HashicorpCloudWaypointWaypointServiceCreateApplicationTemplateBody struct {

	// application template
	ApplicationTemplate *HashicorpCloudWaypointApplicationTemplate `json:"application_template,omitempty"`

	// Global references the entire server. This is used in some APIs
	// as a way to read/write values that are server-global.
	Namespace interface{} `json:"namespace,omitempty"`

	// If true, will auto-import the readme from the Terraform module used
	// rather than the one specified in application_template.
	UseModuleReadme bool `json:"use_module_readme,omitempty"`
}

// Validate validates this hashicorp cloud waypoint waypoint service create application template body
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationTemplateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationTemplateBody) validateApplicationTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationTemplate) { // not required
		return nil
	}

	if m.ApplicationTemplate != nil {
		if err := m.ApplicationTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application_template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint waypoint service create application template body based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationTemplateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationTemplateBody) contextValidateApplicationTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.ApplicationTemplate != nil {

		if swag.IsZero(m.ApplicationTemplate) { // not required
			return nil
		}

		if err := m.ApplicationTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application_template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationTemplateBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationTemplateBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceCreateApplicationTemplateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
