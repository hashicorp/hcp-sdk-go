// Code generated by go-swagger; DO NOT EDIT.

package activation_service

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

// ActivationServiceGetActivationDetailsReader is a Reader for the ActivationServiceGetActivationDetails structure.
type ActivationServiceGetActivationDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActivationServiceGetActivationDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActivationServiceGetActivationDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActivationServiceGetActivationDetailsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActivationServiceGetActivationDetailsOK creates a ActivationServiceGetActivationDetailsOK with default headers values
func NewActivationServiceGetActivationDetailsOK() *ActivationServiceGetActivationDetailsOK {
	return &ActivationServiceGetActivationDetailsOK{}
}

/*
ActivationServiceGetActivationDetailsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ActivationServiceGetActivationDetailsOK struct {
	Payload *models.Billing20201105GetActivationDetailsResponse
}

// IsSuccess returns true when this activation service get activation details o k response has a 2xx status code
func (o *ActivationServiceGetActivationDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this activation service get activation details o k response has a 3xx status code
func (o *ActivationServiceGetActivationDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this activation service get activation details o k response has a 4xx status code
func (o *ActivationServiceGetActivationDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this activation service get activation details o k response has a 5xx status code
func (o *ActivationServiceGetActivationDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this activation service get activation details o k response a status code equal to that given
func (o *ActivationServiceGetActivationDetailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the activation service get activation details o k response
func (o *ActivationServiceGetActivationDetailsOK) Code() int {
	return 200
}

func (o *ActivationServiceGetActivationDetailsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing/2020-11-05/activations/{activation_code}][%d] activationServiceGetActivationDetailsOK %s", 200, payload)
}

func (o *ActivationServiceGetActivationDetailsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing/2020-11-05/activations/{activation_code}][%d] activationServiceGetActivationDetailsOK %s", 200, payload)
}

func (o *ActivationServiceGetActivationDetailsOK) GetPayload() *models.Billing20201105GetActivationDetailsResponse {
	return o.Payload
}

func (o *ActivationServiceGetActivationDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105GetActivationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivationServiceGetActivationDetailsDefault creates a ActivationServiceGetActivationDetailsDefault with default headers values
func NewActivationServiceGetActivationDetailsDefault(code int) *ActivationServiceGetActivationDetailsDefault {
	return &ActivationServiceGetActivationDetailsDefault{
		_statusCode: code,
	}
}

/*
ActivationServiceGetActivationDetailsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ActivationServiceGetActivationDetailsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this activation service get activation details default response has a 2xx status code
func (o *ActivationServiceGetActivationDetailsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this activation service get activation details default response has a 3xx status code
func (o *ActivationServiceGetActivationDetailsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this activation service get activation details default response has a 4xx status code
func (o *ActivationServiceGetActivationDetailsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this activation service get activation details default response has a 5xx status code
func (o *ActivationServiceGetActivationDetailsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this activation service get activation details default response a status code equal to that given
func (o *ActivationServiceGetActivationDetailsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the activation service get activation details default response
func (o *ActivationServiceGetActivationDetailsDefault) Code() int {
	return o._statusCode
}

func (o *ActivationServiceGetActivationDetailsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing/2020-11-05/activations/{activation_code}][%d] ActivationService_GetActivationDetails default %s", o._statusCode, payload)
}

func (o *ActivationServiceGetActivationDetailsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing/2020-11-05/activations/{activation_code}][%d] ActivationService_GetActivationDetails default %s", o._statusCode, payload)
}

func (o *ActivationServiceGetActivationDetailsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ActivationServiceGetActivationDetailsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
