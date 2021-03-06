// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudResourcemanagerOrganizationSetIamPolicyResponse see OrganizationService.SetIamPolicy
//
// swagger:model hashicorp.cloud.resourcemanager.OrganizationSetIamPolicyResponse
type HashicorpCloudResourcemanagerOrganizationSetIamPolicyResponse struct {

	// Policy is the updated IAM policy for the organization.
	Policy *HashicorpCloudResourcemanagerPolicy `json:"policy,omitempty"`
}

// Validate validates this hashicorp cloud resourcemanager organization set iam policy response
func (m *HashicorpCloudResourcemanagerOrganizationSetIamPolicyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudResourcemanagerOrganizationSetIamPolicyResponse) validatePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerOrganizationSetIamPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerOrganizationSetIamPolicyResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudResourcemanagerOrganizationSetIamPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
