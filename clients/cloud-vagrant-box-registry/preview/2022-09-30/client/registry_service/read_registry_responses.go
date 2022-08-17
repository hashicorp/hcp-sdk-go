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

// ReadRegistryReader is a Reader for the ReadRegistry structure.
type ReadRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadRegistryOK creates a ReadRegistryOK with default headers values
func NewReadRegistryOK() *ReadRegistryOK {
	return &ReadRegistryOK{}
}

/*ReadRegistryOK handles this case with default header values.

A successful response.
*/
type ReadRegistryOK struct {
	Payload *models.HashicorpCloudVagrant20220930ReadRegistryResponse
}

func (o *ReadRegistryOK) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}][%d] readRegistryOK  %+v", 200, o.Payload)
}

func (o *ReadRegistryOK) GetPayload() *models.HashicorpCloudVagrant20220930ReadRegistryResponse {
	return o.Payload
}

func (o *ReadRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930ReadRegistryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
