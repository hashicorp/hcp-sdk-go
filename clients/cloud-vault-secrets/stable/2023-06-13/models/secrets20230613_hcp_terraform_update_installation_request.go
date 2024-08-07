// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20230613HcpTerraformUpdateInstallationRequest secrets 20230613 hcp terraform update installation request
//
// swagger:model secrets_20230613HcpTerraformUpdateInstallationRequest
type Secrets20230613HcpTerraformUpdateInstallationRequest struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this secrets 20230613 hcp terraform update installation request
func (m *Secrets20230613HcpTerraformUpdateInstallationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20230613 hcp terraform update installation request based on context it is used
func (m *Secrets20230613HcpTerraformUpdateInstallationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20230613HcpTerraformUpdateInstallationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20230613HcpTerraformUpdateInstallationRequest) UnmarshalBinary(b []byte) error {
	var res Secrets20230613HcpTerraformUpdateInstallationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
