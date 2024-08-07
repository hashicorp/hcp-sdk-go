// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128GcpServiceAccountKeyRequest secrets 20231128 gcp service account key request
//
// swagger:model secrets_20231128GcpServiceAccountKeyRequest
type Secrets20231128GcpServiceAccountKeyRequest struct {

	// credentials
	Credentials string `json:"credentials,omitempty"`
}

// Validate validates this secrets 20231128 gcp service account key request
func (m *Secrets20231128GcpServiceAccountKeyRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 gcp service account key request based on context it is used
func (m *Secrets20231128GcpServiceAccountKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128GcpServiceAccountKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128GcpServiceAccountKeyRequest) UnmarshalBinary(b []byte) error {
	var res Secrets20231128GcpServiceAccountKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
