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

// GetGcpDynamicSecretReader is a Reader for the GetGcpDynamicSecret structure.
type GetGcpDynamicSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGcpDynamicSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGcpDynamicSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetGcpDynamicSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGcpDynamicSecretOK creates a GetGcpDynamicSecretOK with default headers values
func NewGetGcpDynamicSecretOK() *GetGcpDynamicSecretOK {
	return &GetGcpDynamicSecretOK{}
}

/*
GetGcpDynamicSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetGcpDynamicSecretOK struct {
	Payload *models.Secrets20231128GetGcpDynamicSecretResponse
}

// IsSuccess returns true when this get gcp dynamic secret o k response has a 2xx status code
func (o *GetGcpDynamicSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gcp dynamic secret o k response has a 3xx status code
func (o *GetGcpDynamicSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gcp dynamic secret o k response has a 4xx status code
func (o *GetGcpDynamicSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gcp dynamic secret o k response has a 5xx status code
func (o *GetGcpDynamicSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gcp dynamic secret o k response a status code equal to that given
func (o *GetGcpDynamicSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get gcp dynamic secret o k response
func (o *GetGcpDynamicSecretOK) Code() int {
	return 200
}

func (o *GetGcpDynamicSecretOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret/{name}][%d] getGcpDynamicSecretOK %s", 200, payload)
}

func (o *GetGcpDynamicSecretOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret/{name}][%d] getGcpDynamicSecretOK %s", 200, payload)
}

func (o *GetGcpDynamicSecretOK) GetPayload() *models.Secrets20231128GetGcpDynamicSecretResponse {
	return o.Payload
}

func (o *GetGcpDynamicSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128GetGcpDynamicSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGcpDynamicSecretDefault creates a GetGcpDynamicSecretDefault with default headers values
func NewGetGcpDynamicSecretDefault(code int) *GetGcpDynamicSecretDefault {
	return &GetGcpDynamicSecretDefault{
		_statusCode: code,
	}
}

/*
GetGcpDynamicSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetGcpDynamicSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this get gcp dynamic secret default response has a 2xx status code
func (o *GetGcpDynamicSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get gcp dynamic secret default response has a 3xx status code
func (o *GetGcpDynamicSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get gcp dynamic secret default response has a 4xx status code
func (o *GetGcpDynamicSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get gcp dynamic secret default response has a 5xx status code
func (o *GetGcpDynamicSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get gcp dynamic secret default response a status code equal to that given
func (o *GetGcpDynamicSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get gcp dynamic secret default response
func (o *GetGcpDynamicSecretDefault) Code() int {
	return o._statusCode
}

func (o *GetGcpDynamicSecretDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret/{name}][%d] GetGcpDynamicSecret default %s", o._statusCode, payload)
}

func (o *GetGcpDynamicSecretDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret/{name}][%d] GetGcpDynamicSecret default %s", o._statusCode, payload)
}

func (o *GetGcpDynamicSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *GetGcpDynamicSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
