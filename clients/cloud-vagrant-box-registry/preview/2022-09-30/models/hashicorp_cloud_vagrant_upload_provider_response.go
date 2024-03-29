// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVagrantUploadProviderResponse hashicorp cloud vagrant upload provider response
//
// swagger:model hashicorp.cloud.vagrant.UploadProviderResponse
type HashicorpCloudVagrantUploadProviderResponse struct {

	// An HMAC used to validate completion requests for this upload. This should
	// be sent to CompleteUpload when this request is finished.
	Hmac string `json:"hmac,omitempty"`

	// A pre-signed URL to upload the Provider to.
	URL string `json:"url,omitempty"`
}

// Validate validates this hashicorp cloud vagrant upload provider response
func (m *HashicorpCloudVagrantUploadProviderResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vagrant upload provider response based on context it is used
func (m *HashicorpCloudVagrantUploadProviderResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrantUploadProviderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrantUploadProviderResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrantUploadProviderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
