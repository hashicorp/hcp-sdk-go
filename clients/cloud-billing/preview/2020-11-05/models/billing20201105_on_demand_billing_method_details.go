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

// Billing20201105OnDemandBillingMethodDetails OnDemandBillingMethodDetails contains both the information used to register a
// Billing Account billed on-demand, and additional information for display
// purposes (e.g. credit card details).
//
// swagger:model billing_20201105OnDemandBillingMethodDetails
type Billing20201105OnDemandBillingMethodDetails struct {

	// billing_method contains the information used to register an on-demand
	// Billing Account.
	BillingMethod *Billing20201105OnDemandBillingMethod `json:"billing_method,omitempty"`

	// card_details contains additional information about the credit/debit card
	// for display purposes.
	CardDetails *Billing20201105CardDetails `json:"card_details,omitempty"`
}

// Validate validates this billing 20201105 on demand billing method details
func (m *Billing20201105OnDemandBillingMethodDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105OnDemandBillingMethodDetails) validateBillingMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingMethod) { // not required
		return nil
	}

	if m.BillingMethod != nil {
		if err := m.BillingMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing_method")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105OnDemandBillingMethodDetails) validateCardDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.CardDetails) { // not required
		return nil
	}

	if m.CardDetails != nil {
		if err := m.CardDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("card_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this billing 20201105 on demand billing method details based on the context it is used
func (m *Billing20201105OnDemandBillingMethodDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillingMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCardDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105OnDemandBillingMethodDetails) contextValidateBillingMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.BillingMethod != nil {

		if swag.IsZero(m.BillingMethod) { // not required
			return nil
		}

		if err := m.BillingMethod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing_method")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105OnDemandBillingMethodDetails) contextValidateCardDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.CardDetails != nil {

		if swag.IsZero(m.CardDetails) { // not required
			return nil
		}

		if err := m.CardDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("card_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105OnDemandBillingMethodDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105OnDemandBillingMethodDetails) UnmarshalBinary(b []byte) error {
	var res Billing20201105OnDemandBillingMethodDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
