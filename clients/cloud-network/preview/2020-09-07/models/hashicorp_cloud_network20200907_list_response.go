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

// HashicorpCloudNetwork20200907ListResponse hashicorp cloud network 20200907 list response
//
// swagger:model hashicorp.cloud.network_20200907.ListResponse
type HashicorpCloudNetwork20200907ListResponse struct {

	// networks are the retrieved networks.
	Networks []*HashicorpCloudNetwork20200907Network `json:"networks"`

	// Pagination contains the pagination tokens for a subsequent request.
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 list response
func (m *HashicorpCloudNetwork20200907ListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworks(formats); err != nil {
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

func (m *HashicorpCloudNetwork20200907ListResponse) validateNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudNetwork20200907ListResponse) validatePagination(formats strfmt.Registry) error {

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
func (m *HashicorpCloudNetwork20200907ListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907ListResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907ListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
