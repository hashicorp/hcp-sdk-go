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
)

// HashicorpCloudWaypointNoCodeModuleDefinition hashicorp cloud waypoint no code module definition
//
// swagger:model hashicorp.cloud.waypoint.NoCodeModuleDefinition
type HashicorpCloudWaypointNoCodeModuleDefinition struct {

	// created is the date the module was created
	//
	// date created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// last_updated is the date the module was last updated
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// module_id is the ID of the no-code module.
	ModuleID string `json:"module_id,omitempty"`

	// name is the name of the no-code module.
	Name string `json:"name,omitempty"`

	// pinned_version is the version of the module that is pinned for use in no-code.
	PinnedVersion string `json:"pinned_version,omitempty"`

	// provider is the Terraform provider for the no-code module.
	//
	// ex: AWS
	Provider string `json:"provider,omitempty"`

	// registry_name is the name of the Terraform registry where the module is
	// hosted. This is one of "public" or "private".
	RegistryName string `json:"registry_name,omitempty"`

	// tf_namespace is the name of the Terraform organization that owns the module
	TfNamespace string `json:"tf_namespace,omitempty"`

	// DEPRECATED: Do not use.
	Versions []string `json:"versions"`
}

// Validate validates this hashicorp cloud waypoint no code module definition
func (m *HashicorpCloudWaypointNoCodeModuleDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointNoCodeModuleDefinition) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudWaypointNoCodeModuleDefinition) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hashicorp cloud waypoint no code module definition based on context it is used
func (m *HashicorpCloudWaypointNoCodeModuleDefinition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointNoCodeModuleDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointNoCodeModuleDefinition) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointNoCodeModuleDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
