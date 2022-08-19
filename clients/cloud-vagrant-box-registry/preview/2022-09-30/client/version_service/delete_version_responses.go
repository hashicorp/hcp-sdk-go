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

// DeleteVersionReader is a Reader for the DeleteVersion structure.
type DeleteVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVersionOK creates a DeleteVersionOK with default headers values
func NewDeleteVersionOK() *DeleteVersionOK {
	return &DeleteVersionOK{}
}

/*DeleteVersionOK handles this case with default header values.

A successful response.
*/
type DeleteVersionOK struct {
	Payload models.HashicorpCloudVagrant20220930DeleteVersionResponse
}

func (o *DeleteVersionOK) Error() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}][%d] deleteVersionOK  %+v", 200, o.Payload)
}

func (o *DeleteVersionOK) GetPayload() models.HashicorpCloudVagrant20220930DeleteVersionResponse {
	return o.Payload
}

func (o *DeleteVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
