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

// HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason PartitionPeeringIneligibilityReason indicates a reason why a peering connection is not possible
// with a particular partition
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.PartitionPeeringIneligibilityReason
type HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason string

func NewHashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason(value HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason) *HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason.
func (m HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason) Pointer() *HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason {
	return &m
}

const (

	// HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReasonPARTITIONINELIGIBILITYUNDEFINED captures enum value "PARTITION_INELIGIBILITY_UNDEFINED"
	HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReasonPARTITIONINELIGIBILITYUNDEFINED HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason = "PARTITION_INELIGIBILITY_UNDEFINED"

	// HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReasonPARTITIONINELIGIBILITYPEERINGALREADYEXISTS captures enum value "PARTITION_INELIGIBILITY_PEERING_ALREADY_EXISTS"
	HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReasonPARTITIONINELIGIBILITYPEERINGALREADYEXISTS HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason = "PARTITION_INELIGIBILITY_PEERING_ALREADY_EXISTS"
)

// for schema
var hashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReasonEnum []interface{}

func init() {
	var res []HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason
	if err := json.Unmarshal([]byte(`["PARTITION_INELIGIBILITY_UNDEFINED","PARTITION_INELIGIBILITY_PEERING_ALREADY_EXISTS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReasonEnum = append(hashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReasonEnum, v)
	}
}

func (m HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason) validateHashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReasonEnum(path, location string, value HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReasonEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud global network manager 20220215 partition peering ineligibility reason
func (m HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud global network manager 20220215 partition peering ineligibility reason based on context it is used
func (m HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
