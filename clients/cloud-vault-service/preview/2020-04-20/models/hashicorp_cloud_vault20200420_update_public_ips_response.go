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

// HashicorpCloudVault20200420UpdatePublicIpsResponse hashicorp cloud vault 20200420 update public ips response
//
// swagger:model hashicorp.cloud.vault_20200420.UpdatePublicIpsResponse
type HashicorpCloudVault20200420UpdatePublicIpsResponse struct {

	// operation represents the operation of the request.
	Operation *cloud.HashicorpCloudOperationOperation `json:"operation,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 update public ips response
func (m *HashicorpCloudVault20200420UpdatePublicIpsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420UpdatePublicIpsResponse) validateOperation(formats strfmt.Registry) error {

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
func (m *HashicorpCloudVault20200420UpdatePublicIpsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420UpdatePublicIpsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420UpdatePublicIpsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
