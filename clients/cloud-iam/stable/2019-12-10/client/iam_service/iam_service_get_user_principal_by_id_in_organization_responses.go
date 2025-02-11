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

// IamServiceGetUserPrincipalByIDInOrganizationReader is a Reader for the IamServiceGetUserPrincipalByIDInOrganization structure.
type IamServiceGetUserPrincipalByIDInOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamServiceGetUserPrincipalByIDInOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamServiceGetUserPrincipalByIDInOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamServiceGetUserPrincipalByIDInOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamServiceGetUserPrincipalByIDInOrganizationOK creates a IamServiceGetUserPrincipalByIDInOrganizationOK with default headers values
func NewIamServiceGetUserPrincipalByIDInOrganizationOK() *IamServiceGetUserPrincipalByIDInOrganizationOK {
	return &IamServiceGetUserPrincipalByIDInOrganizationOK{}
}

/*
IamServiceGetUserPrincipalByIDInOrganizationOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamServiceGetUserPrincipalByIDInOrganizationOK struct {
	Payload *models.HashicorpCloudIamUserPrincipalResponse
}

// IsSuccess returns true when this iam service get user principal by Id in organization o k response has a 2xx status code
func (o *IamServiceGetUserPrincipalByIDInOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam service get user principal by Id in organization o k response has a 3xx status code
func (o *IamServiceGetUserPrincipalByIDInOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam service get user principal by Id in organization o k response has a 4xx status code
func (o *IamServiceGetUserPrincipalByIDInOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam service get user principal by Id in organization o k response has a 5xx status code
func (o *IamServiceGetUserPrincipalByIDInOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam service get user principal by Id in organization o k response a status code equal to that given
func (o *IamServiceGetUserPrincipalByIDInOrganizationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the iam service get user principal by Id in organization o k response
func (o *IamServiceGetUserPrincipalByIDInOrganizationOK) Code() int {
	return 200
}

func (o *IamServiceGetUserPrincipalByIDInOrganizationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/user-principals/{user_principal_id}][%d] iamServiceGetUserPrincipalByIdInOrganizationOK %s", 200, payload)
}

func (o *IamServiceGetUserPrincipalByIDInOrganizationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/user-principals/{user_principal_id}][%d] iamServiceGetUserPrincipalByIdInOrganizationOK %s", 200, payload)
}

func (o *IamServiceGetUserPrincipalByIDInOrganizationOK) GetPayload() *models.HashicorpCloudIamUserPrincipalResponse {
	return o.Payload
}

func (o *IamServiceGetUserPrincipalByIDInOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamUserPrincipalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamServiceGetUserPrincipalByIDInOrganizationDefault creates a IamServiceGetUserPrincipalByIDInOrganizationDefault with default headers values
func NewIamServiceGetUserPrincipalByIDInOrganizationDefault(code int) *IamServiceGetUserPrincipalByIDInOrganizationDefault {
	return &IamServiceGetUserPrincipalByIDInOrganizationDefault{
		_statusCode: code,
	}
}

/*
IamServiceGetUserPrincipalByIDInOrganizationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IamServiceGetUserPrincipalByIDInOrganizationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this iam service get user principal by Id in organization default response has a 2xx status code
func (o *IamServiceGetUserPrincipalByIDInOrganizationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam service get user principal by Id in organization default response has a 3xx status code
func (o *IamServiceGetUserPrincipalByIDInOrganizationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam service get user principal by Id in organization default response has a 4xx status code
func (o *IamServiceGetUserPrincipalByIDInOrganizationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam service get user principal by Id in organization default response has a 5xx status code
func (o *IamServiceGetUserPrincipalByIDInOrganizationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam service get user principal by Id in organization default response a status code equal to that given
func (o *IamServiceGetUserPrincipalByIDInOrganizationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the iam service get user principal by Id in organization default response
func (o *IamServiceGetUserPrincipalByIDInOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *IamServiceGetUserPrincipalByIDInOrganizationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/user-principals/{user_principal_id}][%d] IamService_GetUserPrincipalByIdInOrganization default %s", o._statusCode, payload)
}

func (o *IamServiceGetUserPrincipalByIDInOrganizationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/user-principals/{user_principal_id}][%d] IamService_GetUserPrincipalByIdInOrganization default %s", o._statusCode, payload)
}

func (o *IamServiceGetUserPrincipalByIDInOrganizationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *IamServiceGetUserPrincipalByIDInOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
