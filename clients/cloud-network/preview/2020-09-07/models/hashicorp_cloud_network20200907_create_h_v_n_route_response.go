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

// HashicorpCloudNetwork20200907CreateHVNRouteResponse hashicorp cloud network 20200907 create h v n route response
//
// swagger:model hashicorp.cloud.network_20200907.CreateHVNRouteResponse
type HashicorpCloudNetwork20200907CreateHVNRouteResponse struct {

	// operation is the operation that represents the pending creation of the HVN Route.
	Operation *cloud.HashicorpCloudOperationOperation `json:"operation,omitempty"`

	// route is the newly created route
	Route *HashicorpCloudNetwork20200907HVNRoute `json:"route,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 create h v n route response
func (m *HashicorpCloudNetwork20200907CreateHVNRouteResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907CreateHVNRouteResponse) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907CreateHVNRouteResponse) validateRoute(formats strfmt.Registry) error {
	if swag.IsZero(m.Route) { // not required
		return nil
	}

	if m.Route != nil {
		if err := m.Route.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud network 20200907 create h v n route response based on the context it is used
func (m *HashicorpCloudNetwork20200907CreateHVNRouteResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907CreateHVNRouteResponse) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.Operation != nil {
		if err := m.Operation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907CreateHVNRouteResponse) contextValidateRoute(ctx context.Context, formats strfmt.Registry) error {

	if m.Route != nil {
		if err := m.Route.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreateHVNRouteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreateHVNRouteResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907CreateHVNRouteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
