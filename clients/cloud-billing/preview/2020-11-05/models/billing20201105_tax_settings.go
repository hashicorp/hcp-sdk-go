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

// Billing20201105TaxSettings TaxSettings contain a Billing Account's settings related to taxation.
//
// These are oriented to Stripe's tax settings:
//
// https://stripe.com/docs/billing/taxes/tax-rates.
//
// swagger:model billing_20201105TaxSettings
type Billing20201105TaxSettings struct {

	// exemption_status indicates the customer’s tax exemption status which
	// can be one of none, exempt, or reverse.
	ExemptionStatus *TaxSettingsExemptionStatus `json:"exemption_status,omitempty"`

	// id is the tax ID to include on invoices.
	ID string `json:"id,omitempty"`

	// id_type is the tax ID type.
	IDType *TaxSettingsIDType `json:"id_type,omitempty"`
}

// Validate validates this billing 20201105 tax settings
func (m *Billing20201105TaxSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExemptionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105TaxSettings) validateExemptionStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ExemptionStatus) { // not required
		return nil
	}

	if m.ExemptionStatus != nil {
		if err := m.ExemptionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exemption_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exemption_status")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105TaxSettings) validateIDType(formats strfmt.Registry) error {
	if swag.IsZero(m.IDType) { // not required
		return nil
	}

	if m.IDType != nil {
		if err := m.IDType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("id_type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this billing 20201105 tax settings based on the context it is used
func (m *Billing20201105TaxSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExemptionStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIDType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105TaxSettings) contextValidateExemptionStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ExemptionStatus != nil {
		if err := m.ExemptionStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exemption_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exemption_status")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105TaxSettings) contextValidateIDType(ctx context.Context, formats strfmt.Registry) error {

	if m.IDType != nil {
		if err := m.IDType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("id_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105TaxSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105TaxSettings) UnmarshalBinary(b []byte) error {
	var res Billing20201105TaxSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}