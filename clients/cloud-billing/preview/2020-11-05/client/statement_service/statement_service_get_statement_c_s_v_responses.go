// Code generated by go-swagger; DO NOT EDIT.

package statement_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/models"
)

// StatementServiceGetStatementCSVReader is a Reader for the StatementServiceGetStatementCSV structure.
type StatementServiceGetStatementCSVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatementServiceGetStatementCSVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatementServiceGetStatementCSVOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatementServiceGetStatementCSVDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatementServiceGetStatementCSVOK creates a StatementServiceGetStatementCSVOK with default headers values
func NewStatementServiceGetStatementCSVOK() *StatementServiceGetStatementCSVOK {
	return &StatementServiceGetStatementCSVOK{}
}

/*
StatementServiceGetStatementCSVOK describes a response with status code 200, with default header values.

A successful response.
*/
type StatementServiceGetStatementCSVOK struct {
	Payload *models.Billing20201105GetStatementCSVResponse
}

// IsSuccess returns true when this statement service get statement c s v o k response has a 2xx status code
func (o *StatementServiceGetStatementCSVOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this statement service get statement c s v o k response has a 3xx status code
func (o *StatementServiceGetStatementCSVOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this statement service get statement c s v o k response has a 4xx status code
func (o *StatementServiceGetStatementCSVOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this statement service get statement c s v o k response has a 5xx status code
func (o *StatementServiceGetStatementCSVOK) IsServerError() bool {
	return false
}

// IsCode returns true when this statement service get statement c s v o k response a status code equal to that given
func (o *StatementServiceGetStatementCSVOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the statement service get statement c s v o k response
func (o *StatementServiceGetStatementCSVOK) Code() int {
	return 200
}

func (o *StatementServiceGetStatementCSVOK) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements/{statement_id}/csv][%d] statementServiceGetStatementCSVOK  %+v", 200, o.Payload)
}

func (o *StatementServiceGetStatementCSVOK) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements/{statement_id}/csv][%d] statementServiceGetStatementCSVOK  %+v", 200, o.Payload)
}

func (o *StatementServiceGetStatementCSVOK) GetPayload() *models.Billing20201105GetStatementCSVResponse {
	return o.Payload
}

func (o *StatementServiceGetStatementCSVOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105GetStatementCSVResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatementServiceGetStatementCSVDefault creates a StatementServiceGetStatementCSVDefault with default headers values
func NewStatementServiceGetStatementCSVDefault(code int) *StatementServiceGetStatementCSVDefault {
	return &StatementServiceGetStatementCSVDefault{
		_statusCode: code,
	}
}

/*
StatementServiceGetStatementCSVDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StatementServiceGetStatementCSVDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this statement service get statement c s v default response has a 2xx status code
func (o *StatementServiceGetStatementCSVDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this statement service get statement c s v default response has a 3xx status code
func (o *StatementServiceGetStatementCSVDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this statement service get statement c s v default response has a 4xx status code
func (o *StatementServiceGetStatementCSVDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this statement service get statement c s v default response has a 5xx status code
func (o *StatementServiceGetStatementCSVDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this statement service get statement c s v default response a status code equal to that given
func (o *StatementServiceGetStatementCSVDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the statement service get statement c s v default response
func (o *StatementServiceGetStatementCSVDefault) Code() int {
	return o._statusCode
}

func (o *StatementServiceGetStatementCSVDefault) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements/{statement_id}/csv][%d] StatementService_GetStatementCSV default  %+v", o._statusCode, o.Payload)
}

func (o *StatementServiceGetStatementCSVDefault) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/statements/{statement_id}/csv][%d] StatementService_GetStatementCSV default  %+v", o._statusCode, o.Payload)
}

func (o *StatementServiceGetStatementCSVDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *StatementServiceGetStatementCSVDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
