// Code generated by go-swagger; DO NOT EDIT.

package resource_service

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

// ResourceServiceGetResourceReader is a Reader for the ResourceServiceGetResource structure.
type ResourceServiceGetResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceServiceGetResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceServiceGetResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResourceServiceGetResourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResourceServiceGetResourceOK creates a ResourceServiceGetResourceOK with default headers values
func NewResourceServiceGetResourceOK() *ResourceServiceGetResourceOK {
	return &ResourceServiceGetResourceOK{}
}

/*
ResourceServiceGetResourceOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResourceServiceGetResourceOK struct {
	Payload *models.HashicorpCloudResourcemanagerGetResourceResponse
}

// IsSuccess returns true when this resource service get resource o k response has a 2xx status code
func (o *ResourceServiceGetResourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource service get resource o k response has a 3xx status code
func (o *ResourceServiceGetResourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource service get resource o k response has a 4xx status code
func (o *ResourceServiceGetResourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource service get resource o k response has a 5xx status code
func (o *ResourceServiceGetResourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource service get resource o k response a status code equal to that given
func (o *ResourceServiceGetResourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resource service get resource o k response
func (o *ResourceServiceGetResourceOK) Code() int {
	return 200
}

func (o *ResourceServiceGetResourceOK) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/resource][%d] resourceServiceGetResourceOK  %+v", 200, o.Payload)
}

func (o *ResourceServiceGetResourceOK) String() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/resource][%d] resourceServiceGetResourceOK  %+v", 200, o.Payload)
}

func (o *ResourceServiceGetResourceOK) GetPayload() *models.HashicorpCloudResourcemanagerGetResourceResponse {
	return o.Payload
}

func (o *ResourceServiceGetResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerGetResourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceServiceGetResourceDefault creates a ResourceServiceGetResourceDefault with default headers values
func NewResourceServiceGetResourceDefault(code int) *ResourceServiceGetResourceDefault {
	return &ResourceServiceGetResourceDefault{
		_statusCode: code,
	}
}

/*
ResourceServiceGetResourceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResourceServiceGetResourceDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this resource service get resource default response has a 2xx status code
func (o *ResourceServiceGetResourceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resource service get resource default response has a 3xx status code
func (o *ResourceServiceGetResourceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resource service get resource default response has a 4xx status code
func (o *ResourceServiceGetResourceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resource service get resource default response has a 5xx status code
func (o *ResourceServiceGetResourceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resource service get resource default response a status code equal to that given
func (o *ResourceServiceGetResourceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resource service get resource default response
func (o *ResourceServiceGetResourceDefault) Code() int {
	return o._statusCode
}

func (o *ResourceServiceGetResourceDefault) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/resource][%d] ResourceService_GetResource default  %+v", o._statusCode, o.Payload)
}

func (o *ResourceServiceGetResourceDefault) String() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/resource][%d] ResourceService_GetResource default  %+v", o._statusCode, o.Payload)
}

func (o *ResourceServiceGetResourceDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ResourceServiceGetResourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}