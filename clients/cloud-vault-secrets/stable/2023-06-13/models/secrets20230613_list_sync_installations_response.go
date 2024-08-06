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

// Secrets20230613ListSyncInstallationsResponse secrets 20230613 list sync installations response
//
// swagger:model secrets_20230613ListSyncInstallationsResponse
type Secrets20230613ListSyncInstallationsResponse struct {

	// sync installations
	SyncInstallations []*Secrets20230613SyncInstallation `json:"sync_installations"`
}

// Validate validates this secrets 20230613 list sync installations response
func (m *Secrets20230613ListSyncInstallationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyncInstallations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20230613ListSyncInstallationsResponse) validateSyncInstallations(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncInstallations) { // not required
		return nil
	}

	for i := 0; i < len(m.SyncInstallations); i++ {
		if swag.IsZero(m.SyncInstallations[i]) { // not required
			continue
		}

		if m.SyncInstallations[i] != nil {
			if err := m.SyncInstallations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sync_installations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sync_installations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this secrets 20230613 list sync installations response based on the context it is used
func (m *Secrets20230613ListSyncInstallationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSyncInstallations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20230613ListSyncInstallationsResponse) contextValidateSyncInstallations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SyncInstallations); i++ {

		if m.SyncInstallations[i] != nil {

			if swag.IsZero(m.SyncInstallations[i]) { // not required
				return nil
			}

			if err := m.SyncInstallations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sync_installations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sync_installations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20230613ListSyncInstallationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20230613ListSyncInstallationsResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20230613ListSyncInstallationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
