// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudConsul20200826CreateResponse hashicorp cloud consul 20200826 create response
//
// swagger:model hashicorp.cloud.consul_20200826.CreateResponse
type HashicorpCloudConsul20200826CreateResponse struct {

	// cluster is the newly created Consul cluster.
	Cluster *HashicorpCloudConsul20200826Cluster `json:"cluster,omitempty"`

	// operation tracks the asynchronous creation of the Consul cluster.
	Operation *cloud.HashicorpCloudOperationOperation `json:"operation,omitempty"`
}

// Validate validates this hashicorp cloud consul 20200826 create response
func (m *HashicorpCloudConsul20200826CreateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826CreateResponse) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsul20200826CreateResponse) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul 20200826 create response based on the context it is used
func (m *HashicorpCloudConsul20200826CreateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826CreateResponse) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsul20200826CreateResponse) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.Operation != nil {
		if err := m.Operation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826CreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826CreateResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20200826CreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
