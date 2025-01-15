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

// HashicorpCloudWaypointListActionRunsByNamespaceResponse hashicorp cloud waypoint list action runs by namespace response
//
// swagger:model hashicorp.cloud.waypoint.ListActionRunsByNamespaceResponse
type HashicorpCloudWaypointListActionRunsByNamespaceResponse struct {

	// The list of action runs for the namespace
	ActionRuns []*HashicorpCloudWaypointActionRun `json:"action_runs"`

	// pagination
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`
}

// Validate validates this hashicorp cloud waypoint list action runs by namespace response
func (m *HashicorpCloudWaypointListActionRunsByNamespaceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionRuns(formats); err != nil {
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

func (m *HashicorpCloudWaypointListActionRunsByNamespaceResponse) validateActionRuns(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionRuns) { // not required
		return nil
	}

	for i := 0; i < len(m.ActionRuns); i++ {
		if swag.IsZero(m.ActionRuns[i]) { // not required
			continue
		}

		if m.ActionRuns[i] != nil {
			if err := m.ActionRuns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_runs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_runs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointListActionRunsByNamespaceResponse) validatePagination(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint list action runs by namespace response based on the context it is used
func (m *HashicorpCloudWaypointListActionRunsByNamespaceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionRuns(ctx, formats); err != nil {
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

func (m *HashicorpCloudWaypointListActionRunsByNamespaceResponse) contextValidateActionRuns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ActionRuns); i++ {

		if m.ActionRuns[i] != nil {

			if swag.IsZero(m.ActionRuns[i]) { // not required
				return nil
			}

			if err := m.ActionRuns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_runs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_runs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointListActionRunsByNamespaceResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointListActionRunsByNamespaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointListActionRunsByNamespaceResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointListActionRunsByNamespaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
