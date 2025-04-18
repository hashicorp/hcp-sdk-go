// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128MssqlUserPasswordDetails secrets 20231128 mssql user password details
//
// swagger:model secrets_20231128MssqlUserPasswordDetails
type Secrets20231128MssqlUserPasswordDetails struct {

	// usernames
	Usernames []string `json:"usernames"`
}

// Validate validates this secrets 20231128 mssql user password details
func (m *Secrets20231128MssqlUserPasswordDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 mssql user password details based on context it is used
func (m *Secrets20231128MssqlUserPasswordDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128MssqlUserPasswordDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128MssqlUserPasswordDetails) UnmarshalBinary(b []byte) error {
	var res Secrets20231128MssqlUserPasswordDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
