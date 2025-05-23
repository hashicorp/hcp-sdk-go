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

// HashicorpCloudWaypointV20241122UIListActionConfigResponse hashicorp cloud waypoint v20241122 UI list action config response
//
// swagger:model hashicorp.cloud.waypoint.v20241122.UI.ListActionConfigResponse
type HashicorpCloudWaypointV20241122UIListActionConfigResponse struct {

	// action config bundles
	ActionConfigBundles []*HashicorpCloudWaypointV20241122UIActionConfigBundle `json:"action_config_bundles"`

	// pagination
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 UI list action config response
func (m *HashicorpCloudWaypointV20241122UIListActionConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionConfigBundles(formats); err != nil {
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

func (m *HashicorpCloudWaypointV20241122UIListActionConfigResponse) validateActionConfigBundles(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionConfigBundles) { // not required
		return nil
	}

	for i := 0; i < len(m.ActionConfigBundles); i++ {
		if swag.IsZero(m.ActionConfigBundles[i]) { // not required
			continue
		}

		if m.ActionConfigBundles[i] != nil {
			if err := m.ActionConfigBundles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_config_bundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_config_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointV20241122UIListActionConfigResponse) validatePagination(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 UI list action config response based on the context it is used
func (m *HashicorpCloudWaypointV20241122UIListActionConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionConfigBundles(ctx, formats); err != nil {
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

func (m *HashicorpCloudWaypointV20241122UIListActionConfigResponse) contextValidateActionConfigBundles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ActionConfigBundles); i++ {

		if m.ActionConfigBundles[i] != nil {

			if swag.IsZero(m.ActionConfigBundles[i]) { // not required
				return nil
			}

			if err := m.ActionConfigBundles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_config_bundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_config_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointV20241122UIListActionConfigResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122UIListActionConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122UIListActionConfigResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122UIListActionConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
