// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	cloud "github.com/hashicorp/cloud-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudConsul20200826Cluster Cluster represents a single Consul cluster.
//
// swagger:model hashicorp.cloud.consul_20200826.Cluster
type HashicorpCloudConsul20200826Cluster struct {

	// config holds the configuration of the cluster.
	Config *HashicorpCloudConsul20200826ClusterConfig `json:"config,omitempty"`

	// consul_version is the current Consul version the server nodes are running.
	ConsulVersion string `json:"consul_version,omitempty"`

	// created_at is the timestamp of when the cluster was first created.
	// Output only.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// dns_names holds the cluster's public and private DNS names.
	// Output only.
	// Read Only: true
	DNSNames *HashicorpCloudConsul20200826ClusterDNSNames `json:"dns_names,omitempty"`

	// id is ID of the Consul cluster.
	ID string `json:"id,omitempty"`

	// location is the location of the cluster.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// state is the current state of the cluster.
	// Output only.
	// Read Only: true
	State HashicorpCloudConsul20200826ClusterState `json:"state,omitempty"`
}

// Validate validates this hashicorp cloud consul 20200826 cluster
func (m *HashicorpCloudConsul20200826Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826Cluster) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsul20200826Cluster) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudConsul20200826Cluster) validateDNSNames(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSNames) { // not required
		return nil
	}

	if m.DNSNames != nil {
		if err := m.DNSNames.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns_names")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsul20200826Cluster) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsul20200826Cluster) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826Cluster) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20200826Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
