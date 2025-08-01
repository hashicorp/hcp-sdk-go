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

// BillingAccountServiceCreateSetupIntentReader is a Reader for the BillingAccountServiceCreateSetupIntent structure.
type BillingAccountServiceCreateSetupIntentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingAccountServiceCreateSetupIntentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingAccountServiceCreateSetupIntentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBillingAccountServiceCreateSetupIntentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBillingAccountServiceCreateSetupIntentOK creates a BillingAccountServiceCreateSetupIntentOK with default headers values
func NewBillingAccountServiceCreateSetupIntentOK() *BillingAccountServiceCreateSetupIntentOK {
	return &BillingAccountServiceCreateSetupIntentOK{}
}

/*
BillingAccountServiceCreateSetupIntentOK describes a response with status code 200, with default header values.

A successful response.
*/
type BillingAccountServiceCreateSetupIntentOK struct {
	Payload *models.Billing20201105CreateSetupIntentResponse
}

// IsSuccess returns true when this billing account service create setup intent o k response has a 2xx status code
func (o *BillingAccountServiceCreateSetupIntentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this billing account service create setup intent o k response has a 3xx status code
func (o *BillingAccountServiceCreateSetupIntentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing account service create setup intent o k response has a 4xx status code
func (o *BillingAccountServiceCreateSetupIntentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing account service create setup intent o k response has a 5xx status code
func (o *BillingAccountServiceCreateSetupIntentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this billing account service create setup intent o k response a status code equal to that given
func (o *BillingAccountServiceCreateSetupIntentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the billing account service create setup intent o k response
func (o *BillingAccountServiceCreateSetupIntentOK) Code() int {
	return 200
}

func (o *BillingAccountServiceCreateSetupIntentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /billing/2020-11-05/organizations/{organization_id}/setup-intents][%d] billingAccountServiceCreateSetupIntentOK %s", 200, payload)
}

func (o *BillingAccountServiceCreateSetupIntentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /billing/2020-11-05/organizations/{organization_id}/setup-intents][%d] billingAccountServiceCreateSetupIntentOK %s", 200, payload)
}

func (o *BillingAccountServiceCreateSetupIntentOK) GetPayload() *models.Billing20201105CreateSetupIntentResponse {
	return o.Payload
}

func (o *BillingAccountServiceCreateSetupIntentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105CreateSetupIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingAccountServiceCreateSetupIntentDefault creates a BillingAccountServiceCreateSetupIntentDefault with default headers values
func NewBillingAccountServiceCreateSetupIntentDefault(code int) *BillingAccountServiceCreateSetupIntentDefault {
	return &BillingAccountServiceCreateSetupIntentDefault{
		_statusCode: code,
	}
}

/*
BillingAccountServiceCreateSetupIntentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BillingAccountServiceCreateSetupIntentDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this billing account service create setup intent default response has a 2xx status code
func (o *BillingAccountServiceCreateSetupIntentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this billing account service create setup intent default response has a 3xx status code
func (o *BillingAccountServiceCreateSetupIntentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this billing account service create setup intent default response has a 4xx status code
func (o *BillingAccountServiceCreateSetupIntentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this billing account service create setup intent default response has a 5xx status code
func (o *BillingAccountServiceCreateSetupIntentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this billing account service create setup intent default response a status code equal to that given
func (o *BillingAccountServiceCreateSetupIntentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the billing account service create setup intent default response
func (o *BillingAccountServiceCreateSetupIntentDefault) Code() int {
	return o._statusCode
}

func (o *BillingAccountServiceCreateSetupIntentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /billing/2020-11-05/organizations/{organization_id}/setup-intents][%d] BillingAccountService_CreateSetupIntent default %s", o._statusCode, payload)
}

func (o *BillingAccountServiceCreateSetupIntentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /billing/2020-11-05/organizations/{organization_id}/setup-intents][%d] BillingAccountService_CreateSetupIntent default %s", o._statusCode, payload)
}

func (o *BillingAccountServiceCreateSetupIntentDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BillingAccountServiceCreateSetupIntentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
