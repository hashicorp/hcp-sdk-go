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

// HashicorpCloudVaultLink20221107LinkedClusterNode hashicorp cloud vault link 20221107 linked cluster node
//
// swagger:model hashicorp.cloud.vault_link_20221107.LinkedClusterNode
type HashicorpCloudVaultLink20221107LinkedClusterNode struct {

	// alternative_versions is a list of versions that should also be considered for
	// an update as they might come with additional improvements and features.
	AlternativeVersions []string `json:"alternative_versions"`

	// current report
	// Format: date-time
	CurrentReport strfmt.DateTime `json:"current_report,omitempty"`

	// current_version is the node's current version in semantic version format.
	CurrentVersion string `json:"current_version,omitempty"`

	// has_security_flaw will be true if the current version has a security flaw.
	HasSecurityFlaws bool `json:"has_security_flaws,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// internal id
	InternalID string `json:"internal_id,omitempty"`

	// last reported at
	// Format: date-time
	LastReportedAt strfmt.DateTime `json:"last_reported_at,omitempty"`

	// leader status
	LeaderStatus *HashicorpCloudVaultLink20221107LinkedClusterNodeLeaderStatus `json:"leader_status,omitempty"`

	// linked cluster internal id
	LinkedClusterInternalID string `json:"linked_cluster_internal_id,omitempty"`

	// listener addresses
	ListenerAddresses []string `json:"listener_addresses"`

	// log level
	LogLevel *HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel `json:"log_level,omitempty"`

	// node_binary_architecture is the lower-case architecture of the client binary
	// (e.g. amd64, arm, ...).
	NodeBinaryArchitecture string `json:"node_binary_architecture,omitempty"`

	// node_id is the node identification.
	NodeID string `json:"node_id,omitempty"`

	// node_initialized indicates if the node has been initialized.
	NodeInitialized bool `json:"node_initialized,omitempty"`

	// node_os is the lower-case name of the operating system the client is
	// running on (e.g. linux, windows).
	NodeOs string `json:"node_os,omitempty"`

	// node_sealed indicates if the node is sealed.
	NodeSealed bool `json:"node_sealed,omitempty"`

	// node state
	NodeState *HashicorpCloudVaultLink20221107LinkedClusterNodeState `json:"node_state,omitempty"`

	// node_type indicates the node type.
	NodeType string `json:"node_type,omitempty"`

	// recommended_version is the version the product should ideally be updated to.
	RecommendedVersion string `json:"recommended_version,omitempty"`

	// status is the status of the current version. The status may help to
	// determine the urgency of the update.
	Status *HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus `json:"status,omitempty"`

	// status_version is the version of the status message format. To ensure
	// that the version is not omitted by accident the initial version is 1.
	StatusVersion int64 `json:"status_version,omitempty"`

	// storage type
	StorageType string `json:"storage_type,omitempty"`
}

// Validate validates this hashicorp cloud vault link 20221107 linked cluster node
func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentReport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastReportedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeaderStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) validateCurrentReport(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentReport) { // not required
		return nil
	}

	if err := validate.FormatOf("current_report", "body", "date-time", m.CurrentReport.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) validateLastReportedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastReportedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("last_reported_at", "body", "date-time", m.LastReportedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) validateLeaderStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.LeaderStatus) { // not required
		return nil
	}

	if m.LeaderStatus != nil {
		if err := m.LeaderStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("leader_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("leader_status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.LogLevel) { // not required
		return nil
	}

	if m.LogLevel != nil {
		if err := m.LogLevel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_level")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_level")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) validateNodeState(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeState) { // not required
		return nil
	}

	if m.NodeState != nil {
		if err := m.NodeState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_state")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault link 20221107 linked cluster node based on the context it is used
func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLeaderStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) contextValidateLeaderStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.LeaderStatus != nil {
		if err := m.LeaderStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("leader_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("leader_status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) contextValidateLogLevel(ctx context.Context, formats strfmt.Registry) error {

	if m.LogLevel != nil {
		if err := m.LogLevel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_level")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_level")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) contextValidateNodeState(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeState != nil {
		if err := m.NodeState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_state")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVaultLink20221107LinkedClusterNode) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVaultLink20221107LinkedClusterNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
