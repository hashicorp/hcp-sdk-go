// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudPackerRegenerateTFCRunTaskHmacKeyRequest hashicorp cloud packer regenerate t f c run task hmac key request
//
// swagger:model hashicorp.cloud.packer.RegenerateTFCRunTaskHmacKeyRequest
type HashicorpCloudPackerRegenerateTFCRunTaskHmacKeyRequest struct {

	// HCP-specific information like project and organization ID
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud packer regenerate t f c run task hmac key request
func (m *HashicorpCloudPackerRegenerateTFCRunTaskHmacKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerRegenerateTFCRunTaskHmacKeyRequest) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerRegenerateTFCRunTaskHmacKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerRegenerateTFCRunTaskHmacKeyRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerRegenerateTFCRunTaskHmacKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
