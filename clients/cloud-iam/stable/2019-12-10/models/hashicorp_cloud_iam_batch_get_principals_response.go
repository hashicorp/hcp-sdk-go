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

// HashicorpCloudIamBatchGetPrincipalsResponse BatchGetPrincipalsResponse is the batch get principals response.
//
// swagger:model hashicorp.cloud.iam.BatchGetPrincipalsResponse
type HashicorpCloudIamBatchGetPrincipalsResponse struct {

	// principals is the requested principals.
	Principals []*HashicorpCloudIamPrincipal `json:"principals"`
}

// Validate validates this hashicorp cloud iam batch get principals response
func (m *HashicorpCloudIamBatchGetPrincipalsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrincipals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamBatchGetPrincipalsResponse) validatePrincipals(formats strfmt.Registry) error {
	if swag.IsZero(m.Principals) { // not required
		return nil
	}

	for i := 0; i < len(m.Principals); i++ {
		if swag.IsZero(m.Principals[i]) { // not required
			continue
		}

		if m.Principals[i] != nil {
			if err := m.Principals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("principals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("principals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud iam batch get principals response based on the context it is used
func (m *HashicorpCloudIamBatchGetPrincipalsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrincipals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamBatchGetPrincipalsResponse) contextValidatePrincipals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Principals); i++ {

		if m.Principals[i] != nil {

			if swag.IsZero(m.Principals[i]) { // not required
				return nil
			}

			if err := m.Principals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("principals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("principals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamBatchGetPrincipalsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamBatchGetPrincipalsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamBatchGetPrincipalsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}