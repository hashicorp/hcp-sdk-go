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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-06-13/models"
)

// GetVercelInstallationLinkReader is a Reader for the GetVercelInstallationLink structure.
type GetVercelInstallationLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVercelInstallationLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVercelInstallationLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVercelInstallationLinkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVercelInstallationLinkOK creates a GetVercelInstallationLinkOK with default headers values
func NewGetVercelInstallationLinkOK() *GetVercelInstallationLinkOK {
	return &GetVercelInstallationLinkOK{}
}

/*
GetVercelInstallationLinkOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetVercelInstallationLinkOK struct {
	Payload *models.Secrets20230613GetVercelInstallationLinkResponse
}

// IsSuccess returns true when this get vercel installation link o k response has a 2xx status code
func (o *GetVercelInstallationLinkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vercel installation link o k response has a 3xx status code
func (o *GetVercelInstallationLinkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vercel installation link o k response has a 4xx status code
func (o *GetVercelInstallationLinkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vercel installation link o k response has a 5xx status code
func (o *GetVercelInstallationLinkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vercel installation link o k response a status code equal to that given
func (o *GetVercelInstallationLinkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get vercel installation link o k response
func (o *GetVercelInstallationLinkOK) Code() int {
	return 200
}

func (o *GetVercelInstallationLinkOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel/link][%d] getVercelInstallationLinkOK %s", 200, payload)
}

func (o *GetVercelInstallationLinkOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel/link][%d] getVercelInstallationLinkOK %s", 200, payload)
}

func (o *GetVercelInstallationLinkOK) GetPayload() *models.Secrets20230613GetVercelInstallationLinkResponse {
	return o.Payload
}

func (o *GetVercelInstallationLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613GetVercelInstallationLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVercelInstallationLinkDefault creates a GetVercelInstallationLinkDefault with default headers values
func NewGetVercelInstallationLinkDefault(code int) *GetVercelInstallationLinkDefault {
	return &GetVercelInstallationLinkDefault{
		_statusCode: code,
	}
}

/*
GetVercelInstallationLinkDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetVercelInstallationLinkDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get vercel installation link default response has a 2xx status code
func (o *GetVercelInstallationLinkDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get vercel installation link default response has a 3xx status code
func (o *GetVercelInstallationLinkDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get vercel installation link default response has a 4xx status code
func (o *GetVercelInstallationLinkDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get vercel installation link default response has a 5xx status code
func (o *GetVercelInstallationLinkDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get vercel installation link default response a status code equal to that given
func (o *GetVercelInstallationLinkDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get vercel installation link default response
func (o *GetVercelInstallationLinkDefault) Code() int {
	return o._statusCode
}

func (o *GetVercelInstallationLinkDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel/link][%d] GetVercelInstallationLink default %s", o._statusCode, payload)
}

func (o *GetVercelInstallationLinkDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel/link][%d] GetVercelInstallationLink default %s", o._statusCode, payload)
}

func (o *GetVercelInstallationLinkDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetVercelInstallationLinkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
