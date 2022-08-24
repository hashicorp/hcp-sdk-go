// Code generated by go-swagger; DO NOT EDIT.

package registry_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/preview/2022-09-30/models"
)

// ActivateRegistryReader is a Reader for the ActivateRegistry structure.
type ActivateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActivateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActivateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewActivateRegistryOK creates a ActivateRegistryOK with default headers values
func NewActivateRegistryOK() *ActivateRegistryOK {
	return &ActivateRegistryOK{}
}

/* ActivateRegistryOK describes a response with status code 200, with default header values.

A successful response.
*/
type ActivateRegistryOK struct {
	Payload models.HashicorpCloudVagrantActivateRegistryResponse
}

func (o *ActivateRegistryOK) Error() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry/{registry}/activate][%d] activateRegistryOK  %+v", 200, o.Payload)
}
func (o *ActivateRegistryOK) GetPayload() models.HashicorpCloudVagrantActivateRegistryResponse {
	return o.Payload
}

func (o *ActivateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
