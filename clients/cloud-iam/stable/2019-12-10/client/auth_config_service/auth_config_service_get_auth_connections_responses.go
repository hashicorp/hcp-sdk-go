// Code generated by go-swagger; DO NOT EDIT.

package auth_config_service

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

// AuthConfigServiceGetAuthConnectionsReader is a Reader for the AuthConfigServiceGetAuthConnections structure.
type AuthConfigServiceGetAuthConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthConfigServiceGetAuthConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthConfigServiceGetAuthConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuthConfigServiceGetAuthConnectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthConfigServiceGetAuthConnectionsOK creates a AuthConfigServiceGetAuthConnectionsOK with default headers values
func NewAuthConfigServiceGetAuthConnectionsOK() *AuthConfigServiceGetAuthConnectionsOK {
	return &AuthConfigServiceGetAuthConnectionsOK{}
}

/*
AuthConfigServiceGetAuthConnectionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuthConfigServiceGetAuthConnectionsOK struct {
	Payload *models.HashicorpCloudIamGetAuthConnectionsResponse
}

// IsSuccess returns true when this auth config service get auth connections o k response has a 2xx status code
func (o *AuthConfigServiceGetAuthConnectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auth config service get auth connections o k response has a 3xx status code
func (o *AuthConfigServiceGetAuthConnectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth config service get auth connections o k response has a 4xx status code
func (o *AuthConfigServiceGetAuthConnectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auth config service get auth connections o k response has a 5xx status code
func (o *AuthConfigServiceGetAuthConnectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auth config service get auth connections o k response a status code equal to that given
func (o *AuthConfigServiceGetAuthConnectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the auth config service get auth connections o k response
func (o *AuthConfigServiceGetAuthConnectionsOK) Code() int {
	return 200
}

func (o *AuthConfigServiceGetAuthConnectionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] authConfigServiceGetAuthConnectionsOK %s", 200, payload)
}

func (o *AuthConfigServiceGetAuthConnectionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] authConfigServiceGetAuthConnectionsOK %s", 200, payload)
}

func (o *AuthConfigServiceGetAuthConnectionsOK) GetPayload() *models.HashicorpCloudIamGetAuthConnectionsResponse {
	return o.Payload
}

func (o *AuthConfigServiceGetAuthConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamGetAuthConnectionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthConfigServiceGetAuthConnectionsDefault creates a AuthConfigServiceGetAuthConnectionsDefault with default headers values
func NewAuthConfigServiceGetAuthConnectionsDefault(code int) *AuthConfigServiceGetAuthConnectionsDefault {
	return &AuthConfigServiceGetAuthConnectionsDefault{
		_statusCode: code,
	}
}

/*
AuthConfigServiceGetAuthConnectionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AuthConfigServiceGetAuthConnectionsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this auth config service get auth connections default response has a 2xx status code
func (o *AuthConfigServiceGetAuthConnectionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this auth config service get auth connections default response has a 3xx status code
func (o *AuthConfigServiceGetAuthConnectionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this auth config service get auth connections default response has a 4xx status code
func (o *AuthConfigServiceGetAuthConnectionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this auth config service get auth connections default response has a 5xx status code
func (o *AuthConfigServiceGetAuthConnectionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this auth config service get auth connections default response a status code equal to that given
func (o *AuthConfigServiceGetAuthConnectionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the auth config service get auth connections default response
func (o *AuthConfigServiceGetAuthConnectionsDefault) Code() int {
	return o._statusCode
}

func (o *AuthConfigServiceGetAuthConnectionsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] AuthConfigService_GetAuthConnections default %s", o._statusCode, payload)
}

func (o *AuthConfigServiceGetAuthConnectionsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] AuthConfigService_GetAuthConnections default %s", o._statusCode, payload)
}

func (o *AuthConfigServiceGetAuthConnectionsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *AuthConfigServiceGetAuthConnectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
