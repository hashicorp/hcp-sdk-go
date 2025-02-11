// Code generated by go-swagger; DO NOT EDIT.

package invitations_service

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

// InvitationsServiceGetOrganizationNameByInvitationTokenReader is a Reader for the InvitationsServiceGetOrganizationNameByInvitationToken structure.
type InvitationsServiceGetOrganizationNameByInvitationTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInvitationsServiceGetOrganizationNameByInvitationTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInvitationsServiceGetOrganizationNameByInvitationTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInvitationsServiceGetOrganizationNameByInvitationTokenOK creates a InvitationsServiceGetOrganizationNameByInvitationTokenOK with default headers values
func NewInvitationsServiceGetOrganizationNameByInvitationTokenOK() *InvitationsServiceGetOrganizationNameByInvitationTokenOK {
	return &InvitationsServiceGetOrganizationNameByInvitationTokenOK{}
}

/*
InvitationsServiceGetOrganizationNameByInvitationTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type InvitationsServiceGetOrganizationNameByInvitationTokenOK struct {
	Payload *models.HashicorpCloudIamGetOrganizationNameByInvitationTokenResponse
}

// IsSuccess returns true when this invitations service get organization name by invitation token o k response has a 2xx status code
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this invitations service get organization name by invitation token o k response has a 3xx status code
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invitations service get organization name by invitation token o k response has a 4xx status code
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this invitations service get organization name by invitation token o k response has a 5xx status code
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this invitations service get organization name by invitation token o k response a status code equal to that given
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the invitations service get organization name by invitation token o k response
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenOK) Code() int {
	return 200
}

func (o *InvitationsServiceGetOrganizationNameByInvitationTokenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/invitations/{invitation_token}/organization][%d] invitationsServiceGetOrganizationNameByInvitationTokenOK %s", 200, payload)
}

func (o *InvitationsServiceGetOrganizationNameByInvitationTokenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/invitations/{invitation_token}/organization][%d] invitationsServiceGetOrganizationNameByInvitationTokenOK %s", 200, payload)
}

func (o *InvitationsServiceGetOrganizationNameByInvitationTokenOK) GetPayload() *models.HashicorpCloudIamGetOrganizationNameByInvitationTokenResponse {
	return o.Payload
}

func (o *InvitationsServiceGetOrganizationNameByInvitationTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamGetOrganizationNameByInvitationTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvitationsServiceGetOrganizationNameByInvitationTokenDefault creates a InvitationsServiceGetOrganizationNameByInvitationTokenDefault with default headers values
func NewInvitationsServiceGetOrganizationNameByInvitationTokenDefault(code int) *InvitationsServiceGetOrganizationNameByInvitationTokenDefault {
	return &InvitationsServiceGetOrganizationNameByInvitationTokenDefault{
		_statusCode: code,
	}
}

/*
InvitationsServiceGetOrganizationNameByInvitationTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type InvitationsServiceGetOrganizationNameByInvitationTokenDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this invitations service get organization name by invitation token default response has a 2xx status code
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this invitations service get organization name by invitation token default response has a 3xx status code
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this invitations service get organization name by invitation token default response has a 4xx status code
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this invitations service get organization name by invitation token default response has a 5xx status code
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this invitations service get organization name by invitation token default response a status code equal to that given
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the invitations service get organization name by invitation token default response
func (o *InvitationsServiceGetOrganizationNameByInvitationTokenDefault) Code() int {
	return o._statusCode
}

func (o *InvitationsServiceGetOrganizationNameByInvitationTokenDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/invitations/{invitation_token}/organization][%d] InvitationsService_GetOrganizationNameByInvitationToken default %s", o._statusCode, payload)
}

func (o *InvitationsServiceGetOrganizationNameByInvitationTokenDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/invitations/{invitation_token}/organization][%d] InvitationsService_GetOrganizationNameByInvitationToken default %s", o._statusCode, payload)
}

func (o *InvitationsServiceGetOrganizationNameByInvitationTokenDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *InvitationsServiceGetOrganizationNameByInvitationTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
