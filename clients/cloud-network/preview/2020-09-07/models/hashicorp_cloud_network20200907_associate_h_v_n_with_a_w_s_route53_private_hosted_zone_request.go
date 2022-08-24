// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest See NetworkService.AssociateHVNWithAWSRoute53PrivateHostedZone for more details.
//
// swagger:model hashicorp.cloud.network_20200907.AssociateHVNWithAWSRoute53PrivateHostedZoneRequest
type HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest struct {

	// hosted_zone_id is the ID of the hosted zone in the external AWS account.
	HostedZoneID string `json:"hosted_zone_id,omitempty"`

	// hvn_id is the id of the HVN the private hosted zone should be associated with.
	HvnID string `json:"hvn_id,omitempty"`

	// location is the location of the HVN.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 associate h v n with a w s route53 private hosted zone request
func (m *HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud network 20200907 associate h v n with a w s route53 private hosted zone request based on the context it is used
func (m *HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
