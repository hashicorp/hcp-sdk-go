// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPackerRegistryConfig The HCP Packer Registry configuration
//
// swagger:model hashicorp.cloud.packer.RegistryConfig
type HashicorpCloudPackerRegistryConfig struct {

	// A registry is activated when the system correctly starts billing for it.
	Activated bool `json:"activated,omitempty"`

	// The information about the billing deprovision
	BillingDeprovision *HashicorpCloudPackerRegistryBillingDeprovision `json:"billing_deprovision,omitempty"`

	// feature_tier is the feature tier for the registry.
	FeatureTier HashicorpCloudPackerRegistryConfigTier `json:"feature_tier,omitempty"`
}

// Validate validates this hashicorp cloud packer registry config
func (m *HashicorpCloudPackerRegistryConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingDeprovision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatureTier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerRegistryConfig) validateBillingDeprovision(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingDeprovision) { // not required
		return nil
	}

	if m.BillingDeprovision != nil {
		if err := m.BillingDeprovision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_deprovision")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerRegistryConfig) validateFeatureTier(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerRegistryConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerRegistryConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerRegistryConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
