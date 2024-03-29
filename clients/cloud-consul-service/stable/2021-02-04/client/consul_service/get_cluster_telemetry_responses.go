// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/stable/2021-02-04/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// GetClusterTelemetryReader is a Reader for the GetClusterTelemetry structure.
type GetClusterTelemetryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterTelemetryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterTelemetryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterTelemetryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterTelemetryOK creates a GetClusterTelemetryOK with default headers values
func NewGetClusterTelemetryOK() *GetClusterTelemetryOK {
	return &GetClusterTelemetryOK{}
}

/*
GetClusterTelemetryOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetClusterTelemetryOK struct {
	Payload *models.HashicorpCloudConsul20210204GetClusterTelemetryResponse
}

// IsSuccess returns true when this get cluster telemetry o k response has a 2xx status code
func (o *GetClusterTelemetryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster telemetry o k response has a 3xx status code
func (o *GetClusterTelemetryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster telemetry o k response has a 4xx status code
func (o *GetClusterTelemetryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster telemetry o k response has a 5xx status code
func (o *GetClusterTelemetryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster telemetry o k response a status code equal to that given
func (o *GetClusterTelemetryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterTelemetryOK) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/telemetry][%d] getClusterTelemetryOK  %+v", 200, o.Payload)
}

func (o *GetClusterTelemetryOK) String() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/telemetry][%d] getClusterTelemetryOK  %+v", 200, o.Payload)
}

func (o *GetClusterTelemetryOK) GetPayload() *models.HashicorpCloudConsul20210204GetClusterTelemetryResponse {
	return o.Payload
}

func (o *GetClusterTelemetryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20210204GetClusterTelemetryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterTelemetryDefault creates a GetClusterTelemetryDefault with default headers values
func NewGetClusterTelemetryDefault(code int) *GetClusterTelemetryDefault {
	return &GetClusterTelemetryDefault{
		_statusCode: code,
	}
}

/*
GetClusterTelemetryDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetClusterTelemetryDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get cluster telemetry default response
func (o *GetClusterTelemetryDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get cluster telemetry default response has a 2xx status code
func (o *GetClusterTelemetryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster telemetry default response has a 3xx status code
func (o *GetClusterTelemetryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster telemetry default response has a 4xx status code
func (o *GetClusterTelemetryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster telemetry default response has a 5xx status code
func (o *GetClusterTelemetryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster telemetry default response a status code equal to that given
func (o *GetClusterTelemetryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetClusterTelemetryDefault) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/telemetry][%d] GetClusterTelemetry default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterTelemetryDefault) String() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/telemetry][%d] GetClusterTelemetry default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterTelemetryDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetClusterTelemetryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
