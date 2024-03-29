// Code generated by go-swagger; DO NOT EDIT.

package provider_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/preview/2022-09-30/models"
)

// GetProviderReader is a Reader for the GetProvider structure.
type GetProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProviderOK creates a GetProviderOK with default headers values
func NewGetProviderOK() *GetProviderOK {
	return &GetProviderOK{}
}

/*
GetProviderOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetProviderOK struct {
	Payload *models.HashicorpCloudVagrantGetProviderResponse
}

// IsSuccess returns true when this get provider o k response has a 2xx status code
func (o *GetProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get provider o k response has a 3xx status code
func (o *GetProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get provider o k response has a 4xx status code
func (o *GetProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get provider o k response has a 5xx status code
func (o *GetProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get provider o k response a status code equal to that given
func (o *GetProviderOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetProviderOK) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}][%d] getProviderOK  %+v", 200, o.Payload)
}

func (o *GetProviderOK) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}][%d] getProviderOK  %+v", 200, o.Payload)
}

func (o *GetProviderOK) GetPayload() *models.HashicorpCloudVagrantGetProviderResponse {
	return o.Payload
}

func (o *GetProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrantGetProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
