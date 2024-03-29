// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPacker20230101CreateChannelBody hashicorp cloud packer 20230101 create channel body
//
// swagger:model hashicorp.cloud.packer_20230101.CreateChannelBody
type HashicorpCloudPacker20230101CreateChannelBody struct {

	// Human-readable name for the channel.
	Name string `json:"name,omitempty"`

	// Whether this channel's access is restricted to users with write permission in the HCP Packer registry.
	Restricted bool `json:"restricted,omitempty"`

	// Fingerprint of the version. The fingerprint is set by Packer when you
	// call `packer build`.
	VersionFingerprint string `json:"version_fingerprint,omitempty"`
}

// Validate validates this hashicorp cloud packer 20230101 create channel body
func (m *HashicorpCloudPacker20230101CreateChannelBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud packer 20230101 create channel body based on context it is used
func (m *HashicorpCloudPacker20230101CreateChannelBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101CreateChannelBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101CreateChannelBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20230101CreateChannelBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
