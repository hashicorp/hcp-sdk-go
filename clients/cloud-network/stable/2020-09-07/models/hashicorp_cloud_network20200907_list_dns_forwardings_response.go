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

// HashicorpCloudNetwork20200907ListDNSForwardingsResponse ListDNSForwardingsResponse is a response type for ListDNSForwardings endpoint
//
// swagger:model hashicorp.cloud.network_20200907.ListDNSForwardingsResponse
type HashicorpCloudNetwork20200907ListDNSForwardingsResponse struct {

	// DNSForwardings is a list of dnsForwardings
	DNSForwardings []*HashicorpCloudNetwork20200907DNSForwardingResponse `json:"dns_forwardings"`

	// Pagination contains the pagination tokens for a subsequent request.
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 list DNS forwardings response
func (m *HashicorpCloudNetwork20200907ListDNSForwardingsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSForwardings(formats); err != nil {
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

func (m *HashicorpCloudNetwork20200907ListDNSForwardingsResponse) validateDNSForwardings(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSForwardings) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSForwardings); i++ {
		if swag.IsZero(m.DNSForwardings[i]) { // not required
			continue
		}

		if m.DNSForwardings[i] != nil {
			if err := m.DNSForwardings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns_forwardings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dns_forwardings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudNetwork20200907ListDNSForwardingsResponse) validatePagination(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud network 20200907 list DNS forwardings response based on the context it is used
func (m *HashicorpCloudNetwork20200907ListDNSForwardingsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNSForwardings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907ListDNSForwardingsResponse) contextValidateDNSForwardings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSForwardings); i++ {

		if m.DNSForwardings[i] != nil {

			if swag.IsZero(m.DNSForwardings[i]) { // not required
				return nil
			}

			if err := m.DNSForwardings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns_forwardings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dns_forwardings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudNetwork20200907ListDNSForwardingsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907ListDNSForwardingsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907ListDNSForwardingsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907ListDNSForwardingsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
