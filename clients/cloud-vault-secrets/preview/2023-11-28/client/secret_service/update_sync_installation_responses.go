// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
)

// UpdateSyncInstallationReader is a Reader for the UpdateSyncInstallation structure.
type UpdateSyncInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSyncInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSyncInstallationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateSyncInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSyncInstallationOK creates a UpdateSyncInstallationOK with default headers values
func NewUpdateSyncInstallationOK() *UpdateSyncInstallationOK {
	return &UpdateSyncInstallationOK{}
}

/*
UpdateSyncInstallationOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateSyncInstallationOK struct {
	Payload *models.Secrets20231128UpdateSyncInstallationResponse
}

// IsSuccess returns true when this update sync installation o k response has a 2xx status code
func (o *UpdateSyncInstallationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update sync installation o k response has a 3xx status code
func (o *UpdateSyncInstallationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update sync installation o k response has a 4xx status code
func (o *UpdateSyncInstallationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update sync installation o k response has a 5xx status code
func (o *UpdateSyncInstallationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update sync installation o k response a status code equal to that given
func (o *UpdateSyncInstallationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update sync installation o k response
func (o *UpdateSyncInstallationOK) Code() int {
	return 200
}

func (o *UpdateSyncInstallationOK) Error() string {
	return fmt.Sprintf("[PATCH /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/installations/name][%d] updateSyncInstallationOK  %+v", 200, o.Payload)
}

func (o *UpdateSyncInstallationOK) String() string {
	return fmt.Sprintf("[PATCH /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/installations/name][%d] updateSyncInstallationOK  %+v", 200, o.Payload)
}

func (o *UpdateSyncInstallationOK) GetPayload() *models.Secrets20231128UpdateSyncInstallationResponse {
	return o.Payload
}

func (o *UpdateSyncInstallationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128UpdateSyncInstallationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSyncInstallationDefault creates a UpdateSyncInstallationDefault with default headers values
func NewUpdateSyncInstallationDefault(code int) *UpdateSyncInstallationDefault {
	return &UpdateSyncInstallationDefault{
		_statusCode: code,
	}
}

/*
UpdateSyncInstallationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateSyncInstallationDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this update sync installation default response has a 2xx status code
func (o *UpdateSyncInstallationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update sync installation default response has a 3xx status code
func (o *UpdateSyncInstallationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update sync installation default response has a 4xx status code
func (o *UpdateSyncInstallationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update sync installation default response has a 5xx status code
func (o *UpdateSyncInstallationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update sync installation default response a status code equal to that given
func (o *UpdateSyncInstallationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update sync installation default response
func (o *UpdateSyncInstallationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSyncInstallationDefault) Error() string {
	return fmt.Sprintf("[PATCH /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/installations/name][%d] UpdateSyncInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSyncInstallationDefault) String() string {
	return fmt.Sprintf("[PATCH /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/installations/name][%d] UpdateSyncInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSyncInstallationDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *UpdateSyncInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}