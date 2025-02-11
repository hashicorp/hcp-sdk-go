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

// SSOManagementServiceDeleteSSOConfigurationReader is a Reader for the SSOManagementServiceDeleteSSOConfiguration structure.
type SSOManagementServiceDeleteSSOConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SSOManagementServiceDeleteSSOConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSSOManagementServiceDeleteSSOConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSSOManagementServiceDeleteSSOConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSSOManagementServiceDeleteSSOConfigurationOK creates a SSOManagementServiceDeleteSSOConfigurationOK with default headers values
func NewSSOManagementServiceDeleteSSOConfigurationOK() *SSOManagementServiceDeleteSSOConfigurationOK {
	return &SSOManagementServiceDeleteSSOConfigurationOK{}
}

/*
SSOManagementServiceDeleteSSOConfigurationOK describes a response with status code 200, with default header values.

A successful response.
*/
type SSOManagementServiceDeleteSSOConfigurationOK struct {
	Payload models.HashicorpCloudIamDeleteSSOConfigurationResponse
}

// IsSuccess returns true when this s s o management service delete s s o configuration o k response has a 2xx status code
func (o *SSOManagementServiceDeleteSSOConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s s o management service delete s s o configuration o k response has a 3xx status code
func (o *SSOManagementServiceDeleteSSOConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s s o management service delete s s o configuration o k response has a 4xx status code
func (o *SSOManagementServiceDeleteSSOConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s s o management service delete s s o configuration o k response has a 5xx status code
func (o *SSOManagementServiceDeleteSSOConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s s o management service delete s s o configuration o k response a status code equal to that given
func (o *SSOManagementServiceDeleteSSOConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s s o management service delete s s o configuration o k response
func (o *SSOManagementServiceDeleteSSOConfigurationOK) Code() int {
	return 200
}

func (o *SSOManagementServiceDeleteSSOConfigurationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] sSOManagementServiceDeleteSSOConfigurationOK %s", 200, payload)
}

func (o *SSOManagementServiceDeleteSSOConfigurationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] sSOManagementServiceDeleteSSOConfigurationOK %s", 200, payload)
}

func (o *SSOManagementServiceDeleteSSOConfigurationOK) GetPayload() models.HashicorpCloudIamDeleteSSOConfigurationResponse {
	return o.Payload
}

func (o *SSOManagementServiceDeleteSSOConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSOManagementServiceDeleteSSOConfigurationDefault creates a SSOManagementServiceDeleteSSOConfigurationDefault with default headers values
func NewSSOManagementServiceDeleteSSOConfigurationDefault(code int) *SSOManagementServiceDeleteSSOConfigurationDefault {
	return &SSOManagementServiceDeleteSSOConfigurationDefault{
		_statusCode: code,
	}
}

/*
SSOManagementServiceDeleteSSOConfigurationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SSOManagementServiceDeleteSSOConfigurationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this s s o management service delete s s o configuration default response has a 2xx status code
func (o *SSOManagementServiceDeleteSSOConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s s o management service delete s s o configuration default response has a 3xx status code
func (o *SSOManagementServiceDeleteSSOConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s s o management service delete s s o configuration default response has a 4xx status code
func (o *SSOManagementServiceDeleteSSOConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s s o management service delete s s o configuration default response has a 5xx status code
func (o *SSOManagementServiceDeleteSSOConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s s o management service delete s s o configuration default response a status code equal to that given
func (o *SSOManagementServiceDeleteSSOConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s s o management service delete s s o configuration default response
func (o *SSOManagementServiceDeleteSSOConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *SSOManagementServiceDeleteSSOConfigurationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] SSOManagementService_DeleteSSOConfiguration default %s", o._statusCode, payload)
}

func (o *SSOManagementServiceDeleteSSOConfigurationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] SSOManagementService_DeleteSSOConfiguration default %s", o._statusCode, payload)
}

func (o *SSOManagementServiceDeleteSSOConfigurationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *SSOManagementServiceDeleteSSOConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
