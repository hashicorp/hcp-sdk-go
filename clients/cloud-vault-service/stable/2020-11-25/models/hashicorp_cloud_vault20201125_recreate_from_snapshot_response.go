// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudVault20201125RecreateFromSnapshotResponse hashicorp cloud vault 20201125 recreate from snapshot response
//
// swagger:model hashicorp.cloud.vault_20201125.RecreateFromSnapshotResponse
type HashicorpCloudVault20201125RecreateFromSnapshotResponse struct {

	// operation
	Operation *cloud.HashicorpCloudOperationOperation `json:"operation,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 recreate from snapshot response
func (m *HashicorpCloudVault20201125RecreateFromSnapshotResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125RecreateFromSnapshotResponse) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125RecreateFromSnapshotResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125RecreateFromSnapshotResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125RecreateFromSnapshotResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
