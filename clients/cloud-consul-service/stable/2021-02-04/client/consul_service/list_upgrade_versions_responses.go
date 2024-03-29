// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/stable/2021-02-04/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ListUpgradeVersionsReader is a Reader for the ListUpgradeVersions structure.
type ListUpgradeVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUpgradeVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUpgradeVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListUpgradeVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListUpgradeVersionsOK creates a ListUpgradeVersionsOK with default headers values
func NewListUpgradeVersionsOK() *ListUpgradeVersionsOK {
	return &ListUpgradeVersionsOK{}
}

/*
ListUpgradeVersionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListUpgradeVersionsOK struct {
	Payload *models.HashicorpCloudConsul20210204ListUpgradeVersionsResponse
}

// IsSuccess returns true when this list upgrade versions o k response has a 2xx status code
func (o *ListUpgradeVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list upgrade versions o k response has a 3xx status code
func (o *ListUpgradeVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list upgrade versions o k response has a 4xx status code
func (o *ListUpgradeVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list upgrade versions o k response has a 5xx status code
func (o *ListUpgradeVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list upgrade versions o k response a status code equal to that given
func (o *ListUpgradeVersionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListUpgradeVersionsOK) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/upgrade-versions][%d] listUpgradeVersionsOK  %+v", 200, o.Payload)
}

func (o *ListUpgradeVersionsOK) String() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/upgrade-versions][%d] listUpgradeVersionsOK  %+v", 200, o.Payload)
}

func (o *ListUpgradeVersionsOK) GetPayload() *models.HashicorpCloudConsul20210204ListUpgradeVersionsResponse {
	return o.Payload
}

func (o *ListUpgradeVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20210204ListUpgradeVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUpgradeVersionsDefault creates a ListUpgradeVersionsDefault with default headers values
func NewListUpgradeVersionsDefault(code int) *ListUpgradeVersionsDefault {
	return &ListUpgradeVersionsDefault{
		_statusCode: code,
	}
}

/*
ListUpgradeVersionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListUpgradeVersionsDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the list upgrade versions default response
func (o *ListUpgradeVersionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list upgrade versions default response has a 2xx status code
func (o *ListUpgradeVersionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list upgrade versions default response has a 3xx status code
func (o *ListUpgradeVersionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list upgrade versions default response has a 4xx status code
func (o *ListUpgradeVersionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list upgrade versions default response has a 5xx status code
func (o *ListUpgradeVersionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list upgrade versions default response a status code equal to that given
func (o *ListUpgradeVersionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListUpgradeVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/upgrade-versions][%d] ListUpgradeVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ListUpgradeVersionsDefault) String() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/upgrade-versions][%d] ListUpgradeVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ListUpgradeVersionsDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ListUpgradeVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
