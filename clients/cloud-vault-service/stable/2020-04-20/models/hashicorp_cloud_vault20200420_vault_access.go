// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20200420VaultAccess hashicorp cloud vault 20200420 vault access
//
// swagger:model hashicorp.cloud.vault_20200420.VaultAccess
type HashicorpCloudVault20200420VaultAccess struct {

	// recovery key
	RecoveryKey string `json:"recovery_key,omitempty"`

	// root token
	RootToken string `json:"root_token,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 vault access
func (m *HashicorpCloudVault20200420VaultAccess) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20200420 vault access based on context it is used
func (m *HashicorpCloudVault20200420VaultAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420VaultAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420VaultAccess) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420VaultAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
