// Code generated by go-swagger; DO NOT EDIT.

package authorization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// AuthorizationServiceBatchTestIamPermissionsReader is a Reader for the AuthorizationServiceBatchTestIamPermissions structure.
type AuthorizationServiceBatchTestIamPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthorizationServiceBatchTestIamPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthorizationServiceBatchTestIamPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuthorizationServiceBatchTestIamPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthorizationServiceBatchTestIamPermissionsOK creates a AuthorizationServiceBatchTestIamPermissionsOK with default headers values
func NewAuthorizationServiceBatchTestIamPermissionsOK() *AuthorizationServiceBatchTestIamPermissionsOK {
	return &AuthorizationServiceBatchTestIamPermissionsOK{}
}

/*
AuthorizationServiceBatchTestIamPermissionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuthorizationServiceBatchTestIamPermissionsOK struct {
	Payload *models.HashicorpCloudResourcemanagerBatchAuthorizationTestIamPermissionsResponse
}

// IsSuccess returns true when this authorization service batch test iam permissions o k response has a 2xx status code
func (o *AuthorizationServiceBatchTestIamPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authorization service batch test iam permissions o k response has a 3xx status code
func (o *AuthorizationServiceBatchTestIamPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorization service batch test iam permissions o k response has a 4xx status code
func (o *AuthorizationServiceBatchTestIamPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authorization service batch test iam permissions o k response has a 5xx status code
func (o *AuthorizationServiceBatchTestIamPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authorization service batch test iam permissions o k response a status code equal to that given
func (o *AuthorizationServiceBatchTestIamPermissionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the authorization service batch test iam permissions o k response
func (o *AuthorizationServiceBatchTestIamPermissionsOK) Code() int {
	return 200
}

func (o *AuthorizationServiceBatchTestIamPermissionsOK) Error() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/authorization/batch-test-iam-permissions][%d] authorizationServiceBatchTestIamPermissionsOK  %+v", 200, o.Payload)
}

func (o *AuthorizationServiceBatchTestIamPermissionsOK) String() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/authorization/batch-test-iam-permissions][%d] authorizationServiceBatchTestIamPermissionsOK  %+v", 200, o.Payload)
}

func (o *AuthorizationServiceBatchTestIamPermissionsOK) GetPayload() *models.HashicorpCloudResourcemanagerBatchAuthorizationTestIamPermissionsResponse {
	return o.Payload
}

func (o *AuthorizationServiceBatchTestIamPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerBatchAuthorizationTestIamPermissionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthorizationServiceBatchTestIamPermissionsDefault creates a AuthorizationServiceBatchTestIamPermissionsDefault with default headers values
func NewAuthorizationServiceBatchTestIamPermissionsDefault(code int) *AuthorizationServiceBatchTestIamPermissionsDefault {
	return &AuthorizationServiceBatchTestIamPermissionsDefault{
		_statusCode: code,
	}
}

/*
AuthorizationServiceBatchTestIamPermissionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AuthorizationServiceBatchTestIamPermissionsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this authorization service batch test iam permissions default response has a 2xx status code
func (o *AuthorizationServiceBatchTestIamPermissionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this authorization service batch test iam permissions default response has a 3xx status code
func (o *AuthorizationServiceBatchTestIamPermissionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this authorization service batch test iam permissions default response has a 4xx status code
func (o *AuthorizationServiceBatchTestIamPermissionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this authorization service batch test iam permissions default response has a 5xx status code
func (o *AuthorizationServiceBatchTestIamPermissionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this authorization service batch test iam permissions default response a status code equal to that given
func (o *AuthorizationServiceBatchTestIamPermissionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the authorization service batch test iam permissions default response
func (o *AuthorizationServiceBatchTestIamPermissionsDefault) Code() int {
	return o._statusCode
}

func (o *AuthorizationServiceBatchTestIamPermissionsDefault) Error() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/authorization/batch-test-iam-permissions][%d] AuthorizationService_BatchTestIamPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *AuthorizationServiceBatchTestIamPermissionsDefault) String() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/authorization/batch-test-iam-permissions][%d] AuthorizationService_BatchTestIamPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *AuthorizationServiceBatchTestIamPermissionsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *AuthorizationServiceBatchTestIamPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
