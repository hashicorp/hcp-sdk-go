// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVagrantUpdateProviderResponse hashicorp cloud vagrant update provider response
//
// swagger:model hashicorp.cloud.vagrant.UpdateProviderResponse
type HashicorpCloudVagrantUpdateProviderResponse struct {

	// external
	External *HashicorpCloudVagrantExternalProvider `json:"external,omitempty"`

	// hosted
	Hosted *HashicorpCloudVagrantHostedProvider `json:"hosted,omitempty"`
}

// Validate validates this hashicorp cloud vagrant update provider response
func (m *HashicorpCloudVagrantUpdateProviderResponse) Validate(formats strfmt.Registry) error {
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

func (m *HashicorpCloudVagrantUpdateProviderResponse) validateExternal(formats strfmt.Registry) error {
	if swag.IsZero(m.External) { // not required
		return nil
	}

	if m.External != nil {
		if err := m.External.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVagrantUpdateProviderResponse) validateHosted(formats strfmt.Registry) error {
	if swag.IsZero(m.Hosted) { // not required
		return nil
	}

	if m.Hosted != nil {
		if err := m.Hosted.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hosted")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hosted")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vagrant update provider response based on the context it is used
func (m *HashicorpCloudVagrantUpdateProviderResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHosted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrantUpdateProviderResponse) contextValidateExternal(ctx context.Context, formats strfmt.Registry) error {

	if m.External != nil {
		if err := m.External.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVagrantUpdateProviderResponse) contextValidateHosted(ctx context.Context, formats strfmt.Registry) error {

	if m.Hosted != nil {
		if err := m.Hosted.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hosted")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hosted")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrantUpdateProviderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrantUpdateProviderResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrantUpdateProviderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
