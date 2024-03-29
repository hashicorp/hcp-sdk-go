// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulTelemetry20230414HTTPMetrics hashicorp cloud consul telemetry 20230414 Http metrics
//
// swagger:model hashicorp.cloud.consul_telemetry_20230414.HttpMetrics
type HashicorpCloudConsulTelemetry20230414HTTPMetrics struct {

	// duration_ms_avg is the average latency of response from a service in milliseconds.
	DurationMsAvg float32 `json:"duration_ms_avg,omitempty"`

	// duration_ms_p50 is the P50 latency of response from a service in milliseconds.
	DurationMsP50 float32 `json:"duration_ms_p50,omitempty"`

	// duration_ms_p90 is the P90 latency of response from a service in milliseconds.
	DurationMsP90 float32 `json:"duration_ms_p90,omitempty"`

	// error_ratio is the ratio of HTTP responses from a service that have a 5XX response code.
	ErrorRatio float32 `json:"error_ratio,omitempty"`

	// req_rate is the per-second rate of requests to a service.
	ReqRate float32 `json:"req_rate,omitempty"`
}

// Validate validates this hashicorp cloud consul telemetry 20230414 Http metrics
func (m *HashicorpCloudConsulTelemetry20230414HTTPMetrics) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud consul telemetry 20230414 Http metrics based on context it is used
func (m *HashicorpCloudConsulTelemetry20230414HTTPMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414HTTPMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414HTTPMetrics) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulTelemetry20230414HTTPMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
