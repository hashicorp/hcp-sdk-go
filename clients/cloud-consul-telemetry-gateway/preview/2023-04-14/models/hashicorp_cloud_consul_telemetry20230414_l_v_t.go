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

// HashicorpCloudConsulTelemetry20230414LVT hashicorp cloud consul telemetry 20230414 l v t
//
// swagger:model hashicorp.cloud.consul_telemetry_20230414.LVT
type HashicorpCloudConsulTelemetry20230414LVT string

func NewHashicorpCloudConsulTelemetry20230414LVT(value HashicorpCloudConsulTelemetry20230414LVT) *HashicorpCloudConsulTelemetry20230414LVT {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudConsulTelemetry20230414LVT.
func (m HashicorpCloudConsulTelemetry20230414LVT) Pointer() *HashicorpCloudConsulTelemetry20230414LVT {
	return &m
}

const (

	// HashicorpCloudConsulTelemetry20230414LVTLVTUNSPECIFIED captures enum value "LVT_UNSPECIFIED"
	HashicorpCloudConsulTelemetry20230414LVTLVTUNSPECIFIED HashicorpCloudConsulTelemetry20230414LVT = "LVT_UNSPECIFIED"

	// HashicorpCloudConsulTelemetry20230414LVTLVTSERVERIDS captures enum value "LVT_SERVER_IDS"
	HashicorpCloudConsulTelemetry20230414LVTLVTSERVERIDS HashicorpCloudConsulTelemetry20230414LVT = "LVT_SERVER_IDS"
)

// for schema
var hashicorpCloudConsulTelemetry20230414LVTEnum []interface{}

func init() {
	var res []HashicorpCloudConsulTelemetry20230414LVT
	if err := json.Unmarshal([]byte(`["LVT_UNSPECIFIED","LVT_SERVER_IDS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudConsulTelemetry20230414LVTEnum = append(hashicorpCloudConsulTelemetry20230414LVTEnum, v)
	}
}

func (m HashicorpCloudConsulTelemetry20230414LVT) validateHashicorpCloudConsulTelemetry20230414LVTEnum(path, location string, value HashicorpCloudConsulTelemetry20230414LVT) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudConsulTelemetry20230414LVTEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud consul telemetry 20230414 l v t
func (m HashicorpCloudConsulTelemetry20230414LVT) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudConsulTelemetry20230414LVTEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud consul telemetry 20230414 l v t based on context it is used
func (m HashicorpCloudConsulTelemetry20230414LVT) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}