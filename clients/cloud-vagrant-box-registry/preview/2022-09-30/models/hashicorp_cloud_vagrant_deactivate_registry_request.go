// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVagrantDeactivateRegistryRequest hashicorp cloud vagrant deactivate registry request
//
// swagger:model hashicorp.cloud.vagrant.DeactivateRegistryRequest
type HashicorpCloudVagrantDeactivateRegistryRequest struct {

	// The name of the Registry to deactivate.
	Registry string `json:"registry,omitempty"`
}

// Validate validates this hashicorp cloud vagrant deactivate registry request
func (m *HashicorpCloudVagrantDeactivateRegistryRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrantDeactivateRegistryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrantDeactivateRegistryRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrantDeactivateRegistryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
