// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudGlobalNetworkManager20220215ACLInfo hashicorp cloud global network manager 20220215 ACL info
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.ACLInfo
type HashicorpCloudGlobalNetworkManager20220215ACLInfo struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 ACL info
func (m *HashicorpCloudGlobalNetworkManager20220215ACLInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud global network manager 20220215 ACL info based on context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215ACLInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ACLInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ACLInfo) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215ACLInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
