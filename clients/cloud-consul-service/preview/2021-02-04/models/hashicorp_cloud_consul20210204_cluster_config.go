// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudConsul20210204ClusterConfig ClusterConfig holds the configuration for a Consul cluster.
//
// swagger:model hashicorp.cloud.consul_20210204.ClusterConfig
type HashicorpCloudConsul20210204ClusterConfig struct {

	// auto_hvn_to_hvn_peering is only used together with the field consul_config.primary when
	// creating secondary clusters in a federation. Enable auto_hvn_to_hvn_peering if the
	// secondary HVN should be peered automatically to every other cluster's HVN in the federation.
	// If left disabled, the peering has to be done manually.
	AutoHvnToHvnPeering bool `json:"auto_hvn_to_hvn_peering,omitempty"`

	// capacity_config contains the configuration for the cluster size settings.
	CapacityConfig *HashicorpCloudConsul20210204CapacityConfig `json:"capacity_config,omitempty"`

	// consul_config allows altering certain fields relating to the Consul
	// cluster’s configuration.
	ConsulConfig *HashicorpCloudConsul20210204ConsulConfig `json:"consul_config,omitempty"`

	// maintenance_config contains the configuration for maintenance events such as
	// auto upgrades.
	MaintenanceConfig HashicorpCloudConsul20210204MaintenanceConfig `json:"maintenance_config,omitempty"`

	// NetworkConfig contains the network to launch the Consul cluster into.
	NetworkConfig *HashicorpCloudConsul20210204NetworkConfig `json:"network_config,omitempty"`

	// primary is readonly and contains a link to the primary consul cluster in a federation.
	// If this link points to itself, this cluster is the primary of a federation.
	// If the link points to another cluster, this cluster is a secondary in a federation.
	// Use consul_config.primary to federate clusters. The difference between these two fields
	// is that this field is present on primaries and secondaries. Whereas consul_config.primary
	// is only present on secondaries.
	// Output only.
	// Read Only: true
	Primary *cloud.HashicorpCloudLocationLink `json:"primary,omitempty"`

	// snapshot_config contains the configuration for how often to snapshot and
	// how many to maintain.
	SnapshotConfig HashicorpCloudConsul20210204SnapshotConfig `json:"snapshot_config,omitempty"`

	// tier is the type of tier this Consul cluster should be provisioned as.
	Tier *HashicorpCloudConsul20210204ClusterConfigTier `json:"tier,omitempty"`
}

// Validate validates this hashicorp cloud consul 20210204 cluster config
func (m *HashicorpCloudConsul20210204ClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacityConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsulConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20210204ClusterConfig) validateCapacityConfig(formats strfmt.Registry) error {
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

func (m *HashicorpCloudConsul20210204ClusterConfig) validateConsulConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsulConfig) { // not required
		return nil
	}

	if m.ConsulConfig != nil {
		if err := m.ConsulConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consul_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsul20210204ClusterConfig) validateNetworkConfig(formats strfmt.Registry) error {
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

func (m *HashicorpCloudConsul20210204ClusterConfig) validatePrimary(formats strfmt.Registry) error {
	if swag.IsZero(m.Primary) { // not required
		return nil
	}

	if m.Primary != nil {
		if err := m.Primary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsul20210204ClusterConfig) validateTier(formats strfmt.Registry) error {
	if swag.IsZero(m.Tier) { // not required
		return nil
	}

	if m.Tier != nil {
		if err := m.Tier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tier")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul 20210204 cluster config based on the context it is used
func (m *HashicorpCloudConsul20210204ClusterConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacityConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsulConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20210204ClusterConfig) contextValidateCapacityConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.CapacityConfig != nil {
		if err := m.CapacityConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsul20210204ClusterConfig) contextValidateConsulConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ConsulConfig != nil {
		if err := m.ConsulConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consul_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsul20210204ClusterConfig) contextValidateNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkConfig != nil {
		if err := m.NetworkConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsul20210204ClusterConfig) contextValidatePrimary(ctx context.Context, formats strfmt.Registry) error {

	if m.Primary != nil {
		if err := m.Primary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsul20210204ClusterConfig) contextValidateTier(ctx context.Context, formats strfmt.Registry) error {

	if m.Tier != nil {
		if err := m.Tier.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tier")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20210204ClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20210204ClusterConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20210204ClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
