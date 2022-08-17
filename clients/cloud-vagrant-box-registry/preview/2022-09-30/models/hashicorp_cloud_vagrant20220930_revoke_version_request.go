// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVagrant20220930RevokeVersionRequest hashicorp cloud vagrant 20220930 revoke version request
//
// swagger:model hashicorp.cloud.vagrant_20220930.RevokeVersionRequest
type HashicorpCloudVagrant20220930RevokeVersionRequest struct {

	// The name segment of the Box. As an example, this field would represent the
	// "vagrant" in "hashicorp/vagrant".
	Box string `json:"box,omitempty"`

	// The Registry segment of the Box. As an example, this field would represent
	// the "hashicorp" in "hashicorp/vagrant".
	Registry string `json:"registry,omitempty"`

	// The name of the Version to revoke.
	Version string `json:"version,omitempty"`
}

// Validate validates this hashicorp cloud vagrant 20220930 revoke version request
func (m *HashicorpCloudVagrant20220930RevokeVersionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930RevokeVersionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930RevokeVersionRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrant20220930RevokeVersionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
