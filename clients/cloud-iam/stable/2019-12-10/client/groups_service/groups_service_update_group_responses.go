// Code generated by go-swagger; DO NOT EDIT.

package groups_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// GroupsServiceUpdateGroupReader is a Reader for the GroupsServiceUpdateGroup structure.
type GroupsServiceUpdateGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsServiceUpdateGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsServiceUpdateGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupsServiceUpdateGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupsServiceUpdateGroupOK creates a GroupsServiceUpdateGroupOK with default headers values
func NewGroupsServiceUpdateGroupOK() *GroupsServiceUpdateGroupOK {
	return &GroupsServiceUpdateGroupOK{}
}

/*
GroupsServiceUpdateGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type GroupsServiceUpdateGroupOK struct {
	Payload interface{}
}

// IsSuccess returns true when this groups service update group o k response has a 2xx status code
func (o *GroupsServiceUpdateGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this groups service update group o k response has a 3xx status code
func (o *GroupsServiceUpdateGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups service update group o k response has a 4xx status code
func (o *GroupsServiceUpdateGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this groups service update group o k response has a 5xx status code
func (o *GroupsServiceUpdateGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this groups service update group o k response a status code equal to that given
func (o *GroupsServiceUpdateGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the groups service update group o k response
func (o *GroupsServiceUpdateGroupOK) Code() int {
	return 200
}

func (o *GroupsServiceUpdateGroupOK) Error() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/{resource_name}][%d] groupsServiceUpdateGroupOK  %+v", 200, o.Payload)
}

func (o *GroupsServiceUpdateGroupOK) String() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/{resource_name}][%d] groupsServiceUpdateGroupOK  %+v", 200, o.Payload)
}

func (o *GroupsServiceUpdateGroupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GroupsServiceUpdateGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupsServiceUpdateGroupDefault creates a GroupsServiceUpdateGroupDefault with default headers values
func NewGroupsServiceUpdateGroupDefault(code int) *GroupsServiceUpdateGroupDefault {
	return &GroupsServiceUpdateGroupDefault{
		_statusCode: code,
	}
}

/*
GroupsServiceUpdateGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GroupsServiceUpdateGroupDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this groups service update group default response has a 2xx status code
func (o *GroupsServiceUpdateGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this groups service update group default response has a 3xx status code
func (o *GroupsServiceUpdateGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this groups service update group default response has a 4xx status code
func (o *GroupsServiceUpdateGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this groups service update group default response has a 5xx status code
func (o *GroupsServiceUpdateGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this groups service update group default response a status code equal to that given
func (o *GroupsServiceUpdateGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the groups service update group default response
func (o *GroupsServiceUpdateGroupDefault) Code() int {
	return o._statusCode
}

func (o *GroupsServiceUpdateGroupDefault) Error() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/{resource_name}][%d] GroupsService_UpdateGroup default  %+v", o._statusCode, o.Payload)
}

func (o *GroupsServiceUpdateGroupDefault) String() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/{resource_name}][%d] GroupsService_UpdateGroup default  %+v", o._statusCode, o.Payload)
}

func (o *GroupsServiceUpdateGroupDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *GroupsServiceUpdateGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GroupsServiceUpdateGroupBody groups service update group body
swagger:model GroupsServiceUpdateGroupBody
*/
type GroupsServiceUpdateGroupBody struct {

	// group is the group being updated.
	Group *models.HashicorpCloudIamGroup `json:"group,omitempty"`

	// update_mask is the list of group fields being updated.
	UpdateMask string `json:"update_mask,omitempty"`
}

// Validate validates this groups service update group body
func (o *GroupsServiceUpdateGroupBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GroupsServiceUpdateGroupBody) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(o.Group) { // not required
		return nil
	}

	if o.Group != nil {
		if err := o.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "group")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this groups service update group body based on the context it is used
func (o *GroupsServiceUpdateGroupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GroupsServiceUpdateGroupBody) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if o.Group != nil {

		if swag.IsZero(o.Group) { // not required
			return nil
		}

		if err := o.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GroupsServiceUpdateGroupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GroupsServiceUpdateGroupBody) UnmarshalBinary(b []byte) error {
	var res GroupsServiceUpdateGroupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}