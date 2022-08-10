// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPacker20220411RevokedAncestor hashicorp cloud packer 20220411 revoked ancestor
//
// swagger:model hashicorp.cloud.packer_20220411.RevokedAncestor
type HashicorpCloudPacker20220411RevokedAncestor struct {

	// The ancestor bucket slug.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// The URL to get the iteration's revoked ancestor.
	Href string `json:"href,omitempty"`

	// The ancestor iteration fingerprint.
	IterationFingerprint string `json:"iteration_fingerprint,omitempty"`

	// The ancestor iteration ULID.
	IterationID string `json:"iteration_id,omitempty"`

	// The ancestor iteration incremental version.
	IterationIncrementalVersion int32 `json:"iteration_incremental_version,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 revoked ancestor
func (m *HashicorpCloudPacker20220411RevokedAncestor) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411RevokedAncestor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411RevokedAncestor) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411RevokedAncestor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
