// Code generated by go-swagger; DO NOT EDIT.

package billing_account_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/stable/2020-11-05/models"
)

// BillingAccountServiceGetReader is a Reader for the BillingAccountServiceGet structure.
type BillingAccountServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingAccountServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingAccountServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBillingAccountServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBillingAccountServiceGetOK creates a BillingAccountServiceGetOK with default headers values
func NewBillingAccountServiceGetOK() *BillingAccountServiceGetOK {
	return &BillingAccountServiceGetOK{}
}

/*
BillingAccountServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type BillingAccountServiceGetOK struct {
	Payload *models.Billing20201105GetBillingAccountResponse
}

// IsSuccess returns true when this billing account service get o k response has a 2xx status code
func (o *BillingAccountServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this billing account service get o k response has a 3xx status code
func (o *BillingAccountServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing account service get o k response has a 4xx status code
func (o *BillingAccountServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing account service get o k response has a 5xx status code
func (o *BillingAccountServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this billing account service get o k response a status code equal to that given
func (o *BillingAccountServiceGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the billing account service get o k response
func (o *BillingAccountServiceGetOK) Code() int {
	return 200
}

func (o *BillingAccountServiceGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{id}][%d] billingAccountServiceGetOK %s", 200, payload)
}

func (o *BillingAccountServiceGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{id}][%d] billingAccountServiceGetOK %s", 200, payload)
}

func (o *BillingAccountServiceGetOK) GetPayload() *models.Billing20201105GetBillingAccountResponse {
	return o.Payload
}

func (o *BillingAccountServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105GetBillingAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingAccountServiceGetDefault creates a BillingAccountServiceGetDefault with default headers values
func NewBillingAccountServiceGetDefault(code int) *BillingAccountServiceGetDefault {
	return &BillingAccountServiceGetDefault{
		_statusCode: code,
	}
}

/*
BillingAccountServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BillingAccountServiceGetDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this billing account service get default response has a 2xx status code
func (o *BillingAccountServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this billing account service get default response has a 3xx status code
func (o *BillingAccountServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this billing account service get default response has a 4xx status code
func (o *BillingAccountServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this billing account service get default response has a 5xx status code
func (o *BillingAccountServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this billing account service get default response a status code equal to that given
func (o *BillingAccountServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the billing account service get default response
func (o *BillingAccountServiceGetDefault) Code() int {
	return o._statusCode
}

func (o *BillingAccountServiceGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{id}][%d] BillingAccountService_Get default %s", o._statusCode, payload)
}

func (o *BillingAccountServiceGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{id}][%d] BillingAccountService_Get default %s", o._statusCode, payload)
}

func (o *BillingAccountServiceGetDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BillingAccountServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
