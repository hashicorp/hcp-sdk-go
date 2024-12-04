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

// GetGitHubRepositoriesReader is a Reader for the GetGitHubRepositories structure.
type GetGitHubRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGitHubRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGitHubRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetGitHubRepositoriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGitHubRepositoriesOK creates a GetGitHubRepositoriesOK with default headers values
func NewGetGitHubRepositoriesOK() *GetGitHubRepositoriesOK {
	return &GetGitHubRepositoriesOK{}
}

/*
GetGitHubRepositoriesOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetGitHubRepositoriesOK struct {
	Payload *models.Secrets20230613GetGitHubRepositoriesResponse
}

// IsSuccess returns true when this get git hub repositories o k response has a 2xx status code
func (o *GetGitHubRepositoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get git hub repositories o k response has a 3xx status code
func (o *GetGitHubRepositoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get git hub repositories o k response has a 4xx status code
func (o *GetGitHubRepositoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get git hub repositories o k response has a 5xx status code
func (o *GetGitHubRepositoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get git hub repositories o k response a status code equal to that given
func (o *GetGitHubRepositoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get git hub repositories o k response
func (o *GetGitHubRepositoriesOK) Code() int {
	return 200
}

func (o *GetGitHubRepositoriesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/github/repositories][%d] getGitHubRepositoriesOK %s", 200, payload)
}

func (o *GetGitHubRepositoriesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/github/repositories][%d] getGitHubRepositoriesOK %s", 200, payload)
}

func (o *GetGitHubRepositoriesOK) GetPayload() *models.Secrets20230613GetGitHubRepositoriesResponse {
	return o.Payload
}

func (o *GetGitHubRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613GetGitHubRepositoriesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGitHubRepositoriesDefault creates a GetGitHubRepositoriesDefault with default headers values
func NewGetGitHubRepositoriesDefault(code int) *GetGitHubRepositoriesDefault {
	return &GetGitHubRepositoriesDefault{
		_statusCode: code,
	}
}

/*
GetGitHubRepositoriesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetGitHubRepositoriesDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get git hub repositories default response has a 2xx status code
func (o *GetGitHubRepositoriesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get git hub repositories default response has a 3xx status code
func (o *GetGitHubRepositoriesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get git hub repositories default response has a 4xx status code
func (o *GetGitHubRepositoriesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get git hub repositories default response has a 5xx status code
func (o *GetGitHubRepositoriesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get git hub repositories default response a status code equal to that given
func (o *GetGitHubRepositoriesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get git hub repositories default response
func (o *GetGitHubRepositoriesDefault) Code() int {
	return o._statusCode
}

func (o *GetGitHubRepositoriesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/github/repositories][%d] GetGitHubRepositories default %s", o._statusCode, payload)
}

func (o *GetGitHubRepositoriesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/github/repositories][%d] GetGitHubRepositories default %s", o._statusCode, payload)
}

func (o *GetGitHubRepositoriesDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetGitHubRepositoriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
