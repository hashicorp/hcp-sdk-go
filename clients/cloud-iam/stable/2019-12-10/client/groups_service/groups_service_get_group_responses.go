// Code generated by go-swagger; DO NOT EDIT.

package groups_service

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

// GroupsServiceGetGroupReader is a Reader for the GroupsServiceGetGroup structure.
type GroupsServiceGetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsServiceGetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsServiceGetGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupsServiceGetGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupsServiceGetGroupOK creates a GroupsServiceGetGroupOK with default headers values
func NewGroupsServiceGetGroupOK() *GroupsServiceGetGroupOK {
	return &GroupsServiceGetGroupOK{}
}

/*
GroupsServiceGetGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type GroupsServiceGetGroupOK struct {
	Payload *models.HashicorpCloudIamGetGroupResponse
}

// IsSuccess returns true when this groups service get group o k response has a 2xx status code
func (o *GroupsServiceGetGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this groups service get group o k response has a 3xx status code
func (o *GroupsServiceGetGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups service get group o k response has a 4xx status code
func (o *GroupsServiceGetGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this groups service get group o k response has a 5xx status code
func (o *GroupsServiceGetGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this groups service get group o k response a status code equal to that given
func (o *GroupsServiceGetGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the groups service get group o k response
func (o *GroupsServiceGetGroupOK) Code() int {
	return 200
}

func (o *GroupsServiceGetGroupOK) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/{resource_name}][%d] groupsServiceGetGroupOK  %+v", 200, o.Payload)
}

func (o *GroupsServiceGetGroupOK) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/{resource_name}][%d] groupsServiceGetGroupOK  %+v", 200, o.Payload)
}

func (o *GroupsServiceGetGroupOK) GetPayload() *models.HashicorpCloudIamGetGroupResponse {
	return o.Payload
}

func (o *GroupsServiceGetGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamGetGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupsServiceGetGroupDefault creates a GroupsServiceGetGroupDefault with default headers values
func NewGroupsServiceGetGroupDefault(code int) *GroupsServiceGetGroupDefault {
	return &GroupsServiceGetGroupDefault{
		_statusCode: code,
	}
}

/*
GroupsServiceGetGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GroupsServiceGetGroupDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this groups service get group default response has a 2xx status code
func (o *GroupsServiceGetGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this groups service get group default response has a 3xx status code
func (o *GroupsServiceGetGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this groups service get group default response has a 4xx status code
func (o *GroupsServiceGetGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this groups service get group default response has a 5xx status code
func (o *GroupsServiceGetGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this groups service get group default response a status code equal to that given
func (o *GroupsServiceGetGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the groups service get group default response
func (o *GroupsServiceGetGroupDefault) Code() int {
	return o._statusCode
}

func (o *GroupsServiceGetGroupDefault) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/{resource_name}][%d] GroupsService_GetGroup default  %+v", o._statusCode, o.Payload)
}

func (o *GroupsServiceGetGroupDefault) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/{resource_name}][%d] GroupsService_GetGroup default  %+v", o._statusCode, o.Payload)
}

func (o *GroupsServiceGetGroupDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *GroupsServiceGetGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
