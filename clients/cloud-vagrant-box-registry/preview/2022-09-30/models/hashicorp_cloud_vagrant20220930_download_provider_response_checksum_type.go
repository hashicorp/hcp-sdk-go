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

// HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType ChecksumTypes describe supported checksum types.
//
// swagger:model hashicorp.cloud.vagrant_20220930.DownloadProviderResponse.ChecksumType
type HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType string

const (

	// HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeNONE captures enum value "NONE"
	HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeNONE HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType = "NONE"

	// HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeMD5 captures enum value "MD5"
	HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeMD5 HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType = "MD5"

	// HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeSHA1 captures enum value "SHA1"
	HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeSHA1 HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType = "SHA1"

	// HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeSHA256 captures enum value "SHA256"
	HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeSHA256 HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType = "SHA256"

	// HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeSHA384 captures enum value "SHA384"
	HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeSHA384 HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType = "SHA384"

	// HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeSHA512 captures enum value "SHA512"
	HashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeSHA512 HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType = "SHA512"
)

// for schema
var hashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeEnum []interface{}

func init() {
	var res []HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType
	if err := json.Unmarshal([]byte(`["NONE","MD5","SHA1","SHA256","SHA384","SHA512"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeEnum = append(hashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeEnum, v)
	}
}

func (m HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType) validateHashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeEnum(path, location string, value HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vagrant 20220930 download provider response checksum type
func (m HashicorpCloudVagrant20220930DownloadProviderResponseChecksumType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVagrant20220930DownloadProviderResponseChecksumTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}