// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecretServiceLockProjectBody --- Locks ---
//
// swagger:model SecretServiceLockProjectBody
type SecretServiceLockProjectBody struct {

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this secret service lock project body
func (m *SecretServiceLockProjectBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret service lock project body based on context it is used
func (m *SecretServiceLockProjectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceLockProjectBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceLockProjectBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceLockProjectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
