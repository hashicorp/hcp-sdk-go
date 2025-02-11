// Code generated by go-swagger; DO NOT EDIT.

package boundary_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-boundary-service/stable/2021-12-21/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// BoundaryServiceUpdateControllerConfigurationReader is a Reader for the BoundaryServiceUpdateControllerConfiguration structure.
type BoundaryServiceUpdateControllerConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BoundaryServiceUpdateControllerConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBoundaryServiceUpdateControllerConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBoundaryServiceUpdateControllerConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBoundaryServiceUpdateControllerConfigurationOK creates a BoundaryServiceUpdateControllerConfigurationOK with default headers values
func NewBoundaryServiceUpdateControllerConfigurationOK() *BoundaryServiceUpdateControllerConfigurationOK {
	return &BoundaryServiceUpdateControllerConfigurationOK{}
}

/*
BoundaryServiceUpdateControllerConfigurationOK describes a response with status code 200, with default header values.

A successful response.
*/
type BoundaryServiceUpdateControllerConfigurationOK struct {
	Payload models.HashicorpCloudBoundary20211221UpdateControllerConfigurationResponse
}

// IsSuccess returns true when this boundary service update controller configuration o k response has a 2xx status code
func (o *BoundaryServiceUpdateControllerConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this boundary service update controller configuration o k response has a 3xx status code
func (o *BoundaryServiceUpdateControllerConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this boundary service update controller configuration o k response has a 4xx status code
func (o *BoundaryServiceUpdateControllerConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this boundary service update controller configuration o k response has a 5xx status code
func (o *BoundaryServiceUpdateControllerConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this boundary service update controller configuration o k response a status code equal to that given
func (o *BoundaryServiceUpdateControllerConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the boundary service update controller configuration o k response
func (o *BoundaryServiceUpdateControllerConfigurationOK) Code() int {
	return 200
}

func (o *BoundaryServiceUpdateControllerConfigurationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/controller-configuration][%d] boundaryServiceUpdateControllerConfigurationOK %s", 200, payload)
}

func (o *BoundaryServiceUpdateControllerConfigurationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/controller-configuration][%d] boundaryServiceUpdateControllerConfigurationOK %s", 200, payload)
}

func (o *BoundaryServiceUpdateControllerConfigurationOK) GetPayload() models.HashicorpCloudBoundary20211221UpdateControllerConfigurationResponse {
	return o.Payload
}

func (o *BoundaryServiceUpdateControllerConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBoundaryServiceUpdateControllerConfigurationDefault creates a BoundaryServiceUpdateControllerConfigurationDefault with default headers values
func NewBoundaryServiceUpdateControllerConfigurationDefault(code int) *BoundaryServiceUpdateControllerConfigurationDefault {
	return &BoundaryServiceUpdateControllerConfigurationDefault{
		_statusCode: code,
	}
}

/*
BoundaryServiceUpdateControllerConfigurationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BoundaryServiceUpdateControllerConfigurationDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this boundary service update controller configuration default response has a 2xx status code
func (o *BoundaryServiceUpdateControllerConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this boundary service update controller configuration default response has a 3xx status code
func (o *BoundaryServiceUpdateControllerConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this boundary service update controller configuration default response has a 4xx status code
func (o *BoundaryServiceUpdateControllerConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this boundary service update controller configuration default response has a 5xx status code
func (o *BoundaryServiceUpdateControllerConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this boundary service update controller configuration default response a status code equal to that given
func (o *BoundaryServiceUpdateControllerConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the boundary service update controller configuration default response
func (o *BoundaryServiceUpdateControllerConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *BoundaryServiceUpdateControllerConfigurationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/controller-configuration][%d] BoundaryService_UpdateControllerConfiguration default %s", o._statusCode, payload)
}

func (o *BoundaryServiceUpdateControllerConfigurationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/controller-configuration][%d] BoundaryService_UpdateControllerConfiguration default %s", o._statusCode, payload)
}

func (o *BoundaryServiceUpdateControllerConfigurationDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *BoundaryServiceUpdateControllerConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
