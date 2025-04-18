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

// HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBody hashicorp cloud waypoint v20241122 waypoint service UI bulk update action assign for app body
//
// swagger:model hashicorp.cloud.waypoint.v20241122.WaypointService.UI_BulkUpdateActionAssignForAppBody
type HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBody struct {

	// applications
	Applications []*HashicorpCloudWaypointV20241122Application `json:"applications"`

	// namespace
	Namespace *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespace `json:"namespace,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service UI bulk update action assign for app body
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplications(formats); err != nil {
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

func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBody) validateApplications(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBody) validateNamespace(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service UI bulk update action assign for app body based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplications(ctx, formats); err != nil {
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

func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBody) contextValidateApplications(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBody) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespace The namespace this app belongs to
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespace
type HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespace struct {

	// id
	ID string `json:"id,omitempty"`

	// location
	Location *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespaceLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service UI bulk update action assign for app body namespace
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespace) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service UI bulk update action assign for app body namespace based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespace) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespace) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespaceLocation When used via an API request, this is populated and used to populate id.
//
// When used via an API request, this is populated and used to populate id.
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespaceLocation
type HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespaceLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *cloud.HashicorpCloudLocationRegion `json:"region,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service UI bulk update action assign for app body namespace location
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespaceLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespaceLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service UI bulk update action assign for app body namespace location based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespaceLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespaceLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespaceLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespaceLocation) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUIBulkUpdateActionAssignForAppBodyNamespaceLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
