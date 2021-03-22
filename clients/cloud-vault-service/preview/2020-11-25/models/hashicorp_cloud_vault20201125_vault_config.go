// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125VaultConfig hashicorp cloud vault 20201125 vault config
//
// swagger:model hashicorp.cloud.vault_20201125.VaultConfig
type HashicorpCloudVault20201125VaultConfig struct {

	// initial_version is the initial Vault version to use when creating the
	// cluster. Once the cluster is created, this value is no longer used.
	InitialVersion string `json:"initial_version,omitempty"`

	// max_lease_ttl is the max lease ttl for this Vault cluster.
	MaxLeaseTTL string `json:"max_lease_ttl,omitempty"`

	// namespace is the name of the customer namespace.
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 vault config
func (m *HashicorpCloudVault20201125VaultConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125VaultConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125VaultConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125VaultConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
