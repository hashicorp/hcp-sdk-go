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

// HashicorpCloudVault20200420ListSnapshotsResponse ListSnapshotsResponse is a response from listing snapshots.
//
// swagger:model hashicorp.cloud.vault_20200420.ListSnapshotsResponse
type HashicorpCloudVault20200420ListSnapshotsResponse struct {

	// Pagination contains the pagination tokens for a subsequent request.
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`

	// snapshots is a list of available snapshots.
	Snapshots []*HashicorpCloudVault20200420Snapshot `json:"snapshots"`
}

// Validate validates this hashicorp cloud vault 20200420 list snapshots response
func (m *HashicorpCloudVault20200420ListSnapshotsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420ListSnapshotsResponse) validatePagination(formats strfmt.Registry) error {
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

func (m *HashicorpCloudVault20200420ListSnapshotsResponse) validateSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.Snapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.Snapshots); i++ {
		if swag.IsZero(m.Snapshots[i]) { // not required
			continue
		}

		if m.Snapshots[i] != nil {
			if err := m.Snapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20200420 list snapshots response based on the context it is used
func (m *HashicorpCloudVault20200420ListSnapshotsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420ListSnapshotsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudVault20200420ListSnapshotsResponse) contextValidateSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Snapshots); i++ {

		if m.Snapshots[i] != nil {

			if swag.IsZero(m.Snapshots[i]) { // not required
				return nil
			}

			if err := m.Snapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420ListSnapshotsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420ListSnapshotsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420ListSnapshotsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
