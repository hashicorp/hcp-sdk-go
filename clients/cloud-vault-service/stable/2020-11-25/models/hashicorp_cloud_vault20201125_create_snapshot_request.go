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

// HashicorpCloudVault20201125CreateSnapshotRequest hashicorp cloud vault 20201125 create snapshot request
//
// swagger:model hashicorp.cloud.vault_20201125.CreateSnapshotRequest
type HashicorpCloudVault20201125CreateSnapshotRequest struct {

	// name is the user provided name of the snapshot.
	Name string `json:"name,omitempty"`

	// resource specifies the link to the resource to snapshot
	Resource *HashicorpCloudInternalLocationLink `json:"resource,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 create snapshot request
func (m *HashicorpCloudVault20201125CreateSnapshotRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125CreateSnapshotRequest) validateResource(formats strfmt.Registry) error {
	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20201125 create snapshot request based on the context it is used
func (m *HashicorpCloudVault20201125CreateSnapshotRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125CreateSnapshotRequest) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {
		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125CreateSnapshotRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125CreateSnapshotRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125CreateSnapshotRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
