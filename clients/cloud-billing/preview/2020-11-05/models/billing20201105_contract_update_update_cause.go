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

// Billing20201105ContractUpdateUpdateCause update_cause defines possible reasons for the update to the contract.
//
// swagger:model billing_20201105ContractUpdateUpdateCause
type Billing20201105ContractUpdateUpdateCause string

func NewBilling20201105ContractUpdateUpdateCause(value Billing20201105ContractUpdateUpdateCause) *Billing20201105ContractUpdateUpdateCause {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Billing20201105ContractUpdateUpdateCause.
func (m Billing20201105ContractUpdateUpdateCause) Pointer() *Billing20201105ContractUpdateUpdateCause {
	return &m
}

const (

	// Billing20201105ContractUpdateUpdateCauseADDON captures enum value "ADDON"
	Billing20201105ContractUpdateUpdateCauseADDON Billing20201105ContractUpdateUpdateCause = "ADDON"

	// Billing20201105ContractUpdateUpdateCauseEARLYRENEWAL captures enum value "EARLY_RENEWAL"
	Billing20201105ContractUpdateUpdateCauseEARLYRENEWAL Billing20201105ContractUpdateUpdateCause = "EARLY_RENEWAL"
)

// for schema
var billing20201105ContractUpdateUpdateCauseEnum []interface{}

func init() {
	var res []Billing20201105ContractUpdateUpdateCause
	if err := json.Unmarshal([]byte(`["ADDON","EARLY_RENEWAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billing20201105ContractUpdateUpdateCauseEnum = append(billing20201105ContractUpdateUpdateCauseEnum, v)
	}
}

func (m Billing20201105ContractUpdateUpdateCause) validateBilling20201105ContractUpdateUpdateCauseEnum(path, location string, value Billing20201105ContractUpdateUpdateCause) error {
	if err := validate.EnumCase(path, location, value, billing20201105ContractUpdateUpdateCauseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this billing 20201105 contract update update cause
func (m Billing20201105ContractUpdateUpdateCause) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBilling20201105ContractUpdateUpdateCauseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this billing 20201105 contract update update cause based on context it is used
func (m Billing20201105ContractUpdateUpdateCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
