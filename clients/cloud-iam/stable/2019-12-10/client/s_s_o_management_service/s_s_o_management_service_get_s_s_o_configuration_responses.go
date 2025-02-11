// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_management_service

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

// SSOManagementServiceGetSSOConfigurationReader is a Reader for the SSOManagementServiceGetSSOConfiguration structure.
type SSOManagementServiceGetSSOConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SSOManagementServiceGetSSOConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSSOManagementServiceGetSSOConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSSOManagementServiceGetSSOConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSSOManagementServiceGetSSOConfigurationOK creates a SSOManagementServiceGetSSOConfigurationOK with default headers values
func NewSSOManagementServiceGetSSOConfigurationOK() *SSOManagementServiceGetSSOConfigurationOK {
	return &SSOManagementServiceGetSSOConfigurationOK{}
}

/*
SSOManagementServiceGetSSOConfigurationOK describes a response with status code 200, with default header values.

A successful response.
*/
type SSOManagementServiceGetSSOConfigurationOK struct {
	Payload *models.HashicorpCloudIamGetSSOConfigurationResponse
}

// IsSuccess returns true when this s s o management service get s s o configuration o k response has a 2xx status code
func (o *SSOManagementServiceGetSSOConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s s o management service get s s o configuration o k response has a 3xx status code
func (o *SSOManagementServiceGetSSOConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s s o management service get s s o configuration o k response has a 4xx status code
func (o *SSOManagementServiceGetSSOConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s s o management service get s s o configuration o k response has a 5xx status code
func (o *SSOManagementServiceGetSSOConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s s o management service get s s o configuration o k response a status code equal to that given
func (o *SSOManagementServiceGetSSOConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s s o management service get s s o configuration o k response
func (o *SSOManagementServiceGetSSOConfigurationOK) Code() int {
	return 200
}

func (o *SSOManagementServiceGetSSOConfigurationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] sSOManagementServiceGetSSOConfigurationOK %s", 200, payload)
}

func (o *SSOManagementServiceGetSSOConfigurationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] sSOManagementServiceGetSSOConfigurationOK %s", 200, payload)
}

func (o *SSOManagementServiceGetSSOConfigurationOK) GetPayload() *models.HashicorpCloudIamGetSSOConfigurationResponse {
	return o.Payload
}

func (o *SSOManagementServiceGetSSOConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamGetSSOConfigurationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSOManagementServiceGetSSOConfigurationDefault creates a SSOManagementServiceGetSSOConfigurationDefault with default headers values
func NewSSOManagementServiceGetSSOConfigurationDefault(code int) *SSOManagementServiceGetSSOConfigurationDefault {
	return &SSOManagementServiceGetSSOConfigurationDefault{
		_statusCode: code,
	}
}

/*
SSOManagementServiceGetSSOConfigurationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SSOManagementServiceGetSSOConfigurationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this s s o management service get s s o configuration default response has a 2xx status code
func (o *SSOManagementServiceGetSSOConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s s o management service get s s o configuration default response has a 3xx status code
func (o *SSOManagementServiceGetSSOConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s s o management service get s s o configuration default response has a 4xx status code
func (o *SSOManagementServiceGetSSOConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s s o management service get s s o configuration default response has a 5xx status code
func (o *SSOManagementServiceGetSSOConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s s o management service get s s o configuration default response a status code equal to that given
func (o *SSOManagementServiceGetSSOConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s s o management service get s s o configuration default response
func (o *SSOManagementServiceGetSSOConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *SSOManagementServiceGetSSOConfigurationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] SSOManagementService_GetSSOConfiguration default %s", o._statusCode, payload)
}

func (o *SSOManagementServiceGetSSOConfigurationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] SSOManagementService_GetSSOConfiguration default %s", o._statusCode, payload)
}

func (o *SSOManagementServiceGetSSOConfigurationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *SSOManagementServiceGetSSOConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
