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

// HashicorpCloudWaypointV20241122ActionRunResponseStatus This is the status of the request that the action ran.
//
//   - NONE: default/zero value; we have no status yet
//   - UNKNOWN: a status we didn't understand
//   - SUCCESS: action submitted successfully
//   - ERROR: it did not work
//
// swagger:model hashicorp.cloud.waypoint.v20241122.ActionRun.ResponseStatus
type HashicorpCloudWaypointV20241122ActionRunResponseStatus string

func NewHashicorpCloudWaypointV20241122ActionRunResponseStatus(value HashicorpCloudWaypointV20241122ActionRunResponseStatus) *HashicorpCloudWaypointV20241122ActionRunResponseStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudWaypointV20241122ActionRunResponseStatus.
func (m HashicorpCloudWaypointV20241122ActionRunResponseStatus) Pointer() *HashicorpCloudWaypointV20241122ActionRunResponseStatus {
	return &m
}

const (

	// HashicorpCloudWaypointV20241122ActionRunResponseStatusNONE captures enum value "NONE"
	HashicorpCloudWaypointV20241122ActionRunResponseStatusNONE HashicorpCloudWaypointV20241122ActionRunResponseStatus = "NONE"

	// HashicorpCloudWaypointV20241122ActionRunResponseStatusUNKNOWN captures enum value "UNKNOWN"
	HashicorpCloudWaypointV20241122ActionRunResponseStatusUNKNOWN HashicorpCloudWaypointV20241122ActionRunResponseStatus = "UNKNOWN"

	// HashicorpCloudWaypointV20241122ActionRunResponseStatusSUCCESS captures enum value "SUCCESS"
	HashicorpCloudWaypointV20241122ActionRunResponseStatusSUCCESS HashicorpCloudWaypointV20241122ActionRunResponseStatus = "SUCCESS"

	// HashicorpCloudWaypointV20241122ActionRunResponseStatusERROR captures enum value "ERROR"
	HashicorpCloudWaypointV20241122ActionRunResponseStatusERROR HashicorpCloudWaypointV20241122ActionRunResponseStatus = "ERROR"
)

// for schema
var hashicorpCloudWaypointV20241122ActionRunResponseStatusEnum []interface{}

func init() {
	var res []HashicorpCloudWaypointV20241122ActionRunResponseStatus
	if err := json.Unmarshal([]byte(`["NONE","UNKNOWN","SUCCESS","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudWaypointV20241122ActionRunResponseStatusEnum = append(hashicorpCloudWaypointV20241122ActionRunResponseStatusEnum, v)
	}
}

func (m HashicorpCloudWaypointV20241122ActionRunResponseStatus) validateHashicorpCloudWaypointV20241122ActionRunResponseStatusEnum(path, location string, value HashicorpCloudWaypointV20241122ActionRunResponseStatus) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudWaypointV20241122ActionRunResponseStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud waypoint v20241122 action run response status
func (m HashicorpCloudWaypointV20241122ActionRunResponseStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudWaypointV20241122ActionRunResponseStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint v20241122 action run response status based on context it is used
func (m HashicorpCloudWaypointV20241122ActionRunResponseStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
