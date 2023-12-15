// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulTelemetry20230414Service Service is a single service within the service topology.
//
// swagger:model hashicorp.cloud.consul_telemetry_20230414.Service
type HashicorpCloudConsulTelemetry20230414Service struct {

	// Count of critical service instances.
	ChecksCritical int32 `json:"checks_critical,omitempty"`

	// Count of passing service instances.
	ChecksPassing int32 `json:"checks_passing,omitempty"`

	// Count of warning service instances.
	ChecksWarning int32 `json:"checks_warning,omitempty"`

	// Count of total service instances.
	InstanceCount int32 `json:"instance_count,omitempty"`

	// The kind of service.
	Kind string `json:"kind,omitempty"`

	// Name of the service.
	Name string `json:"name,omitempty"`

	// Namespace where the service is deployed.
	Namespace string `json:"namespace,omitempty"`

	// Partition where the service is deployed.
	Partition string `json:"partition,omitempty"`

	// The sameness group the service belongs to.
	SamenessGroup string `json:"sameness_group,omitempty"`
}

// Validate validates this hashicorp cloud consul telemetry 20230414 service
func (m *HashicorpCloudConsulTelemetry20230414Service) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud consul telemetry 20230414 service based on context it is used
func (m *HashicorpCloudConsulTelemetry20230414Service) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414Service) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414Service) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulTelemetry20230414Service
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
