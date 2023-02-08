// Code generated by go-swagger; DO NOT EDIT.

package vault_link_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-link-service/preview/2022-11-07/models"
)

// RotateLinkedClusterCredentialsReader is a Reader for the RotateLinkedClusterCredentials structure.
type RotateLinkedClusterCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RotateLinkedClusterCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRotateLinkedClusterCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRotateLinkedClusterCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRotateLinkedClusterCredentialsOK creates a RotateLinkedClusterCredentialsOK with default headers values
func NewRotateLinkedClusterCredentialsOK() *RotateLinkedClusterCredentialsOK {
	return &RotateLinkedClusterCredentialsOK{}
}

/*
RotateLinkedClusterCredentialsOK describes a response with status code 200, with default header values.

A successful response.
*/
type RotateLinkedClusterCredentialsOK struct {
	Payload *models.HashicorpCloudVaultLink20221107RotateLinkedClusterCredentialsResponse
}

// IsSuccess returns true when this rotate linked cluster credentials o k response has a 2xx status code
func (o *RotateLinkedClusterCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rotate linked cluster credentials o k response has a 3xx status code
func (o *RotateLinkedClusterCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rotate linked cluster credentials o k response has a 4xx status code
func (o *RotateLinkedClusterCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rotate linked cluster credentials o k response has a 5xx status code
func (o *RotateLinkedClusterCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rotate linked cluster credentials o k response a status code equal to that given
func (o *RotateLinkedClusterCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *RotateLinkedClusterCredentialsOK) Error() string {
	return fmt.Sprintf("[POST /vault-link/2022-11-07/organizations/{location.organizationId}/projects/{location.projectId}/rotate-credentials/{clusterId}][%d] rotateLinkedClusterCredentialsOK  %+v", 200, o.Payload)
}

func (o *RotateLinkedClusterCredentialsOK) String() string {
	return fmt.Sprintf("[POST /vault-link/2022-11-07/organizations/{location.organizationId}/projects/{location.projectId}/rotate-credentials/{clusterId}][%d] rotateLinkedClusterCredentialsOK  %+v", 200, o.Payload)
}

func (o *RotateLinkedClusterCredentialsOK) GetPayload() *models.HashicorpCloudVaultLink20221107RotateLinkedClusterCredentialsResponse {
	return o.Payload
}

func (o *RotateLinkedClusterCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVaultLink20221107RotateLinkedClusterCredentialsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateLinkedClusterCredentialsDefault creates a RotateLinkedClusterCredentialsDefault with default headers values
func NewRotateLinkedClusterCredentialsDefault(code int) *RotateLinkedClusterCredentialsDefault {
	return &RotateLinkedClusterCredentialsDefault{
		_statusCode: code,
	}
}

/*
RotateLinkedClusterCredentialsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RotateLinkedClusterCredentialsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the rotate linked cluster credentials default response
func (o *RotateLinkedClusterCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this rotate linked cluster credentials default response has a 2xx status code
func (o *RotateLinkedClusterCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rotate linked cluster credentials default response has a 3xx status code
func (o *RotateLinkedClusterCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rotate linked cluster credentials default response has a 4xx status code
func (o *RotateLinkedClusterCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rotate linked cluster credentials default response has a 5xx status code
func (o *RotateLinkedClusterCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rotate linked cluster credentials default response a status code equal to that given
func (o *RotateLinkedClusterCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RotateLinkedClusterCredentialsDefault) Error() string {
	return fmt.Sprintf("[POST /vault-link/2022-11-07/organizations/{location.organizationId}/projects/{location.projectId}/rotate-credentials/{clusterId}][%d] RotateLinkedClusterCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *RotateLinkedClusterCredentialsDefault) String() string {
	return fmt.Sprintf("[POST /vault-link/2022-11-07/organizations/{location.organizationId}/projects/{location.projectId}/rotate-credentials/{clusterId}][%d] RotateLinkedClusterCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *RotateLinkedClusterCredentialsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *RotateLinkedClusterCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
