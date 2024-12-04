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

// ResourceServiceGetIamPolicyReader is a Reader for the ResourceServiceGetIamPolicy structure.
type ResourceServiceGetIamPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceServiceGetIamPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceServiceGetIamPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResourceServiceGetIamPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResourceServiceGetIamPolicyOK creates a ResourceServiceGetIamPolicyOK with default headers values
func NewResourceServiceGetIamPolicyOK() *ResourceServiceGetIamPolicyOK {
	return &ResourceServiceGetIamPolicyOK{}
}

/*
ResourceServiceGetIamPolicyOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResourceServiceGetIamPolicyOK struct {
	Payload *models.HashicorpCloudResourcemanagerResourceGetIamPolicyResponse
}

// IsSuccess returns true when this resource service get iam policy o k response has a 2xx status code
func (o *ResourceServiceGetIamPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource service get iam policy o k response has a 3xx status code
func (o *ResourceServiceGetIamPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource service get iam policy o k response has a 4xx status code
func (o *ResourceServiceGetIamPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource service get iam policy o k response has a 5xx status code
func (o *ResourceServiceGetIamPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource service get iam policy o k response a status code equal to that given
func (o *ResourceServiceGetIamPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resource service get iam policy o k response
func (o *ResourceServiceGetIamPolicyOK) Code() int {
	return 200
}

func (o *ResourceServiceGetIamPolicyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /2019-12-10/resource-manager/resources/iam-policy][%d] resourceServiceGetIamPolicyOK %s", 200, payload)
}

func (o *ResourceServiceGetIamPolicyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /2019-12-10/resource-manager/resources/iam-policy][%d] resourceServiceGetIamPolicyOK %s", 200, payload)
}

func (o *ResourceServiceGetIamPolicyOK) GetPayload() *models.HashicorpCloudResourcemanagerResourceGetIamPolicyResponse {
	return o.Payload
}

func (o *ResourceServiceGetIamPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerResourceGetIamPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceServiceGetIamPolicyDefault creates a ResourceServiceGetIamPolicyDefault with default headers values
func NewResourceServiceGetIamPolicyDefault(code int) *ResourceServiceGetIamPolicyDefault {
	return &ResourceServiceGetIamPolicyDefault{
		_statusCode: code,
	}
}

/*
ResourceServiceGetIamPolicyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResourceServiceGetIamPolicyDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this resource service get iam policy default response has a 2xx status code
func (o *ResourceServiceGetIamPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resource service get iam policy default response has a 3xx status code
func (o *ResourceServiceGetIamPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resource service get iam policy default response has a 4xx status code
func (o *ResourceServiceGetIamPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resource service get iam policy default response has a 5xx status code
func (o *ResourceServiceGetIamPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resource service get iam policy default response a status code equal to that given
func (o *ResourceServiceGetIamPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resource service get iam policy default response
func (o *ResourceServiceGetIamPolicyDefault) Code() int {
	return o._statusCode
}

func (o *ResourceServiceGetIamPolicyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /2019-12-10/resource-manager/resources/iam-policy][%d] ResourceService_GetIamPolicy default %s", o._statusCode, payload)
}

func (o *ResourceServiceGetIamPolicyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /2019-12-10/resource-manager/resources/iam-policy][%d] ResourceService_GetIamPolicy default %s", o._statusCode, payload)
}

func (o *ResourceServiceGetIamPolicyDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ResourceServiceGetIamPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
