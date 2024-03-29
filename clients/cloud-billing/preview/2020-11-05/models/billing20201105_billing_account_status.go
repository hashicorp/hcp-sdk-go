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

// Billing20201105BillingAccountStatus BillingAccountStatus is the status of the billing account (e.g. active,
// canceled).
//
// swagger:model billing_20201105BillingAccountStatus
type Billing20201105BillingAccountStatus string

func NewBilling20201105BillingAccountStatus(value Billing20201105BillingAccountStatus) *Billing20201105BillingAccountStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Billing20201105BillingAccountStatus.
func (m Billing20201105BillingAccountStatus) Pointer() *Billing20201105BillingAccountStatus {
	return &m
}

const (

	// Billing20201105BillingAccountStatusACTIVE captures enum value "ACTIVE"
	Billing20201105BillingAccountStatusACTIVE Billing20201105BillingAccountStatus = "ACTIVE"

	// Billing20201105BillingAccountStatusDELETING captures enum value "DELETING"
	Billing20201105BillingAccountStatusDELETING Billing20201105BillingAccountStatus = "DELETING"

	// Billing20201105BillingAccountStatusDELINQUENT captures enum value "DELINQUENT"
	Billing20201105BillingAccountStatusDELINQUENT Billing20201105BillingAccountStatus = "DELINQUENT"

	// Billing20201105BillingAccountStatusINCOMPLETE captures enum value "INCOMPLETE"
	Billing20201105BillingAccountStatusINCOMPLETE Billing20201105BillingAccountStatus = "INCOMPLETE"

	// Billing20201105BillingAccountStatusTRIAL captures enum value "TRIAL"
	Billing20201105BillingAccountStatusTRIAL Billing20201105BillingAccountStatus = "TRIAL"
)

// for schema
var billing20201105BillingAccountStatusEnum []interface{}

func init() {
	var res []Billing20201105BillingAccountStatus
	if err := json.Unmarshal([]byte(`["ACTIVE","DELETING","DELINQUENT","INCOMPLETE","TRIAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billing20201105BillingAccountStatusEnum = append(billing20201105BillingAccountStatusEnum, v)
	}
}

func (m Billing20201105BillingAccountStatus) validateBilling20201105BillingAccountStatusEnum(path, location string, value Billing20201105BillingAccountStatus) error {
	if err := validate.EnumCase(path, location, value, billing20201105BillingAccountStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this billing 20201105 billing account status
func (m Billing20201105BillingAccountStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBilling20201105BillingAccountStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this billing 20201105 billing account status based on context it is used
func (m Billing20201105BillingAccountStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
