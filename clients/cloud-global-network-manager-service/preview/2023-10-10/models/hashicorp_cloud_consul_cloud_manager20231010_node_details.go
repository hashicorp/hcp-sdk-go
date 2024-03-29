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

// HashicorpCloudConsulCloudManager20231010NodeDetails hashicorp cloud consul cloud manager 20231010 node details
//
// swagger:model hashicorp.cloud.consul_cloud_manager_20231010.NodeDetails
type HashicorpCloudConsulCloudManager20231010NodeDetails struct {

	// address of the node
	Address string `json:"address,omitempty"`

	// name of the cluster the node is part of
	ClusterName string `json:"cluster_name,omitempty"`

	// id of the node
	ID string `json:"id,omitempty"`

	// name of the node
	Name string `json:"name,omitempty"`

	// resource_id of the cluster to which the node belongs
	ParentResourceID string `json:"parent_resource_id,omitempty"`

	// resource_name of the cluster to which the node belongs
	ParentResourceName string `json:"parent_resource_name,omitempty"`

	// name of the partition the node belongs to
	Partition string `json:"partition,omitempty"`

	// service data of all the services running in the node
	ServiceData *HashicorpCloudConsulCloudManager20231010NodeServiceData `json:"service_data,omitempty"`

	// tagged addresses of the node
	TaggedAddresses map[string]string `json:"tagged_addresses,omitempty"`
}

// Validate validates this hashicorp cloud consul cloud manager 20231010 node details
func (m *HashicorpCloudConsulCloudManager20231010NodeDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010NodeDetails) validateServiceData(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceData) { // not required
		return nil
	}

	if m.ServiceData != nil {
		if err := m.ServiceData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul cloud manager 20231010 node details based on the context it is used
func (m *HashicorpCloudConsulCloudManager20231010NodeDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010NodeDetails) contextValidateServiceData(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceData != nil {

		if swag.IsZero(m.ServiceData) { // not required
			return nil
		}

		if err := m.ServiceData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulCloudManager20231010NodeDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulCloudManager20231010NodeDetails) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulCloudManager20231010NodeDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
