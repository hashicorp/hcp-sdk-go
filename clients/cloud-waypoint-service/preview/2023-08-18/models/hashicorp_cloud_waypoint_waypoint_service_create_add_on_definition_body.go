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

// HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody CreateAddOnDefinitionRequest is the request used to create an Add-on
// definition
//
// swagger:model hashicorp.cloud.waypoint.WaypointService.CreateAddOnDefinitionBody
type HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody struct {

	// Longer description of the Add-on
	Description string `json:"description,omitempty"`

	// A list of descriptive labels for an Add-on
	Labels []string `json:"labels"`

	// Name of the Add-on definition
	Name string `json:"name,omitempty"`

	// Global references the entire server. This is used in some APIs
	// as a way to read/write values that are server-global.
	Namespace interface{} `json:"namespace,omitempty"`

	// A templated README describing the Add-on
	// Format: byte
	ReadmeMarkdownTemplate strfmt.Base64 `json:"readme_markdown_template,omitempty"`

	// Short description of the Add-on
	Summary string `json:"summary,omitempty"`

	// The TF project
	TerraformCloudWorkspaceDetails *HashicorpCloudWaypointTerraformCloudWorkspaceDetails `json:"terraform_cloud_workspace_details,omitempty"`

	// Terraform No Code module used for provisioning the Add-on
	TerraformNocodeModule *HashicorpCloudWaypointTerraformNocodeModule `json:"terraform_nocode_module,omitempty"`

	// variable_options is the collection of input variables which may be set for an add-on.
	VariableOptions []*HashicorpCloudWaypointTFModuleVariable `json:"variable_options"`
}

// Validate validates this hashicorp cloud waypoint waypoint service create add on definition body
func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTerraformCloudWorkspaceDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerraformNocodeModule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariableOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody) validateTerraformCloudWorkspaceDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.TerraformCloudWorkspaceDetails) { // not required
		return nil
	}

	if m.TerraformCloudWorkspaceDetails != nil {
		if err := m.TerraformCloudWorkspaceDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraform_cloud_workspace_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraform_cloud_workspace_details")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody) validateTerraformNocodeModule(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody) validateVariableOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.VariableOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.VariableOptions); i++ {
		if swag.IsZero(m.VariableOptions[i]) { // not required
			continue
		}

		if m.VariableOptions[i] != nil {
			if err := m.VariableOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variable_options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variable_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint waypoint service create add on definition body based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTerraformCloudWorkspaceDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerraformNocodeModule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariableOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody) contextValidateTerraformCloudWorkspaceDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.TerraformCloudWorkspaceDetails != nil {

		if swag.IsZero(m.TerraformCloudWorkspaceDetails) { // not required
			return nil
		}

		if err := m.TerraformCloudWorkspaceDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraform_cloud_workspace_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraform_cloud_workspace_details")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody) contextValidateTerraformNocodeModule(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody) contextValidateVariableOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VariableOptions); i++ {

		if m.VariableOptions[i] != nil {

			if swag.IsZero(m.VariableOptions[i]) { // not required
				return nil
			}

			if err := m.VariableOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variable_options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variable_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceCreateAddOnDefinitionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
