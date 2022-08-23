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

// HashicorpCloudVagrant20220930ListProvidersResponse hashicorp cloud vagrant 20220930 list providers response
//
// swagger:model hashicorp.cloud.vagrant_20220930.ListProvidersResponse
type HashicorpCloudVagrant20220930ListProvidersResponse struct {

	// Pagination tokens for a subsequent request.
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`

	// The returned Providers. Note that the message here will be either one of
	// HostedProvider or ExternalProvider.
	Providers []*cloud.GoogleProtobufAny `json:"providers"`
}

// Validate validates this hashicorp cloud vagrant 20220930 list providers response
func (m *HashicorpCloudVagrant20220930ListProvidersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrant20220930ListProvidersResponse) validatePagination(formats strfmt.Registry) error {

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

func (m *HashicorpCloudVagrant20220930ListProvidersResponse) validateProviders(formats strfmt.Registry) error {

	if swag.IsZero(m.Providers) { // not required
		return nil
	}

	for i := 0; i < len(m.Providers); i++ {
		if swag.IsZero(m.Providers[i]) { // not required
			continue
		}

		if m.Providers[i] != nil {
			if err := m.Providers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("providers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930ListProvidersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930ListProvidersResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrant20220930ListProvidersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}