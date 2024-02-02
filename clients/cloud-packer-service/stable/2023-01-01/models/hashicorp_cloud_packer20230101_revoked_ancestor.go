// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPacker20230101RevokedAncestor hashicorp cloud packer 20230101 revoked ancestor
//
// swagger:model hashicorp.cloud.packer_20230101.RevokedAncestor
type HashicorpCloudPacker20230101RevokedAncestor struct {

	// The revoked ancestor bucket name.
	BucketName string `json:"bucket_name,omitempty"`

	// The URL to get the revoked ancestor.
	Href string `json:"href,omitempty"`

	// The revoked ancestor version fingerprint.
	VersionFingerprint string `json:"version_fingerprint,omitempty"`

	// The revoked ancestor version ULID.
	VersionID string `json:"version_id,omitempty"`

	// The revoked ancestor version name.
	VersionName string `json:"version_name,omitempty"`
}

// Validate validates this hashicorp cloud packer 20230101 revoked ancestor
func (m *HashicorpCloudPacker20230101RevokedAncestor) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud packer 20230101 revoked ancestor based on context it is used
func (m *HashicorpCloudPacker20230101RevokedAncestor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101RevokedAncestor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101RevokedAncestor) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20230101RevokedAncestor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}