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
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/preview/2020-04-20/models"
)

// CreateSnapshotReader is a Reader for the CreateSnapshot structure.
type CreateSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSnapshotOK creates a CreateSnapshotOK with default headers values
func NewCreateSnapshotOK() *CreateSnapshotOK {
	return &CreateSnapshotOK{}
}

/*CreateSnapshotOK handles this case with default header values.

A successful response.
*/
type CreateSnapshotOK struct {
	Payload *models.HashicorpCloudVault20200420CreateSnapshotResponse
}

func (o *CreateSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /vault/2020-04-20/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots][%d] createSnapshotOK  %+v", 200, o.Payload)
}

func (o *CreateSnapshotOK) GetPayload() *models.HashicorpCloudVault20200420CreateSnapshotResponse {
	return o.Payload
}

func (o *CreateSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20200420CreateSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotDefault creates a CreateSnapshotDefault with default headers values
func NewCreateSnapshotDefault(code int) *CreateSnapshotDefault {
	return &CreateSnapshotDefault{
		_statusCode: code,
	}
}

/*CreateSnapshotDefault handles this case with default header values.

An unexpected error response.
*/
type CreateSnapshotDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the create snapshot default response
func (o *CreateSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *CreateSnapshotDefault) Error() string {
	return fmt.Sprintf("[POST /vault/2020-04-20/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots][%d] CreateSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSnapshotDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CreateSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
