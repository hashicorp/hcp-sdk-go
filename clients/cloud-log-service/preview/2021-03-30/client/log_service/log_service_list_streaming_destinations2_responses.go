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

// LogServiceListStreamingDestinations2Reader is a Reader for the LogServiceListStreamingDestinations2 structure.
type LogServiceListStreamingDestinations2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogServiceListStreamingDestinations2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogServiceListStreamingDestinations2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogServiceListStreamingDestinations2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogServiceListStreamingDestinations2OK creates a LogServiceListStreamingDestinations2OK with default headers values
func NewLogServiceListStreamingDestinations2OK() *LogServiceListStreamingDestinations2OK {
	return &LogServiceListStreamingDestinations2OK{}
}

/*
LogServiceListStreamingDestinations2OK describes a response with status code 200, with default header values.

A successful response.
*/
type LogServiceListStreamingDestinations2OK struct {
	Payload *models.LogService20210330ListStreamingDestinationsResponse
}

// IsSuccess returns true when this log service list streaming destinations2 o k response has a 2xx status code
func (o *LogServiceListStreamingDestinations2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log service list streaming destinations2 o k response has a 3xx status code
func (o *LogServiceListStreamingDestinations2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log service list streaming destinations2 o k response has a 4xx status code
func (o *LogServiceListStreamingDestinations2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this log service list streaming destinations2 o k response has a 5xx status code
func (o *LogServiceListStreamingDestinations2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this log service list streaming destinations2 o k response a status code equal to that given
func (o *LogServiceListStreamingDestinations2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the log service list streaming destinations2 o k response
func (o *LogServiceListStreamingDestinations2OK) Code() int {
	return 200
}

func (o *LogServiceListStreamingDestinations2OK) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/resources/destinations][%d] logServiceListStreamingDestinations2OK  %+v", 200, o.Payload)
}

func (o *LogServiceListStreamingDestinations2OK) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/resources/destinations][%d] logServiceListStreamingDestinations2OK  %+v", 200, o.Payload)
}

func (o *LogServiceListStreamingDestinations2OK) GetPayload() *models.LogService20210330ListStreamingDestinationsResponse {
	return o.Payload
}

func (o *LogServiceListStreamingDestinations2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogService20210330ListStreamingDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogServiceListStreamingDestinations2Default creates a LogServiceListStreamingDestinations2Default with default headers values
func NewLogServiceListStreamingDestinations2Default(code int) *LogServiceListStreamingDestinations2Default {
	return &LogServiceListStreamingDestinations2Default{
		_statusCode: code,
	}
}

/*
LogServiceListStreamingDestinations2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LogServiceListStreamingDestinations2Default struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this log service list streaming destinations2 default response has a 2xx status code
func (o *LogServiceListStreamingDestinations2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this log service list streaming destinations2 default response has a 3xx status code
func (o *LogServiceListStreamingDestinations2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this log service list streaming destinations2 default response has a 4xx status code
func (o *LogServiceListStreamingDestinations2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this log service list streaming destinations2 default response has a 5xx status code
func (o *LogServiceListStreamingDestinations2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this log service list streaming destinations2 default response a status code equal to that given
func (o *LogServiceListStreamingDestinations2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the log service list streaming destinations2 default response
func (o *LogServiceListStreamingDestinations2Default) Code() int {
	return o._statusCode
}

func (o *LogServiceListStreamingDestinations2Default) Error() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/resources/destinations][%d] LogService_ListStreamingDestinations2 default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceListStreamingDestinations2Default) String() string {
	return fmt.Sprintf("[GET /logs/2021-03-30/organizations/{location.organization_id}/resources/destinations][%d] LogService_ListStreamingDestinations2 default  %+v", o._statusCode, o.Payload)
}

func (o *LogServiceListStreamingDestinations2Default) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *LogServiceListStreamingDestinations2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
