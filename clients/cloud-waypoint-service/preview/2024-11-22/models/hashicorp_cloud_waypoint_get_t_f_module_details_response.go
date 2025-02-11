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

// HashicorpCloudWaypointGetTFModuleDetailsResponse hashicorp cloud waypoint get t f module details response
//
// swagger:model hashicorp.cloud.waypoint.GetTFModuleDetailsResponse
type HashicorpCloudWaypointGetTFModuleDetailsResponse struct {

	// module details
	ModuleDetails *HashicorpCloudWaypointTFModuleDetails `json:"module_details,omitempty"`
}

// Validate validates this hashicorp cloud waypoint get t f module details response
func (m *HashicorpCloudWaypointGetTFModuleDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModuleDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointGetTFModuleDetailsResponse) validateModuleDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ModuleDetails) { // not required
		return nil
	}

	if m.ModuleDetails != nil {
		if err := m.ModuleDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint get t f module details response based on the context it is used
func (m *HashicorpCloudWaypointGetTFModuleDetailsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModuleDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointGetTFModuleDetailsResponse) contextValidateModuleDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ModuleDetails != nil {

		if swag.IsZero(m.ModuleDetails) { // not required
			return nil
		}

		if err := m.ModuleDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointGetTFModuleDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointGetTFModuleDetailsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointGetTFModuleDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
