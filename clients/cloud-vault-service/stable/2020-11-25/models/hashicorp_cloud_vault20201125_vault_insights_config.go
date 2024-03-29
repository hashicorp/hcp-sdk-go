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

// HashicorpCloudVault20201125VaultInsightsConfig hashicorp cloud vault 20201125 vault insights config
//
// swagger:model hashicorp.cloud.vault_20201125.VaultInsightsConfig
type HashicorpCloudVault20201125VaultInsightsConfig struct {

	// enabled controls the streaming of audit logs to Vault Insights
	Enabled bool `json:"enabled,omitempty"`

	// last_changed indicates the last time when enabled was changed
	// Format: date-time
	LastChanged strfmt.DateTime `json:"last_changed,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 vault insights config
func (m *HashicorpCloudVault20201125VaultInsightsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastChanged(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125VaultInsightsConfig) validateLastChanged(formats strfmt.Registry) error {
	if swag.IsZero(m.LastChanged) { // not required
		return nil
	}

	if err := validate.FormatOf("last_changed", "body", "date-time", m.LastChanged.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hashicorp cloud vault 20201125 vault insights config based on context it is used
func (m *HashicorpCloudVault20201125VaultInsightsConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125VaultInsightsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125VaultInsightsConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125VaultInsightsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
