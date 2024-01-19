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

// HashicorpCloudPacker20230101RegistryBillingDeprovision hashicorp cloud packer 20230101 registry billing deprovision
//
// swagger:model hashicorp.cloud.packer_20230101.RegistryBillingDeprovision
type HashicorpCloudPacker20230101RegistryBillingDeprovision struct {

	// Datetime when the registry was deactivated of billing.
	// Format: date-time
	At strfmt.DateTime `json:"at,omitempty"`

	// Reason for registry deactivation.
	Reason *HashicorpCloudPacker20230101RegistryBillingDeprovisionReason `json:"reason,omitempty"`
}

// Validate validates this hashicorp cloud packer 20230101 registry billing deprovision
func (m *HashicorpCloudPacker20230101RegistryBillingDeprovision) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101RegistryBillingDeprovision) validateAt(formats strfmt.Registry) error {
	if swag.IsZero(m.At) { // not required
		return nil
	}

	if err := validate.FormatOf("at", "body", "date-time", m.At.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPacker20230101RegistryBillingDeprovision) validateReason(formats strfmt.Registry) error {
	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if m.Reason != nil {
		if err := m.Reason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reason")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer 20230101 registry billing deprovision based on the context it is used
func (m *HashicorpCloudPacker20230101RegistryBillingDeprovision) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101RegistryBillingDeprovision) contextValidateReason(ctx context.Context, formats strfmt.Registry) error {

	if m.Reason != nil {

		if swag.IsZero(m.Reason) { // not required
			return nil
		}

		if err := m.Reason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reason")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101RegistryBillingDeprovision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101RegistryBillingDeprovision) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20230101RegistryBillingDeprovision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
