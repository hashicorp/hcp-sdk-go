// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125InputClusterTemplateInput hashicorp cloud vault 20201125 input cluster template input
//
// swagger:model hashicorp.cloud.vault_20201125.InputCluster.TemplateInput
type HashicorpCloudVault20201125InputClusterTemplateInput struct {

	// id
	ID string `json:"id,omitempty"`

	// params
	Params map[string]string `json:"params,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 input cluster template input
func (m *HashicorpCloudVault20201125InputClusterTemplateInput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20201125 input cluster template input based on context it is used
func (m *HashicorpCloudVault20201125InputClusterTemplateInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125InputClusterTemplateInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125InputClusterTemplateInput) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125InputClusterTemplateInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
