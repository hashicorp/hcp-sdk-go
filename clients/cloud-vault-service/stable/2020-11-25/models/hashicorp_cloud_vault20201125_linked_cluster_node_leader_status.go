// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus - LEADER: Deprecated values
//   - ACTIVE: Valid values
//
// swagger:model hashicorp.cloud.vault_20201125.LinkedClusterNode.LeaderStatus
type HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus string

func NewHashicorpCloudVault20201125LinkedClusterNodeLeaderStatus(value HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus) *HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus.
func (m HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus) Pointer() *HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus {
	return &m
}

const (

	// HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusLEADERSTATUSCLUSTERSTATEINVALID captures enum value "LEADER_STATUS_CLUSTER_STATE_INVALID"
	HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusLEADERSTATUSCLUSTERSTATEINVALID HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus = "LEADER_STATUS_CLUSTER_STATE_INVALID"

	// HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusLEADER captures enum value "LEADER"
	HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusLEADER HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus = "LEADER"

	// HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusFOLLOWER captures enum value "FOLLOWER"
	HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusFOLLOWER HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus = "FOLLOWER"

	// HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusACTIVE captures enum value "ACTIVE"
	HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusACTIVE HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus = "ACTIVE"

	// HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusSTANDBY captures enum value "STANDBY"
	HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusSTANDBY HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus = "STANDBY"

	// HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusPERFSTANDBY captures enum value "PERF_STANDBY"
	HashicorpCloudVault20201125LinkedClusterNodeLeaderStatusPERFSTANDBY HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus = "PERF_STANDBY"
)

// for schema
var hashicorpCloudVault20201125LinkedClusterNodeLeaderStatusEnum []interface{}

func init() {
	var res []HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus
	if err := json.Unmarshal([]byte(`["LEADER_STATUS_CLUSTER_STATE_INVALID","LEADER","FOLLOWER","ACTIVE","STANDBY","PERF_STANDBY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVault20201125LinkedClusterNodeLeaderStatusEnum = append(hashicorpCloudVault20201125LinkedClusterNodeLeaderStatusEnum, v)
	}
}

func (m HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus) validateHashicorpCloudVault20201125LinkedClusterNodeLeaderStatusEnum(path, location string, value HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVault20201125LinkedClusterNodeLeaderStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vault 20201125 linked cluster node leader status
func (m HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVault20201125LinkedClusterNodeLeaderStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20201125 linked cluster node leader status based on context it is used
func (m HashicorpCloudVault20201125LinkedClusterNodeLeaderStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
