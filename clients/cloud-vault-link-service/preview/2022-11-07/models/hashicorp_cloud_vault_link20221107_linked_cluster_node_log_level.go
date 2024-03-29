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

// HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel hashicorp cloud vault link 20221107 linked cluster node log level
//
// swagger:model hashicorp.cloud.vault_link_20221107.LinkedClusterNode.LogLevel
type HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel string

func NewHashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel(value HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel) *HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel.
func (m HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel) Pointer() *HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel {
	return &m
}

const (

	// HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelLOGLEVELCLUSTERSTATEINVALID captures enum value "LOG_LEVEL_CLUSTER_STATE_INVALID"
	HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelLOGLEVELCLUSTERSTATEINVALID HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel = "LOG_LEVEL_CLUSTER_STATE_INVALID"

	// HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelTRACE captures enum value "TRACE"
	HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelTRACE HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel = "TRACE"

	// HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelDEBUG captures enum value "DEBUG"
	HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelDEBUG HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel = "DEBUG"

	// HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelINFO captures enum value "INFO"
	HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelINFO HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel = "INFO"

	// HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelWARN captures enum value "WARN"
	HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelWARN HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel = "WARN"

	// HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelERROR captures enum value "ERROR"
	HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelERROR HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel = "ERROR"
)

// for schema
var hashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelEnum []interface{}

func init() {
	var res []HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel
	if err := json.Unmarshal([]byte(`["LOG_LEVEL_CLUSTER_STATE_INVALID","TRACE","DEBUG","INFO","WARN","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelEnum = append(hashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelEnum, v)
	}
}

func (m HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel) validateHashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelEnum(path, location string, value HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vault link 20221107 linked cluster node log level
func (m HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVaultLink20221107LinkedClusterNodeLogLevelEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud vault link 20221107 linked cluster node log level based on context it is used
func (m HashicorpCloudVaultLink20221107LinkedClusterNodeLogLevel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
