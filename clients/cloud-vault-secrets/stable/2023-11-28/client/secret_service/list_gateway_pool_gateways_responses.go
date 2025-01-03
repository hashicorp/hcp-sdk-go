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

// ListGatewayPoolGatewaysReader is a Reader for the ListGatewayPoolGateways structure.
type ListGatewayPoolGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGatewayPoolGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGatewayPoolGatewaysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGatewayPoolGatewaysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGatewayPoolGatewaysOK creates a ListGatewayPoolGatewaysOK with default headers values
func NewListGatewayPoolGatewaysOK() *ListGatewayPoolGatewaysOK {
	return &ListGatewayPoolGatewaysOK{}
}

/*
ListGatewayPoolGatewaysOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListGatewayPoolGatewaysOK struct {
	Payload *models.Secrets20231128ListGatewayPoolGatewaysResponse
}

// IsSuccess returns true when this list gateway pool gateways o k response has a 2xx status code
func (o *ListGatewayPoolGatewaysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list gateway pool gateways o k response has a 3xx status code
func (o *ListGatewayPoolGatewaysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list gateway pool gateways o k response has a 4xx status code
func (o *ListGatewayPoolGatewaysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list gateway pool gateways o k response has a 5xx status code
func (o *ListGatewayPoolGatewaysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list gateway pool gateways o k response a status code equal to that given
func (o *ListGatewayPoolGatewaysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list gateway pool gateways o k response
func (o *ListGatewayPoolGatewaysOK) Code() int {
	return 200
}

func (o *ListGatewayPoolGatewaysOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools/{gateway_pool_name}/gateways][%d] listGatewayPoolGatewaysOK %s", 200, payload)
}

func (o *ListGatewayPoolGatewaysOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools/{gateway_pool_name}/gateways][%d] listGatewayPoolGatewaysOK %s", 200, payload)
}

func (o *ListGatewayPoolGatewaysOK) GetPayload() *models.Secrets20231128ListGatewayPoolGatewaysResponse {
	return o.Payload
}

func (o *ListGatewayPoolGatewaysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128ListGatewayPoolGatewaysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGatewayPoolGatewaysDefault creates a ListGatewayPoolGatewaysDefault with default headers values
func NewListGatewayPoolGatewaysDefault(code int) *ListGatewayPoolGatewaysDefault {
	return &ListGatewayPoolGatewaysDefault{
		_statusCode: code,
	}
}

/*
ListGatewayPoolGatewaysDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListGatewayPoolGatewaysDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this list gateway pool gateways default response has a 2xx status code
func (o *ListGatewayPoolGatewaysDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list gateway pool gateways default response has a 3xx status code
func (o *ListGatewayPoolGatewaysDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list gateway pool gateways default response has a 4xx status code
func (o *ListGatewayPoolGatewaysDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list gateway pool gateways default response has a 5xx status code
func (o *ListGatewayPoolGatewaysDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list gateway pool gateways default response a status code equal to that given
func (o *ListGatewayPoolGatewaysDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list gateway pool gateways default response
func (o *ListGatewayPoolGatewaysDefault) Code() int {
	return o._statusCode
}

func (o *ListGatewayPoolGatewaysDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools/{gateway_pool_name}/gateways][%d] ListGatewayPoolGateways default %s", o._statusCode, payload)
}

func (o *ListGatewayPoolGatewaysDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/gateway-pools/{gateway_pool_name}/gateways][%d] ListGatewayPoolGateways default %s", o._statusCode, payload)
}

func (o *ListGatewayPoolGatewaysDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ListGatewayPoolGatewaysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
