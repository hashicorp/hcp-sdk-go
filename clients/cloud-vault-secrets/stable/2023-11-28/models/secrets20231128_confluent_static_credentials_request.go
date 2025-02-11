// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128ConfluentStaticCredentialsRequest secrets 20231128 confluent static credentials request
//
// swagger:model secrets_20231128ConfluentStaticCredentialsRequest
type Secrets20231128ConfluentStaticCredentialsRequest struct {

	// cloud api key id
	CloudAPIKeyID string `json:"cloud_api_key_id,omitempty"`

	// cloud api secret
	CloudAPISecret string `json:"cloud_api_secret,omitempty"`
}

// Validate validates this secrets 20231128 confluent static credentials request
func (m *Secrets20231128ConfluentStaticCredentialsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 confluent static credentials request based on context it is used
func (m *Secrets20231128ConfluentStaticCredentialsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128ConfluentStaticCredentialsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128ConfluentStaticCredentialsRequest) UnmarshalBinary(b []byte) error {
	var res Secrets20231128ConfluentStaticCredentialsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
