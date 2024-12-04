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

// UpdateGcpServiceAccountKeyRotatingSecretReader is a Reader for the UpdateGcpServiceAccountKeyRotatingSecret structure.
type UpdateGcpServiceAccountKeyRotatingSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGcpServiceAccountKeyRotatingSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGcpServiceAccountKeyRotatingSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateGcpServiceAccountKeyRotatingSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateGcpServiceAccountKeyRotatingSecretOK creates a UpdateGcpServiceAccountKeyRotatingSecretOK with default headers values
func NewUpdateGcpServiceAccountKeyRotatingSecretOK() *UpdateGcpServiceAccountKeyRotatingSecretOK {
	return &UpdateGcpServiceAccountKeyRotatingSecretOK{}
}

/*
UpdateGcpServiceAccountKeyRotatingSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateGcpServiceAccountKeyRotatingSecretOK struct {
	Payload *models.Secrets20231128UpdateGcpServiceAccountKeyRotatingSecretResponse
}

// IsSuccess returns true when this update gcp service account key rotating secret o k response has a 2xx status code
func (o *UpdateGcpServiceAccountKeyRotatingSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update gcp service account key rotating secret o k response has a 3xx status code
func (o *UpdateGcpServiceAccountKeyRotatingSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gcp service account key rotating secret o k response has a 4xx status code
func (o *UpdateGcpServiceAccountKeyRotatingSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update gcp service account key rotating secret o k response has a 5xx status code
func (o *UpdateGcpServiceAccountKeyRotatingSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update gcp service account key rotating secret o k response a status code equal to that given
func (o *UpdateGcpServiceAccountKeyRotatingSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update gcp service account key rotating secret o k response
func (o *UpdateGcpServiceAccountKeyRotatingSecretOK) Code() int {
	return 200
}

func (o *UpdateGcpServiceAccountKeyRotatingSecretOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/gcp/secret/{name}][%d] updateGcpServiceAccountKeyRotatingSecretOK %s", 200, payload)
}

func (o *UpdateGcpServiceAccountKeyRotatingSecretOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/gcp/secret/{name}][%d] updateGcpServiceAccountKeyRotatingSecretOK %s", 200, payload)
}

func (o *UpdateGcpServiceAccountKeyRotatingSecretOK) GetPayload() *models.Secrets20231128UpdateGcpServiceAccountKeyRotatingSecretResponse {
	return o.Payload
}

func (o *UpdateGcpServiceAccountKeyRotatingSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128UpdateGcpServiceAccountKeyRotatingSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGcpServiceAccountKeyRotatingSecretDefault creates a UpdateGcpServiceAccountKeyRotatingSecretDefault with default headers values
func NewUpdateGcpServiceAccountKeyRotatingSecretDefault(code int) *UpdateGcpServiceAccountKeyRotatingSecretDefault {
	return &UpdateGcpServiceAccountKeyRotatingSecretDefault{
		_statusCode: code,
	}
}

/*
UpdateGcpServiceAccountKeyRotatingSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateGcpServiceAccountKeyRotatingSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this update gcp service account key rotating secret default response has a 2xx status code
func (o *UpdateGcpServiceAccountKeyRotatingSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update gcp service account key rotating secret default response has a 3xx status code
func (o *UpdateGcpServiceAccountKeyRotatingSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update gcp service account key rotating secret default response has a 4xx status code
func (o *UpdateGcpServiceAccountKeyRotatingSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update gcp service account key rotating secret default response has a 5xx status code
func (o *UpdateGcpServiceAccountKeyRotatingSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update gcp service account key rotating secret default response a status code equal to that given
func (o *UpdateGcpServiceAccountKeyRotatingSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update gcp service account key rotating secret default response
func (o *UpdateGcpServiceAccountKeyRotatingSecretDefault) Code() int {
	return o._statusCode
}

func (o *UpdateGcpServiceAccountKeyRotatingSecretDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/gcp/secret/{name}][%d] UpdateGcpServiceAccountKeyRotatingSecret default %s", o._statusCode, payload)
}

func (o *UpdateGcpServiceAccountKeyRotatingSecretDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/gcp/secret/{name}][%d] UpdateGcpServiceAccountKeyRotatingSecret default %s", o._statusCode, payload)
}

func (o *UpdateGcpServiceAccountKeyRotatingSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *UpdateGcpServiceAccountKeyRotatingSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
