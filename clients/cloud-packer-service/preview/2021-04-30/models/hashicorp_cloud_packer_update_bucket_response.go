// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPackerUpdateBucketResponse hashicorp cloud packer update bucket response
//
// swagger:model hashicorp.cloud.packer.UpdateBucketResponse
type HashicorpCloudPackerUpdateBucketResponse struct {

	// bucket
	Bucket *HashicorpCloudPackerBucket `json:"bucket,omitempty"`
}

// Validate validates this hashicorp cloud packer update bucket response
func (m *HashicorpCloudPackerUpdateBucketResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerUpdateBucketResponse) validateBucket(formats strfmt.Registry) error {
	if swag.IsZero(m.Bucket) { // not required
		return nil
	}

	if m.Bucket != nil {
		if err := m.Bucket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer update bucket response based on the context it is used
func (m *HashicorpCloudPackerUpdateBucketResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBucket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerUpdateBucketResponse) contextValidateBucket(ctx context.Context, formats strfmt.Registry) error {

	if m.Bucket != nil {
		if err := m.Bucket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerUpdateBucketResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerUpdateBucketResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerUpdateBucketResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
