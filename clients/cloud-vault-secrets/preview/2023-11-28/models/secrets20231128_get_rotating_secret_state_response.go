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

// Secrets20231128GetRotatingSecretStateResponse secrets 20231128 get rotating secret state response
//
// swagger:model secrets_20231128GetRotatingSecretStateResponse
type Secrets20231128GetRotatingSecretStateResponse struct {

	// state
	State *Secrets20231128RotatingSecretState `json:"state,omitempty"`
}

// Validate validates this secrets 20231128 get rotating secret state response
func (m *Secrets20231128GetRotatingSecretStateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128GetRotatingSecretStateResponse) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secrets 20231128 get rotating secret state response based on the context it is used
func (m *Secrets20231128GetRotatingSecretStateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128GetRotatingSecretStateResponse) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {

		if swag.IsZero(m.State) { // not required
			return nil
		}

		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128GetRotatingSecretStateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128GetRotatingSecretStateResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128GetRotatingSecretStateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
