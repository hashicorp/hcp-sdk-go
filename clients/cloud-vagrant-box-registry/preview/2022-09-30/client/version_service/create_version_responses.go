// Code generated by go-swagger; DO NOT EDIT.

package version_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/preview/2022-09-30/models"
)

// CreateVersionReader is a Reader for the CreateVersion structure.
type CreateVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVersionOK creates a CreateVersionOK with default headers values
func NewCreateVersionOK() *CreateVersionOK {
	return &CreateVersionOK{}
}

/*CreateVersionOK handles this case with default header values.

A successful response.
*/
type CreateVersionOK struct {
	Payload *models.HashicorpCloudVagrant20220930CreateVersionResponse
}

func (o *CreateVersionOK) Error() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions][%d] createVersionOK  %+v", 200, o.Payload)
}

func (o *CreateVersionOK) GetPayload() *models.HashicorpCloudVagrant20220930CreateVersionResponse {
	return o.Payload
}

func (o *CreateVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930CreateVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
