// Code generated by go-swagger; DO NOT EDIT.

package iam_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// IamServiceCreateUserPrincipalReader is a Reader for the IamServiceCreateUserPrincipal structure.
type IamServiceCreateUserPrincipalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamServiceCreateUserPrincipalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamServiceCreateUserPrincipalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamServiceCreateUserPrincipalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamServiceCreateUserPrincipalOK creates a IamServiceCreateUserPrincipalOK with default headers values
func NewIamServiceCreateUserPrincipalOK() *IamServiceCreateUserPrincipalOK {
	return &IamServiceCreateUserPrincipalOK{}
}

/*
IamServiceCreateUserPrincipalOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamServiceCreateUserPrincipalOK struct {
	Payload *models.HashicorpCloudIamUserPrincipalResponse
}

// IsSuccess returns true when this iam service create user principal o k response has a 2xx status code
func (o *IamServiceCreateUserPrincipalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam service create user principal o k response has a 3xx status code
func (o *IamServiceCreateUserPrincipalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam service create user principal o k response has a 4xx status code
func (o *IamServiceCreateUserPrincipalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam service create user principal o k response has a 5xx status code
func (o *IamServiceCreateUserPrincipalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam service create user principal o k response a status code equal to that given
func (o *IamServiceCreateUserPrincipalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the iam service create user principal o k response
func (o *IamServiceCreateUserPrincipalOK) Code() int {
	return 200
}

func (o *IamServiceCreateUserPrincipalOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/user-principals][%d] iamServiceCreateUserPrincipalOK %s", 200, payload)
}

func (o *IamServiceCreateUserPrincipalOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/user-principals][%d] iamServiceCreateUserPrincipalOK %s", 200, payload)
}

func (o *IamServiceCreateUserPrincipalOK) GetPayload() *models.HashicorpCloudIamUserPrincipalResponse {
	return o.Payload
}

func (o *IamServiceCreateUserPrincipalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamUserPrincipalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamServiceCreateUserPrincipalDefault creates a IamServiceCreateUserPrincipalDefault with default headers values
func NewIamServiceCreateUserPrincipalDefault(code int) *IamServiceCreateUserPrincipalDefault {
	return &IamServiceCreateUserPrincipalDefault{
		_statusCode: code,
	}
}

/*
IamServiceCreateUserPrincipalDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IamServiceCreateUserPrincipalDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this iam service create user principal default response has a 2xx status code
func (o *IamServiceCreateUserPrincipalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam service create user principal default response has a 3xx status code
func (o *IamServiceCreateUserPrincipalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam service create user principal default response has a 4xx status code
func (o *IamServiceCreateUserPrincipalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam service create user principal default response has a 5xx status code
func (o *IamServiceCreateUserPrincipalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam service create user principal default response a status code equal to that given
func (o *IamServiceCreateUserPrincipalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the iam service create user principal default response
func (o *IamServiceCreateUserPrincipalDefault) Code() int {
	return o._statusCode
}

func (o *IamServiceCreateUserPrincipalDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/user-principals][%d] IamService_CreateUserPrincipal default %s", o._statusCode, payload)
}

func (o *IamServiceCreateUserPrincipalDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/user-principals][%d] IamService_CreateUserPrincipal default %s", o._statusCode, payload)
}

func (o *IamServiceCreateUserPrincipalDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *IamServiceCreateUserPrincipalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
