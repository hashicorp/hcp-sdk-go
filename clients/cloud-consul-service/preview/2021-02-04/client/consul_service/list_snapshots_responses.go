// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/preview/2021-02-04/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ListSnapshotsReader is a Reader for the ListSnapshots structure.
type ListSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSnapshotsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSnapshotsOK creates a ListSnapshotsOK with default headers values
func NewListSnapshotsOK() *ListSnapshotsOK {
	return &ListSnapshotsOK{}
}

/* ListSnapshotsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListSnapshotsOK struct {
	Payload *models.HashicorpCloudConsul20210204ListSnapshotsResponse
}

func (o *ListSnapshotsOK) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots][%d] listSnapshotsOK  %+v", 200, o.Payload)
}
func (o *ListSnapshotsOK) GetPayload() *models.HashicorpCloudConsul20210204ListSnapshotsResponse {
	return o.Payload
}

func (o *ListSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20210204ListSnapshotsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSnapshotsDefault creates a ListSnapshotsDefault with default headers values
func NewListSnapshotsDefault(code int) *ListSnapshotsDefault {
	return &ListSnapshotsDefault{
		_statusCode: code,
	}
}

/* ListSnapshotsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListSnapshotsDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the list snapshots default response
func (o *ListSnapshotsDefault) Code() int {
	return o._statusCode
}

func (o *ListSnapshotsDefault) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots][%d] ListSnapshots default  %+v", o._statusCode, o.Payload)
}
func (o *ListSnapshotsDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ListSnapshotsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
