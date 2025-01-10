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

// Secrets20231128LockStatus secrets 20231128 lock status
//
// swagger:model secrets_20231128LockStatus
type Secrets20231128LockStatus struct {

	// lock origin
	LockOrigin *Secrets20231128LockOrigin `json:"lock_origin,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this secrets 20231128 lock status
func (m *Secrets20231128LockStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLockOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128LockStatus) validateLockOrigin(formats strfmt.Registry) error {
	if swag.IsZero(m.LockOrigin) { // not required
		return nil
	}

	if m.LockOrigin != nil {
		if err := m.LockOrigin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lock_origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lock_origin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secrets 20231128 lock status based on the context it is used
func (m *Secrets20231128LockStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLockOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128LockStatus) contextValidateLockOrigin(ctx context.Context, formats strfmt.Registry) error {

	if m.LockOrigin != nil {

		if swag.IsZero(m.LockOrigin) { // not required
			return nil
		}

		if err := m.LockOrigin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lock_origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lock_origin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128LockStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128LockStatus) UnmarshalBinary(b []byte) error {
	var res Secrets20231128LockStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
