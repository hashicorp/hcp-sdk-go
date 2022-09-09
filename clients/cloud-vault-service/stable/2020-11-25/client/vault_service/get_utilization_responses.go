// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-11-25/models"
)

// GetUtilizationReader is a Reader for the GetUtilization structure.
type GetUtilizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUtilizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUtilizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUtilizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUtilizationOK creates a GetUtilizationOK with default headers values
func NewGetUtilizationOK() *GetUtilizationOK {
	return &GetUtilizationOK{}
}

/*
GetUtilizationOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetUtilizationOK struct {
	Payload *models.HashicorpCloudVault20201125GetUtilizationResponse
}

// IsSuccess returns true when this get utilization o k response has a 2xx status code
func (o *GetUtilizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get utilization o k response has a 3xx status code
func (o *GetUtilizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get utilization o k response has a 4xx status code
func (o *GetUtilizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get utilization o k response has a 5xx status code
func (o *GetUtilizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get utilization o k response a status code equal to that given
func (o *GetUtilizationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUtilizationOK) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/utilization][%d] getUtilizationOK  %+v", 200, o.Payload)
}

func (o *GetUtilizationOK) String() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/utilization][%d] getUtilizationOK  %+v", 200, o.Payload)
}

func (o *GetUtilizationOK) GetPayload() *models.HashicorpCloudVault20201125GetUtilizationResponse {
	return o.Payload
}

func (o *GetUtilizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125GetUtilizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUtilizationDefault creates a GetUtilizationDefault with default headers values
func NewGetUtilizationDefault(code int) *GetUtilizationDefault {
	return &GetUtilizationDefault{
		_statusCode: code,
	}
}

/*
GetUtilizationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetUtilizationDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get utilization default response
func (o *GetUtilizationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get utilization default response has a 2xx status code
func (o *GetUtilizationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get utilization default response has a 3xx status code
func (o *GetUtilizationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get utilization default response has a 4xx status code
func (o *GetUtilizationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get utilization default response has a 5xx status code
func (o *GetUtilizationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get utilization default response a status code equal to that given
func (o *GetUtilizationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetUtilizationDefault) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/utilization][%d] GetUtilization default  %+v", o._statusCode, o.Payload)
}

func (o *GetUtilizationDefault) String() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/utilization][%d] GetUtilization default  %+v", o._statusCode, o.Payload)
}

func (o *GetUtilizationDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetUtilizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
