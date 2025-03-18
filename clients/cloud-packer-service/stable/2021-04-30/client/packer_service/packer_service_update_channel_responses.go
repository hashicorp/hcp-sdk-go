// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceUpdateChannelReader is a Reader for the PackerServiceUpdateChannel structure.
type PackerServiceUpdateChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceUpdateChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceUpdateChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceUpdateChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceUpdateChannelOK creates a PackerServiceUpdateChannelOK with default headers values
func NewPackerServiceUpdateChannelOK() *PackerServiceUpdateChannelOK {
	return &PackerServiceUpdateChannelOK{}
}

/*
PackerServiceUpdateChannelOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceUpdateChannelOK struct {
	Payload *models.HashicorpCloudPackerUpdateChannelResponse
}

// IsSuccess returns true when this packer service update channel o k response has a 2xx status code
func (o *PackerServiceUpdateChannelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service update channel o k response has a 3xx status code
func (o *PackerServiceUpdateChannelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service update channel o k response has a 4xx status code
func (o *PackerServiceUpdateChannelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service update channel o k response has a 5xx status code
func (o *PackerServiceUpdateChannelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service update channel o k response a status code equal to that given
func (o *PackerServiceUpdateChannelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service update channel o k response
func (o *PackerServiceUpdateChannelOK) Code() int {
	return 200
}

func (o *PackerServiceUpdateChannelOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] packerServiceUpdateChannelOK %s", 200, payload)
}

func (o *PackerServiceUpdateChannelOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] packerServiceUpdateChannelOK %s", 200, payload)
}

func (o *PackerServiceUpdateChannelOK) GetPayload() *models.HashicorpCloudPackerUpdateChannelResponse {
	return o.Payload
}

func (o *PackerServiceUpdateChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerUpdateChannelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceUpdateChannelDefault creates a PackerServiceUpdateChannelDefault with default headers values
func NewPackerServiceUpdateChannelDefault(code int) *PackerServiceUpdateChannelDefault {
	return &PackerServiceUpdateChannelDefault{
		_statusCode: code,
	}
}

/*
PackerServiceUpdateChannelDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceUpdateChannelDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service update channel default response has a 2xx status code
func (o *PackerServiceUpdateChannelDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service update channel default response has a 3xx status code
func (o *PackerServiceUpdateChannelDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service update channel default response has a 4xx status code
func (o *PackerServiceUpdateChannelDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service update channel default response has a 5xx status code
func (o *PackerServiceUpdateChannelDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service update channel default response a status code equal to that given
func (o *PackerServiceUpdateChannelDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service update channel default response
func (o *PackerServiceUpdateChannelDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceUpdateChannelDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] PackerService_UpdateChannel default %s", o._statusCode, payload)
}

func (o *PackerServiceUpdateChannelDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] PackerService_UpdateChannel default %s", o._statusCode, payload)
}

func (o *PackerServiceUpdateChannelDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceUpdateChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PackerServiceUpdateChannelBody packer service update channel body
swagger:model PackerServiceUpdateChannelBody
*/
type PackerServiceUpdateChannelBody struct {

	// Fingerprint of the iteration set by Packer when you call `packer build`.
	// Refer to the Packer documentation for more information on how this value is set.
	// The fingerprint can be used as an identifier for the iteration.
	Fingerprint string `json:"fingerprint,omitempty"`

	// The human-readable version number assigned to this iteration.
	IncrementalVersion int32 `json:"incremental_version,omitempty"`

	// ULID of the iteration. This was created and set by the
	// HCP Packer registry when the iteration was created.
	IterationID string `json:"iteration_id,omitempty"`

	// Location represents a target for an operation in HCP.
	Location interface{} `json:"location,omitempty"`

	// When set, the service will only update the channel with attributes listed in the mask.
	// For an empty mask list, all attributes will be updated according to their given value.
	//
	// One or more of the iteration identifiers can be listed when updating the channel assignment.
	// In this case, if paths for multiple iteration identifiers are present in the mask, the
	// values of the equivalent fields must belong to the same iteration.
	// NOTE: This is different from the behavior without a mask or with an empty mask, where the
	// first non-zero identifier (in the order iterationId, fingerprint, incrementalVersion) is the
	// only identifier used, and others are ignored.
	//
	// Examples of usage:
	// * "incrementalVersion,iterationId,fingerprint"
	// * "restriction"
	// * "restriction,iterationId"
	Mask string `json:"mask,omitempty"`

	// When set, will set the channel access in HCP Packer registry. The channel is unrestricted by default;
	Restriction *models.HashicorpCloudPackerUpdateChannelRequestRestriction `json:"restriction,omitempty"`
}

// Validate validates this packer service update channel body
func (o *PackerServiceUpdateChannelBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRestriction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceUpdateChannelBody) validateRestriction(formats strfmt.Registry) error {
	if swag.IsZero(o.Restriction) { // not required
		return nil
	}

	if o.Restriction != nil {
		if err := o.Restriction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "restriction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "restriction")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this packer service update channel body based on the context it is used
func (o *PackerServiceUpdateChannelBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRestriction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceUpdateChannelBody) contextValidateRestriction(ctx context.Context, formats strfmt.Registry) error {

	if o.Restriction != nil {

		if swag.IsZero(o.Restriction) { // not required
			return nil
		}

		if err := o.Restriction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "restriction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "restriction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PackerServiceUpdateChannelBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PackerServiceUpdateChannelBody) UnmarshalBinary(b []byte) error {
	var res PackerServiceUpdateChannelBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
