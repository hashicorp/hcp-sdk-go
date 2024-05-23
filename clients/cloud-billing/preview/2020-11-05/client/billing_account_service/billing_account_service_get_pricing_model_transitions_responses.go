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

// BillingAccountServiceGetPricingModelTransitionsReader is a Reader for the BillingAccountServiceGetPricingModelTransitions structure.
type BillingAccountServiceGetPricingModelTransitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingAccountServiceGetPricingModelTransitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingAccountServiceGetPricingModelTransitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBillingAccountServiceGetPricingModelTransitionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBillingAccountServiceGetPricingModelTransitionsOK creates a BillingAccountServiceGetPricingModelTransitionsOK with default headers values
func NewBillingAccountServiceGetPricingModelTransitionsOK() *BillingAccountServiceGetPricingModelTransitionsOK {
	return &BillingAccountServiceGetPricingModelTransitionsOK{}
}

/*
BillingAccountServiceGetPricingModelTransitionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type BillingAccountServiceGetPricingModelTransitionsOK struct {
	Payload *models.Billing20201105GetPricingModelTransitionsResponse
}

// IsSuccess returns true when this billing account service get pricing model transitions o k response has a 2xx status code
func (o *BillingAccountServiceGetPricingModelTransitionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this billing account service get pricing model transitions o k response has a 3xx status code
func (o *BillingAccountServiceGetPricingModelTransitionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing account service get pricing model transitions o k response has a 4xx status code
func (o *BillingAccountServiceGetPricingModelTransitionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing account service get pricing model transitions o k response has a 5xx status code
func (o *BillingAccountServiceGetPricingModelTransitionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this billing account service get pricing model transitions o k response a status code equal to that given
func (o *BillingAccountServiceGetPricingModelTransitionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the billing account service get pricing model transitions o k response
func (o *BillingAccountServiceGetPricingModelTransitionsOK) Code() int {
	return 200
}

func (o *BillingAccountServiceGetPricingModelTransitionsOK) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/pricing-model-transitions][%d] billingAccountServiceGetPricingModelTransitionsOK  %+v", 200, o.Payload)
}

func (o *BillingAccountServiceGetPricingModelTransitionsOK) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/pricing-model-transitions][%d] billingAccountServiceGetPricingModelTransitionsOK  %+v", 200, o.Payload)
}

func (o *BillingAccountServiceGetPricingModelTransitionsOK) GetPayload() *models.Billing20201105GetPricingModelTransitionsResponse {
	return o.Payload
}

func (o *BillingAccountServiceGetPricingModelTransitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105GetPricingModelTransitionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingAccountServiceGetPricingModelTransitionsDefault creates a BillingAccountServiceGetPricingModelTransitionsDefault with default headers values
func NewBillingAccountServiceGetPricingModelTransitionsDefault(code int) *BillingAccountServiceGetPricingModelTransitionsDefault {
	return &BillingAccountServiceGetPricingModelTransitionsDefault{
		_statusCode: code,
	}
}

/*
BillingAccountServiceGetPricingModelTransitionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BillingAccountServiceGetPricingModelTransitionsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this billing account service get pricing model transitions default response has a 2xx status code
func (o *BillingAccountServiceGetPricingModelTransitionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this billing account service get pricing model transitions default response has a 3xx status code
func (o *BillingAccountServiceGetPricingModelTransitionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this billing account service get pricing model transitions default response has a 4xx status code
func (o *BillingAccountServiceGetPricingModelTransitionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this billing account service get pricing model transitions default response has a 5xx status code
func (o *BillingAccountServiceGetPricingModelTransitionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this billing account service get pricing model transitions default response a status code equal to that given
func (o *BillingAccountServiceGetPricingModelTransitionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the billing account service get pricing model transitions default response
func (o *BillingAccountServiceGetPricingModelTransitionsDefault) Code() int {
	return o._statusCode
}

func (o *BillingAccountServiceGetPricingModelTransitionsDefault) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/pricing-model-transitions][%d] BillingAccountService_GetPricingModelTransitions default  %+v", o._statusCode, o.Payload)
}

func (o *BillingAccountServiceGetPricingModelTransitionsDefault) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/pricing-model-transitions][%d] BillingAccountService_GetPricingModelTransitions default  %+v", o._statusCode, o.Payload)
}

func (o *BillingAccountServiceGetPricingModelTransitionsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BillingAccountServiceGetPricingModelTransitionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
