// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Secrets20231128Sync Syncs
//
// swagger:model secrets_20231128Sync
type Secrets20231128Sync struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// integration name
	IntegrationName string `json:"integration_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// resource id
	ResourceID string `json:"resource_id,omitempty"`

	// resource name
	ResourceName string `json:"resource_name,omitempty"`

	// sync config gitlab
	SyncConfigGitlab *Secrets20231128SyncConfigGitlab `json:"sync_config_gitlab,omitempty"`

	// sync config hcp terraform
	SyncConfigHcpTerraform *Secrets20231128SyncConfigHcpTerraform `json:"sync_config_hcp_terraform,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// updated by id
	UpdatedByID string `json:"updated_by_id,omitempty"`
}

// Validate validates this secrets 20231128 sync
func (m *Secrets20231128Sync) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncConfigGitlab(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncConfigHcpTerraform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128Sync) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128Sync) validateSyncConfigGitlab(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncConfigGitlab) { // not required
		return nil
	}

	if m.SyncConfigGitlab != nil {
		if err := m.SyncConfigGitlab.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync_config_gitlab")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync_config_gitlab")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Sync) validateSyncConfigHcpTerraform(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncConfigHcpTerraform) { // not required
		return nil
	}

	if m.SyncConfigHcpTerraform != nil {
		if err := m.SyncConfigHcpTerraform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync_config_hcp_terraform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync_config_hcp_terraform")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Sync) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this secrets 20231128 sync based on the context it is used
func (m *Secrets20231128Sync) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSyncConfigGitlab(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSyncConfigHcpTerraform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128Sync) contextValidateSyncConfigGitlab(ctx context.Context, formats strfmt.Registry) error {

	if m.SyncConfigGitlab != nil {

		if swag.IsZero(m.SyncConfigGitlab) { // not required
			return nil
		}

		if err := m.SyncConfigGitlab.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync_config_gitlab")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync_config_gitlab")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Sync) contextValidateSyncConfigHcpTerraform(ctx context.Context, formats strfmt.Registry) error {

	if m.SyncConfigHcpTerraform != nil {

		if swag.IsZero(m.SyncConfigHcpTerraform) { // not required
			return nil
		}

		if err := m.SyncConfigHcpTerraform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync_config_hcp_terraform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync_config_hcp_terraform")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128Sync) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128Sync) UnmarshalBinary(b []byte) error {
	var res Secrets20231128Sync
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
