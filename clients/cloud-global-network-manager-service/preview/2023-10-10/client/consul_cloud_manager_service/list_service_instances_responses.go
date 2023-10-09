// Code generated by go-swagger; DO NOT EDIT.

package consul_cloud_manager_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-global-network-manager-service/preview/2023-10-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ListServiceInstancesReader is a Reader for the ListServiceInstances structure.
type ListServiceInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListServiceInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListServiceInstancesOK creates a ListServiceInstancesOK with default headers values
func NewListServiceInstancesOK() *ListServiceInstancesOK {
	return &ListServiceInstancesOK{}
}

/*
ListServiceInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListServiceInstancesOK struct {
	Payload *models.HashicorpCloudGlobalNetworkManager20220215ListServiceInstancesResponse
}

// IsSuccess returns true when this list service instances o k response has a 2xx status code
func (o *ListServiceInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list service instances o k response has a 3xx status code
func (o *ListServiceInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list service instances o k response has a 4xx status code
func (o *ListServiceInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list service instances o k response has a 5xx status code
func (o *ListServiceInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list service instances o k response a status code equal to that given
func (o *ListServiceInstancesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListServiceInstancesOK) Error() string {
	return fmt.Sprintf("[GET /2023-10-10/{cluster_resource_name}/service/{service_name}/instances][%d] listServiceInstancesOK  %+v", 200, o.Payload)
}

func (o *ListServiceInstancesOK) String() string {
	return fmt.Sprintf("[GET /2023-10-10/{cluster_resource_name}/service/{service_name}/instances][%d] listServiceInstancesOK  %+v", 200, o.Payload)
}

func (o *ListServiceInstancesOK) GetPayload() *models.HashicorpCloudGlobalNetworkManager20220215ListServiceInstancesResponse {
	return o.Payload
}

func (o *ListServiceInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudGlobalNetworkManager20220215ListServiceInstancesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceInstancesDefault creates a ListServiceInstancesDefault with default headers values
func NewListServiceInstancesDefault(code int) *ListServiceInstancesDefault {
	return &ListServiceInstancesDefault{
		_statusCode: code,
	}
}

/*
ListServiceInstancesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListServiceInstancesDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the list service instances default response
func (o *ListServiceInstancesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list service instances default response has a 2xx status code
func (o *ListServiceInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list service instances default response has a 3xx status code
func (o *ListServiceInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list service instances default response has a 4xx status code
func (o *ListServiceInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list service instances default response has a 5xx status code
func (o *ListServiceInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list service instances default response a status code equal to that given
func (o *ListServiceInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListServiceInstancesDefault) Error() string {
	return fmt.Sprintf("[GET /2023-10-10/{cluster_resource_name}/service/{service_name}/instances][%d] ListServiceInstances default  %+v", o._statusCode, o.Payload)
}

func (o *ListServiceInstancesDefault) String() string {
	return fmt.Sprintf("[GET /2023-10-10/{cluster_resource_name}/service/{service_name}/instances][%d] ListServiceInstances default  %+v", o._statusCode, o.Payload)
}

func (o *ListServiceInstancesDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ListServiceInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
