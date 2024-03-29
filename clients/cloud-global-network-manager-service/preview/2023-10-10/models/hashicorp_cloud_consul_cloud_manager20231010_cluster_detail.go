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

// HashicorpCloudConsulCloudManager20231010ClusterDetail hashicorp cloud consul cloud manager 20231010 cluster detail
//
// swagger:model hashicorp.cloud.consul_cloud_manager_20231010.ClusterDetail
type HashicorpCloudConsulCloudManager20231010ClusterDetail struct {

	// Configuration for the cluster. Returned only for HashiCorp-managed clusters on HCP.
	Config *HashicorpCloudConsulCloudManager20231010ClusterConfig `json:"config,omitempty"`

	// Consul version of the cluster
	ConsulVersion string `json:"consul_version,omitempty"`

	// Metadata of the cluster
	Meta *HashicorpCloudConsulCloudManager20231010Meta `json:"meta,omitempty"`

	// Name of the cluster.
	Name string `json:"name,omitempty"`

	// resource id of the parent
	ParentResourceID string `json:"parent_resource_id,omitempty"`

	// resource name of the parent
	ParentResourceName string `json:"parent_resource_name,omitempty"`

	// Region where the cluster is deployed
	Region *HashicorpCloudConsulCloudManager20231010Region `json:"region,omitempty"`

	// Resource ID of the cluster.
	ResourceID string `json:"resource_id,omitempty"`

	// Resource name of the cluster.
	ResourceName string `json:"resource_name,omitempty"`

	// Current state of the cluster.
	State string `json:"state,omitempty"`

	// Summary of the cluster
	Summary *HashicorpCloudConsulCloudManager20231010ClusterSummary `json:"summary,omitempty"`

	// Type of the cluster.
	Type string `json:"type,omitempty"`
}

// Validate validates this hashicorp cloud consul cloud manager 20231010 cluster detail
func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) validateSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.Summary) { // not required
		return nil
	}

	if m.Summary != nil {
		if err := m.Summary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul cloud manager 20231010 cluster detail based on the context it is used
func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {

		if swag.IsZero(m.Config) { // not required
			return nil
		}

		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {

		if swag.IsZero(m.Meta) { // not required
			return nil
		}

		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {

		if swag.IsZero(m.Region) { // not required
			return nil
		}

		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) contextValidateSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.Summary != nil {

		if swag.IsZero(m.Summary) { // not required
			return nil
		}

		if err := m.Summary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulCloudManager20231010ClusterDetail) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulCloudManager20231010ClusterDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
