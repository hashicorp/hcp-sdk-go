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
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-04-20/models"
)

// GetAdminTokenReader is a Reader for the GetAdminToken structure.
type GetAdminTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdminTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdminTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAdminTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAdminTokenOK creates a GetAdminTokenOK with default headers values
func NewGetAdminTokenOK() *GetAdminTokenOK {
	return &GetAdminTokenOK{}
}

/*
GetAdminTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAdminTokenOK struct {
	Payload *models.HashicorpCloudVault20200420GetAdminTokenResponse
}

// IsSuccess returns true when this get admin token o k response has a 2xx status code
func (o *GetAdminTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get admin token o k response has a 3xx status code
func (o *GetAdminTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admin token o k response has a 4xx status code
func (o *GetAdminTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get admin token o k response has a 5xx status code
func (o *GetAdminTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get admin token o k response a status code equal to that given
func (o *GetAdminTokenOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAdminTokenOK) Error() string {
	return fmt.Sprintf("[GET /vault/2020-04-20/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/admintoken][%d] getAdminTokenOK  %+v", 200, o.Payload)
}

func (o *GetAdminTokenOK) String() string {
	return fmt.Sprintf("[GET /vault/2020-04-20/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/admintoken][%d] getAdminTokenOK  %+v", 200, o.Payload)
}

func (o *GetAdminTokenOK) GetPayload() *models.HashicorpCloudVault20200420GetAdminTokenResponse {
	return o.Payload
}

func (o *GetAdminTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20200420GetAdminTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminTokenDefault creates a GetAdminTokenDefault with default headers values
func NewGetAdminTokenDefault(code int) *GetAdminTokenDefault {
	return &GetAdminTokenDefault{
		_statusCode: code,
	}
}

/*
GetAdminTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetAdminTokenDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get admin token default response
func (o *GetAdminTokenDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get admin token default response has a 2xx status code
func (o *GetAdminTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get admin token default response has a 3xx status code
func (o *GetAdminTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get admin token default response has a 4xx status code
func (o *GetAdminTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get admin token default response has a 5xx status code
func (o *GetAdminTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get admin token default response a status code equal to that given
func (o *GetAdminTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAdminTokenDefault) Error() string {
	return fmt.Sprintf("[GET /vault/2020-04-20/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/admintoken][%d] GetAdminToken default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdminTokenDefault) String() string {
	return fmt.Sprintf("[GET /vault/2020-04-20/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/admintoken][%d] GetAdminToken default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdminTokenDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetAdminTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
