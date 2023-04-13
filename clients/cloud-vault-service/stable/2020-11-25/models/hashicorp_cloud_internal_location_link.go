// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudInternalLocationLink Link is used to uniquely reference any resource within HashiCorp Cloud.
// This can be conceptually considered a "foreign key".
//
// swagger:model hashicorp.cloud.internal.location.Link
type HashicorpCloudInternalLocationLink struct {

	// description is a human-friendly description for this link. This is
	// used primarily for informational purposes such as error messages.
	Description string `json:"description,omitempty"`

	// id is the ID of the resource. This may be a slug ID or a UUID depending on
	// the resource, and as such has no guarantee that it is globally unique. If a
	// globally unique identifier for the resource is required, refer to
	// internal_id.
	ID string `json:"id,omitempty"`

	// internal_id is a globally unique identifier for the resource. In the case
	// that a resource has a user specifiable identifier, the internal_id will
	// differ.
	InternalID string `json:"internal_id,omitempty"`

	// location is the location where this resource is.
	Location *HashicorpCloudInternalLocationLocation `json:"location,omitempty"`

	// type is the unique type of the resource. Each service publishes a
	// unique set of types. The type value is recommended to be formatted
	// in "<org>.<type>" such as "hashicorp.hvn". This is to prevent conflicts
	// in the future, but any string value will work.
	Type string `json:"type,omitempty"`

	// uuid is being deprecated in favor of the id field.
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this hashicorp cloud internal location link
func (m *HashicorpCloudInternalLocationLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudInternalLocationLink) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud internal location link based on the context it is used
func (m *HashicorpCloudInternalLocationLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudInternalLocationLink) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudInternalLocationLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudInternalLocationLink) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudInternalLocationLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
