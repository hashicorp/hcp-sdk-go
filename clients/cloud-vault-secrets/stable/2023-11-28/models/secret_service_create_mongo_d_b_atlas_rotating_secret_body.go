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

// SecretServiceCreateMongoDBAtlasRotatingSecretBody secret service create mongo d b atlas rotating secret body
//
// swagger:model SecretServiceCreateMongoDBAtlasRotatingSecretBody
type SecretServiceCreateMongoDBAtlasRotatingSecretBody struct {

	// integration name
	IntegrationName string `json:"integration_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// rotation policy name
	RotationPolicyName string `json:"rotation_policy_name,omitempty"`

	// secret details
	SecretDetails *Secrets20231128MongoDBAtlasSecretDetails `json:"secret_details,omitempty"`
}

// Validate validates this secret service create mongo d b atlas rotating secret body
func (m *SecretServiceCreateMongoDBAtlasRotatingSecretBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecretDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateMongoDBAtlasRotatingSecretBody) validateSecretDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.SecretDetails) { // not required
		return nil
	}

	if m.SecretDetails != nil {
		if err := m.SecretDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secret_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secret service create mongo d b atlas rotating secret body based on the context it is used
func (m *SecretServiceCreateMongoDBAtlasRotatingSecretBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecretDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateMongoDBAtlasRotatingSecretBody) contextValidateSecretDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.SecretDetails != nil {

		if swag.IsZero(m.SecretDetails) { // not required
			return nil
		}

		if err := m.SecretDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secret_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceCreateMongoDBAtlasRotatingSecretBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceCreateMongoDBAtlasRotatingSecretBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceCreateMongoDBAtlasRotatingSecretBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
