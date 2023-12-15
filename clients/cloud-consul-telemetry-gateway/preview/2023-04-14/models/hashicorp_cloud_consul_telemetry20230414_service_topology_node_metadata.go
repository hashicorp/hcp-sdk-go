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

// HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata hashicorp cloud consul telemetry 20230414 service topology node metadata
//
// swagger:model hashicorp.cloud.consul_telemetry_20230414.ServiceTopology.Node.Metadata
type HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata struct {

	// cluster is the cluster to which this service belongs.
	Cluster *HashicorpCloudConsulTelemetry20230414Cluster `json:"cluster,omitempty"`

	// http_metrics describe the http requests into this service from all others in the mesh.
	HTTPMetrics *HashicorpCloudConsulTelemetry20230414HTTPMetrics `json:"http_metrics,omitempty"`

	// service is the service of this node. These fields are populated from GNM. This includes health information.
	Service *HashicorpCloudConsulTelemetry20230414Service `json:"service,omitempty"`

	// tcp_metrics describe the tcp connections into this service from all others in the mesh.
	TCPMetrics *HashicorpCloudConsulTelemetry20230414TCPMetrics `json:"tcp_metrics,omitempty"`
}

// Validate validates this hashicorp cloud consul telemetry 20230414 service topology node metadata
func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTCPMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) validateCluster(formats strfmt.Registry) error {
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

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) validateHTTPMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPMetrics) { // not required
		return nil
	}

	if m.HTTPMetrics != nil {
		if err := m.HTTPMetrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_metrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_metrics")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) validateTCPMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.TCPMetrics) { // not required
		return nil
	}

	if m.TCPMetrics != nil {
		if err := m.TCPMetrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tcp_metrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tcp_metrics")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul telemetry 20230414 service topology node metadata based on the context it is used
func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTCPMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {

		if swag.IsZero(m.Cluster) { // not required
			return nil
		}

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

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) contextValidateHTTPMetrics(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPMetrics != nil {

		if swag.IsZero(m.HTTPMetrics) { // not required
			return nil
		}

		if err := m.HTTPMetrics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_metrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_metrics")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {

		if swag.IsZero(m.Service) { // not required
			return nil
		}

		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) contextValidateTCPMetrics(ctx context.Context, formats strfmt.Registry) error {

	if m.TCPMetrics != nil {

		if swag.IsZero(m.TCPMetrics) { // not required
			return nil
		}

		if err := m.TCPMetrics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tcp_metrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tcp_metrics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulTelemetry20230414ServiceTopologyNodeMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
