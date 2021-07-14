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

// HashicorpCloudPackerBucket hashicorp cloud packer bucket
//
// swagger:model hashicorp.cloud.packer.Bucket
type HashicorpCloudPackerBucket struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// "This image is a hardened platform for other teams"
	Description string `json:"description,omitempty"`

	// ULID
	ID string `json:"id,omitempty"`

	// number of total iterations (not just versions), for UI.
	IterationCount string `json:"iteration_count,omitempty"`

	// Slot for unstructured metadata tags
	// for example {"repo": "https://github.com/hashicorp/common"}
	Labels map[string]string `json:"labels,omitempty"`

	// latest iteration regardless of completion status
	LatestIteration *HashicorpCloudPackerIteration `json:"latest_iteration,omitempty"`

	// latest completed version
	LatestVersion int32 `json:"latest_version,omitempty"`

	// location
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// ["aws", "gcp"]
	Platforms []string `json:"platforms"`

	// base-ubuntu-18-secure
	Slug string `json:"slug,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this hashicorp cloud packer bucket
func (m *HashicorpCloudPackerBucket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestIteration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
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

func (m *HashicorpCloudPackerBucket) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPackerBucket) validateLatestIteration(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestIteration) { // not required
		return nil
	}

	if m.LatestIteration != nil {
		if err := m.LatestIteration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latest_iteration")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerBucket) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerBucket) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer bucket based on the context it is used
func (m *HashicorpCloudPackerBucket) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLatestIteration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerBucket) contextValidateLatestIteration(ctx context.Context, formats strfmt.Registry) error {

	if m.LatestIteration != nil {
		if err := m.LatestIteration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latest_iteration")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerBucket) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerBucket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerBucket) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerBucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
