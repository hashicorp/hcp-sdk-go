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

// HashicorpCloudPackerCreateChannelResponse hashicorp cloud packer create channel response
//
// swagger:model hashicorp.cloud.packer.CreateChannelResponse
type HashicorpCloudPackerCreateChannelResponse struct {

	// Information about the channel that was created.
	Channel *HashicorpCloudPackerChannel `json:"channel,omitempty"`
}

// Validate validates this hashicorp cloud packer create channel response
func (m *HashicorpCloudPackerCreateChannelResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerCreateChannelResponse) validateChannel(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud packer create channel response based on the context it is used
func (m *HashicorpCloudPackerCreateChannelResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChannel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerCreateChannelResponse) contextValidateChannel(ctx context.Context, formats strfmt.Registry) error {

	if m.Channel != nil {
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
func (m *HashicorpCloudPackerCreateChannelResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerCreateChannelResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerCreateChannelResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
