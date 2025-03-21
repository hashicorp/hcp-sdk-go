// Code generated by go-swagger; DO NOT EDIT.

package iam_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// IamServiceGetUserPrincipalsByIDsInOrganizationReader is a Reader for the IamServiceGetUserPrincipalsByIDsInOrganization structure.
type IamServiceGetUserPrincipalsByIDsInOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamServiceGetUserPrincipalsByIDsInOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamServiceGetUserPrincipalsByIDsInOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamServiceGetUserPrincipalsByIDsInOrganizationOK creates a IamServiceGetUserPrincipalsByIDsInOrganizationOK with default headers values
func NewIamServiceGetUserPrincipalsByIDsInOrganizationOK() *IamServiceGetUserPrincipalsByIDsInOrganizationOK {
	return &IamServiceGetUserPrincipalsByIDsInOrganizationOK{}
}

/*
IamServiceGetUserPrincipalsByIDsInOrganizationOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamServiceGetUserPrincipalsByIDsInOrganizationOK struct {
	Payload *models.HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationResponse
}

// IsSuccess returns true when this iam service get user principals by i ds in organization o k response has a 2xx status code
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam service get user principals by i ds in organization o k response has a 3xx status code
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam service get user principals by i ds in organization o k response has a 4xx status code
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam service get user principals by i ds in organization o k response has a 5xx status code
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam service get user principals by i ds in organization o k response a status code equal to that given
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the iam service get user principals by i ds in organization o k response
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationOK) Code() int {
	return 200
}

func (o *IamServiceGetUserPrincipalsByIDsInOrganizationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/user-principals/batch-fetch][%d] iamServiceGetUserPrincipalsByIDsInOrganizationOK %s", 200, payload)
}

func (o *IamServiceGetUserPrincipalsByIDsInOrganizationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/user-principals/batch-fetch][%d] iamServiceGetUserPrincipalsByIDsInOrganizationOK %s", 200, payload)
}

func (o *IamServiceGetUserPrincipalsByIDsInOrganizationOK) GetPayload() *models.HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationResponse {
	return o.Payload
}

func (o *IamServiceGetUserPrincipalsByIDsInOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamServiceGetUserPrincipalsByIDsInOrganizationDefault creates a IamServiceGetUserPrincipalsByIDsInOrganizationDefault with default headers values
func NewIamServiceGetUserPrincipalsByIDsInOrganizationDefault(code int) *IamServiceGetUserPrincipalsByIDsInOrganizationDefault {
	return &IamServiceGetUserPrincipalsByIDsInOrganizationDefault{
		_statusCode: code,
	}
}

/*
IamServiceGetUserPrincipalsByIDsInOrganizationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IamServiceGetUserPrincipalsByIDsInOrganizationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this iam service get user principals by i ds in organization default response has a 2xx status code
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam service get user principals by i ds in organization default response has a 3xx status code
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam service get user principals by i ds in organization default response has a 4xx status code
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam service get user principals by i ds in organization default response has a 5xx status code
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam service get user principals by i ds in organization default response a status code equal to that given
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the iam service get user principals by i ds in organization default response
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *IamServiceGetUserPrincipalsByIDsInOrganizationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/user-principals/batch-fetch][%d] IamService_GetUserPrincipalsByIDsInOrganization default %s", o._statusCode, payload)
}

func (o *IamServiceGetUserPrincipalsByIDsInOrganizationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/user-principals/batch-fetch][%d] IamService_GetUserPrincipalsByIDsInOrganization default %s", o._statusCode, payload)
}

func (o *IamServiceGetUserPrincipalsByIDsInOrganizationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *IamServiceGetUserPrincipalsByIDsInOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
IamServiceGetUserPrincipalsByIDsInOrganizationBody GetUserPrincipalsByIDsInOrganizationRequest is a request for users by ID in a given organization
swagger:model IamServiceGetUserPrincipalsByIDsInOrganizationBody
*/
type IamServiceGetUserPrincipalsByIDsInOrganizationBody struct {

	// ids is a list of user IDs to look up
	Ids []string `json:"ids"`
}

// Validate validates this iam service get user principals by i ds in organization body
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this iam service get user principals by i ds in organization body based on context it is used
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationBody) UnmarshalBinary(b []byte) error {
	var res IamServiceGetUserPrincipalsByIDsInOrganizationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
