// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsul20210204GetClientConfigResponse GetClientConfigResponse the client config files for a Consul agent.
//
// swagger:model hashicorp.cloud.consul_20210204.GetClientConfigResponse
type HashicorpCloudConsul20210204GetClientConfigResponse struct {

	// ca_file is the body of an intermediate certificate authority file.
	// Format: byte
	CaFile strfmt.Base64 `json:"ca_file,omitempty"`

	// consul_config_file is the body of JSON config file.
	// Format: byte
	ConsulConfigFile strfmt.Base64 `json:"consul_config_file,omitempty"`

	// file_bundle_zip is the contents of a zip archive containing the other
	// two files, if the request specified it.
	// Format: byte
	FileBundleZip strfmt.Base64 `json:"file_bundle_zip,omitempty"`
}

// Validate validates this hashicorp cloud consul 20210204 get client config response
func (m *HashicorpCloudConsul20210204GetClientConfigResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud consul 20210204 get client config response based on context it is used
func (m *HashicorpCloudConsul20210204GetClientConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20210204GetClientConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20210204GetClientConfigResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20210204GetClientConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
