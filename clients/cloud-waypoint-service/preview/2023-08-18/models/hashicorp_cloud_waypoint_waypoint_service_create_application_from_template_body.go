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

// HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody hashicorp cloud waypoint waypoint service create application from template body
//
// swagger:model hashicorp.cloud.waypoint.WaypointService.CreateApplicationFromTemplateBody
type HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody struct {

	// Attach any action config refs to the application
	ActionCfgRefs []*HashicorpCloudWaypointActionCfgRef `json:"action_cfg_refs"`

	// application_template is the name of the application template on which the
	// new application will be based
	ApplicationTemplate *HashicorpCloudWaypointRefApplicationTemplate `json:"application_template,omitempty"`

	// name is the name of the new application
	Name string `json:"name,omitempty"`

	// namespace
	Namespace *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBodyNamespace `json:"namespace,omitempty"`

	// variables is the series of input variables which have been set by the
	// application developer for the new application being created. This may be empty.
	Variables []*HashicorpCloudWaypointInputVariable `json:"variables"`
}

// Validate validates this hashicorp cloud waypoint waypoint service create application from template body
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionCfgRefs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationTemplate(formats); err != nil {
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

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) validateActionCfgRefs(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionCfgRefs) { // not required
		return nil
	}

	for i := 0; i < len(m.ActionCfgRefs); i++ {
		if swag.IsZero(m.ActionCfgRefs[i]) { // not required
			continue
		}

		if m.ActionCfgRefs[i] != nil {
			if err := m.ActionCfgRefs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_cfg_refs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_cfg_refs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) validateApplicationTemplate(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) validateNamespace(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) validateVariables(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint waypoint service create application from template body based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionCfgRefs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplicationTemplate(ctx, formats); err != nil {
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

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) contextValidateActionCfgRefs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ActionCfgRefs); i++ {

		if m.ActionCfgRefs[i] != nil {

			if swag.IsZero(m.ActionCfgRefs[i]) { // not required
				return nil
			}

			if err := m.ActionCfgRefs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_cfg_refs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_cfg_refs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) contextValidateApplicationTemplate(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) contextValidateVariables(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBodyNamespace Global references the entire server. This is used in some APIs
// as a way to read/write values that are server-global.
//
// swagger:model HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBodyNamespace
type HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBodyNamespace struct {

	// When used via an API request, this is populated and used to populate id.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud waypoint waypoint service create application from template body namespace
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBodyNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBodyNamespace) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint waypoint service create application from template body namespace based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBodyNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBodyNamespace) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBodyNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBodyNamespace) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBodyNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
