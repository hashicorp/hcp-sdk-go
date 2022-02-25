// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPackerRegistryTFCRunTaskConfig hashicorp cloud packer registry t f c run task config
//
// swagger:model hashicorp.cloud.packer.RegistryTFCRunTaskConfig
type HashicorpCloudPackerRegistryTFCRunTaskConfig struct {

	// Unique per registry API id for running HCP Packer run tasks.
	APIID string `json:"api_id,omitempty"`

	// Encrypted HMAC key used by Terraform Cloud to sign the requests to the HCP Packer run task API.
	HmacKey string `json:"hmac_key,omitempty"`
}

// Validate validates this hashicorp cloud packer registry t f c run task config
func (m *HashicorpCloudPackerRegistryTFCRunTaskConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerRegistryTFCRunTaskConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerRegistryTFCRunTaskConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerRegistryTFCRunTaskConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
