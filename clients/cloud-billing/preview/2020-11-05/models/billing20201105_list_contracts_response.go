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

// Billing20201105ListContractsResponse ListContractsResponse returns a list of Contracts.
//
// swagger:model billing_20201105ListContractsResponse
type Billing20201105ListContractsResponse struct {

	// contracts is the list of fetched contracts.
	Contracts []*Billing20201105Contract `json:"contracts"`
}

// Validate validates this billing 20201105 list contracts response
func (m *Billing20201105ListContractsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContracts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105ListContractsResponse) validateContracts(formats strfmt.Registry) error {
	if swag.IsZero(m.Contracts) { // not required
		return nil
	}

	for i := 0; i < len(m.Contracts); i++ {
		if swag.IsZero(m.Contracts[i]) { // not required
			continue
		}

		if m.Contracts[i] != nil {
			if err := m.Contracts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contracts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contracts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this billing 20201105 list contracts response based on the context it is used
func (m *Billing20201105ListContractsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContracts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105ListContractsResponse) contextValidateContracts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Contracts); i++ {

		if m.Contracts[i] != nil {

			if swag.IsZero(m.Contracts[i]) { // not required
				return nil
			}

			if err := m.Contracts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contracts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contracts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105ListContractsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105ListContractsResponse) UnmarshalBinary(b []byte) error {
	var res Billing20201105ListContractsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
