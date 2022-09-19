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

// DeleteBoxReader is a Reader for the DeleteBox structure.
type DeleteBoxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBoxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteBoxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteBoxOK creates a DeleteBoxOK with default headers values
func NewDeleteBoxOK() *DeleteBoxOK {
	return &DeleteBoxOK{}
}

/*
DeleteBoxOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteBoxOK struct {
	Payload models.HashicorpCloudVagrantDeleteBoxResponse
}

// IsSuccess returns true when this delete box o k response has a 2xx status code
func (o *DeleteBoxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete box o k response has a 3xx status code
func (o *DeleteBoxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete box o k response has a 4xx status code
func (o *DeleteBoxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete box o k response has a 5xx status code
func (o *DeleteBoxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete box o k response a status code equal to that given
func (o *DeleteBoxOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteBoxOK) Error() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}][%d] deleteBoxOK  %+v", 200, o.Payload)
}

func (o *DeleteBoxOK) String() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}][%d] deleteBoxOK  %+v", 200, o.Payload)
}

func (o *DeleteBoxOK) GetPayload() models.HashicorpCloudVagrantDeleteBoxResponse {
	return o.Payload
}

func (o *DeleteBoxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
