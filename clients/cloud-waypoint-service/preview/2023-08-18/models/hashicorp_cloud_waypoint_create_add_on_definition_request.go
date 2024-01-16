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

// HashicorpCloudWaypointCreateAddOnDefinitionRequest CreateAddOnDefinitionRequest is the request used to create an Add-on
// definition
//
// swagger:model hashicorp.cloud.waypoint.CreateAddOnDefinitionRequest
type HashicorpCloudWaypointCreateAddOnDefinitionRequest struct {

	// Longer description of the Add-on
	Description string `json:"description,omitempty"`

	// A list of descriptive labels for an Add-on
	Labels []string `json:"labels"`

	// Name of the Add-on definition
	Name string `json:"name,omitempty"`

	// namespace
	Namespace *HashicorpCloudWaypointRefNamespace `json:"namespace,omitempty"`

	// A templated README describing the Add-on
	// Format: byte
	ReadmeMarkdownTemplate strfmt.Base64 `json:"readme_markdown_template,omitempty"`

	// Short description of the Add-on
	Summary string `json:"summary,omitempty"`

	// Terraform No Code module used for provisioning the Add-on
	TerraformNocodeModule *HashicorpCloudWaypointTerraformNocodeModule `json:"terraform_nocode_module,omitempty"`

	// The TF variable set to apply to the Add-on's No Code workspace
	TfVariableSetIds []string `json:"tf_variable_set_ids"`
}

// Validate validates this hashicorp cloud waypoint create add on definition request
func (m *HashicorpCloudWaypointCreateAddOnDefinitionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespace(formats); err != nil {
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

func (m *HashicorpCloudWaypointCreateAddOnDefinitionRequest) validateNamespace(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointCreateAddOnDefinitionRequest) validateTerraformNocodeModule(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint create add on definition request based on the context it is used
func (m *HashicorpCloudWaypointCreateAddOnDefinitionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
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

func (m *HashicorpCloudWaypointCreateAddOnDefinitionRequest) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointCreateAddOnDefinitionRequest) contextValidateTerraformNocodeModule(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointCreateAddOnDefinitionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointCreateAddOnDefinitionRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointCreateAddOnDefinitionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
