// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudWaypointWaypointServiceCreateTFCConfigBody hashicorp cloud waypoint waypoint service create t f c config body
//
// swagger:model hashicorp.cloud.waypoint.WaypointService.CreateTFCConfigBody
type HashicorpCloudWaypointWaypointServiceCreateTFCConfigBody struct {

	// namespace
	Namespace *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBodyNamespace `json:"namespace,omitempty"`

	// tfc config
	TfcConfig *HashicorpCloudWaypointTFCConfig `json:"tfc_config,omitempty"`
}

// Validate validates this hashicorp cloud waypoint waypoint service create t f c config body
func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTfcConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBody) validateNamespace(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespace) { // not required
		return nil
	}

	if m.Namespace != nil {
		if err := m.Namespace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBody) validateTfcConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint waypoint service create t f c config body based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTfcConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBody) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

	if m.Namespace != nil {

		if swag.IsZero(m.Namespace) { // not required
			return nil
		}

		if err := m.Namespace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBody) contextValidateTfcConfig(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceCreateTFCConfigBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointWaypointServiceCreateTFCConfigBodyNamespace Global references the entire server. This is used in some APIs
// as a way to read/write values that are server-global.
//
// swagger:model HashicorpCloudWaypointWaypointServiceCreateTFCConfigBodyNamespace
type HashicorpCloudWaypointWaypointServiceCreateTFCConfigBodyNamespace struct {

	// When used via an API request, this is populated and used to populate id.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud waypoint waypoint service create t f c config body namespace
func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBodyNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBodyNamespace) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint waypoint service create t f c config body namespace based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBodyNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBodyNamespace) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {

		if swag.IsZero(m.Location) { // not required
			return nil
		}

		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBodyNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceCreateTFCConfigBodyNamespace) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceCreateTFCConfigBodyNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
