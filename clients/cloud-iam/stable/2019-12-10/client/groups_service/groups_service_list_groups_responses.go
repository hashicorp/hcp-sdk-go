// Code generated by go-swagger; DO NOT EDIT.

package groups_service

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

// GroupsServiceListGroupsReader is a Reader for the GroupsServiceListGroups structure.
type GroupsServiceListGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsServiceListGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsServiceListGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupsServiceListGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupsServiceListGroupsOK creates a GroupsServiceListGroupsOK with default headers values
func NewGroupsServiceListGroupsOK() *GroupsServiceListGroupsOK {
	return &GroupsServiceListGroupsOK{}
}

/*
GroupsServiceListGroupsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GroupsServiceListGroupsOK struct {
	Payload *models.HashicorpCloudIamListGroupsResponse
}

// IsSuccess returns true when this groups service list groups o k response has a 2xx status code
func (o *GroupsServiceListGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this groups service list groups o k response has a 3xx status code
func (o *GroupsServiceListGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups service list groups o k response has a 4xx status code
func (o *GroupsServiceListGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this groups service list groups o k response has a 5xx status code
func (o *GroupsServiceListGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this groups service list groups o k response a status code equal to that given
func (o *GroupsServiceListGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the groups service list groups o k response
func (o *GroupsServiceListGroupsOK) Code() int {
	return 200
}

func (o *GroupsServiceListGroupsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/iam/{parent_resource_name}/groups][%d] groupsServiceListGroupsOK %s", 200, payload)
}

func (o *GroupsServiceListGroupsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/iam/{parent_resource_name}/groups][%d] groupsServiceListGroupsOK %s", 200, payload)
}

func (o *GroupsServiceListGroupsOK) GetPayload() *models.HashicorpCloudIamListGroupsResponse {
	return o.Payload
}

func (o *GroupsServiceListGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamListGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupsServiceListGroupsDefault creates a GroupsServiceListGroupsDefault with default headers values
func NewGroupsServiceListGroupsDefault(code int) *GroupsServiceListGroupsDefault {
	return &GroupsServiceListGroupsDefault{
		_statusCode: code,
	}
}

/*
GroupsServiceListGroupsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GroupsServiceListGroupsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this groups service list groups default response has a 2xx status code
func (o *GroupsServiceListGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this groups service list groups default response has a 3xx status code
func (o *GroupsServiceListGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this groups service list groups default response has a 4xx status code
func (o *GroupsServiceListGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this groups service list groups default response has a 5xx status code
func (o *GroupsServiceListGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this groups service list groups default response a status code equal to that given
func (o *GroupsServiceListGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the groups service list groups default response
func (o *GroupsServiceListGroupsDefault) Code() int {
	return o._statusCode
}

func (o *GroupsServiceListGroupsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/iam/{parent_resource_name}/groups][%d] GroupsService_ListGroups default %s", o._statusCode, payload)
}

func (o *GroupsServiceListGroupsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/iam/{parent_resource_name}/groups][%d] GroupsService_ListGroups default %s", o._statusCode, payload)
}

func (o *GroupsServiceListGroupsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *GroupsServiceListGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
