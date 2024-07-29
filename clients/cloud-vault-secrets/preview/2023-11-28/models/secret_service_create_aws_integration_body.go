// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecretServiceCreateAwsIntegrationBody secret service create aws integration body
//
// swagger:model SecretServiceCreateAwsIntegrationBody
type SecretServiceCreateAwsIntegrationBody struct {

	// capabilities
	Capabilities []*Secrets20231128Capability `json:"capabilities"`

	// federated workload identity
	FederatedWorkloadIdentity *Secrets20231128AwsFederatedWorkloadIdentityRequest `json:"federated_workload_identity,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// static credentials
	StaticCredentials *Secrets20231128AwsStaticCredentialsRequest `json:"static_credentials,omitempty"`
}

// Validate validates this secret service create aws integration body
func (m *SecretServiceCreateAwsIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFederatedWorkloadIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateAwsIntegrationBody) validateCapabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.Capabilities); i++ {
		if swag.IsZero(m.Capabilities[i]) { // not required
			continue
		}

		if m.Capabilities[i] != nil {
			if err := m.Capabilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("capabilities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("capabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SecretServiceCreateAwsIntegrationBody) validateFederatedWorkloadIdentity(formats strfmt.Registry) error {
	if swag.IsZero(m.FederatedWorkloadIdentity) { // not required
		return nil
	}

	if m.FederatedWorkloadIdentity != nil {
		if err := m.FederatedWorkloadIdentity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("federated_workload_identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("federated_workload_identity")
			}
			return err
		}
	}

	return nil
}

func (m *SecretServiceCreateAwsIntegrationBody) validateStaticCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticCredentials) { // not required
		return nil
	}

	if m.StaticCredentials != nil {
		if err := m.StaticCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("static_credentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secret service create aws integration body based on the context it is used
func (m *SecretServiceCreateAwsIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFederatedWorkloadIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateAwsIntegrationBody) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Capabilities); i++ {

		if m.Capabilities[i] != nil {

			if swag.IsZero(m.Capabilities[i]) { // not required
				return nil
			}

			if err := m.Capabilities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("capabilities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("capabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SecretServiceCreateAwsIntegrationBody) contextValidateFederatedWorkloadIdentity(ctx context.Context, formats strfmt.Registry) error {

	if m.FederatedWorkloadIdentity != nil {

		if swag.IsZero(m.FederatedWorkloadIdentity) { // not required
			return nil
		}

		if err := m.FederatedWorkloadIdentity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("federated_workload_identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("federated_workload_identity")
			}
			return err
		}
	}

	return nil
}

func (m *SecretServiceCreateAwsIntegrationBody) contextValidateStaticCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.StaticCredentials != nil {

		if swag.IsZero(m.StaticCredentials) { // not required
			return nil
		}

		if err := m.StaticCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("static_credentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceCreateAwsIntegrationBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceCreateAwsIntegrationBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceCreateAwsIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
