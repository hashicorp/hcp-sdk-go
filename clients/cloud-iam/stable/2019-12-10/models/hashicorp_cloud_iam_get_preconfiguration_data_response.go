// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamGetPreconfigurationDataResponse GetPreconfigurationDataResponse is the response containing the
// preconfiguration data.
//
// swagger:model hashicorp.cloud.iam.GetPreconfigurationDataResponse
type HashicorpCloudIamGetPreconfigurationDataResponse struct {

	// AssertionConsumerUrl has the Assertion Consumer URL (or Post Back in
	// Auth0 parlance) for the IdP.
	AssertionConsumerURL string `json:"assertion_consumer_url,omitempty"`

	// domain_txt_verification_record is the expected value of the TXT ownership
	// verification record we require organizations to add to their email
	// domain's DNS to prove that they own/control the domain.
	DomainTxtVerificationRecord string `json:"domain_txt_verification_record,omitempty"`

	// EntityUrl has the name of the audience the IdP needs to set up on their
	// side.
	EntityID string `json:"entity_id,omitempty"`
}

// Validate validates this hashicorp cloud iam get preconfiguration data response
func (m *HashicorpCloudIamGetPreconfigurationDataResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam get preconfiguration data response based on context it is used
func (m *HashicorpCloudIamGetPreconfigurationDataResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamGetPreconfigurationDataResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamGetPreconfigurationDataResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamGetPreconfigurationDataResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
