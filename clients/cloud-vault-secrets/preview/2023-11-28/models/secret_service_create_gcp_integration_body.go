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

// SecretServiceCreateGcpIntegrationBody secret service create gcp integration body
//
// swagger:model SecretServiceCreateGcpIntegrationBody
type SecretServiceCreateGcpIntegrationBody struct {

	// capabilities
	Capabilities []*Secrets20231128Capability `json:"capabilities"`

	// federated workload identity
	FederatedWorkloadIdentity *Secrets20231128GcpFederatedWorkloadIdentityRequest `json:"federated_workload_identity,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// service account key
	ServiceAccountKey *Secrets20231128GcpServiceAccountKeyRequest `json:"service_account_key,omitempty"`
}

// Validate validates this secret service create gcp integration body
func (m *SecretServiceCreateGcpIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFederatedWorkloadIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAccountKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateGcpIntegrationBody) validateCapabilities(formats strfmt.Registry) error {
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

func (m *SecretServiceCreateGcpIntegrationBody) validateFederatedWorkloadIdentity(formats strfmt.Registry) error {
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

func (m *SecretServiceCreateGcpIntegrationBody) validateServiceAccountKey(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceAccountKey) { // not required
		return nil
	}

	if m.ServiceAccountKey != nil {
		if err := m.ServiceAccountKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_account_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_account_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secret service create gcp integration body based on the context it is used
func (m *SecretServiceCreateGcpIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFederatedWorkloadIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceAccountKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateGcpIntegrationBody) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SecretServiceCreateGcpIntegrationBody) contextValidateFederatedWorkloadIdentity(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SecretServiceCreateGcpIntegrationBody) contextValidateServiceAccountKey(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceAccountKey != nil {

		if swag.IsZero(m.ServiceAccountKey) { // not required
			return nil
		}

		if err := m.ServiceAccountKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_account_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_account_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceCreateGcpIntegrationBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceCreateGcpIntegrationBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceCreateGcpIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
