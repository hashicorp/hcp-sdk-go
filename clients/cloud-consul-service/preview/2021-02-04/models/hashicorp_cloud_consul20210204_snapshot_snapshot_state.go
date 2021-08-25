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

// HashicorpCloudConsul20210204SnapshotSnapshotState SnapshotState represents the different states a snapshot can be in.
//
//  - QUEUED: QUEUED is used for snapshots that haven’t started and are waiting for a cluster operation to finish before starting.
//  - CREATING: CREATING is used for snapshots that are creating.
//  - CREATING_FAILED: CREATING_FAILED is used for snapshots that failed creating.
//  - READY: READY is used for snapshots that are ready to be restored.
//  - DELETING: DELETING is used for snapshots that are deleting.
//  - DELETING_FAILED: DELETING_FAILED is used for snapshots that failed deleting.
//  - RESTORING: RESTORING is used for snapshots that are being restored.
//
// Progress can be tracked using cluster operations.
//
// swagger:model hashicorp.cloud.consul_20210204.Snapshot.SnapshotState
type HashicorpCloudConsul20210204SnapshotSnapshotState string

func NewHashicorpCloudConsul20210204SnapshotSnapshotState(value HashicorpCloudConsul20210204SnapshotSnapshotState) *HashicorpCloudConsul20210204SnapshotSnapshotState {
	v := value
	return &v
}

const (

	// HashicorpCloudConsul20210204SnapshotSnapshotStateSTATEUNSET captures enum value "STATE_UNSET"
	HashicorpCloudConsul20210204SnapshotSnapshotStateSTATEUNSET HashicorpCloudConsul20210204SnapshotSnapshotState = "STATE_UNSET"

	// HashicorpCloudConsul20210204SnapshotSnapshotStateQUEUED captures enum value "QUEUED"
	HashicorpCloudConsul20210204SnapshotSnapshotStateQUEUED HashicorpCloudConsul20210204SnapshotSnapshotState = "QUEUED"

	// HashicorpCloudConsul20210204SnapshotSnapshotStateCREATING captures enum value "CREATING"
	HashicorpCloudConsul20210204SnapshotSnapshotStateCREATING HashicorpCloudConsul20210204SnapshotSnapshotState = "CREATING"

	// HashicorpCloudConsul20210204SnapshotSnapshotStateCREATINGFAILED captures enum value "CREATING_FAILED"
	HashicorpCloudConsul20210204SnapshotSnapshotStateCREATINGFAILED HashicorpCloudConsul20210204SnapshotSnapshotState = "CREATING_FAILED"

	// HashicorpCloudConsul20210204SnapshotSnapshotStateREADY captures enum value "READY"
	HashicorpCloudConsul20210204SnapshotSnapshotStateREADY HashicorpCloudConsul20210204SnapshotSnapshotState = "READY"

	// HashicorpCloudConsul20210204SnapshotSnapshotStateDELETING captures enum value "DELETING"
	HashicorpCloudConsul20210204SnapshotSnapshotStateDELETING HashicorpCloudConsul20210204SnapshotSnapshotState = "DELETING"

	// HashicorpCloudConsul20210204SnapshotSnapshotStateDELETINGFAILED captures enum value "DELETING_FAILED"
	HashicorpCloudConsul20210204SnapshotSnapshotStateDELETINGFAILED HashicorpCloudConsul20210204SnapshotSnapshotState = "DELETING_FAILED"

	// HashicorpCloudConsul20210204SnapshotSnapshotStateRESTORING captures enum value "RESTORING"
	HashicorpCloudConsul20210204SnapshotSnapshotStateRESTORING HashicorpCloudConsul20210204SnapshotSnapshotState = "RESTORING"
)

// for schema
var hashicorpCloudConsul20210204SnapshotSnapshotStateEnum []interface{}

func init() {
	var res []HashicorpCloudConsul20210204SnapshotSnapshotState
	if err := json.Unmarshal([]byte(`["STATE_UNSET","QUEUED","CREATING","CREATING_FAILED","READY","DELETING","DELETING_FAILED","RESTORING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudConsul20210204SnapshotSnapshotStateEnum = append(hashicorpCloudConsul20210204SnapshotSnapshotStateEnum, v)
	}
}

func (m HashicorpCloudConsul20210204SnapshotSnapshotState) validateHashicorpCloudConsul20210204SnapshotSnapshotStateEnum(path, location string, value HashicorpCloudConsul20210204SnapshotSnapshotState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudConsul20210204SnapshotSnapshotStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud consul 20210204 snapshot snapshot state
func (m HashicorpCloudConsul20210204SnapshotSnapshotState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudConsul20210204SnapshotSnapshotStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud consul 20210204 snapshot snapshot state based on context it is used
func (m HashicorpCloudConsul20210204SnapshotSnapshotState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
