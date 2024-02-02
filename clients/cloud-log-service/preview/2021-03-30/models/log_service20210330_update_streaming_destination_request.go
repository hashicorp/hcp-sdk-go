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

// LogService20210330UpdateStreamingDestinationRequest UpdateStreamingDestinationRequest represents a request to update a destination record name or confiuration.
//
// swagger:model log_service_20210330UpdateStreamingDestinationRequest
type LogService20210330UpdateStreamingDestinationRequest struct {

	// destination represents the destination record.
	Destination *LogService20210330Destination `json:"destination,omitempty"`

	// mask represnts the requested fields to be updated.
	Mask *ProtobufFieldMask `json:"mask,omitempty"`

	// source_channel denotes where the update request was made from. For example terraform or the hcp ui.
	SourceChannel string `json:"source_channel,omitempty"`
}

// Validate validates this log service 20210330 update streaming destination request
func (m *LogService20210330UpdateStreamingDestinationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330UpdateStreamingDestinationRequest) validateDestination(formats strfmt.Registry) error {
	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	if m.Destination != nil {
		if err := m.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330UpdateStreamingDestinationRequest) validateMask(formats strfmt.Registry) error {
	if swag.IsZero(m.Mask) { // not required
		return nil
	}

	if m.Mask != nil {
		if err := m.Mask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mask")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mask")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this log service 20210330 update streaming destination request based on the context it is used
func (m *LogService20210330UpdateStreamingDestinationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330UpdateStreamingDestinationRequest) contextValidateDestination(ctx context.Context, formats strfmt.Registry) error {

	if m.Destination != nil {

		if swag.IsZero(m.Destination) { // not required
			return nil
		}

		if err := m.Destination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330UpdateStreamingDestinationRequest) contextValidateMask(ctx context.Context, formats strfmt.Registry) error {

	if m.Mask != nil {

		if swag.IsZero(m.Mask) { // not required
			return nil
		}

		if err := m.Mask.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mask")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mask")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogService20210330UpdateStreamingDestinationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogService20210330UpdateStreamingDestinationRequest) UnmarshalBinary(b []byte) error {
	var res LogService20210330UpdateStreamingDestinationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}