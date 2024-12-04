// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-06-13/models"
)

// CreateAzureKvSyncIntegrationReader is a Reader for the CreateAzureKvSyncIntegration structure.
type CreateAzureKvSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAzureKvSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAzureKvSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateAzureKvSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAzureKvSyncIntegrationOK creates a CreateAzureKvSyncIntegrationOK with default headers values
func NewCreateAzureKvSyncIntegrationOK() *CreateAzureKvSyncIntegrationOK {
	return &CreateAzureKvSyncIntegrationOK{}
}

/*
CreateAzureKvSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateAzureKvSyncIntegrationOK struct {
	Payload *models.Secrets20230613CreateSyncIntegrationResponse
}

// IsSuccess returns true when this create azure kv sync integration o k response has a 2xx status code
func (o *CreateAzureKvSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create azure kv sync integration o k response has a 3xx status code
func (o *CreateAzureKvSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azure kv sync integration o k response has a 4xx status code
func (o *CreateAzureKvSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create azure kv sync integration o k response has a 5xx status code
func (o *CreateAzureKvSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create azure kv sync integration o k response a status code equal to that given
func (o *CreateAzureKvSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create azure kv sync integration o k response
func (o *CreateAzureKvSyncIntegrationOK) Code() int {
	return 200
}

func (o *CreateAzureKvSyncIntegrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/azure-kv][%d] createAzureKvSyncIntegrationOK %s", 200, payload)
}

func (o *CreateAzureKvSyncIntegrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/azure-kv][%d] createAzureKvSyncIntegrationOK %s", 200, payload)
}

func (o *CreateAzureKvSyncIntegrationOK) GetPayload() *models.Secrets20230613CreateSyncIntegrationResponse {
	return o.Payload
}

func (o *CreateAzureKvSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613CreateSyncIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureKvSyncIntegrationDefault creates a CreateAzureKvSyncIntegrationDefault with default headers values
func NewCreateAzureKvSyncIntegrationDefault(code int) *CreateAzureKvSyncIntegrationDefault {
	return &CreateAzureKvSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
CreateAzureKvSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateAzureKvSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this create azure kv sync integration default response has a 2xx status code
func (o *CreateAzureKvSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create azure kv sync integration default response has a 3xx status code
func (o *CreateAzureKvSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create azure kv sync integration default response has a 4xx status code
func (o *CreateAzureKvSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create azure kv sync integration default response has a 5xx status code
func (o *CreateAzureKvSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create azure kv sync integration default response a status code equal to that given
func (o *CreateAzureKvSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create azure kv sync integration default response
func (o *CreateAzureKvSyncIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *CreateAzureKvSyncIntegrationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/azure-kv][%d] CreateAzureKvSyncIntegration default %s", o._statusCode, payload)
}

func (o *CreateAzureKvSyncIntegrationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/azure-kv][%d] CreateAzureKvSyncIntegration default %s", o._statusCode, payload)
}

func (o *CreateAzureKvSyncIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateAzureKvSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateAzureKvSyncIntegrationBody create azure kv sync integration body
swagger:model CreateAzureKvSyncIntegrationBody
*/
type CreateAzureKvSyncIntegrationBody struct {

	// azure kv connection details
	AzureKvConnectionDetails *models.Secrets20230613AzureKvConnectionDetailsRequest `json:"azure_kv_connection_details,omitempty"`

	// location
	Location *CreateAzureKvSyncIntegrationParamsBodyLocation `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create azure kv sync integration body
func (o *CreateAzureKvSyncIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAzureKvConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAzureKvSyncIntegrationBody) validateAzureKvConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.AzureKvConnectionDetails) { // not required
		return nil
	}

	if o.AzureKvConnectionDetails != nil {
		if err := o.AzureKvConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "azure_kv_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "azure_kv_connection_details")
			}
			return err
		}
	}

	return nil
}

func (o *CreateAzureKvSyncIntegrationBody) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(o.Location) { // not required
		return nil
	}

	if o.Location != nil {
		if err := o.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create azure kv sync integration body based on the context it is used
func (o *CreateAzureKvSyncIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAzureKvConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAzureKvSyncIntegrationBody) contextValidateAzureKvConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if o.AzureKvConnectionDetails != nil {

		if swag.IsZero(o.AzureKvConnectionDetails) { // not required
			return nil
		}

		if err := o.AzureKvConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "azure_kv_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "azure_kv_connection_details")
			}
			return err
		}
	}

	return nil
}

func (o *CreateAzureKvSyncIntegrationBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if o.Location != nil {

		if swag.IsZero(o.Location) { // not required
			return nil
		}

		if err := o.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateAzureKvSyncIntegrationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAzureKvSyncIntegrationBody) UnmarshalBinary(b []byte) error {
	var res CreateAzureKvSyncIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateAzureKvSyncIntegrationParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model CreateAzureKvSyncIntegrationParamsBodyLocation
*/
type CreateAzureKvSyncIntegrationParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this create azure kv sync integration params body location
func (o *CreateAzureKvSyncIntegrationParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAzureKvSyncIntegrationParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(o.Region) { // not required
		return nil
	}

	if o.Region != nil {
		if err := o.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create azure kv sync integration params body location based on the context it is used
func (o *CreateAzureKvSyncIntegrationParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAzureKvSyncIntegrationParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if o.Region != nil {

		if swag.IsZero(o.Region) { // not required
			return nil
		}

		if err := o.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateAzureKvSyncIntegrationParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAzureKvSyncIntegrationParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res CreateAzureKvSyncIntegrationParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
