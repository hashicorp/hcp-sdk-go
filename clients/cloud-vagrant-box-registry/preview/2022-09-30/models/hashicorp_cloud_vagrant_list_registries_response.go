// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudVagrantListRegistriesResponse hashicorp cloud vagrant list registries response
//
// swagger:model hashicorp.cloud.vagrant.ListRegistriesResponse
type HashicorpCloudVagrantListRegistriesResponse struct {

	// Pagination tokens for a subsequent request.
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`

	// The returned Registry records for this page.
	Registries []*HashicorpCloudVagrantRegistry `json:"registries"`
}

// Validate validates this hashicorp cloud vagrant list registries response
func (m *HashicorpCloudVagrantListRegistriesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrantListRegistriesResponse) validatePagination(formats strfmt.Registry) error {

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

func (m *HashicorpCloudVagrantListRegistriesResponse) validateRegistries(formats strfmt.Registry) error {

	if swag.IsZero(m.Registries) { // not required
		return nil
	}

	for i := 0; i < len(m.Registries); i++ {
		if swag.IsZero(m.Registries[i]) { // not required
			continue
		}

		if m.Registries[i] != nil {
			if err := m.Registries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrantListRegistriesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrantListRegistriesResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrantListRegistriesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
