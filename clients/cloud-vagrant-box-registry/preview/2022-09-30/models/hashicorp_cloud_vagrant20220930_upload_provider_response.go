// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVagrant20220930UploadProviderResponse hashicorp cloud vagrant 20220930 upload provider response
//
// swagger:model hashicorp.cloud.vagrant_20220930.UploadProviderResponse
type HashicorpCloudVagrant20220930UploadProviderResponse struct {

	// An HMAC used to validate completion requests for this upload. This should
	// be sent to CompleteUpload when this request is finished.
	Hmac string `json:"hmac,omitempty"`

	// A pre-signed URL to upload the Provider to.
	URL string `json:"url,omitempty"`
}

// Validate validates this hashicorp cloud vagrant 20220930 upload provider response
func (m *HashicorpCloudVagrant20220930UploadProviderResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930UploadProviderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930UploadProviderResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrant20220930UploadProviderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}