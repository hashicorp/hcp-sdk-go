// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPackerImageCreateBody This message is used for build/create calls; it removes elements from the
// Image message that the user cannot set.
//
// swagger:model hashicorp.cloud.packer.ImageCreateBody
type HashicorpCloudPackerImageCreateBody struct {

	// ID or URL of the remote cloud image as given by a build.
	ImageID string `json:"image_id,omitempty"`

	// Cloud-specific region as provided by `packer build`. For example,
	// "ap-east-1".
	Region string `json:"region,omitempty"`
}

// Validate validates this hashicorp cloud packer image create body
func (m *HashicorpCloudPackerImageCreateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerImageCreateBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerImageCreateBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerImageCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
