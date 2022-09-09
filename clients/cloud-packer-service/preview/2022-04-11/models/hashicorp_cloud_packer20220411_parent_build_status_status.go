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

// HashicorpCloudPacker20220411ParentBuildStatusStatus  - UNDETERMINED: Status cannot be determined because either a channel wasn't used when
// building the child build, the configured channel no longer exists
// in the parent bucket, or the parent is not tracked by HCP Packer.
//   - UP_TO_DATE: The child build is built from the iteration that is currently
//
// assigned to the channel that is used in its configuration.
//   - OUT_OF_DATE: The child build is built from an iteration that is different from
//
// the one currently assigned to the channel
// that is used in its configuration.
//
// swagger:model hashicorp.cloud.packer_20220411.ParentBuildStatus.Status
type HashicorpCloudPacker20220411ParentBuildStatusStatus string

func NewHashicorpCloudPacker20220411ParentBuildStatusStatus(value HashicorpCloudPacker20220411ParentBuildStatusStatus) *HashicorpCloudPacker20220411ParentBuildStatusStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudPacker20220411ParentBuildStatusStatus.
func (m HashicorpCloudPacker20220411ParentBuildStatusStatus) Pointer() *HashicorpCloudPacker20220411ParentBuildStatusStatus {
	return &m
}

const (

	// HashicorpCloudPacker20220411ParentBuildStatusStatusUNDETERMINED captures enum value "UNDETERMINED"
	HashicorpCloudPacker20220411ParentBuildStatusStatusUNDETERMINED HashicorpCloudPacker20220411ParentBuildStatusStatus = "UNDETERMINED"

	// HashicorpCloudPacker20220411ParentBuildStatusStatusUPTODATE captures enum value "UP_TO_DATE"
	HashicorpCloudPacker20220411ParentBuildStatusStatusUPTODATE HashicorpCloudPacker20220411ParentBuildStatusStatus = "UP_TO_DATE"

	// HashicorpCloudPacker20220411ParentBuildStatusStatusOUTOFDATE captures enum value "OUT_OF_DATE"
	HashicorpCloudPacker20220411ParentBuildStatusStatusOUTOFDATE HashicorpCloudPacker20220411ParentBuildStatusStatus = "OUT_OF_DATE"
)

// for schema
var hashicorpCloudPacker20220411ParentBuildStatusStatusEnum []interface{}

func init() {
	var res []HashicorpCloudPacker20220411ParentBuildStatusStatus
	if err := json.Unmarshal([]byte(`["UNDETERMINED","UP_TO_DATE","OUT_OF_DATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudPacker20220411ParentBuildStatusStatusEnum = append(hashicorpCloudPacker20220411ParentBuildStatusStatusEnum, v)
	}
}

func (m HashicorpCloudPacker20220411ParentBuildStatusStatus) validateHashicorpCloudPacker20220411ParentBuildStatusStatusEnum(path, location string, value HashicorpCloudPacker20220411ParentBuildStatusStatus) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudPacker20220411ParentBuildStatusStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud packer 20220411 parent build status status
func (m HashicorpCloudPacker20220411ParentBuildStatusStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudPacker20220411ParentBuildStatusStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud packer 20220411 parent build status status based on context it is used
func (m HashicorpCloudPacker20220411ParentBuildStatusStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
