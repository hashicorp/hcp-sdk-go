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

// HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata hashicorp cloud consul telemetry 20230414 service topology edge metadata
//
// swagger:model hashicorp.cloud.consul_telemetry_20230414.ServiceTopology.Edge.Metadata
type HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata struct {

	// http_metrics describe the requests between the two services.
	HTTPMetrics *HashicorpCloudConsulTelemetry20230414HTTPMetrics `json:"http_metrics,omitempty"`

	// intentions describe the consul intentions that apply to the call from the downstream to
	// upstream service covered by this edge.
	Intentions *HashicorpCloudConsulTelemetry20230414Intentions `json:"intentions,omitempty"`

	// tcp_metrics describe the connections between the two services.
	TCPMetrics *HashicorpCloudConsulTelemetry20230414TCPMetrics `json:"tcp_metrics,omitempty"`
}

// Validate validates this hashicorp cloud consul telemetry 20230414 service topology edge metadata
func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntentions(formats); err != nil {
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

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata) validateHTTPMetrics(formats strfmt.Registry) error {
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

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata) validateIntentions(formats strfmt.Registry) error {
	if swag.IsZero(m.Intentions) { // not required
		return nil
	}

	if m.Intentions != nil {
		if err := m.Intentions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("intentions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("intentions")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata) validateTCPMetrics(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud consul telemetry 20230414 service topology edge metadata based on the context it is used
func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHTTPMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIntentions(ctx, formats); err != nil {
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

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata) contextValidateHTTPMetrics(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata) contextValidateIntentions(ctx context.Context, formats strfmt.Registry) error {

	if m.Intentions != nil {

		if swag.IsZero(m.Intentions) { // not required
			return nil
		}

		if err := m.Intentions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("intentions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("intentions")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata) contextValidateTCPMetrics(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulTelemetry20230414ServiceTopologyEdgeMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
