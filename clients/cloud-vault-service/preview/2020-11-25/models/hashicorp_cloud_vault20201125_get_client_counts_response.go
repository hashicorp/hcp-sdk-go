// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125GetClientCountsResponse hashicorp cloud vault 20201125 get client counts response
//
// swagger:model hashicorp.cloud.vault_20201125.GetClientCountsResponse
type HashicorpCloudVault20201125GetClientCountsResponse struct {

	// clients
	Clients string `json:"clients,omitempty"`

	// distinct entities
	DistinctEntities string `json:"distinct_entities,omitempty"`

	// non entity tokens
	NonEntityTokens string `json:"non_entity_tokens,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 get client counts response
func (m *HashicorpCloudVault20201125GetClientCountsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetClientCountsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetClientCountsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125GetClientCountsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
