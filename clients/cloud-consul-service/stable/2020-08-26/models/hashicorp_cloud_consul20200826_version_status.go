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

// HashicorpCloudConsul20200826VersionStatus  - AVAILABLE: AVAILABLE represents a version that is generally available
// but no longer the preferred/recommended version.
//   - RECOMMENDED: RECOMMENDED represents a version that is generally available
//
// and recommended by HashiCorp.
//   - PREVIEW: PREVIEW represents a version that is not generally available.
//
// swagger:model hashicorp.cloud.consul_20200826.Version.Status
type HashicorpCloudConsul20200826VersionStatus string

func NewHashicorpCloudConsul20200826VersionStatus(value HashicorpCloudConsul20200826VersionStatus) *HashicorpCloudConsul20200826VersionStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudConsul20200826VersionStatus.
func (m HashicorpCloudConsul20200826VersionStatus) Pointer() *HashicorpCloudConsul20200826VersionStatus {
	return &m
}

const (

	// HashicorpCloudConsul20200826VersionStatusAVAILABLE captures enum value "AVAILABLE"
	HashicorpCloudConsul20200826VersionStatusAVAILABLE HashicorpCloudConsul20200826VersionStatus = "AVAILABLE"

	// HashicorpCloudConsul20200826VersionStatusRECOMMENDED captures enum value "RECOMMENDED"
	HashicorpCloudConsul20200826VersionStatusRECOMMENDED HashicorpCloudConsul20200826VersionStatus = "RECOMMENDED"

	// HashicorpCloudConsul20200826VersionStatusPREVIEW captures enum value "PREVIEW"
	HashicorpCloudConsul20200826VersionStatusPREVIEW HashicorpCloudConsul20200826VersionStatus = "PREVIEW"
)

// for schema
var hashicorpCloudConsul20200826VersionStatusEnum []interface{}

func init() {
	var res []HashicorpCloudConsul20200826VersionStatus
	if err := json.Unmarshal([]byte(`["AVAILABLE","RECOMMENDED","PREVIEW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudConsul20200826VersionStatusEnum = append(hashicorpCloudConsul20200826VersionStatusEnum, v)
	}
}

func (m HashicorpCloudConsul20200826VersionStatus) validateHashicorpCloudConsul20200826VersionStatusEnum(path, location string, value HashicorpCloudConsul20200826VersionStatus) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudConsul20200826VersionStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud consul 20200826 version status
func (m HashicorpCloudConsul20200826VersionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudConsul20200826VersionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud consul 20200826 version status based on context it is used
func (m HashicorpCloudConsul20200826VersionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
