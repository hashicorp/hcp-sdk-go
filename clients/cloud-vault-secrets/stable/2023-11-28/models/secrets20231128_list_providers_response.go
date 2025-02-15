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

// Secrets20231128ListProvidersResponse secrets 20231128 list providers response
//
// swagger:model secrets_20231128ListProvidersResponse
type Secrets20231128ListProvidersResponse struct {

	// providers
	Providers []*Secrets20231128Provider `json:"providers"`
}

// Validate validates this secrets 20231128 list providers response
func (m *Secrets20231128ListProvidersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProviders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128ListProvidersResponse) validateProviders(formats strfmt.Registry) error {
	if swag.IsZero(m.Providers) { // not required
		return nil
	}

	for i := 0; i < len(m.Providers); i++ {
		if swag.IsZero(m.Providers[i]) { // not required
			continue
		}

		if m.Providers[i] != nil {
			if err := m.Providers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("providers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("providers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this secrets 20231128 list providers response based on the context it is used
func (m *Secrets20231128ListProvidersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProviders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128ListProvidersResponse) contextValidateProviders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Providers); i++ {

		if m.Providers[i] != nil {

			if swag.IsZero(m.Providers[i]) { // not required
				return nil
			}

			if err := m.Providers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("providers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("providers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128ListProvidersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128ListProvidersResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128ListProvidersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
