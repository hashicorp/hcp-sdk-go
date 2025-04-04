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

// HashicorpCloudWaypointV20241122GetTFRunStatusResponse GetTFRunStatusResponse is the response containing the url & status of the
// latest run for a Terraform workspace, as well as the variables set on the
// workspace.
//
// swagger:model hashicorp.cloud.waypoint.v20241122.GetTFRunStatusResponse
type HashicorpCloudWaypointV20241122GetTFRunStatusResponse struct {

	// the variables set on the TF workspace
	InputVariables []*HashicorpCloudWaypointV20241122InputVariable `json:"input_variables"`

	// the state of the TF run
	State *HashicorpCloudWaypointV20241122TerraformTFRunState `json:"state,omitempty"`

	// the URL of the TF run
	URL string `json:"url,omitempty"`
}

// Validate validates this hashicorp cloud waypoint v20241122 get t f run status response
func (m *HashicorpCloudWaypointV20241122GetTFRunStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputVariables(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122GetTFRunStatusResponse) validateInputVariables(formats strfmt.Registry) error {
	if swag.IsZero(m.InputVariables) { // not required
		return nil
	}

	for i := 0; i < len(m.InputVariables); i++ {
		if swag.IsZero(m.InputVariables[i]) { // not required
			continue
		}

		if m.InputVariables[i] != nil {
			if err := m.InputVariables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input_variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("input_variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointV20241122GetTFRunStatusResponse) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint v20241122 get t f run status response based on the context it is used
func (m *HashicorpCloudWaypointV20241122GetTFRunStatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointV20241122GetTFRunStatusResponse) contextValidateInputVariables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InputVariables); i++ {

		if m.InputVariables[i] != nil {

			if swag.IsZero(m.InputVariables[i]) { // not required
				return nil
			}

			if err := m.InputVariables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input_variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("input_variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointV20241122GetTFRunStatusResponse) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {

		if swag.IsZero(m.State) { // not required
			return nil
		}

		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122GetTFRunStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointV20241122GetTFRunStatusResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointV20241122GetTFRunStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
