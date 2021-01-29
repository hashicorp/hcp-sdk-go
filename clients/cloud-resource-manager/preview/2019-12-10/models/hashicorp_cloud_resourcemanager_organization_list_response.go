// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/cloud-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudResourcemanagerOrganizationListResponse see OrganizationService.List
//
// swagger:model hashicorp.cloud.resourcemanager.OrganizationListResponse
type HashicorpCloudResourcemanagerOrganizationListResponse struct {

	// Organizations is the organizations the caller has access to.
	Organizations []*HashicorpCloudResourcemanagerOrganization `json:"organizations"`

	// Pagination contains the pagination tokens for a subsequent request.
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`
}

// Validate validates this hashicorp cloud resourcemanager organization list response
func (m *HashicorpCloudResourcemanagerOrganizationListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudResourcemanagerOrganizationListResponse) validateOrganizations(formats strfmt.Registry) error {

	if swag.IsZero(m.Organizations) { // not required
		return nil
	}

	for i := 0; i < len(m.Organizations); i++ {
		if swag.IsZero(m.Organizations[i]) { // not required
			continue
		}

		if m.Organizations[i] != nil {
			if err := m.Organizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudResourcemanagerOrganizationListResponse) validatePagination(formats strfmt.Registry) error {

	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerOrganizationListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerOrganizationListResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudResourcemanagerOrganizationListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
