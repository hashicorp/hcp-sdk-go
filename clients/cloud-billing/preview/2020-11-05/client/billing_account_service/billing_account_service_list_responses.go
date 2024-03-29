// Code generated by go-swagger; DO NOT EDIT.

package billing_account_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/models"
)

// BillingAccountServiceListReader is a Reader for the BillingAccountServiceList structure.
type BillingAccountServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingAccountServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingAccountServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBillingAccountServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBillingAccountServiceListOK creates a BillingAccountServiceListOK with default headers values
func NewBillingAccountServiceListOK() *BillingAccountServiceListOK {
	return &BillingAccountServiceListOK{}
}

/*
BillingAccountServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type BillingAccountServiceListOK struct {
	Payload *models.Billing20201105ListBillingAccountsResponse
}

// IsSuccess returns true when this billing account service list o k response has a 2xx status code
func (o *BillingAccountServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this billing account service list o k response has a 3xx status code
func (o *BillingAccountServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing account service list o k response has a 4xx status code
func (o *BillingAccountServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing account service list o k response has a 5xx status code
func (o *BillingAccountServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this billing account service list o k response a status code equal to that given
func (o *BillingAccountServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *BillingAccountServiceListOK) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts][%d] billingAccountServiceListOK  %+v", 200, o.Payload)
}

func (o *BillingAccountServiceListOK) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts][%d] billingAccountServiceListOK  %+v", 200, o.Payload)
}

func (o *BillingAccountServiceListOK) GetPayload() *models.Billing20201105ListBillingAccountsResponse {
	return o.Payload
}

func (o *BillingAccountServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105ListBillingAccountsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingAccountServiceListDefault creates a BillingAccountServiceListDefault with default headers values
func NewBillingAccountServiceListDefault(code int) *BillingAccountServiceListDefault {
	return &BillingAccountServiceListDefault{
		_statusCode: code,
	}
}

/*
BillingAccountServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BillingAccountServiceListDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the billing account service list default response
func (o *BillingAccountServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this billing account service list default response has a 2xx status code
func (o *BillingAccountServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this billing account service list default response has a 3xx status code
func (o *BillingAccountServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this billing account service list default response has a 4xx status code
func (o *BillingAccountServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this billing account service list default response has a 5xx status code
func (o *BillingAccountServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this billing account service list default response a status code equal to that given
func (o *BillingAccountServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *BillingAccountServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts][%d] BillingAccountService_List default  %+v", o._statusCode, o.Payload)
}

func (o *BillingAccountServiceListDefault) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts][%d] BillingAccountService_List default  %+v", o._statusCode, o.Payload)
}

func (o *BillingAccountServiceListDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *BillingAccountServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
