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
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudVault20201125Cluster Cluster represents a single Vault cluster.
//
// swagger:model hashicorp.cloud.vault_20201125.Cluster
type HashicorpCloudVault20201125Cluster struct {

	// config holds the configuration of the cluster.
	Config *HashicorpCloudVault20201125ClusterConfig `json:"config,omitempty"`

	// created_at is the timestamp of when the cluster was first created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// CreationMetadata related to the creation of the cluster including any template output
	CreationMetadata *HashicorpCloudVault20201125ClusterCreationMetadata `json:"creation_metadata,omitempty"`

	// currently_deployed_version is the version of the current Vault deployment.
	CurrentVersion string `json:"current_version,omitempty"`

	// DNSNames holds all of the cluster's DNS names.
	DNSNames *HashicorpCloudVault20201125ClusterDNSNames `json:"dns_names,omitempty"`

	// id is ID of the Vault cluster.
	ID string `json:"id,omitempty"`

	// location is the location of the cluster.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// notifications is the list of notifications currently valid for the cluster.
	Notifications []*HashicorpCloudVault20201125ClusterNotification `json:"notifications"`

	// performance_replication_info contains the performance replication information.
	PerformanceReplicationInfo *HashicorpCloudVault20201125ClusterPerformanceReplicationInfo `json:"performance_replication_info,omitempty"`

	// resource_id is UUID of the Vault cluster.
	ResourceID string `json:"resource_id,omitempty"`

	// state is the current state of the cluster.
	State *HashicorpCloudVault20201125ClusterState `json:"state,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 cluster
func (m *HashicorpCloudVault20201125Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceReplicationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125Cluster) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) validateCreationMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationMetadata) { // not required
		return nil
	}

	if m.CreationMetadata != nil {
		if err := m.CreationMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creation_metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creation_metadata")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) validateDNSNames(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSNames) { // not required
		return nil
	}

	if m.DNSNames != nil {
		if err := m.DNSNames.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns_names")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns_names")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) validateNotifications(formats strfmt.Registry) error {
	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	for i := 0; i < len(m.Notifications); i++ {
		if swag.IsZero(m.Notifications[i]) { // not required
			continue
		}

		if m.Notifications[i] != nil {
			if err := m.Notifications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notifications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) validatePerformanceReplicationInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceReplicationInfo) { // not required
		return nil
	}

	if m.PerformanceReplicationInfo != nil {
		if err := m.PerformanceReplicationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("performance_replication_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("performance_replication_info")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20201125 cluster based on the context it is used
func (m *HashicorpCloudVault20201125Cluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreationMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNSNames(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerformanceReplicationInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125Cluster) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) contextValidateCreationMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.CreationMetadata != nil {
		if err := m.CreationMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creation_metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creation_metadata")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) contextValidateDNSNames(ctx context.Context, formats strfmt.Registry) error {

	if m.DNSNames != nil {
		if err := m.DNSNames.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns_names")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns_names")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Notifications); i++ {

		if m.Notifications[i] != nil {
			if err := m.Notifications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notifications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) contextValidatePerformanceReplicationInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.PerformanceReplicationInfo != nil {
		if err := m.PerformanceReplicationInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("performance_replication_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("performance_replication_info")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125Cluster) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125Cluster) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
