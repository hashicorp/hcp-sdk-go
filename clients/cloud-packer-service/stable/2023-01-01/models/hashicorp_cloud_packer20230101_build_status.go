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

// HashicorpCloudPacker20230101BuildStatus  - BUILD_UNSET: UNSET is a sentinel zero value so that an uninitialized value can be
// detected.
//   - BUILD_RUNNING: Running means the `packer build` is currently running.
//   - BUILD_DONE: Done means the `packer build` has finished successfully.
//   - BUILD_CANCELLED: Cancelled means the `packer build` was cancelled by a user.
//   - BUILD_FAILED: Failed means the `packer build` failed and therefore artifact creation failed.
//
// swagger:model hashicorp.cloud.packer_20230101.BuildStatus
type HashicorpCloudPacker20230101BuildStatus string

func NewHashicorpCloudPacker20230101BuildStatus(value HashicorpCloudPacker20230101BuildStatus) *HashicorpCloudPacker20230101BuildStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudPacker20230101BuildStatus.
func (m HashicorpCloudPacker20230101BuildStatus) Pointer() *HashicorpCloudPacker20230101BuildStatus {
	return &m
}

const (

	// HashicorpCloudPacker20230101BuildStatusBUILDUNSET captures enum value "BUILD_UNSET"
	HashicorpCloudPacker20230101BuildStatusBUILDUNSET HashicorpCloudPacker20230101BuildStatus = "BUILD_UNSET"

	// HashicorpCloudPacker20230101BuildStatusBUILDRUNNING captures enum value "BUILD_RUNNING"
	HashicorpCloudPacker20230101BuildStatusBUILDRUNNING HashicorpCloudPacker20230101BuildStatus = "BUILD_RUNNING"

	// HashicorpCloudPacker20230101BuildStatusBUILDDONE captures enum value "BUILD_DONE"
	HashicorpCloudPacker20230101BuildStatusBUILDDONE HashicorpCloudPacker20230101BuildStatus = "BUILD_DONE"

	// HashicorpCloudPacker20230101BuildStatusBUILDCANCELLED captures enum value "BUILD_CANCELLED"
	HashicorpCloudPacker20230101BuildStatusBUILDCANCELLED HashicorpCloudPacker20230101BuildStatus = "BUILD_CANCELLED"

	// HashicorpCloudPacker20230101BuildStatusBUILDFAILED captures enum value "BUILD_FAILED"
	HashicorpCloudPacker20230101BuildStatusBUILDFAILED HashicorpCloudPacker20230101BuildStatus = "BUILD_FAILED"
)

// for schema
var hashicorpCloudPacker20230101BuildStatusEnum []interface{}

func init() {
	var res []HashicorpCloudPacker20230101BuildStatus
	if err := json.Unmarshal([]byte(`["BUILD_UNSET","BUILD_RUNNING","BUILD_DONE","BUILD_CANCELLED","BUILD_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudPacker20230101BuildStatusEnum = append(hashicorpCloudPacker20230101BuildStatusEnum, v)
	}
}

func (m HashicorpCloudPacker20230101BuildStatus) validateHashicorpCloudPacker20230101BuildStatusEnum(path, location string, value HashicorpCloudPacker20230101BuildStatus) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudPacker20230101BuildStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud packer 20230101 build status
func (m HashicorpCloudPacker20230101BuildStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudPacker20230101BuildStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud packer 20230101 build status based on context it is used
func (m HashicorpCloudPacker20230101BuildStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
