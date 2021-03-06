// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20200420ClusterConfig hashicorp cloud vault 20200420 cluster config
//
// swagger:model hashicorp.cloud.vault_20200420.ClusterConfig
type HashicorpCloudVault20200420ClusterConfig struct {

	// audit config
	AuditConfig *HashicorpCloudVault20200420AuditConfig `json:"audit_config,omitempty"`

	// capacity config
	CapacityConfig *HashicorpCloudVault20200420CapacityConfig `json:"capacity_config,omitempty"`

	// maintenance config
	MaintenanceConfig HashicorpCloudVault20200420MaintenanceConfig `json:"maintenance_config,omitempty"`

	// network config
	NetworkConfig *HashicorpCloudVault20200420NetworkConfig `json:"network_config,omitempty"`

	// public ips disabled
	PublicIpsDisabled bool `json:"public_ips_disabled,omitempty"`

	// snapshot config
	SnapshotConfig *HashicorpCloudVault20200420SnapshotConfig `json:"snapshot_config,omitempty"`

	// vault access
	VaultAccess *HashicorpCloudVault20200420VaultAccess `json:"vault_access,omitempty"`

	// vault config
	VaultConfig *HashicorpCloudVault20200420VaultConfig `json:"vault_config,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 cluster config
func (m *HashicorpCloudVault20200420ClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacityConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420ClusterConfig) validateAuditConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditConfig) { // not required
		return nil
	}

	if m.AuditConfig != nil {
		if err := m.AuditConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audit_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20200420ClusterConfig) validateCapacityConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.CapacityConfig) { // not required
		return nil
	}

	if m.CapacityConfig != nil {
		if err := m.CapacityConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20200420ClusterConfig) validateNetworkConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkConfig) { // not required
		return nil
	}

	if m.NetworkConfig != nil {
		if err := m.NetworkConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20200420ClusterConfig) validateSnapshotConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.SnapshotConfig) { // not required
		return nil
	}

	if m.SnapshotConfig != nil {
		if err := m.SnapshotConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20200420ClusterConfig) validateVaultAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.VaultAccess) { // not required
		return nil
	}

	if m.VaultAccess != nil {
		if err := m.VaultAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vault_access")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20200420ClusterConfig) validateVaultConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.VaultConfig) { // not required
		return nil
	}

	if m.VaultConfig != nil {
		if err := m.VaultConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vault_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420ClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420ClusterConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420ClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
