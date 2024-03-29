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

// ReadBoxReader is a Reader for the ReadBox structure.
type ReadBoxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadBoxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadBoxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadBoxOK creates a ReadBoxOK with default headers values
func NewReadBoxOK() *ReadBoxOK {
	return &ReadBoxOK{}
}

/*
ReadBoxOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadBoxOK struct {
	Payload *models.HashicorpCloudVagrantReadBoxResponse
}

// IsSuccess returns true when this read box o k response has a 2xx status code
func (o *ReadBoxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read box o k response has a 3xx status code
func (o *ReadBoxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read box o k response has a 4xx status code
func (o *ReadBoxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read box o k response has a 5xx status code
func (o *ReadBoxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read box o k response a status code equal to that given
func (o *ReadBoxOK) IsCode(code int) bool {
	return code == 200
}

func (o *ReadBoxOK) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}][%d] readBoxOK  %+v", 200, o.Payload)
}

func (o *ReadBoxOK) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}][%d] readBoxOK  %+v", 200, o.Payload)
}

func (o *ReadBoxOK) GetPayload() *models.HashicorpCloudVagrantReadBoxResponse {
	return o.Payload
}

func (o *ReadBoxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrantReadBoxResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
