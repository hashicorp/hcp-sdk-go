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

// HashicorpCloudVaultLink20221107LinkedCluster hashicorp cloud vault link 20221107 linked cluster
//
// swagger:model hashicorp.cloud.vault_link_20221107.LinkedCluster
type HashicorpCloudVaultLink20221107LinkedCluster struct {

	// autopilot enabled
	AutopilotEnabled bool `json:"autopilot_enabled,omitempty"`

	// This is the self-managed cluster name and it differs from the HCP's internal id and HCP's slug id because
	// it is configured directly in the self-managed cluster.
	ClusterName string `json:"cluster_name,omitempty"`

	// current version
	CurrentVersion string `json:"current_version,omitempty"`

	// ha enabled
	HaEnabled bool `json:"ha_enabled,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// internal id
	InternalID string `json:"internal_id,omitempty"`

	// is sealed
	IsSealed bool `json:"is_sealed,omitempty"`

	// linked at
	// Format: date-time
	LinkedAt strfmt.DateTime `json:"linked_at,omitempty"`

	// location
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// node statuses
	NodeStatuses []*HashicorpCloudVaultLink20221107LinkedClusterNode `json:"node_statuses"`

	// raft quorum status
	RaftQuorumStatus *HashicorpCloudVaultLink20221107RaftQuorumStatus `json:"raft_quorum_status,omitempty"`

	// state
	State *HashicorpCloudVaultLink20221107LinkedClusterState `json:"state,omitempty"`

	// storage type
	StorageType string `json:"storage_type,omitempty"`
}

// Validate validates this hashicorp cloud vault link 20221107 linked cluster
func (m *HashicorpCloudVaultLink20221107LinkedCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinkedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaftQuorumStatus(formats); err != nil {
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

func (m *HashicorpCloudVaultLink20221107LinkedCluster) validateLinkedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LinkedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("linked_at", "body", "date-time", m.LinkedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedCluster) validateLocation(formats strfmt.Registry) error {
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

func (m *HashicorpCloudVaultLink20221107LinkedCluster) validateNodeStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.NodeStatuses); i++ {
		if swag.IsZero(m.NodeStatuses[i]) { // not required
			continue
		}

		if m.NodeStatuses[i] != nil {
			if err := m.NodeStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_statuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_statuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedCluster) validateRaftQuorumStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.RaftQuorumStatus) { // not required
		return nil
	}

	if m.RaftQuorumStatus != nil {
		if err := m.RaftQuorumStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raft_quorum_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("raft_quorum_status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedCluster) validateState(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud vault link 20221107 linked cluster based on the context it is used
func (m *HashicorpCloudVaultLink20221107LinkedCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRaftQuorumStatus(ctx, formats); err != nil {
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

func (m *HashicorpCloudVaultLink20221107LinkedCluster) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudVaultLink20221107LinkedCluster) contextValidateNodeStatuses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeStatuses); i++ {

		if m.NodeStatuses[i] != nil {
			if err := m.NodeStatuses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_statuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_statuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedCluster) contextValidateRaftQuorumStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.RaftQuorumStatus != nil {
		if err := m.RaftQuorumStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raft_quorum_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("raft_quorum_status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedCluster) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudVaultLink20221107LinkedCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVaultLink20221107LinkedCluster) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVaultLink20221107LinkedCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
