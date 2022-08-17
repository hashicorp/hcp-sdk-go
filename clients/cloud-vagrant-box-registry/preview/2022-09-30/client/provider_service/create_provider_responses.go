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

// CreateProviderReader is a Reader for the CreateProvider structure.
type CreateProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProviderOK creates a CreateProviderOK with default headers values
func NewCreateProviderOK() *CreateProviderOK {
	return &CreateProviderOK{}
}

/*CreateProviderOK handles this case with default header values.

A successful response.
*/
type CreateProviderOK struct {
	Payload *models.HashicorpCloudVagrant20220930CreateProviderResponse
}

func (o *CreateProviderOK) Error() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers][%d] createProviderOK  %+v", 200, o.Payload)
}

func (o *CreateProviderOK) GetPayload() *models.HashicorpCloudVagrant20220930CreateProviderResponse {
	return o.Payload
}

func (o *CreateProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930CreateProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
