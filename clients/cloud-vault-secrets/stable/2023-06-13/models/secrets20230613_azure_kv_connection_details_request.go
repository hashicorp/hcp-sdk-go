// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20230613AzureKvConnectionDetailsRequest secrets 20230613 azure kv connection details request
//
// swagger:model secrets_20230613AzureKvConnectionDetailsRequest
type Secrets20230613AzureKvConnectionDetailsRequest struct {

	// client id
	ClientID string `json:"client_id,omitempty"`

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`

	// key vault uri
	KeyVaultURI string `json:"key_vault_uri,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty"`
}

// Validate validates this secrets 20230613 azure kv connection details request
func (m *Secrets20230613AzureKvConnectionDetailsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20230613 azure kv connection details request based on context it is used
func (m *Secrets20230613AzureKvConnectionDetailsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20230613AzureKvConnectionDetailsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20230613AzureKvConnectionDetailsRequest) UnmarshalBinary(b []byte) error {
	var res Secrets20230613AzureKvConnectionDetailsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
