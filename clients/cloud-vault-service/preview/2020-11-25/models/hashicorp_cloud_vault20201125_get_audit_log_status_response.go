// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125GetAuditLogStatusResponse hashicorp cloud vault 20201125 get audit log status response
//
// swagger:model hashicorp.cloud.vault_20201125.GetAuditLogStatusResponse
type HashicorpCloudVault20201125GetAuditLogStatusResponse struct {

	// log
	Log *HashicorpCloudVault20201125AuditLog `json:"log,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 get audit log status response
func (m *HashicorpCloudVault20201125GetAuditLogStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125GetAuditLogStatusResponse) validateLog(formats strfmt.Registry) error {

	if swag.IsZero(m.Log) { // not required
		return nil
	}

	if m.Log != nil {
		if err := m.Log.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetAuditLogStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetAuditLogStatusResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125GetAuditLogStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
