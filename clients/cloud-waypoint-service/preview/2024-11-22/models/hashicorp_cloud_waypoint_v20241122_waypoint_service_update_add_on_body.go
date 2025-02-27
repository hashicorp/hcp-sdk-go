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

// HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBody UpdateAddOnRequest is the request used to update an existing Add-on.
//
// swagger:model hashicorp.cloud.waypoint.v20241122.WaypointService.UpdateAddOnBody
type HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBody struct {

	// existing add on
	ExistingAddOn *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyExistingAddOn `json:"existing_add_on,omitempty"`

	// The new name of the Add-on.
	Name string `json:"name,omitempty"`

	// namespace
	Namespace *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespace `json:"namespace,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service update add on body
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExistingAddOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBody) validateExistingAddOn(formats strfmt.Registry) error {
	if swag.IsZero(m.ExistingAddOn) { // not required
		return nil
	}

	if m.ExistingAddOn != nil {
		if err := m.ExistingAddOn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existing_add_on")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existing_add_on")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBody) validateNamespace(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service update add on body based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExistingAddOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBody) contextValidateExistingAddOn(ctx context.Context, formats strfmt.Registry) error {

	if m.ExistingAddOn != nil {

		if swag.IsZero(m.ExistingAddOn) { // not required
			return nil
		}

		if err := m.ExistingAddOn.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existing_add_on")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existing_add_on")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBody) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyExistingAddOn Add-on ref without ID
//
// # The add-on to be updated
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyExistingAddOn
type HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyExistingAddOn struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service update add on body existing add on
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyExistingAddOn) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint v20241122 waypoint service update add on body existing add on based on context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyExistingAddOn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyExistingAddOn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyExistingAddOn) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyExistingAddOn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespace Global references the entire server. This is used in some APIs
// as a way to read/write values that are server-global.
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespace
type HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespace struct {

	// id
	ID string `json:"id,omitempty"`

	// location
	Location *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespaceLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service update add on body namespace
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespace) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service update add on body namespace based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespace) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespace) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespaceLocation When used via an API request, this is populated and used to populate id.
//
// When used via an API request, this is populated and used to populate id.
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespaceLocation
type HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespaceLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *cloud.HashicorpCloudLocationRegion `json:"region,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service update add on body namespace location
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespaceLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespaceLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service update add on body namespace location based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespaceLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespaceLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespaceLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespaceLocation) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnBodyNamespaceLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
