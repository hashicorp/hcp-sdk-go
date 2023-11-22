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

// HashicorpCloudVault20201125HTTPEncodingCodec hashicorp cloud vault 20201125 HTTP encoding codec
//
// swagger:model hashicorp.cloud.vault_20201125.HTTP.EncodingCodec
type HashicorpCloudVault20201125HTTPEncodingCodec string

func NewHashicorpCloudVault20201125HTTPEncodingCodec(value HashicorpCloudVault20201125HTTPEncodingCodec) *HashicorpCloudVault20201125HTTPEncodingCodec {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudVault20201125HTTPEncodingCodec.
func (m HashicorpCloudVault20201125HTTPEncodingCodec) Pointer() *HashicorpCloudVault20201125HTTPEncodingCodec {
	return &m
}

const (

	// HashicorpCloudVault20201125HTTPEncodingCodecJSON captures enum value "JSON"
	HashicorpCloudVault20201125HTTPEncodingCodecJSON HashicorpCloudVault20201125HTTPEncodingCodec = "JSON"

	// HashicorpCloudVault20201125HTTPEncodingCodecNDJSON captures enum value "NDJSON"
	HashicorpCloudVault20201125HTTPEncodingCodecNDJSON HashicorpCloudVault20201125HTTPEncodingCodec = "NDJSON"
)

// for schema
var hashicorpCloudVault20201125HttpEncodingCodecEnum []interface{}

func init() {
	var res []HashicorpCloudVault20201125HTTPEncodingCodec
	if err := json.Unmarshal([]byte(`["JSON","NDJSON"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVault20201125HttpEncodingCodecEnum = append(hashicorpCloudVault20201125HttpEncodingCodecEnum, v)
	}
}

func (m HashicorpCloudVault20201125HTTPEncodingCodec) validateHashicorpCloudVault20201125HTTPEncodingCodecEnum(path, location string, value HashicorpCloudVault20201125HTTPEncodingCodec) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVault20201125HttpEncodingCodecEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vault 20201125 HTTP encoding codec
func (m HashicorpCloudVault20201125HTTPEncodingCodec) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVault20201125HTTPEncodingCodecEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20201125 HTTP encoding codec based on context it is used
func (m HashicorpCloudVault20201125HTTPEncodingCodec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
