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
	"github.com/go-openapi/validate"
)

// Secrets20231128AzureIntegration secrets 20231128 azure integration
//
// swagger:model secrets_20231128AzureIntegration
type Secrets20231128AzureIntegration struct {

	// capabilities
	Capabilities []*Secrets20231128Capability `json:"capabilities"`

	// client secret
	ClientSecret *Secrets20231128AzureClientSecretResponse `json:"client_secret,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// federated workload identity
	FederatedWorkloadIdentity *Secrets20231128AzureFederatedWorkloadIdentityResponse `json:"federated_workload_identity,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// resource id
	ResourceID string `json:"resource_id,omitempty"`

	// resource name
	ResourceName string `json:"resource_name,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// updated by id
	UpdatedByID string `json:"updated_by_id,omitempty"`

	// used by
	UsedBy map[string]Secrets20231128IntegrationUsage `json:"used_by,omitempty"`
}

// Validate validates this secrets 20231128 azure integration
func (m *Secrets20231128AzureIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFederatedWorkloadIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128AzureIntegration) validateCapabilities(formats strfmt.Registry) error {
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

func (m *Secrets20231128AzureIntegration) validateClientSecret(formats strfmt.Registry) error {
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

func (m *Secrets20231128AzureIntegration) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128AzureIntegration) validateFederatedWorkloadIdentity(formats strfmt.Registry) error {
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

func (m *Secrets20231128AzureIntegration) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128AzureIntegration) validateUsedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.UsedBy) { // not required
		return nil
	}

	for k := range m.UsedBy {

		if err := validate.Required("used_by"+"."+k, "body", m.UsedBy[k]); err != nil {
			return err
		}
		if val, ok := m.UsedBy[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("used_by" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("used_by" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this secrets 20231128 azure integration based on the context it is used
func (m *Secrets20231128AzureIntegration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateUsedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128AzureIntegration) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Secrets20231128AzureIntegration) contextValidateClientSecret(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Secrets20231128AzureIntegration) contextValidateFederatedWorkloadIdentity(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Secrets20231128AzureIntegration) contextValidateUsedBy(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.UsedBy {

		if val, ok := m.UsedBy[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128AzureIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128AzureIntegration) UnmarshalBinary(b []byte) error {
	var res Secrets20231128AzureIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
