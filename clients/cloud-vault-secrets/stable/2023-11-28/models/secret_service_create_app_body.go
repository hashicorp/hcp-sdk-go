// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecretServiceCreateAppBody secret service create app body
//
// swagger:model SecretServiceCreateAppBody
type SecretServiceCreateAppBody struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// sync names
	SyncNames []string `json:"sync_names"`
}

// Validate validates this secret service create app body
func (m *SecretServiceCreateAppBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret service create app body based on context it is used
func (m *SecretServiceCreateAppBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceCreateAppBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceCreateAppBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceCreateAppBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
