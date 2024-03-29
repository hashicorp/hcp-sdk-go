// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudNetwork20200907AWSPeeringTarget PeeringTarget is AWS specific details about a target for network peering
//
// swagger:model hashicorp.cloud.network_20200907.AWS.PeeringTarget
type HashicorpCloudNetwork20200907AWSPeeringTarget struct {

	// AccountId is the Account ID where target ID belongs to
	AccountID string `json:"account_id,omitempty"`

	// Cidr is the IPv4 CIDR block of the target VPC
	Cidr string `json:"cidr,omitempty"`

	// Region is the region of the target VPC
	Region string `json:"region,omitempty"`

	// VpcId is the id of a target VPC
	VpcID string `json:"vpc_id,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 a w s peering target
func (m *HashicorpCloudNetwork20200907AWSPeeringTarget) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud network 20200907 a w s peering target based on context it is used
func (m *HashicorpCloudNetwork20200907AWSPeeringTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907AWSPeeringTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907AWSPeeringTarget) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907AWSPeeringTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
