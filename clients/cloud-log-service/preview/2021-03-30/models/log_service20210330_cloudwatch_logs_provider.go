// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogService20210330CloudwatchLogsProvider CloudwatchLogsProvider represents the necessary information to configure a cloudwatch destination.
//
// swagger:model log_service_20210330CloudwatchLogsProvider
type LogService20210330CloudwatchLogsProvider struct {

	// external_id represents the external_id to provide when assuming the aws IAM role.
	ExternalID string `json:"external_id,omitempty"`

	// log_group_name represents the name of the log group.
	LogGroupName string `json:"log_group_name,omitempty"`

	// log_stream_name represents the name of the log stream.
	LogStreamName string `json:"log_stream_name,omitempty"`

	// region represents the region the cloudwatch log is set up for.
	Region string `json:"region,omitempty"`

	// role_arn is the role_arn that will be assumed to stream logs.
	RoleArn string `json:"role_arn,omitempty"`
}

// Validate validates this log service 20210330 cloudwatch logs provider
func (m *LogService20210330CloudwatchLogsProvider) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this log service 20210330 cloudwatch logs provider based on context it is used
func (m *LogService20210330CloudwatchLogsProvider) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogService20210330CloudwatchLogsProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogService20210330CloudwatchLogsProvider) UnmarshalBinary(b []byte) error {
	var res LogService20210330CloudwatchLogsProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
