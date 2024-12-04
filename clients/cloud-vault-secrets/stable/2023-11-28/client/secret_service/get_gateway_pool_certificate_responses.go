// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-11-28/models"
)

// GetGatewayPoolCertificateReader is a Reader for the GetGatewayPoolCertificate structure.
type GetGatewayPoolCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGatewayPoolCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGatewayPoolCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetGatewayPoolCertificateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGatewayPoolCertificateOK creates a GetGatewayPoolCertificateOK with default headers values
func NewGetGatewayPoolCertificateOK() *GetGatewayPoolCertificateOK {
	return &GetGatewayPoolCertificateOK{}
}

/*
GetGatewayPoolCertificateOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetGatewayPoolCertificateOK struct {
	Payload *models.Secrets20231128GetGatewayPoolCertificateResponse
}

// IsSuccess returns true when this get gateway pool certificate o k response has a 2xx status code
func (o *GetGatewayPoolCertificateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gateway pool certificate o k response has a 3xx status code
func (o *GetGatewayPoolCertificateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway pool certificate o k response has a 4xx status code
func (o *GetGatewayPoolCertificateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gateway pool certificate o k response has a 5xx status code
func (o *GetGatewayPoolCertificateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway pool certificate o k response a status code equal to that given
func (o *GetGatewayPoolCertificateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get gateway pool certificate o k response
func (o *GetGatewayPoolCertificateOK) Code() int {
	return 200
}

func (o *GetGatewayPoolCertificateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools/{gateway_pool_name}/certificate][%d] getGatewayPoolCertificateOK %s", 200, payload)
}

func (o *GetGatewayPoolCertificateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools/{gateway_pool_name}/certificate][%d] getGatewayPoolCertificateOK %s", 200, payload)
}

func (o *GetGatewayPoolCertificateOK) GetPayload() *models.Secrets20231128GetGatewayPoolCertificateResponse {
	return o.Payload
}

func (o *GetGatewayPoolCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128GetGatewayPoolCertificateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayPoolCertificateDefault creates a GetGatewayPoolCertificateDefault with default headers values
func NewGetGatewayPoolCertificateDefault(code int) *GetGatewayPoolCertificateDefault {
	return &GetGatewayPoolCertificateDefault{
		_statusCode: code,
	}
}

/*
GetGatewayPoolCertificateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetGatewayPoolCertificateDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this get gateway pool certificate default response has a 2xx status code
func (o *GetGatewayPoolCertificateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get gateway pool certificate default response has a 3xx status code
func (o *GetGatewayPoolCertificateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get gateway pool certificate default response has a 4xx status code
func (o *GetGatewayPoolCertificateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get gateway pool certificate default response has a 5xx status code
func (o *GetGatewayPoolCertificateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get gateway pool certificate default response a status code equal to that given
func (o *GetGatewayPoolCertificateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get gateway pool certificate default response
func (o *GetGatewayPoolCertificateDefault) Code() int {
	return o._statusCode
}

func (o *GetGatewayPoolCertificateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools/{gateway_pool_name}/certificate][%d] GetGatewayPoolCertificate default %s", o._statusCode, payload)
}

func (o *GetGatewayPoolCertificateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools/{gateway_pool_name}/certificate][%d] GetGatewayPoolCertificate default %s", o._statusCode, payload)
}

func (o *GetGatewayPoolCertificateDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *GetGatewayPoolCertificateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
