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

// HashicorpCloudIamOrganizationInvitationState State is one of the states that an organization invitation can be in. The public and private
// values for State do not map onto each other in a 1-to-1 way.
//
//   - PENDING: PENDING is the state of an organization invitation that has been created successfully but
//
// has not been accepted yet.
//   - ACCEPTED: ACCEPTED is the state of an organization invitation that has been accepted.
//
// swagger:model hashicorp.cloud.iam.OrganizationInvitation.State
type HashicorpCloudIamOrganizationInvitationState string

func NewHashicorpCloudIamOrganizationInvitationState(value HashicorpCloudIamOrganizationInvitationState) *HashicorpCloudIamOrganizationInvitationState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudIamOrganizationInvitationState.
func (m HashicorpCloudIamOrganizationInvitationState) Pointer() *HashicorpCloudIamOrganizationInvitationState {
	return &m
}

const (

	// HashicorpCloudIamOrganizationInvitationStatePENDING captures enum value "PENDING"
	HashicorpCloudIamOrganizationInvitationStatePENDING HashicorpCloudIamOrganizationInvitationState = "PENDING"

	// HashicorpCloudIamOrganizationInvitationStateACCEPTED captures enum value "ACCEPTED"
	HashicorpCloudIamOrganizationInvitationStateACCEPTED HashicorpCloudIamOrganizationInvitationState = "ACCEPTED"
)

// for schema
var hashicorpCloudIamOrganizationInvitationStateEnum []interface{}

func init() {
	var res []HashicorpCloudIamOrganizationInvitationState
	if err := json.Unmarshal([]byte(`["PENDING","ACCEPTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudIamOrganizationInvitationStateEnum = append(hashicorpCloudIamOrganizationInvitationStateEnum, v)
	}
}

func (m HashicorpCloudIamOrganizationInvitationState) validateHashicorpCloudIamOrganizationInvitationStateEnum(path, location string, value HashicorpCloudIamOrganizationInvitationState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudIamOrganizationInvitationStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud iam organization invitation state
func (m HashicorpCloudIamOrganizationInvitationState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudIamOrganizationInvitationStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud iam organization invitation state based on context it is used
func (m HashicorpCloudIamOrganizationInvitationState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}