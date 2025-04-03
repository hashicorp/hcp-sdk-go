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

// HashicorpCloudWaypointV20241122IntegrationState - UNSPECIFIED: Default state or unknown state
//   - ACTIVE: Integration is active
//   - DISCONNECTED: Integration is disconnected
//
// swagger:model hashicorp.cloud.waypoint.v20241122.Integration.State
type HashicorpCloudWaypointV20241122IntegrationState string

func NewHashicorpCloudWaypointV20241122IntegrationState(value HashicorpCloudWaypointV20241122IntegrationState) *HashicorpCloudWaypointV20241122IntegrationState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudWaypointV20241122IntegrationState.
func (m HashicorpCloudWaypointV20241122IntegrationState) Pointer() *HashicorpCloudWaypointV20241122IntegrationState {
	return &m
}

const (

	// HashicorpCloudWaypointV20241122IntegrationStateUNSPECIFIED captures enum value "UNSPECIFIED"
	HashicorpCloudWaypointV20241122IntegrationStateUNSPECIFIED HashicorpCloudWaypointV20241122IntegrationState = "UNSPECIFIED"

	// HashicorpCloudWaypointV20241122IntegrationStateACTIVE captures enum value "ACTIVE"
	HashicorpCloudWaypointV20241122IntegrationStateACTIVE HashicorpCloudWaypointV20241122IntegrationState = "ACTIVE"

	// HashicorpCloudWaypointV20241122IntegrationStateDISCONNECTED captures enum value "DISCONNECTED"
	HashicorpCloudWaypointV20241122IntegrationStateDISCONNECTED HashicorpCloudWaypointV20241122IntegrationState = "DISCONNECTED"
)

// for schema
var hashicorpCloudWaypointV20241122IntegrationStateEnum []interface{}

func init() {
	var res []HashicorpCloudWaypointV20241122IntegrationState
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","ACTIVE","DISCONNECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudWaypointV20241122IntegrationStateEnum = append(hashicorpCloudWaypointV20241122IntegrationStateEnum, v)
	}
}

func (m HashicorpCloudWaypointV20241122IntegrationState) validateHashicorpCloudWaypointV20241122IntegrationStateEnum(path, location string, value HashicorpCloudWaypointV20241122IntegrationState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudWaypointV20241122IntegrationStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud waypoint v20241122 integration state
func (m HashicorpCloudWaypointV20241122IntegrationState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudWaypointV20241122IntegrationStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint v20241122 integration state based on context it is used
func (m HashicorpCloudWaypointV20241122IntegrationState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
