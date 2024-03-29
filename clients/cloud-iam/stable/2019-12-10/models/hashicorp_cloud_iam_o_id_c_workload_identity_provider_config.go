// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamOIDCWorkloadIdentityProviderConfig OIDCWorkloadIdentityProviderConfig configures an OIDC Workload Identity
// Provider.
//
// swagger:model hashicorp.cloud.iam.OIDCWorkloadIdentityProviderConfig
type HashicorpCloudIamOIDCWorkloadIdentityProviderConfig struct {

	// allowed_audiences is the set of audiences set on the access token that
	// are allowed to exchange identities. The access token must have an
	// audience that is contained in this set. If no audience is set, the
	// default allowed audience will be the resource name of the
	// WorkloadIdentityProvider.
	AllowedAudiences []string `json:"allowed_audiences"`

	// issuer_uri is the OIDC issuer URI
	IssuerURI string `json:"issuer_uri,omitempty"`
}

// Validate validates this hashicorp cloud iam o ID c workload identity provider config
func (m *HashicorpCloudIamOIDCWorkloadIdentityProviderConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam o ID c workload identity provider config based on context it is used
func (m *HashicorpCloudIamOIDCWorkloadIdentityProviderConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamOIDCWorkloadIdentityProviderConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamOIDCWorkloadIdentityProviderConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamOIDCWorkloadIdentityProviderConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
