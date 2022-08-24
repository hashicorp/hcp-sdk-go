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

// HashicorpCloudConsul20210204PlatformType PlatformType refers to the type of platform that the cluster was created with ex. HCP or HCS
// It determines which version list is returned. NOTSET will default to the HCP list.
//
// swagger:model hashicorp.cloud.consul_20210204.PlatformType
type HashicorpCloudConsul20210204PlatformType string

func NewHashicorpCloudConsul20210204PlatformType(value HashicorpCloudConsul20210204PlatformType) *HashicorpCloudConsul20210204PlatformType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudConsul20210204PlatformType.
func (m HashicorpCloudConsul20210204PlatformType) Pointer() *HashicorpCloudConsul20210204PlatformType {
	return &m
}

const (

	// HashicorpCloudConsul20210204PlatformTypeNOTSET captures enum value "NOTSET"
	HashicorpCloudConsul20210204PlatformTypeNOTSET HashicorpCloudConsul20210204PlatformType = "NOTSET"

	// HashicorpCloudConsul20210204PlatformTypeHCP captures enum value "HCP"
	HashicorpCloudConsul20210204PlatformTypeHCP HashicorpCloudConsul20210204PlatformType = "HCP"

	// HashicorpCloudConsul20210204PlatformTypeHCS captures enum value "HCS"
	HashicorpCloudConsul20210204PlatformTypeHCS HashicorpCloudConsul20210204PlatformType = "HCS"
)

// for schema
var hashicorpCloudConsul20210204PlatformTypeEnum []interface{}

func init() {
	var res []HashicorpCloudConsul20210204PlatformType
	if err := json.Unmarshal([]byte(`["NOTSET","HCP","HCS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudConsul20210204PlatformTypeEnum = append(hashicorpCloudConsul20210204PlatformTypeEnum, v)
	}
}

func (m HashicorpCloudConsul20210204PlatformType) validateHashicorpCloudConsul20210204PlatformTypeEnum(path, location string, value HashicorpCloudConsul20210204PlatformType) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudConsul20210204PlatformTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud consul 20210204 platform type
func (m HashicorpCloudConsul20210204PlatformType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudConsul20210204PlatformTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud consul 20210204 platform type based on context it is used
func (m HashicorpCloudConsul20210204PlatformType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
