// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVagrant20220930UpdateProviderRequest hashicorp cloud vagrant 20220930 update provider request
//
// swagger:model hashicorp.cloud.vagrant_20220930.UpdateProviderRequest
type HashicorpCloudVagrant20220930UpdateProviderRequest struct {

	// The name segment of the Box. As an example, this field would represent the
	// "vagrant" in "hashicorp/vagrant".
	Box string `json:"box,omitempty"`

	// external
	External *HashicorpCloudVagrant20220930ExternalProvider `json:"external,omitempty"`

	// hosted
	Hosted *HashicorpCloudVagrant20220930HostedProvider `json:"hosted,omitempty"`

	// The name of the Provider.
	Provider string `json:"provider,omitempty"`

	// The Registry segment of the Box. As an example, this field would represent
	// the "hashicorp" in "hashicorp/vagrant".
	Registry string `json:"registry,omitempty"`

	// The name of the Version for the Provider.
	Version string `json:"version,omitempty"`
}

// Validate validates this hashicorp cloud vagrant 20220930 update provider request
func (m *HashicorpCloudVagrant20220930UpdateProviderRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHosted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrant20220930UpdateProviderRequest) validateExternal(formats strfmt.Registry) error {

	if swag.IsZero(m.External) { // not required
		return nil
	}

	if m.External != nil {
		if err := m.External.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVagrant20220930UpdateProviderRequest) validateHosted(formats strfmt.Registry) error {

	if swag.IsZero(m.Hosted) { // not required
		return nil
	}

	if m.Hosted != nil {
		if err := m.Hosted.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hosted")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930UpdateProviderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930UpdateProviderRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrant20220930UpdateProviderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
