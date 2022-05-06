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

// HashicorpCloudPackerUpdateRegistryRequest hashicorp cloud packer update registry request
//
// swagger:model hashicorp.cloud.packer.UpdateRegistryRequest
type HashicorpCloudPackerUpdateRegistryRequest struct {

	// Activates a deactivated registry. A registry can only be activated
	// if the organization billing account is valid (e.g. valid credits or payment method).
	Activate bool `json:"activate,omitempty"`

	// Feature tier of the Registry.
	FeatureTier HashicorpCloudPackerRegistryConfigTier `json:"feature_tier,omitempty"`

	// location
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud packer update registry request
func (m *HashicorpCloudPackerUpdateRegistryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatureTier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerUpdateRegistryRequest) validateFeatureTier(formats strfmt.Registry) error {

	if swag.IsZero(m.FeatureTier) { // not required
		return nil
	}

	if err := m.FeatureTier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("feature_tier")
		}
		return err
	}

	return nil
}

func (m *HashicorpCloudPackerUpdateRegistryRequest) validateLocation(formats strfmt.Registry) error {

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
func (m *HashicorpCloudPackerUpdateRegistryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerUpdateRegistryRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerUpdateRegistryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}