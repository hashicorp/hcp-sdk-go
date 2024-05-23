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

// Billing20201105ContractStatus Status represents the status of a given contract based on its active from and until dates.
//
// swagger:model billing_20201105ContractStatus
type Billing20201105ContractStatus string

func NewBilling20201105ContractStatus(value Billing20201105ContractStatus) *Billing20201105ContractStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Billing20201105ContractStatus.
func (m Billing20201105ContractStatus) Pointer() *Billing20201105ContractStatus {
	return &m
}

const (

	// Billing20201105ContractStatusUNSET captures enum value "UNSET"
	Billing20201105ContractStatusUNSET Billing20201105ContractStatus = "UNSET"

	// Billing20201105ContractStatusACTIVE captures enum value "ACTIVE"
	Billing20201105ContractStatusACTIVE Billing20201105ContractStatus = "ACTIVE"

	// Billing20201105ContractStatusUPCOMING captures enum value "UPCOMING"
	Billing20201105ContractStatusUPCOMING Billing20201105ContractStatus = "UPCOMING"

	// Billing20201105ContractStatusEXPIRED captures enum value "EXPIRED"
	Billing20201105ContractStatusEXPIRED Billing20201105ContractStatus = "EXPIRED"
)

// for schema
var billing20201105ContractStatusEnum []interface{}

func init() {
	var res []Billing20201105ContractStatus
	if err := json.Unmarshal([]byte(`["UNSET","ACTIVE","UPCOMING","EXPIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billing20201105ContractStatusEnum = append(billing20201105ContractStatusEnum, v)
	}
}

func (m Billing20201105ContractStatus) validateBilling20201105ContractStatusEnum(path, location string, value Billing20201105ContractStatus) error {
	if err := validate.EnumCase(path, location, value, billing20201105ContractStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this billing 20201105 contract status
func (m Billing20201105ContractStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBilling20201105ContractStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this billing 20201105 contract status based on context it is used
func (m Billing20201105ContractStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
