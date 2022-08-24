// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVagrantActivateRegistryRequest hashicorp cloud vagrant activate registry request
//
// swagger:model hashicorp.cloud.vagrant.ActivateRegistryRequest
type HashicorpCloudVagrantActivateRegistryRequest struct {

	// The name of the Registry to activate.
	Registry string `json:"registry,omitempty"`
}

// Validate validates this hashicorp cloud vagrant activate registry request
func (m *HashicorpCloudVagrantActivateRegistryRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vagrant activate registry request based on context it is used
func (m *HashicorpCloudVagrantActivateRegistryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrantActivateRegistryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrantActivateRegistryRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrantActivateRegistryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
