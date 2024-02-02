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

// HashicorpCloudPacker20230101RegistryConfigTier hashicorp cloud packer 20230101 registry config tier
//
// swagger:model hashicorp.cloud.packer_20230101.RegistryConfig.Tier
type HashicorpCloudPacker20230101RegistryConfigTier string

func NewHashicorpCloudPacker20230101RegistryConfigTier(value HashicorpCloudPacker20230101RegistryConfigTier) *HashicorpCloudPacker20230101RegistryConfigTier {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudPacker20230101RegistryConfigTier.
func (m HashicorpCloudPacker20230101RegistryConfigTier) Pointer() *HashicorpCloudPacker20230101RegistryConfigTier {
	return &m
}

const (

	// HashicorpCloudPacker20230101RegistryConfigTierUNSET captures enum value "UNSET"
	HashicorpCloudPacker20230101RegistryConfigTierUNSET HashicorpCloudPacker20230101RegistryConfigTier = "UNSET"

	// HashicorpCloudPacker20230101RegistryConfigTierSTANDARD captures enum value "STANDARD"
	HashicorpCloudPacker20230101RegistryConfigTierSTANDARD HashicorpCloudPacker20230101RegistryConfigTier = "STANDARD"

	// HashicorpCloudPacker20230101RegistryConfigTierPLUS captures enum value "PLUS"
	HashicorpCloudPacker20230101RegistryConfigTierPLUS HashicorpCloudPacker20230101RegistryConfigTier = "PLUS"
)

// for schema
var hashicorpCloudPacker20230101RegistryConfigTierEnum []interface{}

func init() {
	var res []HashicorpCloudPacker20230101RegistryConfigTier
	if err := json.Unmarshal([]byte(`["UNSET","STANDARD","PLUS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudPacker20230101RegistryConfigTierEnum = append(hashicorpCloudPacker20230101RegistryConfigTierEnum, v)
	}
}

func (m HashicorpCloudPacker20230101RegistryConfigTier) validateHashicorpCloudPacker20230101RegistryConfigTierEnum(path, location string, value HashicorpCloudPacker20230101RegistryConfigTier) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudPacker20230101RegistryConfigTierEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud packer 20230101 registry config tier
func (m HashicorpCloudPacker20230101RegistryConfigTier) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudPacker20230101RegistryConfigTierEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud packer 20230101 registry config tier based on context it is used
func (m HashicorpCloudPacker20230101RegistryConfigTier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}