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
	"github.com/go-openapi/validate"
)

// Secrets20231128GetOrganizationLockStatusResponse secrets 20231128 get organization lock status response
//
// swagger:model secrets_20231128GetOrganizationLockStatusResponse
type Secrets20231128GetOrganizationLockStatusResponse struct {

	// locks
	Locks []*Secrets20231128LockStatus `json:"locks"`

	// project locks
	ProjectLocks map[string]Secrets20231128ProjectLocks `json:"project_locks,omitempty"`
}

// Validate validates this secrets 20231128 get organization lock status response
func (m *Secrets20231128GetOrganizationLockStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectLocks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128GetOrganizationLockStatusResponse) validateLocks(formats strfmt.Registry) error {
	if swag.IsZero(m.Locks) { // not required
		return nil
	}

	for i := 0; i < len(m.Locks); i++ {
		if swag.IsZero(m.Locks[i]) { // not required
			continue
		}

		if m.Locks[i] != nil {
			if err := m.Locks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("locks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Secrets20231128GetOrganizationLockStatusResponse) validateProjectLocks(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectLocks) { // not required
		return nil
	}

	for k := range m.ProjectLocks {

		if err := validate.Required("project_locks"+"."+k, "body", m.ProjectLocks[k]); err != nil {
			return err
		}
		if val, ok := m.ProjectLocks[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("project_locks" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("project_locks" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this secrets 20231128 get organization lock status response based on the context it is used
func (m *Secrets20231128GetOrganizationLockStatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectLocks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128GetOrganizationLockStatusResponse) contextValidateLocks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Locks); i++ {

		if m.Locks[i] != nil {

			if swag.IsZero(m.Locks[i]) { // not required
				return nil
			}

			if err := m.Locks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("locks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Secrets20231128GetOrganizationLockStatusResponse) contextValidateProjectLocks(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ProjectLocks {

		if val, ok := m.ProjectLocks[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128GetOrganizationLockStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128GetOrganizationLockStatusResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128GetOrganizationLockStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
