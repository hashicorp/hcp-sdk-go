// Code generated by go-swagger; DO NOT EDIT.

package resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/preview/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ResourceServiceListReader is a Reader for the ResourceServiceList structure.
type ResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResourceServiceListOK creates a ResourceServiceListOK with default headers values
func NewResourceServiceListOK() *ResourceServiceListOK {
	return &ResourceServiceListOK{}
}

/*ResourceServiceListOK handles this case with default header values.

A successful response.
*/
type ResourceServiceListOK struct {
	Payload *models.HashicorpCloudResourcemanagerResourceListResponse
}

func (o *ResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/resources][%d] resourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ResourceServiceListOK) GetPayload() *models.HashicorpCloudResourcemanagerResourceListResponse {
	return o.Payload
}

func (o *ResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerResourceListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceServiceListDefault creates a ResourceServiceListDefault with default headers values
func NewResourceServiceListDefault(code int) *ResourceServiceListDefault {
	return &ResourceServiceListDefault{
		_statusCode: code,
	}
}

/*ResourceServiceListDefault handles this case with default header values.

An unexpected error response.
*/
type ResourceServiceListDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the resource service list default response
func (o *ResourceServiceListDefault) Code() int {
	return o._statusCode
}

func (o *ResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/resources][%d] ResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ResourceServiceListDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
