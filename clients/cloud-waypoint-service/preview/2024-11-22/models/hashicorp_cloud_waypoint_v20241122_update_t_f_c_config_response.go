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

// HashicorpCloudWaypointV20241122UpdateTFCConfigResponse hashicorp cloud waypoint v20241122 update t f c config response
//
// swagger:model hashicorp.cloud.waypoint.v20241122.UpdateTFCConfigResponse
type HashicorpCloudWaypointV20241122UpdateTFCConfigResponse struct {

	// tfc config
	TfcConfig *HashicorpCloudWaypointV20241122TFCConfig `json:"tfc_config,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 update t f c config response
func (m *HashicorpCloudWaypointV20241122UpdateTFCConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTfcConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122UpdateTFCConfigResponse) validateTfcConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.TfcConfig) { // not required
		return nil
	}

	if m.TfcConfig != nil {
		if err := m.TfcConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tfc_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tfc_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint v20241122 update t f c config response based on the context it is used
func (m *HashicorpCloudWaypointV20241122UpdateTFCConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTfcConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122UpdateTFCConfigResponse) contextValidateTfcConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TfcConfig != nil {

		if swag.IsZero(m.TfcConfig) { // not required
			return nil
		}

		if err := m.TfcConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tfc_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tfc_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122UpdateTFCConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122UpdateTFCConfigResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122UpdateTFCConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
