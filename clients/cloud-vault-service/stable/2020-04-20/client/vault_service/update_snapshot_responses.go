// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-04-20/models"
)

// UpdateSnapshotReader is a Reader for the UpdateSnapshot structure.
type UpdateSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSnapshotOK creates a UpdateSnapshotOK with default headers values
func NewUpdateSnapshotOK() *UpdateSnapshotOK {
	return &UpdateSnapshotOK{}
}

/*
UpdateSnapshotOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateSnapshotOK struct {
	Payload *models.HashicorpCloudVault20200420UpdateSnapshotResponse
}

// IsSuccess returns true when this update snapshot o k response has a 2xx status code
func (o *UpdateSnapshotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update snapshot o k response has a 3xx status code
func (o *UpdateSnapshotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update snapshot o k response has a 4xx status code
func (o *UpdateSnapshotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update snapshot o k response has a 5xx status code
func (o *UpdateSnapshotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update snapshot o k response a status code equal to that given
func (o *UpdateSnapshotOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update snapshot o k response
func (o *UpdateSnapshotOK) Code() int {
	return 200
}

func (o *UpdateSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /vault/2020-04-20/organizations/{snapshot.location.organization_id}/projects/{snapshot.location.project_id}/snapshots/{snapshot.snapshot_id}][%d] updateSnapshotOK  %+v", 200, o.Payload)
}

func (o *UpdateSnapshotOK) String() string {
	return fmt.Sprintf("[POST /vault/2020-04-20/organizations/{snapshot.location.organization_id}/projects/{snapshot.location.project_id}/snapshots/{snapshot.snapshot_id}][%d] updateSnapshotOK  %+v", 200, o.Payload)
}

func (o *UpdateSnapshotOK) GetPayload() *models.HashicorpCloudVault20200420UpdateSnapshotResponse {
	return o.Payload
}

func (o *UpdateSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20200420UpdateSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSnapshotDefault creates a UpdateSnapshotDefault with default headers values
func NewUpdateSnapshotDefault(code int) *UpdateSnapshotDefault {
	return &UpdateSnapshotDefault{
		_statusCode: code,
	}
}

/*
UpdateSnapshotDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateSnapshotDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this update snapshot default response has a 2xx status code
func (o *UpdateSnapshotDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update snapshot default response has a 3xx status code
func (o *UpdateSnapshotDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update snapshot default response has a 4xx status code
func (o *UpdateSnapshotDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update snapshot default response has a 5xx status code
func (o *UpdateSnapshotDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update snapshot default response a status code equal to that given
func (o *UpdateSnapshotDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update snapshot default response
func (o *UpdateSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSnapshotDefault) Error() string {
	return fmt.Sprintf("[POST /vault/2020-04-20/organizations/{snapshot.location.organization_id}/projects/{snapshot.location.project_id}/snapshots/{snapshot.snapshot_id}][%d] UpdateSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSnapshotDefault) String() string {
	return fmt.Sprintf("[POST /vault/2020-04-20/organizations/{snapshot.location.organization_id}/projects/{snapshot.location.project_id}/snapshots/{snapshot.snapshot_id}][%d] UpdateSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSnapshotDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *UpdateSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
