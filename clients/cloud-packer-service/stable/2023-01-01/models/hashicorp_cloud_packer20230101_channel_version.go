// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPacker20230101ChannelVersion hashicorp cloud packer 20230101 channel version
//
// swagger:model hashicorp.cloud.packer_20230101.ChannelVersion
type HashicorpCloudPacker20230101ChannelVersion struct {

	// The currently assigned version's build fingerprint.
	Fingerprint string `json:"fingerprint,omitempty"`

	// The currently assigned version's ULID.
	ID string `json:"id,omitempty"`

	// The currently assigned version's name.
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp cloud packer 20230101 channel version
func (m *HashicorpCloudPacker20230101ChannelVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud packer 20230101 channel version based on context it is used
func (m *HashicorpCloudPacker20230101ChannelVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101ChannelVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101ChannelVersion) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20230101ChannelVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
