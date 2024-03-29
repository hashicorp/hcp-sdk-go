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

// HashicorpCloudConsul20200826SnapshotSnapshotState SnapshotState represents the different states a snapshot can be in.
//
//   - QUEUED: QUEUED is used for snapshots that haven’t started and are waiting for a cluster operation to finish before starting.
//   - CREATING: CREATING is used for snapshots that are creating.
//   - CREATING_FAILED: CREATING_FAILED is used for snapshots that failed creating.
//   - READY: READY is used for snapshots that are ready to be restored.
//   - DELETING: DELETING is used for snapshots that are deleting.
//   - DELETING_FAILED: DELETING_FAILED is used for snapshots that failed deleting.
//   - RESTORING: RESTORING is used for snapshots that are being restored.
//
// Progress can be tracked using cluster operations.
//
// swagger:model hashicorp.cloud.consul_20200826.Snapshot.SnapshotState
type HashicorpCloudConsul20200826SnapshotSnapshotState string

func NewHashicorpCloudConsul20200826SnapshotSnapshotState(value HashicorpCloudConsul20200826SnapshotSnapshotState) *HashicorpCloudConsul20200826SnapshotSnapshotState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudConsul20200826SnapshotSnapshotState.
func (m HashicorpCloudConsul20200826SnapshotSnapshotState) Pointer() *HashicorpCloudConsul20200826SnapshotSnapshotState {
	return &m
}

const (

	// HashicorpCloudConsul20200826SnapshotSnapshotStateSTATEUNSET captures enum value "STATE_UNSET"
	HashicorpCloudConsul20200826SnapshotSnapshotStateSTATEUNSET HashicorpCloudConsul20200826SnapshotSnapshotState = "STATE_UNSET"

	// HashicorpCloudConsul20200826SnapshotSnapshotStateQUEUED captures enum value "QUEUED"
	HashicorpCloudConsul20200826SnapshotSnapshotStateQUEUED HashicorpCloudConsul20200826SnapshotSnapshotState = "QUEUED"

	// HashicorpCloudConsul20200826SnapshotSnapshotStateCREATING captures enum value "CREATING"
	HashicorpCloudConsul20200826SnapshotSnapshotStateCREATING HashicorpCloudConsul20200826SnapshotSnapshotState = "CREATING"

	// HashicorpCloudConsul20200826SnapshotSnapshotStateCREATINGFAILED captures enum value "CREATING_FAILED"
	HashicorpCloudConsul20200826SnapshotSnapshotStateCREATINGFAILED HashicorpCloudConsul20200826SnapshotSnapshotState = "CREATING_FAILED"

	// HashicorpCloudConsul20200826SnapshotSnapshotStateREADY captures enum value "READY"
	HashicorpCloudConsul20200826SnapshotSnapshotStateREADY HashicorpCloudConsul20200826SnapshotSnapshotState = "READY"

	// HashicorpCloudConsul20200826SnapshotSnapshotStateDELETING captures enum value "DELETING"
	HashicorpCloudConsul20200826SnapshotSnapshotStateDELETING HashicorpCloudConsul20200826SnapshotSnapshotState = "DELETING"

	// HashicorpCloudConsul20200826SnapshotSnapshotStateDELETINGFAILED captures enum value "DELETING_FAILED"
	HashicorpCloudConsul20200826SnapshotSnapshotStateDELETINGFAILED HashicorpCloudConsul20200826SnapshotSnapshotState = "DELETING_FAILED"

	// HashicorpCloudConsul20200826SnapshotSnapshotStateRESTORING captures enum value "RESTORING"
	HashicorpCloudConsul20200826SnapshotSnapshotStateRESTORING HashicorpCloudConsul20200826SnapshotSnapshotState = "RESTORING"
)

// for schema
var hashicorpCloudConsul20200826SnapshotSnapshotStateEnum []interface{}

func init() {
	var res []HashicorpCloudConsul20200826SnapshotSnapshotState
	if err := json.Unmarshal([]byte(`["STATE_UNSET","QUEUED","CREATING","CREATING_FAILED","READY","DELETING","DELETING_FAILED","RESTORING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudConsul20200826SnapshotSnapshotStateEnum = append(hashicorpCloudConsul20200826SnapshotSnapshotStateEnum, v)
	}
}

func (m HashicorpCloudConsul20200826SnapshotSnapshotState) validateHashicorpCloudConsul20200826SnapshotSnapshotStateEnum(path, location string, value HashicorpCloudConsul20200826SnapshotSnapshotState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudConsul20200826SnapshotSnapshotStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud consul 20200826 snapshot snapshot state
func (m HashicorpCloudConsul20200826SnapshotSnapshotState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudConsul20200826SnapshotSnapshotStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud consul 20200826 snapshot snapshot state based on context it is used
func (m HashicorpCloudConsul20200826SnapshotSnapshotState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
