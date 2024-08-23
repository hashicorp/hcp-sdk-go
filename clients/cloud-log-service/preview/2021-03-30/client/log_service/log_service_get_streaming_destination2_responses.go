// Code generated by go-swagger; DO NOT EDIT.

package log_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-log-service/preview/2021-03-30/models"
)

// LogServiceGetStreamingDestination2Reader is a Reader for the LogServiceGetStreamingDestination2 structure.
type LogServiceGetStreamingDestination2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogServiceGetStreamingDestination2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogServiceGetStreamingDestination2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogServiceGetStreamingDestination2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogServiceGetStreamingDestination2OK creates a LogServiceGetStreamingDestination2OK with default headers values
func NewLogServiceGetStreamingDestination2OK() *LogServiceGetStreamingDestination2OK {
	return &LogServiceGetStreamingDestination2OK{}
}

/*
LogServiceGetStreamingDestination2OK describes a response with status code 200, with default header values.

A successful response.
*/
type LogServiceGetStreamingDestination2OK struct {
	Payload *models.LogService20210330GetStreamingDestinationResponse
}

// IsSuccess returns true when this log service get streaming destination2 o k response has a 2xx status code
func (o *LogServiceGetStreamingDestination2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log service get streaming destination2 o k response has a 3xx status code
func (o *LogServiceGetStreamingDestination2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log service get streaming destination2 o k response has a 4xx status code
func (o *LogServiceGetStreamingDestination2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this log service get streaming destination2 o k response has a 5xx status code
func (o *LogServiceGetStreamingDestination2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this log service get streaming destination2 o k response a status code equal to that given
func (o *LogServiceGetStreamingDestination2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the log service get streaming destination2 o k response
func (o *LogServiceGetStreamingDestination2OK) Code() int {
	return 200
}

func (o *LogServiceGetStreamingDestination2OK) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/resources/destinations/{destination_id}][%d] logServiceGetStreamingDestination2OK  %+v", 200, o.Payload)
}

func (o *LogServiceGetStreamingDestination2OK) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/resources/destinations/{destination_id}][%d] logServiceGetStreamingDestination2OK  %+v", 200, o.Payload)
}

func (o *LogServiceGetStreamingDestination2OK) GetPayload() *models.LogService20210330GetStreamingDestinationResponse {
	return o.Payload
}

func (o *LogServiceGetStreamingDestination2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogService20210330GetStreamingDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogServiceGetStreamingDestination2Default creates a LogServiceGetStreamingDestination2Default with default headers values
func NewLogServiceGetStreamingDestination2Default(code int) *LogServiceGetStreamingDestination2Default {
	return &LogServiceGetStreamingDestination2Default{
		_statusCode: code,
	}
}

/*
LogServiceGetStreamingDestination2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LogServiceGetStreamingDestination2Default struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this log service get streaming destination2 default response has a 2xx status code
func (o *LogServiceGetStreamingDestination2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this log service get streaming destination2 default response has a 3xx status code
func (o *LogServiceGetStreamingDestination2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this log service get streaming destination2 default response has a 4xx status code
func (o *LogServiceGetStreamingDestination2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this log service get streaming destination2 default response has a 5xx status code
func (o *LogServiceGetStreamingDestination2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this log service get streaming destination2 default response a status code equal to that given
func (o *LogServiceGetStreamingDestination2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the log service get streaming destination2 default response
func (o *LogServiceGetStreamingDestination2Default) Code() int {
	return o._statusCode
}

func (o *LogServiceGetStreamingDestination2Default) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/resources/destinations/{destination_id}][%d] LogService_GetStreamingDestination2 default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceGetStreamingDestination2Default) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/resources/destinations/{destination_id}][%d] LogService_GetStreamingDestination2 default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceGetStreamingDestination2Default) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *LogServiceGetStreamingDestination2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}