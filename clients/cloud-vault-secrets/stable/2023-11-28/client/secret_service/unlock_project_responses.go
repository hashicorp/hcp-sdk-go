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

// UnlockProjectReader is a Reader for the UnlockProject structure.
type UnlockProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnlockProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnlockProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnlockProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnlockProjectOK creates a UnlockProjectOK with default headers values
func NewUnlockProjectOK() *UnlockProjectOK {
	return &UnlockProjectOK{}
}

/*
UnlockProjectOK describes a response with status code 200, with default header values.

A successful response.
*/
type UnlockProjectOK struct {
	Payload models.Secrets20231128UnlockProjectResponse
}

// IsSuccess returns true when this unlock project o k response has a 2xx status code
func (o *UnlockProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unlock project o k response has a 3xx status code
func (o *UnlockProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unlock project o k response has a 4xx status code
func (o *UnlockProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unlock project o k response has a 5xx status code
func (o *UnlockProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unlock project o k response a status code equal to that given
func (o *UnlockProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unlock project o k response
func (o *UnlockProjectOK) Code() int {
	return 200
}

func (o *UnlockProjectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/unlock][%d] unlockProjectOK %s", 200, payload)
}

func (o *UnlockProjectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/unlock][%d] unlockProjectOK %s", 200, payload)
}

func (o *UnlockProjectOK) GetPayload() models.Secrets20231128UnlockProjectResponse {
	return o.Payload
}

func (o *UnlockProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnlockProjectDefault creates a UnlockProjectDefault with default headers values
func NewUnlockProjectDefault(code int) *UnlockProjectDefault {
	return &UnlockProjectDefault{
		_statusCode: code,
	}
}

/*
UnlockProjectDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UnlockProjectDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this unlock project default response has a 2xx status code
func (o *UnlockProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unlock project default response has a 3xx status code
func (o *UnlockProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unlock project default response has a 4xx status code
func (o *UnlockProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unlock project default response has a 5xx status code
func (o *UnlockProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unlock project default response a status code equal to that given
func (o *UnlockProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unlock project default response
func (o *UnlockProjectDefault) Code() int {
	return o._statusCode
}

func (o *UnlockProjectDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/unlock][%d] UnlockProject default %s", o._statusCode, payload)
}

func (o *UnlockProjectDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/unlock][%d] UnlockProject default %s", o._statusCode, payload)
}

func (o *UnlockProjectDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *UnlockProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
