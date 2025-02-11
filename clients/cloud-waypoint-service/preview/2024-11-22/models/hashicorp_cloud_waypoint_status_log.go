// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudWaypointStatusLog hashicorp cloud waypoint status log
//
// swagger:model hashicorp.cloud.waypoint.StatusLog
type HashicorpCloudWaypointStatusLog struct {

	// The time the client generated this message
	// Format: date-time
	EmittedAt strfmt.DateTime `json:"emitted_at,omitempty"`

	// Free-form text description of the log
	Log string `json:"log,omitempty"`

	// A map of key/value pairs of whatever
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Validate validates this hashicorp cloud waypoint status log
func (m *HashicorpCloudWaypointStatusLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmittedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointStatusLog) validateEmittedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.EmittedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("emitted_at", "body", "date-time", m.EmittedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hashicorp cloud waypoint status log based on context it is used
func (m *HashicorpCloudWaypointStatusLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointStatusLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointStatusLog) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointStatusLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
