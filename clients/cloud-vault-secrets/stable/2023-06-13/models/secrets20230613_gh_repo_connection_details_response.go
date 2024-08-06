// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20230613GhRepoConnectionDetailsResponse secrets 20230613 gh repo connection details response
//
// swagger:model secrets_20230613GhRepoConnectionDetailsResponse
type Secrets20230613GhRepoConnectionDetailsResponse struct {

	// environment
	Environment string `json:"environment,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// repository
	Repository string `json:"repository,omitempty"`
}

// Validate validates this secrets 20230613 gh repo connection details response
func (m *Secrets20230613GhRepoConnectionDetailsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20230613 gh repo connection details response based on context it is used
func (m *Secrets20230613GhRepoConnectionDetailsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20230613GhRepoConnectionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20230613GhRepoConnectionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20230613GhRepoConnectionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
