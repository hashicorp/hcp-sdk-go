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

// DeleteGcpDynamicSecretReader is a Reader for the DeleteGcpDynamicSecret structure.
type DeleteGcpDynamicSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGcpDynamicSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteGcpDynamicSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteGcpDynamicSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteGcpDynamicSecretOK creates a DeleteGcpDynamicSecretOK with default headers values
func NewDeleteGcpDynamicSecretOK() *DeleteGcpDynamicSecretOK {
	return &DeleteGcpDynamicSecretOK{}
}

/*
DeleteGcpDynamicSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteGcpDynamicSecretOK struct {
	Payload models.Secrets20231128DeleteGcpDynamicSecretResponse
}

// IsSuccess returns true when this delete gcp dynamic secret o k response has a 2xx status code
func (o *DeleteGcpDynamicSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete gcp dynamic secret o k response has a 3xx status code
func (o *DeleteGcpDynamicSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gcp dynamic secret o k response has a 4xx status code
func (o *DeleteGcpDynamicSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete gcp dynamic secret o k response has a 5xx status code
func (o *DeleteGcpDynamicSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete gcp dynamic secret o k response a status code equal to that given
func (o *DeleteGcpDynamicSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete gcp dynamic secret o k response
func (o *DeleteGcpDynamicSecretOK) Code() int {
	return 200
}

func (o *DeleteGcpDynamicSecretOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret/{name}][%d] deleteGcpDynamicSecretOK %s", 200, payload)
}

func (o *DeleteGcpDynamicSecretOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret/{name}][%d] deleteGcpDynamicSecretOK %s", 200, payload)
}

func (o *DeleteGcpDynamicSecretOK) GetPayload() models.Secrets20231128DeleteGcpDynamicSecretResponse {
	return o.Payload
}

func (o *DeleteGcpDynamicSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGcpDynamicSecretDefault creates a DeleteGcpDynamicSecretDefault with default headers values
func NewDeleteGcpDynamicSecretDefault(code int) *DeleteGcpDynamicSecretDefault {
	return &DeleteGcpDynamicSecretDefault{
		_statusCode: code,
	}
}

/*
DeleteGcpDynamicSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteGcpDynamicSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this delete gcp dynamic secret default response has a 2xx status code
func (o *DeleteGcpDynamicSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete gcp dynamic secret default response has a 3xx status code
func (o *DeleteGcpDynamicSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete gcp dynamic secret default response has a 4xx status code
func (o *DeleteGcpDynamicSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete gcp dynamic secret default response has a 5xx status code
func (o *DeleteGcpDynamicSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete gcp dynamic secret default response a status code equal to that given
func (o *DeleteGcpDynamicSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete gcp dynamic secret default response
func (o *DeleteGcpDynamicSecretDefault) Code() int {
	return o._statusCode
}

func (o *DeleteGcpDynamicSecretDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret/{name}][%d] DeleteGcpDynamicSecret default %s", o._statusCode, payload)
}

func (o *DeleteGcpDynamicSecretDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret/{name}][%d] DeleteGcpDynamicSecret default %s", o._statusCode, payload)
}

func (o *DeleteGcpDynamicSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DeleteGcpDynamicSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
