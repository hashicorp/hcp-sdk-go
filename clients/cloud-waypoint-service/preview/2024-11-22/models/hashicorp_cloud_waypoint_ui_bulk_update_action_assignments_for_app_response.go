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

// HashicorpCloudWaypointUIBulkUpdateActionAssignmentsForAppResponse hashicorp cloud waypoint UI bulk update action assignments for app response
//
// swagger:model hashicorp.cloud.waypoint.UI.BulkUpdateActionAssignmentsForAppResponse
type HashicorpCloudWaypointUIBulkUpdateActionAssignmentsForAppResponse struct {

	// The updated applications
	Applications []*HashicorpCloudWaypointApplication `json:"applications"`
}

// Validate validates this hashicorp cloud waypoint UI bulk update action assignments for app response
func (m *HashicorpCloudWaypointUIBulkUpdateActionAssignmentsForAppResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplications(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUIBulkUpdateActionAssignmentsForAppResponse) validateApplications(formats strfmt.Registry) error {
	if swag.IsZero(m.Applications) { // not required
		return nil
	}

	for i := 0; i < len(m.Applications); i++ {
		if swag.IsZero(m.Applications[i]) { // not required
			continue
		}

		if m.Applications[i] != nil {
			if err := m.Applications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint UI bulk update action assignments for app response based on the context it is used
func (m *HashicorpCloudWaypointUIBulkUpdateActionAssignmentsForAppResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUIBulkUpdateActionAssignmentsForAppResponse) contextValidateApplications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Applications); i++ {

		if m.Applications[i] != nil {

			if swag.IsZero(m.Applications[i]) { // not required
				return nil
			}

			if err := m.Applications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointUIBulkUpdateActionAssignmentsForAppResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointUIBulkUpdateActionAssignmentsForAppResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointUIBulkUpdateActionAssignmentsForAppResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
