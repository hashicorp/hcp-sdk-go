// Code generated by go-swagger; DO NOT EDIT.

package box_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/preview/2022-09-30/models"
)

// UpdateBoxReader is a Reader for the UpdateBox structure.
type UpdateBoxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBoxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBoxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateBoxOK creates a UpdateBoxOK with default headers values
func NewUpdateBoxOK() *UpdateBoxOK {
	return &UpdateBoxOK{}
}

/*UpdateBoxOK handles this case with default header values.

A successful response.
*/
type UpdateBoxOK struct {
	Payload *models.HashicorpCloudVagrant20220930UpdateBoxResponse
}

func (o *UpdateBoxOK) Error() string {
	return fmt.Sprintf("[PATCH /vagrant/2022-09-30/registry/{registry}/boxes/{box}][%d] updateBoxOK  %+v", 200, o.Payload)
}

func (o *UpdateBoxOK) GetPayload() *models.HashicorpCloudVagrant20220930UpdateBoxResponse {
	return o.Payload
}

func (o *UpdateBoxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930UpdateBoxResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
