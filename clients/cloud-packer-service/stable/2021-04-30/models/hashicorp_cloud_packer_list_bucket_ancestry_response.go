// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudPackerListBucketAncestryResponse hashicorp cloud packer list bucket ancestry response
//
// swagger:model hashicorp.cloud.packer.ListBucketAncestryResponse
type HashicorpCloudPackerListBucketAncestryResponse struct {

	// pagination
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`

	// relations
	Relations []*HashicorpCloudPackerBucketAncestry `json:"relations"`

	// The total number of ancestral relationships returned for the specified image bucket. These can be parent or child images.
	TotalCount int32 `json:"total_count,omitempty"`
}

// Validate validates this hashicorp cloud packer list bucket ancestry response
func (m *HashicorpCloudPackerListBucketAncestryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerListBucketAncestryResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerListBucketAncestryResponse) validateRelations(formats strfmt.Registry) error {
	if swag.IsZero(m.Relations) { // not required
		return nil
	}

	for i := 0; i < len(m.Relations); i++ {
		if swag.IsZero(m.Relations[i]) { // not required
			continue
		}

		if m.Relations[i] != nil {
			if err := m.Relations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer list bucket ancestry response based on the context it is used
func (m *HashicorpCloudPackerListBucketAncestryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerListBucketAncestryResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {

		if swag.IsZero(m.Pagination) { // not required
			return nil
		}

		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerListBucketAncestryResponse) contextValidateRelations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Relations); i++ {

		if m.Relations[i] != nil {

			if swag.IsZero(m.Relations[i]) { // not required
				return nil
			}

			if err := m.Relations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerListBucketAncestryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerListBucketAncestryResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerListBucketAncestryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
