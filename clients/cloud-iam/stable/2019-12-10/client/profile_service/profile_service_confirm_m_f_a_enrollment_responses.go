// Code generated by go-swagger; DO NOT EDIT.

package profile_service

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

// ProfileServiceConfirmMFAEnrollmentReader is a Reader for the ProfileServiceConfirmMFAEnrollment structure.
type ProfileServiceConfirmMFAEnrollmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfileServiceConfirmMFAEnrollmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProfileServiceConfirmMFAEnrollmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProfileServiceConfirmMFAEnrollmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProfileServiceConfirmMFAEnrollmentOK creates a ProfileServiceConfirmMFAEnrollmentOK with default headers values
func NewProfileServiceConfirmMFAEnrollmentOK() *ProfileServiceConfirmMFAEnrollmentOK {
	return &ProfileServiceConfirmMFAEnrollmentOK{}
}

/*
ProfileServiceConfirmMFAEnrollmentOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProfileServiceConfirmMFAEnrollmentOK struct {
	Payload models.HashicorpCloudIamConfirmMFAEnrollmentResponse
}

// IsSuccess returns true when this profile service confirm m f a enrollment o k response has a 2xx status code
func (o *ProfileServiceConfirmMFAEnrollmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this profile service confirm m f a enrollment o k response has a 3xx status code
func (o *ProfileServiceConfirmMFAEnrollmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this profile service confirm m f a enrollment o k response has a 4xx status code
func (o *ProfileServiceConfirmMFAEnrollmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this profile service confirm m f a enrollment o k response has a 5xx status code
func (o *ProfileServiceConfirmMFAEnrollmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this profile service confirm m f a enrollment o k response a status code equal to that given
func (o *ProfileServiceConfirmMFAEnrollmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the profile service confirm m f a enrollment o k response
func (o *ProfileServiceConfirmMFAEnrollmentOK) Code() int {
	return 200
}

func (o *ProfileServiceConfirmMFAEnrollmentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/me/confirm-mfa-enrollment][%d] profileServiceConfirmMFAEnrollmentOK %s", 200, payload)
}

func (o *ProfileServiceConfirmMFAEnrollmentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/me/confirm-mfa-enrollment][%d] profileServiceConfirmMFAEnrollmentOK %s", 200, payload)
}

func (o *ProfileServiceConfirmMFAEnrollmentOK) GetPayload() models.HashicorpCloudIamConfirmMFAEnrollmentResponse {
	return o.Payload
}

func (o *ProfileServiceConfirmMFAEnrollmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfileServiceConfirmMFAEnrollmentDefault creates a ProfileServiceConfirmMFAEnrollmentDefault with default headers values
func NewProfileServiceConfirmMFAEnrollmentDefault(code int) *ProfileServiceConfirmMFAEnrollmentDefault {
	return &ProfileServiceConfirmMFAEnrollmentDefault{
		_statusCode: code,
	}
}

/*
ProfileServiceConfirmMFAEnrollmentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProfileServiceConfirmMFAEnrollmentDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this profile service confirm m f a enrollment default response has a 2xx status code
func (o *ProfileServiceConfirmMFAEnrollmentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this profile service confirm m f a enrollment default response has a 3xx status code
func (o *ProfileServiceConfirmMFAEnrollmentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this profile service confirm m f a enrollment default response has a 4xx status code
func (o *ProfileServiceConfirmMFAEnrollmentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this profile service confirm m f a enrollment default response has a 5xx status code
func (o *ProfileServiceConfirmMFAEnrollmentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this profile service confirm m f a enrollment default response a status code equal to that given
func (o *ProfileServiceConfirmMFAEnrollmentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the profile service confirm m f a enrollment default response
func (o *ProfileServiceConfirmMFAEnrollmentDefault) Code() int {
	return o._statusCode
}

func (o *ProfileServiceConfirmMFAEnrollmentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/me/confirm-mfa-enrollment][%d] ProfileService_ConfirmMFAEnrollment default %s", o._statusCode, payload)
}

func (o *ProfileServiceConfirmMFAEnrollmentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/me/confirm-mfa-enrollment][%d] ProfileService_ConfirmMFAEnrollment default %s", o._statusCode, payload)
}

func (o *ProfileServiceConfirmMFAEnrollmentDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ProfileServiceConfirmMFAEnrollmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
