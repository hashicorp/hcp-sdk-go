// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsul20210204CreateCustomerMasterACLTokenResponse CreateCustomerMasterACLTokenResponse contains the newly created customer master ACL token.
//
// swagger:model hashicorp.cloud.consul_20210204.CreateCustomerMasterACLTokenResponse
type HashicorpCloudConsul20210204CreateCustomerMasterACLTokenResponse struct {

	// acl_token is the newly created master token.
	ACLToken *HashicorpCloudConsul20210204ACLToken `json:"acl_token,omitempty"`
}

// Validate validates this hashicorp cloud consul 20210204 create customer master ACL token response
func (m *HashicorpCloudConsul20210204CreateCustomerMasterACLTokenResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateACLToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20210204CreateCustomerMasterACLTokenResponse) validateACLToken(formats strfmt.Registry) error {

	if swag.IsZero(m.ACLToken) { // not required
		return nil
	}

	if m.ACLToken != nil {
		if err := m.ACLToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acl_token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20210204CreateCustomerMasterACLTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20210204CreateCustomerMasterACLTokenResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20210204CreateCustomerMasterACLTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
