// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-11-28/models"
)

// GetPostgresRotatingSecretConfigReader is a Reader for the GetPostgresRotatingSecretConfig structure.
type GetPostgresRotatingSecretConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPostgresRotatingSecretConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPostgresRotatingSecretConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPostgresRotatingSecretConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPostgresRotatingSecretConfigOK creates a GetPostgresRotatingSecretConfigOK with default headers values
func NewGetPostgresRotatingSecretConfigOK() *GetPostgresRotatingSecretConfigOK {
	return &GetPostgresRotatingSecretConfigOK{}
}

/*
GetPostgresRotatingSecretConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetPostgresRotatingSecretConfigOK struct {
	Payload *models.Secrets20231128GetPostgresRotatingSecretConfigResponse
}

// IsSuccess returns true when this get postgres rotating secret config o k response has a 2xx status code
func (o *GetPostgresRotatingSecretConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get postgres rotating secret config o k response has a 3xx status code
func (o *GetPostgresRotatingSecretConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get postgres rotating secret config o k response has a 4xx status code
func (o *GetPostgresRotatingSecretConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get postgres rotating secret config o k response has a 5xx status code
func (o *GetPostgresRotatingSecretConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get postgres rotating secret config o k response a status code equal to that given
func (o *GetPostgresRotatingSecretConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get postgres rotating secret config o k response
func (o *GetPostgresRotatingSecretConfigOK) Code() int {
	return 200
}

func (o *GetPostgresRotatingSecretConfigOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/postgres/secret/{name}][%d] getPostgresRotatingSecretConfigOK  %+v", 200, o.Payload)
}

func (o *GetPostgresRotatingSecretConfigOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/postgres/secret/{name}][%d] getPostgresRotatingSecretConfigOK  %+v", 200, o.Payload)
}

func (o *GetPostgresRotatingSecretConfigOK) GetPayload() *models.Secrets20231128GetPostgresRotatingSecretConfigResponse {
	return o.Payload
}

func (o *GetPostgresRotatingSecretConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128GetPostgresRotatingSecretConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPostgresRotatingSecretConfigDefault creates a GetPostgresRotatingSecretConfigDefault with default headers values
func NewGetPostgresRotatingSecretConfigDefault(code int) *GetPostgresRotatingSecretConfigDefault {
	return &GetPostgresRotatingSecretConfigDefault{
		_statusCode: code,
	}
}

/*
GetPostgresRotatingSecretConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetPostgresRotatingSecretConfigDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this get postgres rotating secret config default response has a 2xx status code
func (o *GetPostgresRotatingSecretConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get postgres rotating secret config default response has a 3xx status code
func (o *GetPostgresRotatingSecretConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get postgres rotating secret config default response has a 4xx status code
func (o *GetPostgresRotatingSecretConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get postgres rotating secret config default response has a 5xx status code
func (o *GetPostgresRotatingSecretConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get postgres rotating secret config default response a status code equal to that given
func (o *GetPostgresRotatingSecretConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get postgres rotating secret config default response
func (o *GetPostgresRotatingSecretConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetPostgresRotatingSecretConfigDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/postgres/secret/{name}][%d] GetPostgresRotatingSecretConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetPostgresRotatingSecretConfigDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/postgres/secret/{name}][%d] GetPostgresRotatingSecretConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetPostgresRotatingSecretConfigDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *GetPostgresRotatingSecretConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}