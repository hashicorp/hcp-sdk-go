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

// LogServiceCreateStreamingDestinationFilterReader is a Reader for the LogServiceCreateStreamingDestinationFilter structure.
type LogServiceCreateStreamingDestinationFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogServiceCreateStreamingDestinationFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogServiceCreateStreamingDestinationFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogServiceCreateStreamingDestinationFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogServiceCreateStreamingDestinationFilterOK creates a LogServiceCreateStreamingDestinationFilterOK with default headers values
func NewLogServiceCreateStreamingDestinationFilterOK() *LogServiceCreateStreamingDestinationFilterOK {
	return &LogServiceCreateStreamingDestinationFilterOK{}
}

/*
LogServiceCreateStreamingDestinationFilterOK describes a response with status code 200, with default header values.

A successful response.
*/
type LogServiceCreateStreamingDestinationFilterOK struct {
	Payload interface{}
}

// IsSuccess returns true when this log service create streaming destination filter o k response has a 2xx status code
func (o *LogServiceCreateStreamingDestinationFilterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log service create streaming destination filter o k response has a 3xx status code
func (o *LogServiceCreateStreamingDestinationFilterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log service create streaming destination filter o k response has a 4xx status code
func (o *LogServiceCreateStreamingDestinationFilterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this log service create streaming destination filter o k response has a 5xx status code
func (o *LogServiceCreateStreamingDestinationFilterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this log service create streaming destination filter o k response a status code equal to that given
func (o *LogServiceCreateStreamingDestinationFilterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the log service create streaming destination filter o k response
func (o *LogServiceCreateStreamingDestinationFilterOK) Code() int {
	return 200
}

func (o *LogServiceCreateStreamingDestinationFilterOK) Error() string {
	return fmt.Sprintf("[POST /logs/2021-03-30/organizations/{organization_id}/resources/destinations/{destination_id}/filter][%d] logServiceCreateStreamingDestinationFilterOK  %+v", 200, o.Payload)
}

func (o *LogServiceCreateStreamingDestinationFilterOK) String() string {
	return fmt.Sprintf("[POST /logs/2021-03-30/organizations/{organization_id}/resources/destinations/{destination_id}/filter][%d] logServiceCreateStreamingDestinationFilterOK  %+v", 200, o.Payload)
}

func (o *LogServiceCreateStreamingDestinationFilterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *LogServiceCreateStreamingDestinationFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogServiceCreateStreamingDestinationFilterDefault creates a LogServiceCreateStreamingDestinationFilterDefault with default headers values
func NewLogServiceCreateStreamingDestinationFilterDefault(code int) *LogServiceCreateStreamingDestinationFilterDefault {
	return &LogServiceCreateStreamingDestinationFilterDefault{
		_statusCode: code,
	}
}

/*
LogServiceCreateStreamingDestinationFilterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LogServiceCreateStreamingDestinationFilterDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this log service create streaming destination filter default response has a 2xx status code
func (o *LogServiceCreateStreamingDestinationFilterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this log service create streaming destination filter default response has a 3xx status code
func (o *LogServiceCreateStreamingDestinationFilterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this log service create streaming destination filter default response has a 4xx status code
func (o *LogServiceCreateStreamingDestinationFilterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this log service create streaming destination filter default response has a 5xx status code
func (o *LogServiceCreateStreamingDestinationFilterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this log service create streaming destination filter default response a status code equal to that given
func (o *LogServiceCreateStreamingDestinationFilterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the log service create streaming destination filter default response
func (o *LogServiceCreateStreamingDestinationFilterDefault) Code() int {
	return o._statusCode
}

func (o *LogServiceCreateStreamingDestinationFilterDefault) Error() string {
	return fmt.Sprintf("[POST /logs/2021-03-30/organizations/{organization_id}/resources/destinations/{destination_id}/filter][%d] LogService_CreateStreamingDestinationFilter default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceCreateStreamingDestinationFilterDefault) String() string {
	return fmt.Sprintf("[POST /logs/2021-03-30/organizations/{organization_id}/resources/destinations/{destination_id}/filter][%d] LogService_CreateStreamingDestinationFilter default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceCreateStreamingDestinationFilterDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *LogServiceCreateStreamingDestinationFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
