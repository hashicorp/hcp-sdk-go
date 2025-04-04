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

// DescribeProviderReader is a Reader for the DescribeProvider structure.
type DescribeProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeProviderOK creates a DescribeProviderOK with default headers values
func NewDescribeProviderOK() *DescribeProviderOK {
	return &DescribeProviderOK{}
}

/*
DescribeProviderOK describes a response with status code 200, with default header values.

A successful response.
*/
type DescribeProviderOK struct {
	Payload *models.Secrets20231128DescribeProviderResponse
}

// IsSuccess returns true when this describe provider o k response has a 2xx status code
func (o *DescribeProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe provider o k response has a 3xx status code
func (o *DescribeProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe provider o k response has a 4xx status code
func (o *DescribeProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe provider o k response has a 5xx status code
func (o *DescribeProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe provider o k response a status code equal to that given
func (o *DescribeProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe provider o k response
func (o *DescribeProviderOK) Code() int {
	return 200
}

func (o *DescribeProviderOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/providers/{name}][%d] describeProviderOK %s", 200, payload)
}

func (o *DescribeProviderOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/providers/{name}][%d] describeProviderOK %s", 200, payload)
}

func (o *DescribeProviderOK) GetPayload() *models.Secrets20231128DescribeProviderResponse {
	return o.Payload
}

func (o *DescribeProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128DescribeProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeProviderDefault creates a DescribeProviderDefault with default headers values
func NewDescribeProviderDefault(code int) *DescribeProviderDefault {
	return &DescribeProviderDefault{
		_statusCode: code,
	}
}

/*
DescribeProviderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DescribeProviderDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this describe provider default response has a 2xx status code
func (o *DescribeProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe provider default response has a 3xx status code
func (o *DescribeProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe provider default response has a 4xx status code
func (o *DescribeProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe provider default response has a 5xx status code
func (o *DescribeProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe provider default response a status code equal to that given
func (o *DescribeProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe provider default response
func (o *DescribeProviderDefault) Code() int {
	return o._statusCode
}

func (o *DescribeProviderDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/providers/{name}][%d] DescribeProvider default %s", o._statusCode, payload)
}

func (o *DescribeProviderDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/providers/{name}][%d] DescribeProvider default %s", o._statusCode, payload)
}

func (o *DescribeProviderDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DescribeProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
