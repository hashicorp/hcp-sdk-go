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

// SecretServiceUpdateAzureIntegrationBody secret service update azure integration body
//
// swagger:model SecretServiceUpdateAzureIntegrationBody
type SecretServiceUpdateAzureIntegrationBody struct {

	// capabilities
	Capabilities []*Secrets20231128Capability `json:"capabilities"`

	// client secret
	ClientSecret *Secrets20231128AzureClientSecretRequest `json:"client_secret,omitempty"`

	// federated workload identity
	FederatedWorkloadIdentity *Secrets20231128AzureFederatedWorkloadIdentityRequest `json:"federated_workload_identity,omitempty"`

	// rotate on update
	RotateOnUpdate bool `json:"rotate_on_update,omitempty"`
}

// Validate validates this secret service update azure integration body
func (m *SecretServiceUpdateAzureIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFederatedWorkloadIdentity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceUpdateAzureIntegrationBody) validateCapabilities(formats strfmt.Registry) error {
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

func (m *SecretServiceUpdateAzureIntegrationBody) validateClientSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientSecret) { // not required
		return nil
	}

	if m.ClientSecret != nil {
		if err := m.ClientSecret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client_secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client_secret")
			}
			return err
		}
	}

	return nil
}

func (m *SecretServiceUpdateAzureIntegrationBody) validateFederatedWorkloadIdentity(formats strfmt.Registry) error {
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

// ContextValidate validate this secret service update azure integration body based on the context it is used
func (m *SecretServiceUpdateAzureIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClientSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFederatedWorkloadIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceUpdateAzureIntegrationBody) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SecretServiceUpdateAzureIntegrationBody) contextValidateClientSecret(ctx context.Context, formats strfmt.Registry) error {

	if m.ClientSecret != nil {

		if swag.IsZero(m.ClientSecret) { // not required
			return nil
		}

		if err := m.ClientSecret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client_secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client_secret")
			}
			return err
		}
	}

	return nil
}

func (m *SecretServiceUpdateAzureIntegrationBody) contextValidateFederatedWorkloadIdentity(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *SecretServiceUpdateAzureIntegrationBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceUpdateAzureIntegrationBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceUpdateAzureIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
