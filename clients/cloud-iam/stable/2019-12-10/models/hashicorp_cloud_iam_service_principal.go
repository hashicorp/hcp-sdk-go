// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudIamServicePrincipal ServicePrincipal is the HCP identity for a machine user.
//
// swagger:model hashicorp.cloud.iam.ServicePrincipal
type HashicorpCloudIamServicePrincipal struct {

	// created_at is when the service principal was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// id is an opaque, globally unique public identifier for the service principal.
	ID string `json:"id,omitempty"`

	// managed_by is the optional resource_id of the resource which manages this service principal.
	// The presence of this field denotes that the service principal's lifecycle is not managed by
	// a user.
	ManagedBy *string `json:"managed_by,omitempty"`

	// name is the name for this service principal.
	Name string `json:"name,omitempty"`

	// organization_id is the UUID identifier of the HCP organization that this
	// service principal belongs to.
	OrganizationID string `json:"organization_id,omitempty"`

	// project_id is the UUID identifier of the HCP project that this
	// project service principal belongs to.
	ProjectID string `json:"project_id,omitempty"`

	// resource_name is the resource name for this service principal.
	ResourceName string `json:"resource_name,omitempty"`
}

// Validate validates this hashicorp cloud iam service principal
func (m *HashicorpCloudIamServicePrincipal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamServicePrincipal) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hashicorp cloud iam service principal based on context it is used
func (m *HashicorpCloudIamServicePrincipal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamServicePrincipal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamServicePrincipal) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamServicePrincipal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
