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

// HashicorpCloudWaypointV20241122JobStatus The various status we allow a job to be in.
//
//   - STATUS_UNSPECIFIED: default/zero value; we have no job status
//   - STATUS_UNKNOWN: we don't know the status of the job!
//   - STATUS_RUNNING: the job was marked as running
//   - STATUS_HALTED: the job was halted/failed to be launched
//   - STATUS_SUCCESS: the job was successful
//   - STATUS_ERRORED: the job ran, but errored out
//   - STATUS_QUEUED: the job is queued to run, used by agents
//
// swagger:model hashicorp.cloud.waypoint.v20241122.Job.Status
type HashicorpCloudWaypointV20241122JobStatus string

func NewHashicorpCloudWaypointV20241122JobStatus(value HashicorpCloudWaypointV20241122JobStatus) *HashicorpCloudWaypointV20241122JobStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudWaypointV20241122JobStatus.
func (m HashicorpCloudWaypointV20241122JobStatus) Pointer() *HashicorpCloudWaypointV20241122JobStatus {
	return &m
}

const (

	// HashicorpCloudWaypointV20241122JobStatusSTATUSUNSPECIFIED captures enum value "STATUS_UNSPECIFIED"
	HashicorpCloudWaypointV20241122JobStatusSTATUSUNSPECIFIED HashicorpCloudWaypointV20241122JobStatus = "STATUS_UNSPECIFIED"

	// HashicorpCloudWaypointV20241122JobStatusSTATUSUNKNOWN captures enum value "STATUS_UNKNOWN"
	HashicorpCloudWaypointV20241122JobStatusSTATUSUNKNOWN HashicorpCloudWaypointV20241122JobStatus = "STATUS_UNKNOWN"

	// HashicorpCloudWaypointV20241122JobStatusSTATUSRUNNING captures enum value "STATUS_RUNNING"
	HashicorpCloudWaypointV20241122JobStatusSTATUSRUNNING HashicorpCloudWaypointV20241122JobStatus = "STATUS_RUNNING"

	// HashicorpCloudWaypointV20241122JobStatusSTATUSHALTED captures enum value "STATUS_HALTED"
	HashicorpCloudWaypointV20241122JobStatusSTATUSHALTED HashicorpCloudWaypointV20241122JobStatus = "STATUS_HALTED"

	// HashicorpCloudWaypointV20241122JobStatusSTATUSSUCCESS captures enum value "STATUS_SUCCESS"
	HashicorpCloudWaypointV20241122JobStatusSTATUSSUCCESS HashicorpCloudWaypointV20241122JobStatus = "STATUS_SUCCESS"

	// HashicorpCloudWaypointV20241122JobStatusSTATUSERRORED captures enum value "STATUS_ERRORED"
	HashicorpCloudWaypointV20241122JobStatusSTATUSERRORED HashicorpCloudWaypointV20241122JobStatus = "STATUS_ERRORED"

	// HashicorpCloudWaypointV20241122JobStatusSTATUSQUEUED captures enum value "STATUS_QUEUED"
	HashicorpCloudWaypointV20241122JobStatusSTATUSQUEUED HashicorpCloudWaypointV20241122JobStatus = "STATUS_QUEUED"
)

// for schema
var hashicorpCloudWaypointV20241122JobStatusEnum []interface{}

func init() {
	var res []HashicorpCloudWaypointV20241122JobStatus
	if err := json.Unmarshal([]byte(`["STATUS_UNSPECIFIED","STATUS_UNKNOWN","STATUS_RUNNING","STATUS_HALTED","STATUS_SUCCESS","STATUS_ERRORED","STATUS_QUEUED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudWaypointV20241122JobStatusEnum = append(hashicorpCloudWaypointV20241122JobStatusEnum, v)
	}
}

func (m HashicorpCloudWaypointV20241122JobStatus) validateHashicorpCloudWaypointV20241122JobStatusEnum(path, location string, value HashicorpCloudWaypointV20241122JobStatus) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudWaypointV20241122JobStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud waypoint v20241122 job status
func (m HashicorpCloudWaypointV20241122JobStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudWaypointV20241122JobStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint v20241122 job status based on context it is used
func (m HashicorpCloudWaypointV20241122JobStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
