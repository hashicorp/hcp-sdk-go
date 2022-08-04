// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudVagrantHostedProvider hashicorp cloud vagrant hosted provider
//
// swagger:model hashicorp.cloud.vagrant.HostedProvider
type HashicorpCloudVagrantHostedProvider struct {

	// The checksum for the Provider at the external URL.
	Checksum string `json:"checksum,omitempty"`

	// The algorithm type for the provided checksum.
	ChecksumType HashicorpCloudVagrantHostedProviderChecksumType `json:"checksum_type,omitempty"`

	// The date the record was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The name of the Provider, should be unique within the version.
	ID string `json:"id,omitempty"`

	// The upload state of the Provider.
	State HashicorpCloudVagrantHostedProviderState `json:"state,omitempty"`

	// The date that the record was last updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this hashicorp cloud vagrant hosted provider
func (m *HashicorpCloudVagrantHostedProvider) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChecksumType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrantHostedProvider) validateChecksumType(formats strfmt.Registry) error {

	if swag.IsZero(m.ChecksumType) { // not required
		return nil
	}

	if err := m.ChecksumType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("checksum_type")
		}
		return err
	}

	return nil
}

func (m *HashicorpCloudVagrantHostedProvider) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVagrantHostedProvider) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *HashicorpCloudVagrantHostedProvider) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrantHostedProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrantHostedProvider) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrantHostedProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
