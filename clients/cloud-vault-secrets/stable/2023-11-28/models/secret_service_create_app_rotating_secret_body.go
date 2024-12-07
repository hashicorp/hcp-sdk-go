// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecretServiceCreateAppRotatingSecretBody secret service create app rotating secret body
//
// swagger:model SecretServiceCreateAppRotatingSecretBody
type SecretServiceCreateAppRotatingSecretBody struct {

	// integration name
	IntegrationName string `json:"integration_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// rotation policy name
	RotationPolicyName string `json:"rotation_policy_name,omitempty"`

	// secret details
	SecretDetails interface{} `json:"secret_details,omitempty"`
}

// Validate validates this secret service create app rotating secret body
func (m *SecretServiceCreateAppRotatingSecretBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret service create app rotating secret body based on context it is used
func (m *SecretServiceCreateAppRotatingSecretBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceCreateAppRotatingSecretBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceCreateAppRotatingSecretBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceCreateAppRotatingSecretBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
