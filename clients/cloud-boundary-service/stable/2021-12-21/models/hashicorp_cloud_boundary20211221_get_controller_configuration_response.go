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

// HashicorpCloudBoundary20211221GetControllerConfigurationResponse The GetControllerConfigurationResponse type is the response given when querying a cluster's controller
// configuration.
//
// swagger:model hashicorp.cloud.boundary_20211221.GetControllerConfigurationResponse
type HashicorpCloudBoundary20211221GetControllerConfigurationResponse struct {

	// The controller configuration values.
	Config *HashicorpCloudBoundary20211221ControllerConfiguration `json:"config,omitempty"`
}

// Validate validates this hashicorp cloud boundary 20211221 get controller configuration response
func (m *HashicorpCloudBoundary20211221GetControllerConfigurationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudBoundary20211221GetControllerConfigurationResponse) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud boundary 20211221 get controller configuration response based on the context it is used
func (m *HashicorpCloudBoundary20211221GetControllerConfigurationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudBoundary20211221GetControllerConfigurationResponse) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {

		if swag.IsZero(m.Config) { // not required
			return nil
		}

		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudBoundary20211221GetControllerConfigurationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudBoundary20211221GetControllerConfigurationResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudBoundary20211221GetControllerConfigurationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
