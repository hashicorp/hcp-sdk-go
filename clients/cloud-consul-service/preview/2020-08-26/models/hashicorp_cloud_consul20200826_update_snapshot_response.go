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

// HashicorpCloudConsul20200826UpdateSnapshotResponse UpdateSnapshotResponse is a response to updating a snapshot.
//
// swagger:model hashicorp.cloud.consul_20200826.UpdateSnapshotResponse
type HashicorpCloudConsul20200826UpdateSnapshotResponse struct {

	// snapshot is the updated snapshot.
	Snapshot *HashicorpCloudConsul20200826Snapshot `json:"snapshot,omitempty"`
}

// Validate validates this hashicorp cloud consul 20200826 update snapshot response
func (m *HashicorpCloudConsul20200826UpdateSnapshotResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826UpdateSnapshotResponse) validateSnapshot(formats strfmt.Registry) error {
	if swag.IsZero(m.Snapshot) { // not required
		return nil
	}

	if m.Snapshot != nil {
		if err := m.Snapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul 20200826 update snapshot response based on the context it is used
func (m *HashicorpCloudConsul20200826UpdateSnapshotResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826UpdateSnapshotResponse) contextValidateSnapshot(ctx context.Context, formats strfmt.Registry) error {

	if m.Snapshot != nil {
		if err := m.Snapshot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826UpdateSnapshotResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826UpdateSnapshotResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20200826UpdateSnapshotResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
