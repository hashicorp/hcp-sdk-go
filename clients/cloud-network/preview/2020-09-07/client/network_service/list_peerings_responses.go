// Code generated by go-swagger; DO NOT EDIT.

package network_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/cloud-sdk-go/clients/cloud-network/preview/2020-09-07/models"
)

// ListPeeringsReader is a Reader for the ListPeerings structure.
type ListPeeringsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPeeringsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPeeringsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListPeeringsOK creates a ListPeeringsOK with default headers values
func NewListPeeringsOK() *ListPeeringsOK {
	return &ListPeeringsOK{}
}

/*ListPeeringsOK handles this case with default header values.

A successful response.
*/
type ListPeeringsOK struct {
	Payload *models.HashicorpCloudNetwork20200907ListPeeringsResponse
}

func (o *ListPeeringsOK) Error() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{hvn_id}/peerings][%d] listPeeringsOK  %+v", 200, o.Payload)
}

func (o *ListPeeringsOK) GetPayload() *models.HashicorpCloudNetwork20200907ListPeeringsResponse {
	return o.Payload
}

func (o *ListPeeringsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907ListPeeringsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
