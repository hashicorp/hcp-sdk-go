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
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudWaypointV20241122ListApplicationsResponse hashicorp cloud waypoint v20241122 list applications response
//
// swagger:model hashicorp.cloud.waypoint.v20241122.ListApplicationsResponse
type HashicorpCloudWaypointV20241122ListApplicationsResponse struct {

	// applications
	Applications []*HashicorpCloudWaypointV20241122Application `json:"applications"`

	// pagination
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`

	// Will return -1 if request has a with_total_count = false or is with_total_count parameter is omitted.
	TotalCount string `json:"total_count,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 list applications response
func (m *HashicorpCloudWaypointV20241122ListApplicationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122ListApplicationsResponse) validateApplications(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointV20241122ListApplicationsResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint v20241122 list applications response based on the context it is used
func (m *HashicorpCloudWaypointV20241122ListApplicationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122ListApplicationsResponse) contextValidateApplications(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointV20241122ListApplicationsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {

		if swag.IsZero(m.Pagination) { // not required
			return nil
		}

		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122ListApplicationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122ListApplicationsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122ListApplicationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
