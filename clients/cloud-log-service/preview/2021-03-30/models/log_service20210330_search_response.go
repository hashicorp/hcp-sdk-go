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

// LogService20210330SearchResponse log service 20210330 search response
//
// swagger:model log_service_20210330SearchResponse
type LogService20210330SearchResponse struct {

	// matrix are only returned when running a query that computes some value.
	Matrix *LogService20210330MatrixResponse `json:"matrix,omitempty"`

	// This token allows you to get the next page of results for list requests.
	NextPageToken string `json:"next_page_token,omitempty"`

	// streams is a set of all values (logs) over the queried time range.
	Streams *LogService20210330StreamsResponse `json:"streams,omitempty"`
}

// Validate validates this log service 20210330 search response
func (m *LogService20210330SearchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatrix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330SearchResponse) validateMatrix(formats strfmt.Registry) error {
	if swag.IsZero(m.Matrix) { // not required
		return nil
	}

	if m.Matrix != nil {
		if err := m.Matrix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matrix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("matrix")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330SearchResponse) validateStreams(formats strfmt.Registry) error {
	if swag.IsZero(m.Streams) { // not required
		return nil
	}

	if m.Streams != nil {
		if err := m.Streams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this log service 20210330 search response based on the context it is used
func (m *LogService20210330SearchResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatrix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStreams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330SearchResponse) contextValidateMatrix(ctx context.Context, formats strfmt.Registry) error {

	if m.Matrix != nil {

		if swag.IsZero(m.Matrix) { // not required
			return nil
		}

		if err := m.Matrix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matrix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("matrix")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330SearchResponse) contextValidateStreams(ctx context.Context, formats strfmt.Registry) error {

	if m.Streams != nil {

		if swag.IsZero(m.Streams) { // not required
			return nil
		}

		if err := m.Streams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogService20210330SearchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogService20210330SearchResponse) UnmarshalBinary(b []byte) error {
	var res LogService20210330SearchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
