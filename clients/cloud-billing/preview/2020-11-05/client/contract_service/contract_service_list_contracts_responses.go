// Code generated by go-swagger; DO NOT EDIT.

package contract_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/models"
)

// ContractServiceListContractsReader is a Reader for the ContractServiceListContracts structure.
type ContractServiceListContractsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContractServiceListContractsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContractServiceListContractsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContractServiceListContractsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContractServiceListContractsOK creates a ContractServiceListContractsOK with default headers values
func NewContractServiceListContractsOK() *ContractServiceListContractsOK {
	return &ContractServiceListContractsOK{}
}

/*
ContractServiceListContractsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ContractServiceListContractsOK struct {
	Payload *models.Billing20201105ListContractsResponse
}

// IsSuccess returns true when this contract service list contracts o k response has a 2xx status code
func (o *ContractServiceListContractsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contract service list contracts o k response has a 3xx status code
func (o *ContractServiceListContractsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contract service list contracts o k response has a 4xx status code
func (o *ContractServiceListContractsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contract service list contracts o k response has a 5xx status code
func (o *ContractServiceListContractsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contract service list contracts o k response a status code equal to that given
func (o *ContractServiceListContractsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contract service list contracts o k response
func (o *ContractServiceListContractsOK) Code() int {
	return 200
}

func (o *ContractServiceListContractsOK) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/contracts][%d] contractServiceListContractsOK  %+v", 200, o.Payload)
}

func (o *ContractServiceListContractsOK) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/contracts][%d] contractServiceListContractsOK  %+v", 200, o.Payload)
}

func (o *ContractServiceListContractsOK) GetPayload() *models.Billing20201105ListContractsResponse {
	return o.Payload
}

func (o *ContractServiceListContractsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Billing20201105ListContractsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContractServiceListContractsDefault creates a ContractServiceListContractsDefault with default headers values
func NewContractServiceListContractsDefault(code int) *ContractServiceListContractsDefault {
	return &ContractServiceListContractsDefault{
		_statusCode: code,
	}
}

/*
ContractServiceListContractsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ContractServiceListContractsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this contract service list contracts default response has a 2xx status code
func (o *ContractServiceListContractsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this contract service list contracts default response has a 3xx status code
func (o *ContractServiceListContractsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this contract service list contracts default response has a 4xx status code
func (o *ContractServiceListContractsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this contract service list contracts default response has a 5xx status code
func (o *ContractServiceListContractsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this contract service list contracts default response a status code equal to that given
func (o *ContractServiceListContractsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the contract service list contracts default response
func (o *ContractServiceListContractsDefault) Code() int {
	return o._statusCode
}

func (o *ContractServiceListContractsDefault) Error() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/contracts][%d] ContractService_ListContracts default  %+v", o._statusCode, o.Payload)
}

func (o *ContractServiceListContractsDefault) String() string {
	return fmt.Sprintf("[GET /billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/contracts][%d] ContractService_ListContracts default  %+v", o._statusCode, o.Payload)
}

func (o *ContractServiceListContractsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ContractServiceListContractsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
