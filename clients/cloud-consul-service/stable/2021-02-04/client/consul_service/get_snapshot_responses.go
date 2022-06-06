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

// GetSnapshotReader is a Reader for the GetSnapshot structure.
type GetSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSnapshotOK creates a GetSnapshotOK with default headers values
func NewGetSnapshotOK() *GetSnapshotOK {
	return &GetSnapshotOK{}
}

/*GetSnapshotOK handles this case with default header values.

A successful response.
*/
type GetSnapshotOK struct {
	Payload *models.HashicorpCloudConsul20210204GetSnapshotResponse
}

func (o *GetSnapshotOK) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/snapshots/{snapshot_id}][%d] getSnapshotOK  %+v", 200, o.Payload)
}

func (o *GetSnapshotOK) GetPayload() *models.HashicorpCloudConsul20210204GetSnapshotResponse {
	return o.Payload
}

func (o *GetSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20210204GetSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnapshotDefault creates a GetSnapshotDefault with default headers values
func NewGetSnapshotDefault(code int) *GetSnapshotDefault {
	return &GetSnapshotDefault{
		_statusCode: code,
	}
}

/*GetSnapshotDefault handles this case with default header values.

An unexpected error response.
*/
type GetSnapshotDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get snapshot default response
func (o *GetSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *GetSnapshotDefault) Error() string {
	return fmt.Sprintf("[GET /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/snapshots/{snapshot_id}][%d] GetSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *GetSnapshotDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
