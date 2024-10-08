// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128IntegrationUsage secrets 20231128 integration usage
//
// swagger:model secrets_20231128IntegrationUsage
type Secrets20231128IntegrationUsage struct {

	// apps count
	AppsCount int32 `json:"apps_count,omitempty"`

	// secrets count
	SecretsCount int32 `json:"secrets_count,omitempty"`

	// secrets count by app
	SecretsCountByApp map[string]int32 `json:"secrets_count_by_app,omitempty"`
}

// Validate validates this secrets 20231128 integration usage
func (m *Secrets20231128IntegrationUsage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 integration usage based on context it is used
func (m *Secrets20231128IntegrationUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128IntegrationUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128IntegrationUsage) UnmarshalBinary(b []byte) error {
	var res Secrets20231128IntegrationUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
