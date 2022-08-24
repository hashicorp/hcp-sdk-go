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

// HashicorpCloudPackerChannel hashicorp cloud packer channel
//
// swagger:model hashicorp.cloud.packer.Channel
type HashicorpCloudPackerChannel struct {

	// The user who last updated the channel.
	AuthorID string `json:"author_id,omitempty"`

	// Human-readable name for the bucket this channel is associated with.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// When the channel was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Universally Unique Lexicographically Sortable Identifier (ULID) of the channel.
	ID string `json:"id,omitempty"`

	// The iteration the channel is pointing to.
	Iteration *HashicorpCloudPackerIteration `json:"iteration,omitempty"`

	// A pointer to the iteration currently associated with this channel.
	// Deprecated: look at the Channel.iteration instead.
	Pointer *HashicorpCloudPackerChannelIterationPointer `json:"pointer,omitempty"`

	// Human-readable name for the channel.
	Slug string `json:"slug,omitempty"`

	// When the channel was last updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this hashicorp cloud packer channel
func (m *HashicorpCloudPackerChannel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIteration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePointer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerChannel) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPackerChannel) validateIteration(formats strfmt.Registry) error {
	if swag.IsZero(m.Iteration) { // not required
		return nil
	}

	if m.Iteration != nil {
		if err := m.Iteration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iteration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iteration")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerChannel) validatePointer(formats strfmt.Registry) error {
	if swag.IsZero(m.Pointer) { // not required
		return nil
	}

	if m.Pointer != nil {
		if err := m.Pointer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pointer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pointer")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerChannel) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer channel based on the context it is used
func (m *HashicorpCloudPackerChannel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIteration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePointer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerChannel) contextValidateIteration(ctx context.Context, formats strfmt.Registry) error {

	if m.Iteration != nil {
		if err := m.Iteration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iteration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iteration")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerChannel) contextValidatePointer(ctx context.Context, formats strfmt.Registry) error {

	if m.Pointer != nil {
		if err := m.Pointer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pointer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pointer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerChannel) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
