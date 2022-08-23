// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-11-25/models"
)

// RecreateFromSnapshotReader is a Reader for the RecreateFromSnapshot structure.
type RecreateFromSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecreateFromSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRecreateFromSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRecreateFromSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRecreateFromSnapshotOK creates a RecreateFromSnapshotOK with default headers values
func NewRecreateFromSnapshotOK() *RecreateFromSnapshotOK {
	return &RecreateFromSnapshotOK{}
}

/*RecreateFromSnapshotOK handles this case with default header values.

A successful response.
*/
type RecreateFromSnapshotOK struct {
	Payload *models.HashicorpCloudVault20201125RecreateFromSnapshotResponse
}

func (o *RecreateFromSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/recreatefromsnapshot][%d] recreateFromSnapshotOK  %+v", 200, o.Payload)
}

func (o *RecreateFromSnapshotOK) GetPayload() *models.HashicorpCloudVault20201125RecreateFromSnapshotResponse {
	return o.Payload
}

func (o *RecreateFromSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125RecreateFromSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecreateFromSnapshotDefault creates a RecreateFromSnapshotDefault with default headers values
func NewRecreateFromSnapshotDefault(code int) *RecreateFromSnapshotDefault {
	return &RecreateFromSnapshotDefault{
		_statusCode: code,
	}
}

/*RecreateFromSnapshotDefault handles this case with default header values.

An unexpected error response.
*/
type RecreateFromSnapshotDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the recreate from snapshot default response
func (o *RecreateFromSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *RecreateFromSnapshotDefault) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/recreatefromsnapshot][%d] RecreateFromSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *RecreateFromSnapshotDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *RecreateFromSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}