// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVagrant20220930UploadProviderRequest hashicorp cloud vagrant 20220930 upload provider request
//
// swagger:model hashicorp.cloud.vagrant_20220930.UploadProviderRequest
type HashicorpCloudVagrant20220930UploadProviderRequest struct {

	// The name segment of the Box. As an example, this field would represent the
	// "vagrant" in "hashicorp/vagrant".
	Box string `json:"box,omitempty"`

	// The name of the Provider.
	Provider string `json:"provider,omitempty"`

	// The Registry segment of the Box. As an example, this field would represent
	// the "hashicorp" in "hashicorp/vagrant".
	Registry string `json:"registry,omitempty"`

	// The name of the Version for the Provider.
	Version string `json:"version,omitempty"`
}

// Validate validates this hashicorp cloud vagrant 20220930 upload provider request
func (m *HashicorpCloudVagrant20220930UploadProviderRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930UploadProviderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930UploadProviderRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrant20220930UploadProviderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}