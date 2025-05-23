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

// HashicorpCloudNetwork20200907DNSForwardingRuleState hashicorp cloud network 20200907 DNS forwarding rule state
//
// swagger:model hashicorp.cloud.network_20200907.DNSForwardingRule.State
type HashicorpCloudNetwork20200907DNSForwardingRuleState string

func NewHashicorpCloudNetwork20200907DNSForwardingRuleState(value HashicorpCloudNetwork20200907DNSForwardingRuleState) *HashicorpCloudNetwork20200907DNSForwardingRuleState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudNetwork20200907DNSForwardingRuleState.
func (m HashicorpCloudNetwork20200907DNSForwardingRuleState) Pointer() *HashicorpCloudNetwork20200907DNSForwardingRuleState {
	return &m
}

const (

	// HashicorpCloudNetwork20200907DNSForwardingRuleStateUNSET captures enum value "UNSET"
	HashicorpCloudNetwork20200907DNSForwardingRuleStateUNSET HashicorpCloudNetwork20200907DNSForwardingRuleState = "UNSET"

	// HashicorpCloudNetwork20200907DNSForwardingRuleStateCOMPLETE captures enum value "COMPLETE"
	HashicorpCloudNetwork20200907DNSForwardingRuleStateCOMPLETE HashicorpCloudNetwork20200907DNSForwardingRuleState = "COMPLETE"

	// HashicorpCloudNetwork20200907DNSForwardingRuleStateCREATING captures enum value "CREATING"
	HashicorpCloudNetwork20200907DNSForwardingRuleStateCREATING HashicorpCloudNetwork20200907DNSForwardingRuleState = "CREATING"

	// HashicorpCloudNetwork20200907DNSForwardingRuleStateUPDATING captures enum value "UPDATING"
	HashicorpCloudNetwork20200907DNSForwardingRuleStateUPDATING HashicorpCloudNetwork20200907DNSForwardingRuleState = "UPDATING"

	// HashicorpCloudNetwork20200907DNSForwardingRuleStateFAILED captures enum value "FAILED"
	HashicorpCloudNetwork20200907DNSForwardingRuleStateFAILED HashicorpCloudNetwork20200907DNSForwardingRuleState = "FAILED"

	// HashicorpCloudNetwork20200907DNSForwardingRuleStateDELETING captures enum value "DELETING"
	HashicorpCloudNetwork20200907DNSForwardingRuleStateDELETING HashicorpCloudNetwork20200907DNSForwardingRuleState = "DELETING"
)

// for schema
var hashicorpCloudNetwork20200907DnsForwardingRuleStateEnum []interface{}

func init() {
	var res []HashicorpCloudNetwork20200907DNSForwardingRuleState
	if err := json.Unmarshal([]byte(`["UNSET","COMPLETE","CREATING","UPDATING","FAILED","DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudNetwork20200907DnsForwardingRuleStateEnum = append(hashicorpCloudNetwork20200907DnsForwardingRuleStateEnum, v)
	}
}

func (m HashicorpCloudNetwork20200907DNSForwardingRuleState) validateHashicorpCloudNetwork20200907DNSForwardingRuleStateEnum(path, location string, value HashicorpCloudNetwork20200907DNSForwardingRuleState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudNetwork20200907DnsForwardingRuleStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud network 20200907 DNS forwarding rule state
func (m HashicorpCloudNetwork20200907DNSForwardingRuleState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudNetwork20200907DNSForwardingRuleStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud network 20200907 DNS forwarding rule state based on context it is used
func (m HashicorpCloudNetwork20200907DNSForwardingRuleState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
