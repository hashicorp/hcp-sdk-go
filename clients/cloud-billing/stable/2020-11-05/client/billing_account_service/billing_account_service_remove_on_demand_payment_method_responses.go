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

// BillingAccountServiceRemoveOnDemandPaymentMethodReader is a Reader for the BillingAccountServiceRemoveOnDemandPaymentMethod structure.
type BillingAccountServiceRemoveOnDemandPaymentMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingAccountServiceRemoveOnDemandPaymentMethodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBillingAccountServiceRemoveOnDemandPaymentMethodDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBillingAccountServiceRemoveOnDemandPaymentMethodOK creates a BillingAccountServiceRemoveOnDemandPaymentMethodOK with default headers values
func NewBillingAccountServiceRemoveOnDemandPaymentMethodOK() *BillingAccountServiceRemoveOnDemandPaymentMethodOK {
	return &BillingAccountServiceRemoveOnDemandPaymentMethodOK{}
}

/*
BillingAccountServiceRemoveOnDemandPaymentMethodOK describes a response with status code 200, with default header values.

A successful response.
*/
type BillingAccountServiceRemoveOnDemandPaymentMethodOK struct {
	Payload *models.Billing20201105RemoveOnDemandPaymentMethodResponse
}

// IsSuccess returns true when this billing account service remove on demand payment method o k response has a 2xx status code
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this billing account service remove on demand payment method o k response has a 3xx status code
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing account service remove on demand payment method o k response has a 4xx status code
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing account service remove on demand payment method o k response has a 5xx status code
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this billing account service remove on demand payment method o k response a status code equal to that given
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the billing account service remove on demand payment method o k response
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodOK) Code() int {
	return 200
}

func (o *BillingAccountServiceRemoveOnDemandPaymentMethodOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /billing/2020-11-05/organizations/{organization_id}/billing_accounts/{billing_account_id}/on_demand_payment_method][%d] billingAccountServiceRemoveOnDemandPaymentMethodOK %s", 200, payload)
}

func (o *BillingAccountServiceRemoveOnDemandPaymentMethodOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /billing/2020-11-05/organizations/{organization_id}/billing_accounts/{billing_account_id}/on_demand_payment_method][%d] billingAccountServiceRemoveOnDemandPaymentMethodOK %s", 200, payload)
}

func (o *BillingAccountServiceRemoveOnDemandPaymentMethodOK) GetPayload() *models.Billing20201105RemoveOnDemandPaymentMethodResponse {
	return o.Payload
}

func (o *BillingAccountServiceRemoveOnDemandPaymentMethodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105RemoveOnDemandPaymentMethodResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingAccountServiceRemoveOnDemandPaymentMethodDefault creates a BillingAccountServiceRemoveOnDemandPaymentMethodDefault with default headers values
func NewBillingAccountServiceRemoveOnDemandPaymentMethodDefault(code int) *BillingAccountServiceRemoveOnDemandPaymentMethodDefault {
	return &BillingAccountServiceRemoveOnDemandPaymentMethodDefault{
		_statusCode: code,
	}
}

/*
BillingAccountServiceRemoveOnDemandPaymentMethodDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BillingAccountServiceRemoveOnDemandPaymentMethodDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this billing account service remove on demand payment method default response has a 2xx status code
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this billing account service remove on demand payment method default response has a 3xx status code
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this billing account service remove on demand payment method default response has a 4xx status code
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this billing account service remove on demand payment method default response has a 5xx status code
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this billing account service remove on demand payment method default response a status code equal to that given
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the billing account service remove on demand payment method default response
func (o *BillingAccountServiceRemoveOnDemandPaymentMethodDefault) Code() int {
	return o._statusCode
}

func (o *BillingAccountServiceRemoveOnDemandPaymentMethodDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /billing/2020-11-05/organizations/{organization_id}/billing_accounts/{billing_account_id}/on_demand_payment_method][%d] BillingAccountService_RemoveOnDemandPaymentMethod default %s", o._statusCode, payload)
}

func (o *BillingAccountServiceRemoveOnDemandPaymentMethodDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /billing/2020-11-05/organizations/{organization_id}/billing_accounts/{billing_account_id}/on_demand_payment_method][%d] BillingAccountService_RemoveOnDemandPaymentMethod default %s", o._statusCode, payload)
}

func (o *BillingAccountServiceRemoveOnDemandPaymentMethodDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BillingAccountServiceRemoveOnDemandPaymentMethodDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
