// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VaultRadar20230501CreateIntegrationSubscriptionResponse vault radar 20230501 create integration subscription response
//
// swagger:model vault_radar_20230501CreateIntegrationSubscriptionResponse
type VaultRadar20230501CreateIntegrationSubscriptionResponse struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this vault radar 20230501 create integration subscription response
func (m *VaultRadar20230501CreateIntegrationSubscriptionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vault radar 20230501 create integration subscription response based on context it is used
func (m *VaultRadar20230501CreateIntegrationSubscriptionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VaultRadar20230501CreateIntegrationSubscriptionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VaultRadar20230501CreateIntegrationSubscriptionResponse) UnmarshalBinary(b []byte) error {
	var res VaultRadar20230501CreateIntegrationSubscriptionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
