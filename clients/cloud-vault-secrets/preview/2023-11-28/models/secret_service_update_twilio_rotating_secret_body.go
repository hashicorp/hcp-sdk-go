// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecretServiceUpdateTwilioRotatingSecretBody secret service update twilio rotating secret body
//
// swagger:model SecretServiceUpdateTwilioRotatingSecretBody
type SecretServiceUpdateTwilioRotatingSecretBody struct {

	// rotate on update
	RotateOnUpdate bool `json:"rotate_on_update,omitempty"`

	// rotation policy name
	RotationPolicyName string `json:"rotation_policy_name,omitempty"`
}

// Validate validates this secret service update twilio rotating secret body
func (m *SecretServiceUpdateTwilioRotatingSecretBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret service update twilio rotating secret body based on context it is used
func (m *SecretServiceUpdateTwilioRotatingSecretBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceUpdateTwilioRotatingSecretBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceUpdateTwilioRotatingSecretBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceUpdateTwilioRotatingSecretBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
