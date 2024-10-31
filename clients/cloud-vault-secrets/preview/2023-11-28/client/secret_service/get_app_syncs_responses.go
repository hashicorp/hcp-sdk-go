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

// GetAppSyncsReader is a Reader for the GetAppSyncs structure.
type GetAppSyncsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppSyncsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppSyncsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAppSyncsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppSyncsOK creates a GetAppSyncsOK with default headers values
func NewGetAppSyncsOK() *GetAppSyncsOK {
	return &GetAppSyncsOK{}
}

/*
GetAppSyncsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAppSyncsOK struct {
	Payload *models.Secrets20231128GetAppSyncsResponse
}

// IsSuccess returns true when this get app syncs o k response has a 2xx status code
func (o *GetAppSyncsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app syncs o k response has a 3xx status code
func (o *GetAppSyncsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app syncs o k response has a 4xx status code
func (o *GetAppSyncsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app syncs o k response has a 5xx status code
func (o *GetAppSyncsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app syncs o k response a status code equal to that given
func (o *GetAppSyncsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app syncs o k response
func (o *GetAppSyncsOK) Code() int {
	return 200
}

func (o *GetAppSyncsOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{name}/syncs][%d] getAppSyncsOK  %+v", 200, o.Payload)
}

func (o *GetAppSyncsOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{name}/syncs][%d] getAppSyncsOK  %+v", 200, o.Payload)
}

func (o *GetAppSyncsOK) GetPayload() *models.Secrets20231128GetAppSyncsResponse {
	return o.Payload
}

func (o *GetAppSyncsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128GetAppSyncsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSyncsDefault creates a GetAppSyncsDefault with default headers values
func NewGetAppSyncsDefault(code int) *GetAppSyncsDefault {
	return &GetAppSyncsDefault{
		_statusCode: code,
	}
}

/*
GetAppSyncsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetAppSyncsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this get app syncs default response has a 2xx status code
func (o *GetAppSyncsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get app syncs default response has a 3xx status code
func (o *GetAppSyncsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get app syncs default response has a 4xx status code
func (o *GetAppSyncsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get app syncs default response has a 5xx status code
func (o *GetAppSyncsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get app syncs default response a status code equal to that given
func (o *GetAppSyncsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get app syncs default response
func (o *GetAppSyncsDefault) Code() int {
	return o._statusCode
}

func (o *GetAppSyncsDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{name}/syncs][%d] GetAppSyncs default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppSyncsDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{name}/syncs][%d] GetAppSyncs default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppSyncsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *GetAppSyncsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}