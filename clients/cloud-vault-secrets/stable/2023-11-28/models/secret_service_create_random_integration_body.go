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

// SecretServiceCreateRandomIntegrationBody secret service create random integration body
//
// swagger:model SecretServiceCreateRandomIntegrationBody
type SecretServiceCreateRandomIntegrationBody struct {

	// capabilities
	Capabilities []*Secrets20231128Capability `json:"capabilities"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this secret service create random integration body
func (m *SecretServiceCreateRandomIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateRandomIntegrationBody) validateCapabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.Capabilities); i++ {
		if swag.IsZero(m.Capabilities[i]) { // not required
			continue
		}

		if m.Capabilities[i] != nil {
			if err := m.Capabilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("capabilities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("capabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this secret service create random integration body based on the context it is used
func (m *SecretServiceCreateRandomIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateRandomIntegrationBody) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Capabilities); i++ {

		if m.Capabilities[i] != nil {

			if swag.IsZero(m.Capabilities[i]) { // not required
				return nil
			}

			if err := m.Capabilities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("capabilities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("capabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceCreateRandomIntegrationBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceCreateRandomIntegrationBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceCreateRandomIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
