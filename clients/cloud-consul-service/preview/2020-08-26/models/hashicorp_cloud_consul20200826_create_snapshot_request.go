// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudConsul20200826CreateSnapshotRequest CreateSnapshotRequest is a request to create a new snapshot.
//
// swagger:model hashicorp.cloud.consul_20200826.CreateSnapshotRequest
type HashicorpCloudConsul20200826CreateSnapshotRequest struct {

	// name identifies the name of the snapshot.
	Name string `json:"name,omitempty"`

	// resource identifies what resource to take a snapshot from.
	Resource *cloud.HashicorpCloudLocationLink `json:"resource,omitempty"`
}

// Validate validates this hashicorp cloud consul 20200826 create snapshot request
func (m *HashicorpCloudConsul20200826CreateSnapshotRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826CreateSnapshotRequest) validateResource(formats strfmt.Registry) error {

	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826CreateSnapshotRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826CreateSnapshotRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20200826CreateSnapshotRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
