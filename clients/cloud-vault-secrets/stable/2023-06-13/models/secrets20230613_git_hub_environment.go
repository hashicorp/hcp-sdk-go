// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20230613GitHubEnvironment secrets 20230613 git hub environment
//
// swagger:model secrets_20230613GitHubEnvironment
type Secrets20230613GitHubEnvironment struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this secrets 20230613 git hub environment
func (m *Secrets20230613GitHubEnvironment) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20230613 git hub environment based on context it is used
func (m *Secrets20230613GitHubEnvironment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20230613GitHubEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20230613GitHubEnvironment) UnmarshalBinary(b []byte) error {
	var res Secrets20230613GitHubEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
