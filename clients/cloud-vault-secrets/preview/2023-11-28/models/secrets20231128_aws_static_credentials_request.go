// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128AwsStaticCredentialsRequest secrets 20231128 aws static credentials request
//
// swagger:model secrets_20231128AwsStaticCredentialsRequest
type Secrets20231128AwsStaticCredentialsRequest struct {

	// access key id
	AccessKeyID string `json:"access_key_id,omitempty"`

	// secret access key
	SecretAccessKey string `json:"secret_access_key,omitempty"`
}

// Validate validates this secrets 20231128 aws static credentials request
func (m *Secrets20231128AwsStaticCredentialsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 aws static credentials request based on context it is used
func (m *Secrets20231128AwsStaticCredentialsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128AwsStaticCredentialsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128AwsStaticCredentialsRequest) UnmarshalBinary(b []byte) error {
	var res Secrets20231128AwsStaticCredentialsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
