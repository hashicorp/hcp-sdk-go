// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package buildservice

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2022-12-02/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// GetImageByBuildLabelsReader reads the GetImageByBuildLabels response.
type GetImageByBuildLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageByBuildLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImageByBuildLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetImageByBuildLabelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetImageByBuildLabelsOK creates a GetImageByBuildLabelsOK.
func NewGetImageByBuildLabelsOK() *GetImageByBuildLabelsOK {
	return &GetImageByBuildLabelsOK{}
}

// GetImageByBuildLabelsOK is the success response.
type GetImageByBuildLabelsOK struct {
	Payload *models.GetImageByBuildLabelsResponse
}

func (o *GetImageByBuildLabelsOK) GetPayload() *models.GetImageByBuildLabelsResponse {
	return o.Payload
}

func (o *GetImageByBuildLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.GetImageByBuildLabelsResponse)
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}

func (o *GetImageByBuildLabelsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /packer/2022-12-02/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_name}/channels/{channel_name}/_search/by_labels][%d] buildServiceGetImageByBuildLabelsOK %s", 200, payload)
}

// NewGetImageByBuildLabelsDefault creates a default error response.
func NewGetImageByBuildLabelsDefault(code int) *GetImageByBuildLabelsDefault {
	return &GetImageByBuildLabelsDefault{_statusCode: code}
}

// GetImageByBuildLabelsDefault is the default error response.
type GetImageByBuildLabelsDefault struct {
	_statusCode int
	Payload     *cloud.GoogleRPCStatus
}

func (o *GetImageByBuildLabelsDefault) Code() int {
	return o._statusCode
}

func (o *GetImageByBuildLabelsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *GetImageByBuildLabelsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /packer/2022-12-02/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_name}/channels/{channel_name}/_search/by_labels][%d] BuildService_GetImageByBuildLabels default %s", o._statusCode, payload)
}

func (o *GetImageByBuildLabelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(cloud.GoogleRPCStatus)
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
