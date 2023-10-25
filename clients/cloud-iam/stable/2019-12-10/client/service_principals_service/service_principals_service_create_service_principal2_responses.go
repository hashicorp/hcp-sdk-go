// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ServicePrincipalsServiceCreateServicePrincipal2Reader is a Reader for the ServicePrincipalsServiceCreateServicePrincipal2 structure.
type ServicePrincipalsServiceCreateServicePrincipal2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceCreateServicePrincipal2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceCreateServicePrincipal2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceCreateServicePrincipal2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceCreateServicePrincipal2OK creates a ServicePrincipalsServiceCreateServicePrincipal2OK with default headers values
func NewServicePrincipalsServiceCreateServicePrincipal2OK() *ServicePrincipalsServiceCreateServicePrincipal2OK {
	return &ServicePrincipalsServiceCreateServicePrincipal2OK{}
}

/*
ServicePrincipalsServiceCreateServicePrincipal2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceCreateServicePrincipal2OK struct {
	Payload *models.HashicorpCloudIamCreateServicePrincipalResponse
}

// IsSuccess returns true when this service principals service create service principal2 o k response has a 2xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipal2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service create service principal2 o k response has a 3xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipal2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service create service principal2 o k response has a 4xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipal2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service create service principal2 o k response has a 5xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipal2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service create service principal2 o k response a status code equal to that given
func (o *ServicePrincipalsServiceCreateServicePrincipal2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service principals service create service principal2 o k response
func (o *ServicePrincipalsServiceCreateServicePrincipal2OK) Code() int {
	return 200
}

func (o *ServicePrincipalsServiceCreateServicePrincipal2OK) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/iam/{parent_resource_name_1}/service-principals][%d] servicePrincipalsServiceCreateServicePrincipal2OK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceCreateServicePrincipal2OK) String() string {
	return fmt.Sprintf("[POST /2019-12-10/iam/{parent_resource_name_1}/service-principals][%d] servicePrincipalsServiceCreateServicePrincipal2OK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceCreateServicePrincipal2OK) GetPayload() *models.HashicorpCloudIamCreateServicePrincipalResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceCreateServicePrincipal2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamCreateServicePrincipalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceCreateServicePrincipal2Default creates a ServicePrincipalsServiceCreateServicePrincipal2Default with default headers values
func NewServicePrincipalsServiceCreateServicePrincipal2Default(code int) *ServicePrincipalsServiceCreateServicePrincipal2Default {
	return &ServicePrincipalsServiceCreateServicePrincipal2Default{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceCreateServicePrincipal2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceCreateServicePrincipal2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this service principals service create service principal2 default response has a 2xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipal2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service create service principal2 default response has a 3xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipal2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service create service principal2 default response has a 4xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipal2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service create service principal2 default response has a 5xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipal2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service create service principal2 default response a status code equal to that given
func (o *ServicePrincipalsServiceCreateServicePrincipal2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the service principals service create service principal2 default response
func (o *ServicePrincipalsServiceCreateServicePrincipal2Default) Code() int {
	return o._statusCode
}

func (o *ServicePrincipalsServiceCreateServicePrincipal2Default) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/iam/{parent_resource_name_1}/service-principals][%d] ServicePrincipalsService_CreateServicePrincipal2 default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceCreateServicePrincipal2Default) String() string {
	return fmt.Sprintf("[POST /2019-12-10/iam/{parent_resource_name_1}/service-principals][%d] ServicePrincipalsService_CreateServicePrincipal2 default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceCreateServicePrincipal2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceCreateServicePrincipal2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ServicePrincipalsServiceCreateServicePrincipal2Body CreateServicePrincipalRequest is the request message used when creating a
// service principal.
swagger:model ServicePrincipalsServiceCreateServicePrincipal2Body
*/
type ServicePrincipalsServiceCreateServicePrincipal2Body struct {

	// name is the customer-chosen name for this service principal.
	Name string `json:"name,omitempty"`
}

// Validate validates this service principals service create service principal2 body
func (o *ServicePrincipalsServiceCreateServicePrincipal2Body) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service principals service create service principal2 body based on context it is used
func (o *ServicePrincipalsServiceCreateServicePrincipal2Body) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ServicePrincipalsServiceCreateServicePrincipal2Body) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServicePrincipalsServiceCreateServicePrincipal2Body) UnmarshalBinary(b []byte) error {
	var res ServicePrincipalsServiceCreateServicePrincipal2Body
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}