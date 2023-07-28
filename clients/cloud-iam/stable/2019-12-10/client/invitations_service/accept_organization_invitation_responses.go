// Code generated by go-swagger; DO NOT EDIT.

package invitations_service

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

// AcceptOrganizationInvitationReader is a Reader for the AcceptOrganizationInvitation structure.
type AcceptOrganizationInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptOrganizationInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcceptOrganizationInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAcceptOrganizationInvitationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAcceptOrganizationInvitationOK creates a AcceptOrganizationInvitationOK with default headers values
func NewAcceptOrganizationInvitationOK() *AcceptOrganizationInvitationOK {
	return &AcceptOrganizationInvitationOK{}
}

/*
AcceptOrganizationInvitationOK describes a response with status code 200, with default header values.

A successful response.
*/
type AcceptOrganizationInvitationOK struct {
	Payload *models.HashicorpCloudIamAcceptOrganizationInvitationResponse
}

// IsSuccess returns true when this accept organization invitation o k response has a 2xx status code
func (o *AcceptOrganizationInvitationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this accept organization invitation o k response has a 3xx status code
func (o *AcceptOrganizationInvitationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this accept organization invitation o k response has a 4xx status code
func (o *AcceptOrganizationInvitationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this accept organization invitation o k response has a 5xx status code
func (o *AcceptOrganizationInvitationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this accept organization invitation o k response a status code equal to that given
func (o *AcceptOrganizationInvitationOK) IsCode(code int) bool {
	return code == 200
}

func (o *AcceptOrganizationInvitationOK) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/accept-invitation][%d] acceptOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *AcceptOrganizationInvitationOK) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/accept-invitation][%d] acceptOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *AcceptOrganizationInvitationOK) GetPayload() *models.HashicorpCloudIamAcceptOrganizationInvitationResponse {
	return o.Payload
}

func (o *AcceptOrganizationInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamAcceptOrganizationInvitationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptOrganizationInvitationDefault creates a AcceptOrganizationInvitationDefault with default headers values
func NewAcceptOrganizationInvitationDefault(code int) *AcceptOrganizationInvitationDefault {
	return &AcceptOrganizationInvitationDefault{
		_statusCode: code,
	}
}

/*
AcceptOrganizationInvitationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AcceptOrganizationInvitationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the accept organization invitation default response
func (o *AcceptOrganizationInvitationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this accept organization invitation default response has a 2xx status code
func (o *AcceptOrganizationInvitationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this accept organization invitation default response has a 3xx status code
func (o *AcceptOrganizationInvitationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this accept organization invitation default response has a 4xx status code
func (o *AcceptOrganizationInvitationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this accept organization invitation default response has a 5xx status code
func (o *AcceptOrganizationInvitationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this accept organization invitation default response a status code equal to that given
func (o *AcceptOrganizationInvitationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AcceptOrganizationInvitationDefault) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/accept-invitation][%d] AcceptOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *AcceptOrganizationInvitationDefault) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/accept-invitation][%d] AcceptOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *AcceptOrganizationInvitationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *AcceptOrganizationInvitationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
