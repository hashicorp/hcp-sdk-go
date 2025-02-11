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

// HashicorpCloudWaypointListNoCodeModulesResponse hashicorp cloud waypoint list no code modules response
//
// swagger:model hashicorp.cloud.waypoint.ListNoCodeModulesResponse
type HashicorpCloudWaypointListNoCodeModulesResponse struct {

	// no code modules
	NoCodeModules []*HashicorpCloudWaypointNoCodeModuleDefinition `json:"no_code_modules"`
}

// Validate validates this hashicorp cloud waypoint list no code modules response
func (m *HashicorpCloudWaypointListNoCodeModulesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNoCodeModules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointListNoCodeModulesResponse) validateNoCodeModules(formats strfmt.Registry) error {
	if swag.IsZero(m.NoCodeModules) { // not required
		return nil
	}

	for i := 0; i < len(m.NoCodeModules); i++ {
		if swag.IsZero(m.NoCodeModules[i]) { // not required
			continue
		}

		if m.NoCodeModules[i] != nil {
			if err := m.NoCodeModules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("no_code_modules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("no_code_modules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint list no code modules response based on the context it is used
func (m *HashicorpCloudWaypointListNoCodeModulesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNoCodeModules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointListNoCodeModulesResponse) contextValidateNoCodeModules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NoCodeModules); i++ {

		if m.NoCodeModules[i] != nil {

			if swag.IsZero(m.NoCodeModules[i]) { // not required
				return nil
			}

			if err := m.NoCodeModules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("no_code_modules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("no_code_modules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointListNoCodeModulesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointListNoCodeModulesResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointListNoCodeModulesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
