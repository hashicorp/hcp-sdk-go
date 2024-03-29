// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128GhAppMetadata secrets 20231128 gh app metadata
//
// swagger:model secrets_20231128GhAppMetadata
type Secrets20231128GhAppMetadata struct {

	// installation id
	InstallationID string `json:"installation_id,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// repositories
	Repositories []string `json:"repositories"`
}

// Validate validates this secrets 20231128 gh app metadata
func (m *Secrets20231128GhAppMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 gh app metadata based on context it is used
func (m *Secrets20231128GhAppMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128GhAppMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128GhAppMetadata) UnmarshalBinary(b []byte) error {
	var res Secrets20231128GhAppMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
