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

// Billing20201105CreateBillingAccountResponse CreateBillingAccountResponse is the response from creating a new Billing
// Account for an organization.
//
// swagger:model billing_20201105CreateBillingAccountResponse
type Billing20201105CreateBillingAccountResponse struct {

	// billing_account is the Billing Account that was created.
	BillingAccount *Billing20201105BillingAccount `json:"billing_account,omitempty"`

	// entitlement_billing_method_details contains details about the entitlement contract
	// that can be attached to a billing account to cover charges of it.
	EntitlementBillingMethodDetails *Billing20201105EntitlementBillingMethodDetails `json:"entitlement_billing_method_details,omitempty"`

	// flexible_consumption_billing_method_details contains details of the billing
	// account's flexible consumption billing method.
	// Only present when the billing account is activated for Flexible Consumption.
	FlexibleConsumptionBillingMethodDetails Billing20201105FlexibleConsumptionBillingMethodDetails `json:"flexible_consumption_billing_method_details,omitempty"`

	// on_demand_billing_method_details contains details about the on-demand
	// billing method of the billing account.
	OnDemandBillingMethodDetails *Billing20201105OnDemandBillingMethodDetails `json:"on_demand_billing_method_details,omitempty"`
}

// Validate validates this billing 20201105 create billing account response
func (m *Billing20201105CreateBillingAccountResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlementBillingMethodDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnDemandBillingMethodDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105CreateBillingAccountResponse) validateBillingAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingAccount) { // not required
		return nil
	}

	if m.BillingAccount != nil {
		if err := m.BillingAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing_account")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105CreateBillingAccountResponse) validateEntitlementBillingMethodDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.EntitlementBillingMethodDetails) { // not required
		return nil
	}

	if m.EntitlementBillingMethodDetails != nil {
		if err := m.EntitlementBillingMethodDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entitlement_billing_method_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entitlement_billing_method_details")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105CreateBillingAccountResponse) validateOnDemandBillingMethodDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.OnDemandBillingMethodDetails) { // not required
		return nil
	}

	if m.OnDemandBillingMethodDetails != nil {
		if err := m.OnDemandBillingMethodDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("on_demand_billing_method_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("on_demand_billing_method_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this billing 20201105 create billing account response based on the context it is used
func (m *Billing20201105CreateBillingAccountResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillingAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntitlementBillingMethodDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnDemandBillingMethodDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105CreateBillingAccountResponse) contextValidateBillingAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.BillingAccount != nil {

		if swag.IsZero(m.BillingAccount) { // not required
			return nil
		}

		if err := m.BillingAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing_account")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105CreateBillingAccountResponse) contextValidateEntitlementBillingMethodDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.EntitlementBillingMethodDetails != nil {

		if swag.IsZero(m.EntitlementBillingMethodDetails) { // not required
			return nil
		}

		if err := m.EntitlementBillingMethodDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entitlement_billing_method_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entitlement_billing_method_details")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105CreateBillingAccountResponse) contextValidateOnDemandBillingMethodDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.OnDemandBillingMethodDetails != nil {

		if swag.IsZero(m.OnDemandBillingMethodDetails) { // not required
			return nil
		}

		if err := m.OnDemandBillingMethodDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("on_demand_billing_method_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("on_demand_billing_method_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105CreateBillingAccountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105CreateBillingAccountResponse) UnmarshalBinary(b []byte) error {
	var res Billing20201105CreateBillingAccountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
