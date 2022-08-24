// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20200420GetAdminTokenResponse hashicorp cloud vault 20200420 get admin token response
//
// swagger:model hashicorp.cloud.vault_20200420.GetAdminTokenResponse
type HashicorpCloudVault20200420GetAdminTokenResponse struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 get admin token response
func (m *HashicorpCloudVault20200420GetAdminTokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20200420 get admin token response based on context it is used
func (m *HashicorpCloudVault20200420GetAdminTokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420GetAdminTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420GetAdminTokenResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420GetAdminTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
