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

// HashicorpCloudVagrantHostedProviderState State is an enumeration of possible HostedProvider states.
//
// swagger:model hashicorp.cloud.vagrant.HostedProvider.State
type HashicorpCloudVagrantHostedProviderState string

func NewHashicorpCloudVagrantHostedProviderState(value HashicorpCloudVagrantHostedProviderState) *HashicorpCloudVagrantHostedProviderState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudVagrantHostedProviderState.
func (m HashicorpCloudVagrantHostedProviderState) Pointer() *HashicorpCloudVagrantHostedProviderState {
	return &m
}

const (

	// HashicorpCloudVagrantHostedProviderStatePENDING captures enum value "PENDING"
	HashicorpCloudVagrantHostedProviderStatePENDING HashicorpCloudVagrantHostedProviderState = "PENDING"

	// HashicorpCloudVagrantHostedProviderStateUPLOADING captures enum value "UPLOADING"
	HashicorpCloudVagrantHostedProviderStateUPLOADING HashicorpCloudVagrantHostedProviderState = "UPLOADING"

	// HashicorpCloudVagrantHostedProviderStateCOMPLETE captures enum value "COMPLETE"
	HashicorpCloudVagrantHostedProviderStateCOMPLETE HashicorpCloudVagrantHostedProviderState = "COMPLETE"

	// HashicorpCloudVagrantHostedProviderStateFAILED captures enum value "FAILED"
	HashicorpCloudVagrantHostedProviderStateFAILED HashicorpCloudVagrantHostedProviderState = "FAILED"
)

// for schema
var hashicorpCloudVagrantHostedProviderStateEnum []interface{}

func init() {
	var res []HashicorpCloudVagrantHostedProviderState
	if err := json.Unmarshal([]byte(`["PENDING","UPLOADING","COMPLETE","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVagrantHostedProviderStateEnum = append(hashicorpCloudVagrantHostedProviderStateEnum, v)
	}
}

func (m HashicorpCloudVagrantHostedProviderState) validateHashicorpCloudVagrantHostedProviderStateEnum(path, location string, value HashicorpCloudVagrantHostedProviderState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVagrantHostedProviderStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vagrant hosted provider state
func (m HashicorpCloudVagrantHostedProviderState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVagrantHostedProviderStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud vagrant hosted provider state based on context it is used
func (m HashicorpCloudVagrantHostedProviderState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
