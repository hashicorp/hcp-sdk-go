// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudConsul20210204DeleteSnapshotResponse DeleteSnapshotResponse is a response to deleting a snapshot.
//
// swagger:model hashicorp.cloud.consul_20210204.DeleteSnapshotResponse
type HashicorpCloudConsul20210204DeleteSnapshotResponse struct {

	// operation represents the deletion of the requested snapshot.
	Operation *cloud.HashicorpCloudOperationOperation `json:"operation,omitempty"`
}

// Validate validates this hashicorp cloud consul 20210204 delete snapshot response
func (m *HashicorpCloudConsul20210204DeleteSnapshotResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20210204DeleteSnapshotResponse) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul 20210204 delete snapshot response based on the context it is used
func (m *HashicorpCloudConsul20210204DeleteSnapshotResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20210204DeleteSnapshotResponse) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.Operation != nil {
		if err := m.Operation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20210204DeleteSnapshotResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20210204DeleteSnapshotResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20210204DeleteSnapshotResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
