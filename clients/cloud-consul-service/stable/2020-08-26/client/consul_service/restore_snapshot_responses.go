// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/stable/2020-08-26/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// RestoreSnapshotReader is a Reader for the RestoreSnapshot structure.
type RestoreSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestoreSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestoreSnapshotOK creates a RestoreSnapshotOK with default headers values
func NewRestoreSnapshotOK() *RestoreSnapshotOK {
	return &RestoreSnapshotOK{}
}

/* RestoreSnapshotOK describes a response with status code 200, with default header values.

A successful response.
*/
type RestoreSnapshotOK struct {
	Payload *models.HashicorpCloudConsul20200826RestoreSnapshotResponse
}

func (o *RestoreSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /consul/2020-08-26/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/restore][%d] restoreSnapshotOK  %+v", 200, o.Payload)
}
func (o *RestoreSnapshotOK) GetPayload() *models.HashicorpCloudConsul20200826RestoreSnapshotResponse {
	return o.Payload
}

func (o *RestoreSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20200826RestoreSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreSnapshotDefault creates a RestoreSnapshotDefault with default headers values
func NewRestoreSnapshotDefault(code int) *RestoreSnapshotDefault {
	return &RestoreSnapshotDefault{
		_statusCode: code,
	}
}

/* RestoreSnapshotDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RestoreSnapshotDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the restore snapshot default response
func (o *RestoreSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *RestoreSnapshotDefault) Error() string {
	return fmt.Sprintf("[POST /consul/2020-08-26/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/restore][%d] RestoreSnapshot default  %+v", o._statusCode, o.Payload)
}
func (o *RestoreSnapshotDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *RestoreSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
