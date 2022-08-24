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

// HashicorpCloudConsul20200826SnapshotSnapshotType SnapshotType represents the type of a snapshot.
//
//  - AUTOMATIC: AUTOMATIC is used for snapshots generated automatically, like a snapshot before update.
//  - MANUAL: MANUAL is used for snapshots that are initiated by a user's actions.
//  - SCHEDULED: SCHEDULED is used for snapshots that are generated based on a schedule.
//
// swagger:model hashicorp.cloud.consul_20200826.Snapshot.SnapshotType
type HashicorpCloudConsul20200826SnapshotSnapshotType string

func NewHashicorpCloudConsul20200826SnapshotSnapshotType(value HashicorpCloudConsul20200826SnapshotSnapshotType) *HashicorpCloudConsul20200826SnapshotSnapshotType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudConsul20200826SnapshotSnapshotType.
func (m HashicorpCloudConsul20200826SnapshotSnapshotType) Pointer() *HashicorpCloudConsul20200826SnapshotSnapshotType {
	return &m
}

const (

	// HashicorpCloudConsul20200826SnapshotSnapshotTypeTYPEUNSET captures enum value "TYPE_UNSET"
	HashicorpCloudConsul20200826SnapshotSnapshotTypeTYPEUNSET HashicorpCloudConsul20200826SnapshotSnapshotType = "TYPE_UNSET"

	// HashicorpCloudConsul20200826SnapshotSnapshotTypeAUTOMATIC captures enum value "AUTOMATIC"
	HashicorpCloudConsul20200826SnapshotSnapshotTypeAUTOMATIC HashicorpCloudConsul20200826SnapshotSnapshotType = "AUTOMATIC"

	// HashicorpCloudConsul20200826SnapshotSnapshotTypeMANUAL captures enum value "MANUAL"
	HashicorpCloudConsul20200826SnapshotSnapshotTypeMANUAL HashicorpCloudConsul20200826SnapshotSnapshotType = "MANUAL"

	// HashicorpCloudConsul20200826SnapshotSnapshotTypeSCHEDULED captures enum value "SCHEDULED"
	HashicorpCloudConsul20200826SnapshotSnapshotTypeSCHEDULED HashicorpCloudConsul20200826SnapshotSnapshotType = "SCHEDULED"
)

// for schema
var hashicorpCloudConsul20200826SnapshotSnapshotTypeEnum []interface{}

func init() {
	var res []HashicorpCloudConsul20200826SnapshotSnapshotType
	if err := json.Unmarshal([]byte(`["TYPE_UNSET","AUTOMATIC","MANUAL","SCHEDULED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudConsul20200826SnapshotSnapshotTypeEnum = append(hashicorpCloudConsul20200826SnapshotSnapshotTypeEnum, v)
	}
}

func (m HashicorpCloudConsul20200826SnapshotSnapshotType) validateHashicorpCloudConsul20200826SnapshotSnapshotTypeEnum(path, location string, value HashicorpCloudConsul20200826SnapshotSnapshotType) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudConsul20200826SnapshotSnapshotTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud consul 20200826 snapshot snapshot type
func (m HashicorpCloudConsul20200826SnapshotSnapshotType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudConsul20200826SnapshotSnapshotTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud consul 20200826 snapshot snapshot type based on context it is used
func (m HashicorpCloudConsul20200826SnapshotSnapshotType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
