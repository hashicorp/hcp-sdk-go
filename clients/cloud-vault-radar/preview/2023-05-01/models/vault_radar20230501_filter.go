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

// VaultRadar20230501Filter vault radar 20230501 filter
//
// swagger:model vault_radar_20230501Filter
type VaultRadar20230501Filter struct {

	// exact match
	ExactMatch bool `json:"exact_match,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// op
	Op *FilterFilterOperation `json:"op,omitempty"`

	// value
	Value []*VaultRadar20230501FilterValue `json:"value"`
}

// Validate validates this vault radar 20230501 filter
func (m *VaultRadar20230501Filter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VaultRadar20230501Filter) validateOp(formats strfmt.Registry) error {
	if swag.IsZero(m.Op) { // not required
		return nil
	}

	if m.Op != nil {
		if err := m.Op.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("op")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("op")
			}
			return err
		}
	}

	return nil
}

func (m *VaultRadar20230501Filter) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	for i := 0; i < len(m.Value); i++ {
		if swag.IsZero(m.Value[i]) { // not required
			continue
		}

		if m.Value[i] != nil {
			if err := m.Value[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vault radar 20230501 filter based on the context it is used
func (m *VaultRadar20230501Filter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VaultRadar20230501Filter) contextValidateOp(ctx context.Context, formats strfmt.Registry) error {

	if m.Op != nil {

		if swag.IsZero(m.Op) { // not required
			return nil
		}

		if err := m.Op.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("op")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("op")
			}
			return err
		}
	}

	return nil
}

func (m *VaultRadar20230501Filter) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Value); i++ {

		if m.Value[i] != nil {

			if swag.IsZero(m.Value[i]) { // not required
				return nil
			}

			if err := m.Value[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VaultRadar20230501Filter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VaultRadar20230501Filter) UnmarshalBinary(b []byte) error {
	var res VaultRadar20230501Filter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
