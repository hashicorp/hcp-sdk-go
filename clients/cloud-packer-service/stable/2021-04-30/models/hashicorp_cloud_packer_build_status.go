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

// HashicorpCloudPackerBuildStatus  - UNSET: UNSET is a sentinel zero value so that an uninitialized value can be
// detected.
//   - RUNNING: Running means the `packer build` is currently running.
//   - DONE: Done means the `packer build` has finished successfully.
//   - CANCELLED: Cancelled means the `packer build` was cancelled by a user.
//   - FAILED: Failed means the `packer build` failed and therefore image creation failed.
//
// swagger:model hashicorp.cloud.packer.BuildStatus
type HashicorpCloudPackerBuildStatus string

func NewHashicorpCloudPackerBuildStatus(value HashicorpCloudPackerBuildStatus) *HashicorpCloudPackerBuildStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudPackerBuildStatus.
func (m HashicorpCloudPackerBuildStatus) Pointer() *HashicorpCloudPackerBuildStatus {
	return &m
}

const (

	// HashicorpCloudPackerBuildStatusUNSET captures enum value "UNSET"
	HashicorpCloudPackerBuildStatusUNSET HashicorpCloudPackerBuildStatus = "UNSET"

	// HashicorpCloudPackerBuildStatusRUNNING captures enum value "RUNNING"
	HashicorpCloudPackerBuildStatusRUNNING HashicorpCloudPackerBuildStatus = "RUNNING"

	// HashicorpCloudPackerBuildStatusDONE captures enum value "DONE"
	HashicorpCloudPackerBuildStatusDONE HashicorpCloudPackerBuildStatus = "DONE"

	// HashicorpCloudPackerBuildStatusCANCELLED captures enum value "CANCELLED"
	HashicorpCloudPackerBuildStatusCANCELLED HashicorpCloudPackerBuildStatus = "CANCELLED"

	// HashicorpCloudPackerBuildStatusFAILED captures enum value "FAILED"
	HashicorpCloudPackerBuildStatusFAILED HashicorpCloudPackerBuildStatus = "FAILED"
)

// for schema
var hashicorpCloudPackerBuildStatusEnum []interface{}

func init() {
	var res []HashicorpCloudPackerBuildStatus
	if err := json.Unmarshal([]byte(`["UNSET","RUNNING","DONE","CANCELLED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudPackerBuildStatusEnum = append(hashicorpCloudPackerBuildStatusEnum, v)
	}
}

func (m HashicorpCloudPackerBuildStatus) validateHashicorpCloudPackerBuildStatusEnum(path, location string, value HashicorpCloudPackerBuildStatus) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudPackerBuildStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud packer build status
func (m HashicorpCloudPackerBuildStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudPackerBuildStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud packer build status based on context it is used
func (m HashicorpCloudPackerBuildStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
