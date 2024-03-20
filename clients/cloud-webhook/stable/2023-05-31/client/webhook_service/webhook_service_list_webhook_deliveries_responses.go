// Code generated by go-swagger; DO NOT EDIT.

package webhook_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-webhook/stable/2023-05-31/models"
)

// WebhookServiceListWebhookDeliveriesReader is a Reader for the WebhookServiceListWebhookDeliveries structure.
type WebhookServiceListWebhookDeliveriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebhookServiceListWebhookDeliveriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebhookServiceListWebhookDeliveriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWebhookServiceListWebhookDeliveriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWebhookServiceListWebhookDeliveriesOK creates a WebhookServiceListWebhookDeliveriesOK with default headers values
func NewWebhookServiceListWebhookDeliveriesOK() *WebhookServiceListWebhookDeliveriesOK {
	return &WebhookServiceListWebhookDeliveriesOK{}
}

/*
WebhookServiceListWebhookDeliveriesOK describes a response with status code 200, with default header values.

A successful response.
*/
type WebhookServiceListWebhookDeliveriesOK struct {
	Payload *models.HashicorpCloudWebhookListWebhookDeliveriesResponse
}

// IsSuccess returns true when this webhook service list webhook deliveries o k response has a 2xx status code
func (o *WebhookServiceListWebhookDeliveriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webhook service list webhook deliveries o k response has a 3xx status code
func (o *WebhookServiceListWebhookDeliveriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webhook service list webhook deliveries o k response has a 4xx status code
func (o *WebhookServiceListWebhookDeliveriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webhook service list webhook deliveries o k response has a 5xx status code
func (o *WebhookServiceListWebhookDeliveriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webhook service list webhook deliveries o k response a status code equal to that given
func (o *WebhookServiceListWebhookDeliveriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webhook service list webhook deliveries o k response
func (o *WebhookServiceListWebhookDeliveriesOK) Code() int {
	return 200
}

func (o *WebhookServiceListWebhookDeliveriesOK) Error() string {
	return fmt.Sprintf("[GET /2023-05-31/{parent_resource_name}/deliveries][%d] webhookServiceListWebhookDeliveriesOK  %+v", 200, o.Payload)
}

func (o *WebhookServiceListWebhookDeliveriesOK) String() string {
	return fmt.Sprintf("[GET /2023-05-31/{parent_resource_name}/deliveries][%d] webhookServiceListWebhookDeliveriesOK  %+v", 200, o.Payload)
}

func (o *WebhookServiceListWebhookDeliveriesOK) GetPayload() *models.HashicorpCloudWebhookListWebhookDeliveriesResponse {
	return o.Payload
}

func (o *WebhookServiceListWebhookDeliveriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWebhookListWebhookDeliveriesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookServiceListWebhookDeliveriesDefault creates a WebhookServiceListWebhookDeliveriesDefault with default headers values
func NewWebhookServiceListWebhookDeliveriesDefault(code int) *WebhookServiceListWebhookDeliveriesDefault {
	return &WebhookServiceListWebhookDeliveriesDefault{
		_statusCode: code,
	}
}

/*
WebhookServiceListWebhookDeliveriesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WebhookServiceListWebhookDeliveriesDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this webhook service list webhook deliveries default response has a 2xx status code
func (o *WebhookServiceListWebhookDeliveriesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this webhook service list webhook deliveries default response has a 3xx status code
func (o *WebhookServiceListWebhookDeliveriesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this webhook service list webhook deliveries default response has a 4xx status code
func (o *WebhookServiceListWebhookDeliveriesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this webhook service list webhook deliveries default response has a 5xx status code
func (o *WebhookServiceListWebhookDeliveriesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this webhook service list webhook deliveries default response a status code equal to that given
func (o *WebhookServiceListWebhookDeliveriesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the webhook service list webhook deliveries default response
func (o *WebhookServiceListWebhookDeliveriesDefault) Code() int {
	return o._statusCode
}

func (o *WebhookServiceListWebhookDeliveriesDefault) Error() string {
	return fmt.Sprintf("[GET /2023-05-31/{parent_resource_name}/deliveries][%d] WebhookService_ListWebhookDeliveries default  %+v", o._statusCode, o.Payload)
}

func (o *WebhookServiceListWebhookDeliveriesDefault) String() string {
	return fmt.Sprintf("[GET /2023-05-31/{parent_resource_name}/deliveries][%d] WebhookService_ListWebhookDeliveries default  %+v", o._statusCode, o.Payload)
}

func (o *WebhookServiceListWebhookDeliveriesDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WebhookServiceListWebhookDeliveriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}