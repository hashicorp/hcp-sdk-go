// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpCloudVagrant20220930HostedProviderChecksumType ChecksumTypes describe supported checksum types.
//
// swagger:model hashicorp.cloud.vagrant_20220930.HostedProvider.ChecksumType
type HashicorpCloudVagrant20220930HostedProviderChecksumType string

const (

	// HashicorpCloudVagrant20220930HostedProviderChecksumTypeNONE captures enum value "NONE"
	HashicorpCloudVagrant20220930HostedProviderChecksumTypeNONE HashicorpCloudVagrant20220930HostedProviderChecksumType = "NONE"

	// HashicorpCloudVagrant20220930HostedProviderChecksumTypeMD5 captures enum value "MD5"
	HashicorpCloudVagrant20220930HostedProviderChecksumTypeMD5 HashicorpCloudVagrant20220930HostedProviderChecksumType = "MD5"

	// HashicorpCloudVagrant20220930HostedProviderChecksumTypeSHA1 captures enum value "SHA1"
	HashicorpCloudVagrant20220930HostedProviderChecksumTypeSHA1 HashicorpCloudVagrant20220930HostedProviderChecksumType = "SHA1"

	// HashicorpCloudVagrant20220930HostedProviderChecksumTypeSHA256 captures enum value "SHA256"
	HashicorpCloudVagrant20220930HostedProviderChecksumTypeSHA256 HashicorpCloudVagrant20220930HostedProviderChecksumType = "SHA256"

	// HashicorpCloudVagrant20220930HostedProviderChecksumTypeSHA384 captures enum value "SHA384"
	HashicorpCloudVagrant20220930HostedProviderChecksumTypeSHA384 HashicorpCloudVagrant20220930HostedProviderChecksumType = "SHA384"

	// HashicorpCloudVagrant20220930HostedProviderChecksumTypeSHA512 captures enum value "SHA512"
	HashicorpCloudVagrant20220930HostedProviderChecksumTypeSHA512 HashicorpCloudVagrant20220930HostedProviderChecksumType = "SHA512"
)

// for schema
var hashicorpCloudVagrant20220930HostedProviderChecksumTypeEnum []interface{}

func init() {
	var res []HashicorpCloudVagrant20220930HostedProviderChecksumType
	if err := json.Unmarshal([]byte(`["NONE","MD5","SHA1","SHA256","SHA384","SHA512"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVagrant20220930HostedProviderChecksumTypeEnum = append(hashicorpCloudVagrant20220930HostedProviderChecksumTypeEnum, v)
	}
}

func (m HashicorpCloudVagrant20220930HostedProviderChecksumType) validateHashicorpCloudVagrant20220930HostedProviderChecksumTypeEnum(path, location string, value HashicorpCloudVagrant20220930HostedProviderChecksumType) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVagrant20220930HostedProviderChecksumTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vagrant 20220930 hosted provider checksum type
func (m HashicorpCloudVagrant20220930HostedProviderChecksumType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVagrant20220930HostedProviderChecksumTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}