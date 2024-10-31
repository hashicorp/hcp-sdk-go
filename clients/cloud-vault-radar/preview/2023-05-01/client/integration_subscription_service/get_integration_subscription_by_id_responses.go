// Code generated by go-swagger; DO NOT EDIT.

package integration_subscription_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-radar/preview/2023-05-01/models"
)

// GetIntegrationSubscriptionByIDReader is a Reader for the GetIntegrationSubscriptionByID structure.
type GetIntegrationSubscriptionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationSubscriptionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationSubscriptionByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIntegrationSubscriptionByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIntegrationSubscriptionByIDOK creates a GetIntegrationSubscriptionByIDOK with default headers values
func NewGetIntegrationSubscriptionByIDOK() *GetIntegrationSubscriptionByIDOK {
	return &GetIntegrationSubscriptionByIDOK{}
}

/*
GetIntegrationSubscriptionByIDOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetIntegrationSubscriptionByIDOK struct {
	Payload *models.VaultRadar20230501IntegrationSubscription
}

// IsSuccess returns true when this get integration subscription by Id o k response has a 2xx status code
func (o *GetIntegrationSubscriptionByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integration subscription by Id o k response has a 3xx status code
func (o *GetIntegrationSubscriptionByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration subscription by Id o k response has a 4xx status code
func (o *GetIntegrationSubscriptionByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integration subscription by Id o k response has a 5xx status code
func (o *GetIntegrationSubscriptionByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration subscription by Id o k response a status code equal to that given
func (o *GetIntegrationSubscriptionByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get integration subscription by Id o k response
func (o *GetIntegrationSubscriptionByIDOK) Code() int {
	return 200
}

func (o *GetIntegrationSubscriptionByIDOK) Error() string {
	return fmt.Sprintf("[GET /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions/{id}][%d] getIntegrationSubscriptionByIdOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationSubscriptionByIDOK) String() string {
	return fmt.Sprintf("[GET /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions/{id}][%d] getIntegrationSubscriptionByIdOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationSubscriptionByIDOK) GetPayload() *models.VaultRadar20230501IntegrationSubscription {
	return o.Payload
}

func (o *GetIntegrationSubscriptionByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VaultRadar20230501IntegrationSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationSubscriptionByIDDefault creates a GetIntegrationSubscriptionByIDDefault with default headers values
func NewGetIntegrationSubscriptionByIDDefault(code int) *GetIntegrationSubscriptionByIDDefault {
	return &GetIntegrationSubscriptionByIDDefault{
		_statusCode: code,
	}
}

/*
GetIntegrationSubscriptionByIDDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetIntegrationSubscriptionByIDDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get integration subscription by ID default response has a 2xx status code
func (o *GetIntegrationSubscriptionByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get integration subscription by ID default response has a 3xx status code
func (o *GetIntegrationSubscriptionByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get integration subscription by ID default response has a 4xx status code
func (o *GetIntegrationSubscriptionByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get integration subscription by ID default response has a 5xx status code
func (o *GetIntegrationSubscriptionByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get integration subscription by ID default response a status code equal to that given
func (o *GetIntegrationSubscriptionByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get integration subscription by ID default response
func (o *GetIntegrationSubscriptionByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetIntegrationSubscriptionByIDDefault) Error() string {
	return fmt.Sprintf("[GET /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions/{id}][%d] GetIntegrationSubscriptionByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetIntegrationSubscriptionByIDDefault) String() string {
	return fmt.Sprintf("[GET /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions/{id}][%d] GetIntegrationSubscriptionByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetIntegrationSubscriptionByIDDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetIntegrationSubscriptionByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}