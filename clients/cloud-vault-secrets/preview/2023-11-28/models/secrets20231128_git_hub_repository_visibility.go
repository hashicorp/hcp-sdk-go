// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Secrets20231128GitHubRepositoryVisibility secrets 20231128 git hub repository visibility
//
// swagger:model secrets_20231128GitHubRepositoryVisibility
type Secrets20231128GitHubRepositoryVisibility string

func NewSecrets20231128GitHubRepositoryVisibility(value Secrets20231128GitHubRepositoryVisibility) *Secrets20231128GitHubRepositoryVisibility {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Secrets20231128GitHubRepositoryVisibility.
func (m Secrets20231128GitHubRepositoryVisibility) Pointer() *Secrets20231128GitHubRepositoryVisibility {
	return &m
}

const (

	// Secrets20231128GitHubRepositoryVisibilityUNKNOWN captures enum value "UNKNOWN"
	Secrets20231128GitHubRepositoryVisibilityUNKNOWN Secrets20231128GitHubRepositoryVisibility = "UNKNOWN"

	// Secrets20231128GitHubRepositoryVisibilityPUBLIC captures enum value "PUBLIC"
	Secrets20231128GitHubRepositoryVisibilityPUBLIC Secrets20231128GitHubRepositoryVisibility = "PUBLIC"

	// Secrets20231128GitHubRepositoryVisibilityPRIVATE captures enum value "PRIVATE"
	Secrets20231128GitHubRepositoryVisibilityPRIVATE Secrets20231128GitHubRepositoryVisibility = "PRIVATE"
)

// for schema
var secrets20231128GitHubRepositoryVisibilityEnum []interface{}

func init() {
	var res []Secrets20231128GitHubRepositoryVisibility
	if err := json.Unmarshal([]byte(`["UNKNOWN","PUBLIC","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		secrets20231128GitHubRepositoryVisibilityEnum = append(secrets20231128GitHubRepositoryVisibilityEnum, v)
	}
}

func (m Secrets20231128GitHubRepositoryVisibility) validateSecrets20231128GitHubRepositoryVisibilityEnum(path, location string, value Secrets20231128GitHubRepositoryVisibility) error {
	if err := validate.EnumCase(path, location, value, secrets20231128GitHubRepositoryVisibilityEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this secrets 20231128 git hub repository visibility
func (m Secrets20231128GitHubRepositoryVisibility) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSecrets20231128GitHubRepositoryVisibilityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this secrets 20231128 git hub repository visibility based on context it is used
func (m Secrets20231128GitHubRepositoryVisibility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
