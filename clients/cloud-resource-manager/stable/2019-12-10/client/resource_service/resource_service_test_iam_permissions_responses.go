// Code generated by go-swagger; DO NOT EDIT.

package resource_service

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

// ResourceServiceTestIamPermissionsReader is a Reader for the ResourceServiceTestIamPermissions structure.
type ResourceServiceTestIamPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceServiceTestIamPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceServiceTestIamPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResourceServiceTestIamPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResourceServiceTestIamPermissionsOK creates a ResourceServiceTestIamPermissionsOK with default headers values
func NewResourceServiceTestIamPermissionsOK() *ResourceServiceTestIamPermissionsOK {
	return &ResourceServiceTestIamPermissionsOK{}
}

/*
ResourceServiceTestIamPermissionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResourceServiceTestIamPermissionsOK struct {
	Payload *models.HashicorpCloudResourcemanagerResourceTestIamPermissionsResponse
}

// IsSuccess returns true when this resource service test iam permissions o k response has a 2xx status code
func (o *ResourceServiceTestIamPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource service test iam permissions o k response has a 3xx status code
func (o *ResourceServiceTestIamPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource service test iam permissions o k response has a 4xx status code
func (o *ResourceServiceTestIamPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource service test iam permissions o k response has a 5xx status code
func (o *ResourceServiceTestIamPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource service test iam permissions o k response a status code equal to that given
func (o *ResourceServiceTestIamPermissionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resource service test iam permissions o k response
func (o *ResourceServiceTestIamPermissionsOK) Code() int {
	return 200
}

func (o *ResourceServiceTestIamPermissionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/resources/test-iam-permissions][%d] resourceServiceTestIamPermissionsOK %s", 200, payload)
}

func (o *ResourceServiceTestIamPermissionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/resources/test-iam-permissions][%d] resourceServiceTestIamPermissionsOK %s", 200, payload)
}

func (o *ResourceServiceTestIamPermissionsOK) GetPayload() *models.HashicorpCloudResourcemanagerResourceTestIamPermissionsResponse {
	return o.Payload
}

func (o *ResourceServiceTestIamPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerResourceTestIamPermissionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceServiceTestIamPermissionsDefault creates a ResourceServiceTestIamPermissionsDefault with default headers values
func NewResourceServiceTestIamPermissionsDefault(code int) *ResourceServiceTestIamPermissionsDefault {
	return &ResourceServiceTestIamPermissionsDefault{
		_statusCode: code,
	}
}

/*
ResourceServiceTestIamPermissionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResourceServiceTestIamPermissionsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this resource service test iam permissions default response has a 2xx status code
func (o *ResourceServiceTestIamPermissionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resource service test iam permissions default response has a 3xx status code
func (o *ResourceServiceTestIamPermissionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resource service test iam permissions default response has a 4xx status code
func (o *ResourceServiceTestIamPermissionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resource service test iam permissions default response has a 5xx status code
func (o *ResourceServiceTestIamPermissionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resource service test iam permissions default response a status code equal to that given
func (o *ResourceServiceTestIamPermissionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resource service test iam permissions default response
func (o *ResourceServiceTestIamPermissionsDefault) Code() int {
	return o._statusCode
}

func (o *ResourceServiceTestIamPermissionsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/resources/test-iam-permissions][%d] ResourceService_TestIamPermissions default %s", o._statusCode, payload)
}

func (o *ResourceServiceTestIamPermissionsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/resources/test-iam-permissions][%d] ResourceService_TestIamPermissions default %s", o._statusCode, payload)
}

func (o *ResourceServiceTestIamPermissionsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ResourceServiceTestIamPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
