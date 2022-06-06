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

// ListVersionsReader is a Reader for the ListVersions structure.
type ListVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVersionsOK creates a ListVersionsOK with default headers values
func NewListVersionsOK() *ListVersionsOK {
	return &ListVersionsOK{}
}

/*ListVersionsOK handles this case with default header values.

A successful response.
*/
type ListVersionsOK struct {
	Payload *models.HashicorpCloudConsul20210204ListVersionsResponse
}

func (o *ListVersionsOK) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/versions][%d] listVersionsOK  %+v", 200, o.Payload)
}

func (o *ListVersionsOK) GetPayload() *models.HashicorpCloudConsul20210204ListVersionsResponse {
	return o.Payload
}

func (o *ListVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20210204ListVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVersionsDefault creates a ListVersionsDefault with default headers values
func NewListVersionsDefault(code int) *ListVersionsDefault {
	return &ListVersionsDefault{
		_statusCode: code,
	}
}

/*ListVersionsDefault handles this case with default header values.

An unexpected error response.
*/
type ListVersionsDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the list versions default response
func (o *ListVersionsDefault) Code() int {
	return o._statusCode
}

func (o *ListVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/versions][%d] ListVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ListVersionsDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ListVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
