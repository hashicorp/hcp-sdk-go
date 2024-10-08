// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogService20210330ListDestinationsResponse log service 20210330 list destinations response
//
// swagger:model log_service_20210330ListDestinationsResponse
type LogService20210330ListDestinationsResponse struct {

	// destinations
	Destinations []*LogService20210330StreamingListDestination `json:"destinations"`
}

// Validate validates this log service 20210330 list destinations response
func (m *LogService20210330ListDestinationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330ListDestinationsResponse) validateDestinations(formats strfmt.Registry) error {
	if swag.IsZero(m.Destinations) { // not required
		return nil
	}

	for i := 0; i < len(m.Destinations); i++ {
		if swag.IsZero(m.Destinations[i]) { // not required
			continue
		}

		if m.Destinations[i] != nil {
			if err := m.Destinations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("destinations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("destinations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this log service 20210330 list destinations response based on the context it is used
func (m *LogService20210330ListDestinationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330ListDestinationsResponse) contextValidateDestinations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Destinations); i++ {

		if m.Destinations[i] != nil {

			if swag.IsZero(m.Destinations[i]) { // not required
				return nil
			}

			if err := m.Destinations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("destinations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("destinations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogService20210330ListDestinationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogService20210330ListDestinationsResponse) UnmarshalBinary(b []byte) error {
	var res LogService20210330ListDestinationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
