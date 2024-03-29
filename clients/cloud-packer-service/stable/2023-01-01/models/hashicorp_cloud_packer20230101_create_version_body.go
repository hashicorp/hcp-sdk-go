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

// HashicorpCloudPacker20230101CreateVersionBody hashicorp cloud packer 20230101 create version body
//
// swagger:model hashicorp.cloud.packer_20230101.CreateVersionBody
type HashicorpCloudPacker20230101CreateVersionBody struct {

	// Fingerprint of the version set by Packer when you call `packer build`.
	// Refer to the Packer documentation for more information on how this value is set.
	// The fingerprint can be used as an identifier for the version.
	// A valid fingerprint is 1-40 characters long, begins and ends with a letter or number,
	// and contains only ASCII letters, numbers, hyphens, dots, and underscores.
	Fingerprint string `json:"fingerprint,omitempty"`

	// The type of Packer configuration template used to build this version.
	TemplateType *HashicorpCloudPacker20230101TemplateType `json:"template_type,omitempty"`
}

// Validate validates this hashicorp cloud packer 20230101 create version body
func (m *HashicorpCloudPacker20230101CreateVersionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemplateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101CreateVersionBody) validateTemplateType(formats strfmt.Registry) error {
	if swag.IsZero(m.TemplateType) { // not required
		return nil
	}

	if m.TemplateType != nil {
		if err := m.TemplateType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template_type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer 20230101 create version body based on the context it is used
func (m *HashicorpCloudPacker20230101CreateVersionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTemplateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101CreateVersionBody) contextValidateTemplateType(ctx context.Context, formats strfmt.Registry) error {

	if m.TemplateType != nil {

		if swag.IsZero(m.TemplateType) { // not required
			return nil
		}

		if err := m.TemplateType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101CreateVersionBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101CreateVersionBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20230101CreateVersionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
