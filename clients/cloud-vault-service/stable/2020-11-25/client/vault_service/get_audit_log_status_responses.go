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

// GetAuditLogStatusReader is a Reader for the GetAuditLogStatus structure.
type GetAuditLogStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditLogStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditLogStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAuditLogStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuditLogStatusOK creates a GetAuditLogStatusOK with default headers values
func NewGetAuditLogStatusOK() *GetAuditLogStatusOK {
	return &GetAuditLogStatusOK{}
}

/*
GetAuditLogStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAuditLogStatusOK struct {
	Payload *models.HashicorpCloudVault20201125GetAuditLogStatusResponse
}

// IsSuccess returns true when this get audit log status o k response has a 2xx status code
func (o *GetAuditLogStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get audit log status o k response has a 3xx status code
func (o *GetAuditLogStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audit log status o k response has a 4xx status code
func (o *GetAuditLogStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audit log status o k response has a 5xx status code
func (o *GetAuditLogStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get audit log status o k response a status code equal to that given
func (o *GetAuditLogStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get audit log status o k response
func (o *GetAuditLogStatusOK) Code() int {
	return 200
}

func (o *GetAuditLogStatusOK) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog/{log_id}][%d] getAuditLogStatusOK  %+v", 200, o.Payload)
}

func (o *GetAuditLogStatusOK) String() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog/{log_id}][%d] getAuditLogStatusOK  %+v", 200, o.Payload)
}

func (o *GetAuditLogStatusOK) GetPayload() *models.HashicorpCloudVault20201125GetAuditLogStatusResponse {
	return o.Payload
}

func (o *GetAuditLogStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125GetAuditLogStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditLogStatusDefault creates a GetAuditLogStatusDefault with default headers values
func NewGetAuditLogStatusDefault(code int) *GetAuditLogStatusDefault {
	return &GetAuditLogStatusDefault{
		_statusCode: code,
	}
}

/*
GetAuditLogStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetAuditLogStatusDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this get audit log status default response has a 2xx status code
func (o *GetAuditLogStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get audit log status default response has a 3xx status code
func (o *GetAuditLogStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get audit log status default response has a 4xx status code
func (o *GetAuditLogStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get audit log status default response has a 5xx status code
func (o *GetAuditLogStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get audit log status default response a status code equal to that given
func (o *GetAuditLogStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get audit log status default response
func (o *GetAuditLogStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetAuditLogStatusDefault) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog/{log_id}][%d] GetAuditLogStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetAuditLogStatusDefault) String() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog/{log_id}][%d] GetAuditLogStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetAuditLogStatusDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetAuditLogStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
