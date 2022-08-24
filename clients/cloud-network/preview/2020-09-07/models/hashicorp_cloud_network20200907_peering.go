// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudNetwork20200907Peering Peering represents a network peering between HVN VPC and target VPC
//
// swagger:model hashicorp.cloud.network_20200907.Peering
type HashicorpCloudNetwork20200907Peering struct {

	// CreatedAt is a timestamp when peering was originally created
	//
	// Output only.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// ExpiresAt is the timestamp after which peering will be considered
	// Expired if it doesn't transition into Accepted or Active state before that
	//
	// Output only.
	// Read Only: true
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty"`

	// Hvn is the a link to the HVN where peering belongs to.
	Hvn *cloud.HashicorpCloudLocationLink `json:"hvn,omitempty"`

	// Id is the id of a peering.
	ID string `json:"id,omitempty"`

	// ProviderPeeringId is the id of this peering in the cloud provider
	//
	// Output only.
	// Read Only: true
	ProviderPeeringID string `json:"provider_peering_id,omitempty"`

	// State is the state of a peering
	//
	// Output only.
	// Read Only: true
	State *HashicorpCloudNetwork20200907PeeringState `json:"state,omitempty"`

	// Target is a provider specific details about peering target
	Target *HashicorpCloudNetwork20200907PeeringTarget `json:"target,omitempty"`

	// UpdatedAt is a timestamp when peering was originally created
	//
	// Output only.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 peering
func (m *HashicorpCloudNetwork20200907Peering) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHvn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) validateExpiresAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) validateHvn(formats strfmt.Registry) error {
	if swag.IsZero(m.Hvn) { // not required
		return nil
	}

	if m.Hvn != nil {
		if err := m.Hvn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hvn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hvn")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hashicorp cloud network 20200907 peering based on the context it is used
func (m *HashicorpCloudNetwork20200907Peering) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpiresAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHvn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProviderPeeringID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) contextValidateExpiresAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expires_at", "body", strfmt.DateTime(m.ExpiresAt)); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) contextValidateHvn(ctx context.Context, formats strfmt.Registry) error {

	if m.Hvn != nil {
		if err := m.Hvn.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hvn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hvn")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) contextValidateProviderPeeringID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "provider_peering_id", "body", string(m.ProviderPeeringID)); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {
		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Peering) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907Peering) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907Peering) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907Peering
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
