// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamEnableMFAResponse hashicorp cloud iam enable m f a response
//
// swagger:model hashicorp.cloud.iam.EnableMFAResponse
type HashicorpCloudIamEnableMFAResponse struct {

	// qrcode uri
	QrcodeURI string `json:"qrcode_uri,omitempty"`

	// recovery codes
	RecoveryCodes []string `json:"recovery_codes"`

	// secret
	Secret string `json:"secret,omitempty"`
}

// Validate validates this hashicorp cloud iam enable m f a response
func (m *HashicorpCloudIamEnableMFAResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam enable m f a response based on context it is used
func (m *HashicorpCloudIamEnableMFAResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamEnableMFAResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamEnableMFAResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamEnableMFAResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
