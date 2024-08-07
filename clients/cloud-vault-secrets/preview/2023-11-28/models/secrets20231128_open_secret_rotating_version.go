// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Secrets20231128OpenSecretRotatingVersion secrets 20231128 open secret rotating version
//
// swagger:model secrets_20231128OpenSecretRotatingVersion
type Secrets20231128OpenSecretRotatingVersion struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// expires at
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty"`

	// keys
	Keys []string `json:"keys"`

	// revoked at
	// Format: date-time
	RevokedAt strfmt.DateTime `json:"revoked_at,omitempty"`

	// values
	Values map[string]string `json:"values,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this secrets 20231128 open secret rotating version
func (m *Secrets20231128OpenSecretRotatingVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevokedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128OpenSecretRotatingVersion) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128OpenSecretRotatingVersion) validateExpiresAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128OpenSecretRotatingVersion) validateRevokedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.RevokedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("revoked_at", "body", "date-time", m.RevokedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this secrets 20231128 open secret rotating version based on context it is used
func (m *Secrets20231128OpenSecretRotatingVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128OpenSecretRotatingVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128OpenSecretRotatingVersion) UnmarshalBinary(b []byte) error {
	var res Secrets20231128OpenSecretRotatingVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
