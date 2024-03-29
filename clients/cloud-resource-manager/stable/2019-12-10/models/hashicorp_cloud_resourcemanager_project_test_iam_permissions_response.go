// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudResourcemanagerProjectTestIamPermissionsResponse see ProjectService.TestIamPermissions
//
// swagger:model hashicorp.cloud.resourcemanager.ProjectTestIamPermissionsResponse
type HashicorpCloudResourcemanagerProjectTestIamPermissionsResponse struct {

	// AllowedPermissions are a subset of the request permissions the calling principal has for the project.
	AllowedPermissions []string `json:"allowed_permissions"`
}

// Validate validates this hashicorp cloud resourcemanager project test iam permissions response
func (m *HashicorpCloudResourcemanagerProjectTestIamPermissionsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud resourcemanager project test iam permissions response based on context it is used
func (m *HashicorpCloudResourcemanagerProjectTestIamPermissionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerProjectTestIamPermissionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerProjectTestIamPermissionsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudResourcemanagerProjectTestIamPermissionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
