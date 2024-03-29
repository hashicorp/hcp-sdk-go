// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128VercelProject secrets 20231128 vercel project
//
// swagger:model secrets_20231128VercelProject
type Secrets20231128VercelProject struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this secrets 20231128 vercel project
func (m *Secrets20231128VercelProject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 vercel project based on context it is used
func (m *Secrets20231128VercelProject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128VercelProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128VercelProject) UnmarshalBinary(b []byte) error {
	var res Secrets20231128VercelProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
