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

// HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBody hashicorp cloud waypoint v20241122 waypoint service update t f c config body
//
// swagger:model hashicorp.cloud.waypoint.v20241122.WaypointService.UpdateTFCConfigBody
type HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBody struct {

	// namespace
	Namespace *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespace `json:"namespace,omitempty"`

	// tfc config
	TfcConfig *HashicorpCloudWaypointV20241122TFCConfig `json:"tfc_config,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service update t f c config body
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBody) Validate(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBody) validateNamespace(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBody) validateTfcConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service update t f c config body based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBody) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBody) contextValidateTfcConfig(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespace Global references the entire server. This is used in some APIs
// as a way to read/write values that are server-global.
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespace
type HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespace struct {

	// id
	ID string `json:"id,omitempty"`

	// location
	Location *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespaceLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service update t f c config body namespace
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespace) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service update t f c config body namespace based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespace) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespace) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespaceLocation When used via an API request, this is populated and used to populate id.
//
// When used via an API request, this is populated and used to populate id.
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespaceLocation
type HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespaceLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *cloud.HashicorpCloudLocationRegion `json:"region,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service update t f c config body namespace location
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespaceLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespaceLocation) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service update t f c config body namespace location based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespaceLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespaceLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {

		if swag.IsZero(m.Region) { // not required
			return nil
		}

		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespaceLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespaceLocation) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUpdateTFCConfigBodyNamespaceLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
