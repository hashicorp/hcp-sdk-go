// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

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

// ExchangeWorkloadIdentityTokenReader is a Reader for the ExchangeWorkloadIdentityToken structure.
type ExchangeWorkloadIdentityTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExchangeWorkloadIdentityTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExchangeWorkloadIdentityTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExchangeWorkloadIdentityTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExchangeWorkloadIdentityTokenOK creates a ExchangeWorkloadIdentityTokenOK with default headers values
func NewExchangeWorkloadIdentityTokenOK() *ExchangeWorkloadIdentityTokenOK {
	return &ExchangeWorkloadIdentityTokenOK{}
}

/*
ExchangeWorkloadIdentityTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type ExchangeWorkloadIdentityTokenOK struct {
	Payload *models.HashicorpCloudIamExchangeWorkloadIdentityTokenResponse
}

// IsSuccess returns true when this exchange workload identity token o k response has a 2xx status code
func (o *ExchangeWorkloadIdentityTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this exchange workload identity token o k response has a 3xx status code
func (o *ExchangeWorkloadIdentityTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exchange workload identity token o k response has a 4xx status code
func (o *ExchangeWorkloadIdentityTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this exchange workload identity token o k response has a 5xx status code
func (o *ExchangeWorkloadIdentityTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this exchange workload identity token o k response a status code equal to that given
func (o *ExchangeWorkloadIdentityTokenOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExchangeWorkloadIdentityTokenOK) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/{resource_name}/exchange-token][%d] exchangeWorkloadIdentityTokenOK  %+v", 200, o.Payload)
}

func (o *ExchangeWorkloadIdentityTokenOK) String() string {
	return fmt.Sprintf("[POST /2019-12-10/{resource_name}/exchange-token][%d] exchangeWorkloadIdentityTokenOK  %+v", 200, o.Payload)
}

func (o *ExchangeWorkloadIdentityTokenOK) GetPayload() *models.HashicorpCloudIamExchangeWorkloadIdentityTokenResponse {
	return o.Payload
}

func (o *ExchangeWorkloadIdentityTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamExchangeWorkloadIdentityTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExchangeWorkloadIdentityTokenDefault creates a ExchangeWorkloadIdentityTokenDefault with default headers values
func NewExchangeWorkloadIdentityTokenDefault(code int) *ExchangeWorkloadIdentityTokenDefault {
	return &ExchangeWorkloadIdentityTokenDefault{
		_statusCode: code,
	}
}

/*
ExchangeWorkloadIdentityTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ExchangeWorkloadIdentityTokenDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the exchange workload identity token default response
func (o *ExchangeWorkloadIdentityTokenDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this exchange workload identity token default response has a 2xx status code
func (o *ExchangeWorkloadIdentityTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this exchange workload identity token default response has a 3xx status code
func (o *ExchangeWorkloadIdentityTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this exchange workload identity token default response has a 4xx status code
func (o *ExchangeWorkloadIdentityTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this exchange workload identity token default response has a 5xx status code
func (o *ExchangeWorkloadIdentityTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this exchange workload identity token default response a status code equal to that given
func (o *ExchangeWorkloadIdentityTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExchangeWorkloadIdentityTokenDefault) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/{resource_name}/exchange-token][%d] ExchangeWorkloadIdentityToken default  %+v", o._statusCode, o.Payload)
}

func (o *ExchangeWorkloadIdentityTokenDefault) String() string {
	return fmt.Sprintf("[POST /2019-12-10/{resource_name}/exchange-token][%d] ExchangeWorkloadIdentityToken default  %+v", o._statusCode, o.Payload)
}

func (o *ExchangeWorkloadIdentityTokenDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ExchangeWorkloadIdentityTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
