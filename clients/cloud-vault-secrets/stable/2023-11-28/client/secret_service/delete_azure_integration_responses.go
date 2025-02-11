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

// DeleteAzureIntegrationReader is a Reader for the DeleteAzureIntegration structure.
type DeleteAzureIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAzureIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAzureIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAzureIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAzureIntegrationOK creates a DeleteAzureIntegrationOK with default headers values
func NewDeleteAzureIntegrationOK() *DeleteAzureIntegrationOK {
	return &DeleteAzureIntegrationOK{}
}

/*
DeleteAzureIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAzureIntegrationOK struct {
	Payload models.Secrets20231128DeleteAzureIntegrationResponse
}

// IsSuccess returns true when this delete azure integration o k response has a 2xx status code
func (o *DeleteAzureIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete azure integration o k response has a 3xx status code
func (o *DeleteAzureIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete azure integration o k response has a 4xx status code
func (o *DeleteAzureIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete azure integration o k response has a 5xx status code
func (o *DeleteAzureIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete azure integration o k response a status code equal to that given
func (o *DeleteAzureIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete azure integration o k response
func (o *DeleteAzureIntegrationOK) Code() int {
	return 200
}

func (o *DeleteAzureIntegrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/azure/config/{name}][%d] deleteAzureIntegrationOK %s", 200, payload)
}

func (o *DeleteAzureIntegrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/azure/config/{name}][%d] deleteAzureIntegrationOK %s", 200, payload)
}

func (o *DeleteAzureIntegrationOK) GetPayload() models.Secrets20231128DeleteAzureIntegrationResponse {
	return o.Payload
}

func (o *DeleteAzureIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAzureIntegrationDefault creates a DeleteAzureIntegrationDefault with default headers values
func NewDeleteAzureIntegrationDefault(code int) *DeleteAzureIntegrationDefault {
	return &DeleteAzureIntegrationDefault{
		_statusCode: code,
	}
}

/*
DeleteAzureIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteAzureIntegrationDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this delete azure integration default response has a 2xx status code
func (o *DeleteAzureIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete azure integration default response has a 3xx status code
func (o *DeleteAzureIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete azure integration default response has a 4xx status code
func (o *DeleteAzureIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete azure integration default response has a 5xx status code
func (o *DeleteAzureIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete azure integration default response a status code equal to that given
func (o *DeleteAzureIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete azure integration default response
func (o *DeleteAzureIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAzureIntegrationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/azure/config/{name}][%d] DeleteAzureIntegration default %s", o._statusCode, payload)
}

func (o *DeleteAzureIntegrationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/azure/config/{name}][%d] DeleteAzureIntegration default %s", o._statusCode, payload)
}

func (o *DeleteAzureIntegrationDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DeleteAzureIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
