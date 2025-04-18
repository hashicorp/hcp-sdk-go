// Code generated by go-swagger; DO NOT EDIT.

package authorization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// AuthorizationServiceTestIamPermissionsReader is a Reader for the AuthorizationServiceTestIamPermissions structure.
type AuthorizationServiceTestIamPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthorizationServiceTestIamPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthorizationServiceTestIamPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuthorizationServiceTestIamPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthorizationServiceTestIamPermissionsOK creates a AuthorizationServiceTestIamPermissionsOK with default headers values
func NewAuthorizationServiceTestIamPermissionsOK() *AuthorizationServiceTestIamPermissionsOK {
	return &AuthorizationServiceTestIamPermissionsOK{}
}

/*
AuthorizationServiceTestIamPermissionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuthorizationServiceTestIamPermissionsOK struct {
	Payload *models.HashicorpCloudResourcemanagerAuthorizationTestIamPermissionsResponse
}

// IsSuccess returns true when this authorization service test iam permissions o k response has a 2xx status code
func (o *AuthorizationServiceTestIamPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authorization service test iam permissions o k response has a 3xx status code
func (o *AuthorizationServiceTestIamPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorization service test iam permissions o k response has a 4xx status code
func (o *AuthorizationServiceTestIamPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authorization service test iam permissions o k response has a 5xx status code
func (o *AuthorizationServiceTestIamPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authorization service test iam permissions o k response a status code equal to that given
func (o *AuthorizationServiceTestIamPermissionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the authorization service test iam permissions o k response
func (o *AuthorizationServiceTestIamPermissionsOK) Code() int {
	return 200
}

func (o *AuthorizationServiceTestIamPermissionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/authorization/test-iam-permissions][%d] authorizationServiceTestIamPermissionsOK %s", 200, payload)
}

func (o *AuthorizationServiceTestIamPermissionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/authorization/test-iam-permissions][%d] authorizationServiceTestIamPermissionsOK %s", 200, payload)
}

func (o *AuthorizationServiceTestIamPermissionsOK) GetPayload() *models.HashicorpCloudResourcemanagerAuthorizationTestIamPermissionsResponse {
	return o.Payload
}

func (o *AuthorizationServiceTestIamPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerAuthorizationTestIamPermissionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthorizationServiceTestIamPermissionsDefault creates a AuthorizationServiceTestIamPermissionsDefault with default headers values
func NewAuthorizationServiceTestIamPermissionsDefault(code int) *AuthorizationServiceTestIamPermissionsDefault {
	return &AuthorizationServiceTestIamPermissionsDefault{
		_statusCode: code,
	}
}

/*
AuthorizationServiceTestIamPermissionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AuthorizationServiceTestIamPermissionsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this authorization service test iam permissions default response has a 2xx status code
func (o *AuthorizationServiceTestIamPermissionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this authorization service test iam permissions default response has a 3xx status code
func (o *AuthorizationServiceTestIamPermissionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this authorization service test iam permissions default response has a 4xx status code
func (o *AuthorizationServiceTestIamPermissionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this authorization service test iam permissions default response has a 5xx status code
func (o *AuthorizationServiceTestIamPermissionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this authorization service test iam permissions default response a status code equal to that given
func (o *AuthorizationServiceTestIamPermissionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the authorization service test iam permissions default response
func (o *AuthorizationServiceTestIamPermissionsDefault) Code() int {
	return o._statusCode
}

func (o *AuthorizationServiceTestIamPermissionsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/authorization/test-iam-permissions][%d] AuthorizationService_TestIamPermissions default %s", o._statusCode, payload)
}

func (o *AuthorizationServiceTestIamPermissionsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/authorization/test-iam-permissions][%d] AuthorizationService_TestIamPermissions default %s", o._statusCode, payload)
}

func (o *AuthorizationServiceTestIamPermissionsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *AuthorizationServiceTestIamPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
