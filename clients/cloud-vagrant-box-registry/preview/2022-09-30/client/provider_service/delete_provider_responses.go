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

// DeleteProviderReader is a Reader for the DeleteProvider structure.
type DeleteProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProviderOK creates a DeleteProviderOK with default headers values
func NewDeleteProviderOK() *DeleteProviderOK {
	return &DeleteProviderOK{}
}

/*DeleteProviderOK handles this case with default header values.

A successful response.
*/
type DeleteProviderOK struct {
	Payload models.HashicorpCloudVagrant20220930DeleteProviderResponse
}

func (o *DeleteProviderOK) Error() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}][%d] deleteProviderOK  %+v", 200, o.Payload)
}

func (o *DeleteProviderOK) GetPayload() models.HashicorpCloudVagrant20220930DeleteProviderResponse {
	return o.Payload
}

func (o *DeleteProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
