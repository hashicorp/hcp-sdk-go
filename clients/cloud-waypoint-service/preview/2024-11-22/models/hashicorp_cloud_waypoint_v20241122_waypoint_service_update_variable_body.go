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

// HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody hashicorp cloud waypoint v20241122 waypoint service update variable body
//
// swagger:model hashicorp.cloud.waypoint.v20241122.WaypointService.UpdateVariableBody
type HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody struct {

	// namespace
	Namespace *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespace `json:"namespace,omitempty"`

	// The reference to the variable to update
	Ref *HashicorpCloudWaypointV20241122RefVariable `json:"ref,omitempty"`

	// Optional, will be used for PATCH requests. This includes what fields to
	// update in the request.
	UpdateMask string `json:"update_mask,omitempty"`

	// The variables fields to update
	Variable *HashicorpCloudWaypointV20241122Variable `json:"variable,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service update variable body
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody) validateNamespace(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody) validateRef(formats strfmt.Registry) error {
	if swag.IsZero(m.Ref) { // not required
		return nil
	}

	if m.Ref != nil {
		if err := m.Ref.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ref")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ref")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody) validateVariable(formats strfmt.Registry) error {
	if swag.IsZero(m.Variable) { // not required
		return nil
	}

	if m.Variable != nil {
		if err := m.Variable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("variable")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service update variable body based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody) contextValidateRef(ctx context.Context, formats strfmt.Registry) error {

	if m.Ref != nil {

		if swag.IsZero(m.Ref) { // not required
			return nil
		}

		if err := m.Ref.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ref")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ref")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody) contextValidateVariable(ctx context.Context, formats strfmt.Registry) error {

	if m.Variable != nil {

		if swag.IsZero(m.Variable) { // not required
			return nil
		}

		if err := m.Variable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("variable")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespace The namespace to update the variable in
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespace
type HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespace struct {

	// id
	ID string `json:"id,omitempty"`

	// location
	Location *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespaceLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service update variable body namespace
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespace) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service update variable body namespace based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespace) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespace) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespaceLocation When used via an API request, this is populated and used to populate id.
//
// When used via an API request, this is populated and used to populate id.
//
// swagger:model HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespaceLocation
type HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespaceLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *cloud.HashicorpCloudLocationRegion `json:"region,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 waypoint service update variable body namespace location
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespaceLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespaceLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint v20241122 waypoint service update variable body namespace location based on the context it is used
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespaceLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespaceLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespaceLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespaceLocation) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122WaypointServiceUpdateVariableBodyNamespaceLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
