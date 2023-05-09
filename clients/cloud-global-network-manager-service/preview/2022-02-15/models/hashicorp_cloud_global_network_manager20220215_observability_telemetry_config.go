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

// HashicorpCloudGlobalNetworkManager20220215ObservabilityTelemetryConfig hashicorp cloud global network manager 20220215 observability telemetry config
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.observability.TelemetryConfig
type HashicorpCloudGlobalNetworkManager20220215ObservabilityTelemetryConfig struct {

	// Endpoint is the URL to use when exporting telemetry via OTLP
	Endpoint string `json:"endpoint,omitempty"`

	// Labels to add to each message
	Labels map[string]string `json:"labels,omitempty"`

	// Metrics is the configuration specific to metric data
	Metrics *HashicorpCloudGlobalNetworkManager20220215ObservabilityTelemetryMetricsConfig `json:"metrics,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 observability telemetry config
func (m *HashicorpCloudGlobalNetworkManager20220215ObservabilityTelemetryConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ObservabilityTelemetryConfig) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	if m.Metrics != nil {
		if err := m.Metrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 observability telemetry config based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215ObservabilityTelemetryConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ObservabilityTelemetryConfig) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	if m.Metrics != nil {
		if err := m.Metrics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ObservabilityTelemetryConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ObservabilityTelemetryConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215ObservabilityTelemetryConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}