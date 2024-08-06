// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20230613HcpTerraformMetadata secrets 20230613 hcp terraform metadata
//
// swagger:model secrets_20230613HcpTerraformMetadata
type Secrets20230613HcpTerraformMetadata struct {

	// org
	Org string `json:"org,omitempty"`

	// team
	Team string `json:"team,omitempty"`
}

// Validate validates this secrets 20230613 hcp terraform metadata
func (m *Secrets20230613HcpTerraformMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20230613 hcp terraform metadata based on context it is used
func (m *Secrets20230613HcpTerraformMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20230613HcpTerraformMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20230613HcpTerraformMetadata) UnmarshalBinary(b []byte) error {
	var res Secrets20230613HcpTerraformMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
