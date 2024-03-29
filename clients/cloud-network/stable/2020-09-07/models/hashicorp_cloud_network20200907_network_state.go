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

// HashicorpCloudNetwork20200907NetworkState State is the state of the network resource. Note that this state
// represents the abstract network itself, not necessarilly whether
// network connectivity is currently available or not.
//
// swagger:model hashicorp.cloud.network_20200907.Network.State
type HashicorpCloudNetwork20200907NetworkState string

func NewHashicorpCloudNetwork20200907NetworkState(value HashicorpCloudNetwork20200907NetworkState) *HashicorpCloudNetwork20200907NetworkState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudNetwork20200907NetworkState.
func (m HashicorpCloudNetwork20200907NetworkState) Pointer() *HashicorpCloudNetwork20200907NetworkState {
	return &m
}

const (

	// HashicorpCloudNetwork20200907NetworkStateUNSET captures enum value "UNSET"
	HashicorpCloudNetwork20200907NetworkStateUNSET HashicorpCloudNetwork20200907NetworkState = "UNSET"

	// HashicorpCloudNetwork20200907NetworkStateCREATING captures enum value "CREATING"
	HashicorpCloudNetwork20200907NetworkStateCREATING HashicorpCloudNetwork20200907NetworkState = "CREATING"

	// HashicorpCloudNetwork20200907NetworkStateSTABLE captures enum value "STABLE"
	HashicorpCloudNetwork20200907NetworkStateSTABLE HashicorpCloudNetwork20200907NetworkState = "STABLE"

	// HashicorpCloudNetwork20200907NetworkStateFAILED captures enum value "FAILED"
	HashicorpCloudNetwork20200907NetworkStateFAILED HashicorpCloudNetwork20200907NetworkState = "FAILED"

	// HashicorpCloudNetwork20200907NetworkStateDELETING captures enum value "DELETING"
	HashicorpCloudNetwork20200907NetworkStateDELETING HashicorpCloudNetwork20200907NetworkState = "DELETING"

	// HashicorpCloudNetwork20200907NetworkStateDELETED captures enum value "DELETED"
	HashicorpCloudNetwork20200907NetworkStateDELETED HashicorpCloudNetwork20200907NetworkState = "DELETED"
)

// for schema
var hashicorpCloudNetwork20200907NetworkStateEnum []interface{}

func init() {
	var res []HashicorpCloudNetwork20200907NetworkState
	if err := json.Unmarshal([]byte(`["UNSET","CREATING","STABLE","FAILED","DELETING","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudNetwork20200907NetworkStateEnum = append(hashicorpCloudNetwork20200907NetworkStateEnum, v)
	}
}

func (m HashicorpCloudNetwork20200907NetworkState) validateHashicorpCloudNetwork20200907NetworkStateEnum(path, location string, value HashicorpCloudNetwork20200907NetworkState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudNetwork20200907NetworkStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud network 20200907 network state
func (m HashicorpCloudNetwork20200907NetworkState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudNetwork20200907NetworkStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud network 20200907 network state based on context it is used
func (m HashicorpCloudNetwork20200907NetworkState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
