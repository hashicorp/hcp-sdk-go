// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128RandomParams secrets 20231128 random params
//
// swagger:model secrets_20231128RandomParams
type Secrets20231128RandomParams struct {

	// username template
	UsernameTemplate string `json:"username_template,omitempty"`
}

// Validate validates this secrets 20231128 random params
func (m *Secrets20231128RandomParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 random params based on context it is used
func (m *Secrets20231128RandomParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128RandomParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128RandomParams) UnmarshalBinary(b []byte) error {
	var res Secrets20231128RandomParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
