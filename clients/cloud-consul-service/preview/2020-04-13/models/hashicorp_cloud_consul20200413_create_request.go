// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsul20200413CreateRequest hashicorp cloud consul 20200413 create request
//
// swagger:model hashicorp.cloud.consul_20200413.CreateRequest
type HashicorpCloudConsul20200413CreateRequest struct {

	// Cluster is the Consul Cluster to create.
	Cluster *HashicorpCloudConsul20200413Cluster `json:"cluster,omitempty"`
}

// Validate validates this hashicorp cloud consul 20200413 create request
func (m *HashicorpCloudConsul20200413CreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200413CreateRequest) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20200413CreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20200413CreateRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20200413CreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
