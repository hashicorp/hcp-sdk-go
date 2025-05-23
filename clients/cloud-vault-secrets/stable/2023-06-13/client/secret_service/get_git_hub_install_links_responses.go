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

// GetGitHubInstallLinksReader is a Reader for the GetGitHubInstallLinks structure.
type GetGitHubInstallLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGitHubInstallLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGitHubInstallLinksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetGitHubInstallLinksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGitHubInstallLinksOK creates a GetGitHubInstallLinksOK with default headers values
func NewGetGitHubInstallLinksOK() *GetGitHubInstallLinksOK {
	return &GetGitHubInstallLinksOK{}
}

/*
GetGitHubInstallLinksOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetGitHubInstallLinksOK struct {
	Payload *models.Secrets20230613GetGitHubInstallLinksResponse
}

// IsSuccess returns true when this get git hub install links o k response has a 2xx status code
func (o *GetGitHubInstallLinksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get git hub install links o k response has a 3xx status code
func (o *GetGitHubInstallLinksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get git hub install links o k response has a 4xx status code
func (o *GetGitHubInstallLinksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get git hub install links o k response has a 5xx status code
func (o *GetGitHubInstallLinksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get git hub install links o k response a status code equal to that given
func (o *GetGitHubInstallLinksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get git hub install links o k response
func (o *GetGitHubInstallLinksOK) Code() int {
	return 200
}

func (o *GetGitHubInstallLinksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/github/links][%d] getGitHubInstallLinksOK %s", 200, payload)
}

func (o *GetGitHubInstallLinksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/github/links][%d] getGitHubInstallLinksOK %s", 200, payload)
}

func (o *GetGitHubInstallLinksOK) GetPayload() *models.Secrets20230613GetGitHubInstallLinksResponse {
	return o.Payload
}

func (o *GetGitHubInstallLinksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613GetGitHubInstallLinksResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGitHubInstallLinksDefault creates a GetGitHubInstallLinksDefault with default headers values
func NewGetGitHubInstallLinksDefault(code int) *GetGitHubInstallLinksDefault {
	return &GetGitHubInstallLinksDefault{
		_statusCode: code,
	}
}

/*
GetGitHubInstallLinksDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetGitHubInstallLinksDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get git hub install links default response has a 2xx status code
func (o *GetGitHubInstallLinksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get git hub install links default response has a 3xx status code
func (o *GetGitHubInstallLinksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get git hub install links default response has a 4xx status code
func (o *GetGitHubInstallLinksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get git hub install links default response has a 5xx status code
func (o *GetGitHubInstallLinksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get git hub install links default response a status code equal to that given
func (o *GetGitHubInstallLinksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get git hub install links default response
func (o *GetGitHubInstallLinksDefault) Code() int {
	return o._statusCode
}

func (o *GetGitHubInstallLinksDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/github/links][%d] GetGitHubInstallLinks default %s", o._statusCode, payload)
}

func (o *GetGitHubInstallLinksDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/github/links][%d] GetGitHubInstallLinks default %s", o._statusCode, payload)
}

func (o *GetGitHubInstallLinksDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetGitHubInstallLinksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
