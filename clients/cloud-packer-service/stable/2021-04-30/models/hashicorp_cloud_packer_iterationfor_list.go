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

// HashicorpCloudPackerIterationforList The list endpoint does not return build information.
//
// swagger:model hashicorp.cloud.packer.IterationforList
type HashicorpCloudPackerIterationforList struct {

	// Who created the iteration.
	AuthorID string `json:"author_id,omitempty"`

	// Human-readable name for the bucket.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// Maps the build component type to its status enum, for displaying build
	// status in the iterations view.
	BuildStatuses map[string]string `json:"build_statuses,omitempty"`

	// If true, all builds associated with this iteration have successfully
	// completed and uploaded metadata to the registry. When "complete" is true,
	// This iteration is considered ready to use, and can have channels assigned
	// to it.
	Complete bool `json:"complete,omitempty"`

	// When the iteration was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Fingerprint of the iteration. The fingerprint is set by Packer when you
	// call `packer build`.
	Fingerprint string `json:"fingerprint,omitempty"`

	// If true, this iteration has children iterations. Knowing if an iteration has descendants can help
	// taking decisions such as persist revocation to all its descendants or not.
	HasDescendants bool `json:"has_descendants,omitempty"`

	// Universally Unique Lexicographically Sortable Identifier (ULID) of the iteration.
	ID string `json:"id,omitempty"`

	// The human-readable version number assigned to this iteration. This
	// field will only be set if the iteration is "complete".
	IncrementalVersion int32 `json:"incremental_version,omitempty"`

	// The unique identifier of the iteration that was used as a source
	// for this iteration, if this iteration was built on a base layer.
	// Deprecated: Deprecated: refer to build specific source_build_ulid.
	IterationAncestorID string `json:"iteration_ancestor_id,omitempty"`

	// Who revoked this iteration. For human authors (e.g. HCP Portal) this will be an email address.
	// For machine authors using service principals, this is the customer-chosen name for this service principal.
	RevocationAuthor string `json:"revocation_author,omitempty"`

	// The ancestor iteration from whom this iteration inherited the revocation state.
	RevocationInheritedFrom *HashicorpCloudPackerRevokedAncestor `json:"revocation_inherited_from,omitempty"`

	// A short explanation of why this iteration was revoked.
	RevocationMessage string `json:"revocation_message,omitempty"`

	// Revocation type is 'manual' when self revoked or 'inherited' when inherited from a revoked ancestor.
	RevocationType *HashicorpCloudPackerIterationRevocationType `json:"revocation_type,omitempty"`

	// Timestamp from when the iteration is revoked an no longer trusted to be secure.
	// Format: date-time
	RevokeAt strfmt.DateTime `json:"revoke_at,omitempty"`

	// When the iteration was most recently updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this hashicorp cloud packer iterationfor list
func (m *HashicorpCloudPackerIterationforList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevocationInheritedFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevocationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevokeAt(formats); err != nil {
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

func (m *HashicorpCloudPackerIterationforList) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPackerIterationforList) validateRevocationInheritedFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.RevocationInheritedFrom) { // not required
		return nil
	}

	if m.RevocationInheritedFrom != nil {
		if err := m.RevocationInheritedFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revocation_inherited_from")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revocation_inherited_from")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerIterationforList) validateRevocationType(formats strfmt.Registry) error {
	if swag.IsZero(m.RevocationType) { // not required
		return nil
	}

	if m.RevocationType != nil {
		if err := m.RevocationType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revocation_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revocation_type")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerIterationforList) validateRevokeAt(formats strfmt.Registry) error {
	if swag.IsZero(m.RevokeAt) { // not required
		return nil
	}

	if err := validate.FormatOf("revoke_at", "body", "date-time", m.RevokeAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPackerIterationforList) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer iterationfor list based on the context it is used
func (m *HashicorpCloudPackerIterationforList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRevocationInheritedFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevocationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerIterationforList) contextValidateRevocationInheritedFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.RevocationInheritedFrom != nil {

		if swag.IsZero(m.RevocationInheritedFrom) { // not required
			return nil
		}

		if err := m.RevocationInheritedFrom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revocation_inherited_from")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revocation_inherited_from")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerIterationforList) contextValidateRevocationType(ctx context.Context, formats strfmt.Registry) error {

	if m.RevocationType != nil {

		if swag.IsZero(m.RevocationType) { // not required
			return nil
		}

		if err := m.RevocationType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revocation_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revocation_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerIterationforList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerIterationforList) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerIterationforList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
