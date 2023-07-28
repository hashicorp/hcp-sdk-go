// Code generated by go-swagger; DO NOT EDIT.

package iam_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// SearchPrincipalsReader is a Reader for the SearchPrincipals structure.
type SearchPrincipalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchPrincipalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchPrincipalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchPrincipalsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchPrincipalsOK creates a SearchPrincipalsOK with default headers values
func NewSearchPrincipalsOK() *SearchPrincipalsOK {
	return &SearchPrincipalsOK{}
}

/*
SearchPrincipalsOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchPrincipalsOK struct {
	Payload *models.HashicorpCloudIamSearchPrincipalsResponse
}

// IsSuccess returns true when this search principals o k response has a 2xx status code
func (o *SearchPrincipalsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search principals o k response has a 3xx status code
func (o *SearchPrincipalsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search principals o k response has a 4xx status code
func (o *SearchPrincipalsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search principals o k response has a 5xx status code
func (o *SearchPrincipalsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search principals o k response a status code equal to that given
func (o *SearchPrincipalsOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchPrincipalsOK) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/principals/search][%d] searchPrincipalsOK  %+v", 200, o.Payload)
}

func (o *SearchPrincipalsOK) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/principals/search][%d] searchPrincipalsOK  %+v", 200, o.Payload)
}

func (o *SearchPrincipalsOK) GetPayload() *models.HashicorpCloudIamSearchPrincipalsResponse {
	return o.Payload
}

func (o *SearchPrincipalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamSearchPrincipalsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPrincipalsDefault creates a SearchPrincipalsDefault with default headers values
func NewSearchPrincipalsDefault(code int) *SearchPrincipalsDefault {
	return &SearchPrincipalsDefault{
		_statusCode: code,
	}
}

/*
SearchPrincipalsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SearchPrincipalsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the search principals default response
func (o *SearchPrincipalsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this search principals default response has a 2xx status code
func (o *SearchPrincipalsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search principals default response has a 3xx status code
func (o *SearchPrincipalsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search principals default response has a 4xx status code
func (o *SearchPrincipalsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search principals default response has a 5xx status code
func (o *SearchPrincipalsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search principals default response a status code equal to that given
func (o *SearchPrincipalsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SearchPrincipalsDefault) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/principals/search][%d] SearchPrincipals default  %+v", o._statusCode, o.Payload)
}

func (o *SearchPrincipalsDefault) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/principals/search][%d] SearchPrincipals default  %+v", o._statusCode, o.Payload)
}

func (o *SearchPrincipalsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *SearchPrincipalsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
