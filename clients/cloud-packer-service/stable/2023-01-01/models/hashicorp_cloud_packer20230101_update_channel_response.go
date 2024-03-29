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

// HashicorpCloudPacker20230101UpdateChannelResponse hashicorp cloud packer 20230101 update channel response
//
// swagger:model hashicorp.cloud.packer_20230101.UpdateChannelResponse
type HashicorpCloudPacker20230101UpdateChannelResponse struct {

	// channel
	Channel *HashicorpCloudPacker20230101Channel `json:"channel,omitempty"`
}

// Validate validates this hashicorp cloud packer 20230101 update channel response
func (m *HashicorpCloudPacker20230101UpdateChannelResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101UpdateChannelResponse) validateChannel(formats strfmt.Registry) error {
	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	if m.Channel != nil {
		if err := m.Channel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer 20230101 update channel response based on the context it is used
func (m *HashicorpCloudPacker20230101UpdateChannelResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChannel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101UpdateChannelResponse) contextValidateChannel(ctx context.Context, formats strfmt.Registry) error {

	if m.Channel != nil {

		if swag.IsZero(m.Channel) { // not required
			return nil
		}

		if err := m.Channel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101UpdateChannelResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101UpdateChannelResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20230101UpdateChannelResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
