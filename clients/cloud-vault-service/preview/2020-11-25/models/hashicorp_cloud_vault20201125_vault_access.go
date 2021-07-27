// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125VaultAccess hashicorp cloud vault 20201125 vault access
//
// swagger:model hashicorp.cloud.vault_20201125.VaultAccess
type HashicorpCloudVault20201125VaultAccess struct {

	// recovery key
	RecoveryKey string `json:"recovery_key,omitempty"`

	// root token
	RootToken string `json:"root_token,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 vault access
func (m *HashicorpCloudVault20201125VaultAccess) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125VaultAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125VaultAccess) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125VaultAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
