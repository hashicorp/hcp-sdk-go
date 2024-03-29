// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVagrantCompleteUploadRequest hashicorp cloud vagrant complete upload request
//
// swagger:model hashicorp.cloud.vagrant.CompleteUploadRequest
type HashicorpCloudVagrantCompleteUploadRequest struct {

	// The name segment of the Box. As an example, this field would represent the
	// "vagrant" in "hashicorp/vagrant".
	Box string `json:"box,omitempty"`

	// The HMAC for the upload supplied in the output of the corresponding
	// UploadProvider call.
	Hmac string `json:"hmac,omitempty"`

	// The name of the Provider.
	Provider string `json:"provider,omitempty"`

	// The Registry segment of the Box. As an example, this field would represent
	// the "hashicorp" in "hashicorp/vagrant".
	Registry string `json:"registry,omitempty"`

	// The name of the Version for the Provider.
	Version string `json:"version,omitempty"`
}

// Validate validates this hashicorp cloud vagrant complete upload request
func (m *HashicorpCloudVagrantCompleteUploadRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vagrant complete upload request based on context it is used
func (m *HashicorpCloudVagrantCompleteUploadRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrantCompleteUploadRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrantCompleteUploadRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrantCompleteUploadRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
