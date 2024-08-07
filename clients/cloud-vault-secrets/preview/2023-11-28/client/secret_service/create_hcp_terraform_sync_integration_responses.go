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

// CreateHcpTerraformSyncIntegrationReader is a Reader for the CreateHcpTerraformSyncIntegration structure.
type CreateHcpTerraformSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHcpTerraformSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateHcpTerraformSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateHcpTerraformSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateHcpTerraformSyncIntegrationOK creates a CreateHcpTerraformSyncIntegrationOK with default headers values
func NewCreateHcpTerraformSyncIntegrationOK() *CreateHcpTerraformSyncIntegrationOK {
	return &CreateHcpTerraformSyncIntegrationOK{}
}

/*
CreateHcpTerraformSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateHcpTerraformSyncIntegrationOK struct {
	Payload *models.Secrets20231128CreateSyncIntegrationResponse
}

// IsSuccess returns true when this create hcp terraform sync integration o k response has a 2xx status code
func (o *CreateHcpTerraformSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create hcp terraform sync integration o k response has a 3xx status code
func (o *CreateHcpTerraformSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create hcp terraform sync integration o k response has a 4xx status code
func (o *CreateHcpTerraformSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create hcp terraform sync integration o k response has a 5xx status code
func (o *CreateHcpTerraformSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create hcp terraform sync integration o k response a status code equal to that given
func (o *CreateHcpTerraformSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create hcp terraform sync integration o k response
func (o *CreateHcpTerraformSyncIntegrationOK) Code() int {
	return 200
}

func (o *CreateHcpTerraformSyncIntegrationOK) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/hcp-terraform][%d] createHcpTerraformSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *CreateHcpTerraformSyncIntegrationOK) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/hcp-terraform][%d] createHcpTerraformSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *CreateHcpTerraformSyncIntegrationOK) GetPayload() *models.Secrets20231128CreateSyncIntegrationResponse {
	return o.Payload
}

func (o *CreateHcpTerraformSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128CreateSyncIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHcpTerraformSyncIntegrationDefault creates a CreateHcpTerraformSyncIntegrationDefault with default headers values
func NewCreateHcpTerraformSyncIntegrationDefault(code int) *CreateHcpTerraformSyncIntegrationDefault {
	return &CreateHcpTerraformSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
CreateHcpTerraformSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateHcpTerraformSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this create hcp terraform sync integration default response has a 2xx status code
func (o *CreateHcpTerraformSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create hcp terraform sync integration default response has a 3xx status code
func (o *CreateHcpTerraformSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create hcp terraform sync integration default response has a 4xx status code
func (o *CreateHcpTerraformSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create hcp terraform sync integration default response has a 5xx status code
func (o *CreateHcpTerraformSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create hcp terraform sync integration default response a status code equal to that given
func (o *CreateHcpTerraformSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create hcp terraform sync integration default response
func (o *CreateHcpTerraformSyncIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *CreateHcpTerraformSyncIntegrationDefault) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/hcp-terraform][%d] CreateHcpTerraformSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *CreateHcpTerraformSyncIntegrationDefault) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/hcp-terraform][%d] CreateHcpTerraformSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *CreateHcpTerraformSyncIntegrationDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CreateHcpTerraformSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
