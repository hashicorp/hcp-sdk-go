// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamExchangeWorkloadIdentityTokenResponse ExchangeWorkloadIdentityTokenResponse is the response to the exchange of an
// external token for a service principal access_token.
//
// swagger:model hashicorp.cloud.iam.ExchangeWorkloadIdentityTokenResponse
type HashicorpCloudIamExchangeWorkloadIdentityTokenResponse struct {

	// access_token is an access_token with the subject matching that of the
	// requested service principal.
	AccessToken string `json:"access_token,omitempty"`

	// expires_in is the number of seconds the returned access token expires in.
	ExpiresIn int32 `json:"expires_in,omitempty"`
}

// Validate validates this hashicorp cloud iam exchange workload identity token response
func (m *HashicorpCloudIamExchangeWorkloadIdentityTokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam exchange workload identity token response based on context it is used
func (m *HashicorpCloudIamExchangeWorkloadIdentityTokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamExchangeWorkloadIdentityTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamExchangeWorkloadIdentityTokenResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamExchangeWorkloadIdentityTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
