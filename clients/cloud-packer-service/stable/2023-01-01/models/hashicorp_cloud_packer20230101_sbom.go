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

// HashicorpCloudPacker20230101Sbom hashicorp cloud packer 20230101 sbom
//
// swagger:model hashicorp.cloud.packer_20230101.Sbom
type HashicorpCloudPacker20230101Sbom struct {

	// Format of the SBOM will either be CYCLONEDX or SPDX.
	Format *HashicorpCloudPacker20230101SbomFormat `json:"format,omitempty"`

	// Unique identifier (ULID).
	ID string `json:"id,omitempty"`

	// Name of the SBOM, user-settable.
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp cloud packer 20230101 sbom
func (m *HashicorpCloudPacker20230101Sbom) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101Sbom) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	if m.Format != nil {
		if err := m.Format.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("format")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer 20230101 sbom based on the context it is used
func (m *HashicorpCloudPacker20230101Sbom) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101Sbom) contextValidateFormat(ctx context.Context, formats strfmt.Registry) error {

	if m.Format != nil {

		if swag.IsZero(m.Format) { // not required
			return nil
		}

		if err := m.Format.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("format")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101Sbom) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101Sbom) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20230101Sbom
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
