// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BulkCreateAppKVSecretsRequestKeyValue bulk create app k v secrets request key value
//
// swagger:model BulkCreateAppKVSecretsRequestKeyValue
type BulkCreateAppKVSecretsRequestKeyValue struct {

	// name
	Name string `json:"name,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this bulk create app k v secrets request key value
func (m *BulkCreateAppKVSecretsRequestKeyValue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this bulk create app k v secrets request key value based on context it is used
func (m *BulkCreateAppKVSecretsRequestKeyValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BulkCreateAppKVSecretsRequestKeyValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkCreateAppKVSecretsRequestKeyValue) UnmarshalBinary(b []byte) error {
	var res BulkCreateAppKVSecretsRequestKeyValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
