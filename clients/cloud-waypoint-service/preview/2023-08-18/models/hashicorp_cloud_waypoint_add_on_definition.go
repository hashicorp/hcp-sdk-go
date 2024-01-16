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
)

// HashicorpCloudWaypointAddOnDefinition AddOnDefinition is the template for an add-on
//
// swagger:model hashicorp.cloud.waypoint.AddOnDefinition
type HashicorpCloudWaypointAddOnDefinition struct {

	// Longer description of the Add-on
	Description string `json:"description,omitempty"`

	// Unique identifier of the Add-on definition
	ID string `json:"id,omitempty"`

	// A list of descriptive labels for an Add-on
	Labels []string `json:"labels"`

	// Name of the Add-on definition
	Name string `json:"name,omitempty"`

	// A templated README describing the Add-on
	// Format: byte
	ReadmeMarkdownTemplate strfmt.Base64 `json:"readme_markdown_template,omitempty"`

	// Short description of the Add-on
	Summary string `json:"summary,omitempty"`

	// kv tags
	Tags []*HashicorpCloudWaypointTag `json:"tags"`

	// Terraform No Code module used for provisioning the Add-on
	TerraformNocodeModule *HashicorpCloudWaypointTerraformNocodeModule `json:"terraform_nocode_module,omitempty"`

	// The TF variable set to apply to the Add-on's No Code workspace
	TfVariableSetIds []string `json:"tf_variable_set_ids"`
}

// Validate validates this hashicorp cloud waypoint add on definition
func (m *HashicorpCloudWaypointAddOnDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerraformNocodeModule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointAddOnDefinition) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointAddOnDefinition) validateTerraformNocodeModule(formats strfmt.Registry) error {
	if swag.IsZero(m.TerraformNocodeModule) { // not required
		return nil
	}

	if m.TerraformNocodeModule != nil {
		if err := m.TerraformNocodeModule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraform_nocode_module")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraform_nocode_module")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint add on definition based on the context it is used
func (m *HashicorpCloudWaypointAddOnDefinition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerraformNocodeModule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointAddOnDefinition) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointAddOnDefinition) contextValidateTerraformNocodeModule(ctx context.Context, formats strfmt.Registry) error {

	if m.TerraformNocodeModule != nil {

		if swag.IsZero(m.TerraformNocodeModule) { // not required
			return nil
		}

		if err := m.TerraformNocodeModule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraform_nocode_module")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraform_nocode_module")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointAddOnDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointAddOnDefinition) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointAddOnDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
