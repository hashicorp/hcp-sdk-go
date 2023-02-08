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

// VaultLink20221107LinkedClusterNode vault link 20221107 linked cluster node
//
// swagger:model vault_link_20221107LinkedClusterNode
type VaultLink20221107LinkedClusterNode struct {

	// alternative_versions is a list of versions that should also be considered for
	// an update as they might come with additional improvements and features.
	AlternativeVersions []string `json:"alternativeVersions"`

	// current report
	// Format: date-time
	CurrentReport strfmt.DateTime `json:"currentReport,omitempty"`

	// current_version is the node's current version in semantic version format.
	CurrentVersion string `json:"currentVersion,omitempty"`

	// has_security_flaw will be true if the current version has a security flaw.
	HasSecurityFlaws bool `json:"hasSecurityFlaws,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// internal Id
	InternalID string `json:"internalId,omitempty"`

	// last reported at
	// Format: date-time
	LastReportedAt strfmt.DateTime `json:"lastReportedAt,omitempty"`

	// leader status
	LeaderStatus *LinkedClusterNodeLeaderStatus `json:"leaderStatus,omitempty"`

	// linked cluster internal Id
	LinkedClusterInternalID string `json:"linkedClusterInternalId,omitempty"`

	// listener addresses
	ListenerAddresses []string `json:"listenerAddresses"`

	// log level
	LogLevel *LinkedClusterNodeLogLevel `json:"logLevel,omitempty"`

	// node_binary_architecture is the lower-case architecture of the client binary
	// (e.g. amd64, arm, ...).
	NodeBinaryArchitecture string `json:"nodeBinaryArchitecture,omitempty"`

	// node_id is the node identification.
	NodeID string `json:"nodeId,omitempty"`

	// node_initialized indicates if the node has been initialized.
	NodeInitialized bool `json:"nodeInitialized,omitempty"`

	// node_os is the lower-case name of the operating system the client is
	// running on (e.g. linux, windows).
	NodeOs string `json:"nodeOs,omitempty"`

	// node_sealed indicates if the node is sealed.
	NodeSealed bool `json:"nodeSealed,omitempty"`

	// node state
	NodeState *VaultLink20221107LinkedClusterNodeState `json:"nodeState,omitempty"`

	// node_type indicates the node type.
	NodeType string `json:"nodeType,omitempty"`

	// recommended_version is the version the product should ideally be updated to.
	RecommendedVersion string `json:"recommendedVersion,omitempty"`

	// status is the status of the current version. The status may help to
	// determine the urgency of the update.
	Status *LinkedClusterNodeVersionStatus `json:"status,omitempty"`

	// status_version is the version of the status message format. To ensure
	// that the version is not omitted by accident the initial version is 1.
	StatusVersion int64 `json:"statusVersion,omitempty"`

	// storage type
	StorageType string `json:"storageType,omitempty"`
}

// Validate validates this vault link 20221107 linked cluster node
func (m *VaultLink20221107LinkedClusterNode) Validate(formats strfmt.Registry) error {
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

func (m *VaultLink20221107LinkedClusterNode) validateCurrentReport(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentReport) { // not required
		return nil
	}

	if err := validate.FormatOf("currentReport", "body", "date-time", m.CurrentReport.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VaultLink20221107LinkedClusterNode) validateLastReportedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastReportedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastReportedAt", "body", "date-time", m.LastReportedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VaultLink20221107LinkedClusterNode) validateLeaderStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.LeaderStatus) { // not required
		return nil
	}

	if m.LeaderStatus != nil {
		if err := m.LeaderStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("leaderStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("leaderStatus")
			}
			return err
		}
	}

	return nil
}

func (m *VaultLink20221107LinkedClusterNode) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.LogLevel) { // not required
		return nil
	}

	if m.LogLevel != nil {
		if err := m.LogLevel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logLevel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logLevel")
			}
			return err
		}
	}

	return nil
}

func (m *VaultLink20221107LinkedClusterNode) validateNodeState(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeState) { // not required
		return nil
	}

	if m.NodeState != nil {
		if err := m.NodeState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeState")
			}
			return err
		}
	}

	return nil
}

func (m *VaultLink20221107LinkedClusterNode) validateStatus(formats strfmt.Registry) error {
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

// ContextValidate validate this vault link 20221107 linked cluster node based on the context it is used
func (m *VaultLink20221107LinkedClusterNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *VaultLink20221107LinkedClusterNode) contextValidateLeaderStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.LeaderStatus != nil {
		if err := m.LeaderStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("leaderStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("leaderStatus")
			}
			return err
		}
	}

	return nil
}

func (m *VaultLink20221107LinkedClusterNode) contextValidateLogLevel(ctx context.Context, formats strfmt.Registry) error {

	if m.LogLevel != nil {
		if err := m.LogLevel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logLevel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logLevel")
			}
			return err
		}
	}

	return nil
}

func (m *VaultLink20221107LinkedClusterNode) contextValidateNodeState(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeState != nil {
		if err := m.NodeState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeState")
			}
			return err
		}
	}

	return nil
}

func (m *VaultLink20221107LinkedClusterNode) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VaultLink20221107LinkedClusterNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VaultLink20221107LinkedClusterNode) UnmarshalBinary(b []byte) error {
	var res VaultLink20221107LinkedClusterNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
