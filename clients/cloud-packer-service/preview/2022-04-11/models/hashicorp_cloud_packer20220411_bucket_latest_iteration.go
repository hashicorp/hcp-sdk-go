// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudPacker20220411BucketLatestIteration A simplified Iteration used in Bucket to represent a bucket's latest iteration.
// This iteration representation does not contain ancestry information to avoid repetition.
//
// swagger:model hashicorp.cloud.packer_20220411.BucketLatestIteration
type HashicorpCloudPacker20220411BucketLatestIteration struct {

	// The name of the person who created this iteration.
	AuthorID string `json:"author_id,omitempty"`

	// Human-readable name for the bucket that this iteration is associated with.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// A list of all the builds associated with this iteration.
	Builds []*HashicorpCloudPacker20220411LatestIterationBuild `json:"builds"`

	// If true, all builds associated with this iteration have successfully
	// completed and uploaded metadata to the registry. When "complete" is true,
	// This iteration is considered ready to use, and can have channels assigned
	// to it.
	Complete bool `json:"complete,omitempty"`

	// When the iteration was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Fingerprint of the iteration. The fingerprint is set by Packer when you
	// call `packer build`. It will most often correspond to a git commit sha,
	// but can be manually overridden by setting the environment variable
	// `HCP_PACKER_BUILD_FINGERPRINT`.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Universally Unique Lexicographically Sortable Identifier (ULID) of the iteration.
	ID string `json:"id,omitempty"`

	// The human-readable version number assigned to this iteration. This
	// field will only be set if the iteration is "complete".
	IncrementalVersion int32 `json:"incremental_version,omitempty"`

	// The unique identifier of the iteration that was used as a source
	// for this iteration, if this iteration was built on a base layer.
	// Deprecated: refer to 'parent' for source image information.
	IterationAncestorID string `json:"iteration_ancestor_id,omitempty"`

	// Who revoked this iteration. For human authors (e.g. HCP Portal) this will be an email address.
	// For machine authors using service principals, this is the customer-chosen name for this service principal.
	RevocationAuthor string `json:"revocation_author,omitempty"`

	// The ancestor iteration from whom this iteration inherited the revocation state.
	RevocationInheritedFrom *HashicorpCloudPacker20220411RevokedAncestor `json:"revocation_inherited_from,omitempty"`

	// A short explanation of why this iteration was revoked.
	RevocationMessage string `json:"revocation_message,omitempty"`

	// Revocation type is 'manual' when self revoked or 'inherited' when inherited from a revoked ancestor.
	RevocationType *HashicorpCloudPacker20220411IterationRevocationType `json:"revocation_type,omitempty"`

	// Timestamp from when the iteration is revoked an no longer trusted to be secure.
	// Format: date-time
	RevokeAt strfmt.DateTime `json:"revoke_at,omitempty"`

	// When the iteration was last updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 bucket latest iteration
func (m *HashicorpCloudPacker20220411BucketLatestIteration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuilds(formats); err != nil {
		res = append(res, err)
	}

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

func (m *HashicorpCloudPacker20220411BucketLatestIteration) validateBuilds(formats strfmt.Registry) error {
	if swag.IsZero(m.Builds) { // not required
		return nil
	}

	for i := 0; i < len(m.Builds); i++ {
		if swag.IsZero(m.Builds[i]) { // not required
			continue
		}

		if m.Builds[i] != nil {
			if err := m.Builds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("builds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("builds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudPacker20220411BucketLatestIteration) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPacker20220411BucketLatestIteration) validateRevocationInheritedFrom(formats strfmt.Registry) error {
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

func (m *HashicorpCloudPacker20220411BucketLatestIteration) validateRevocationType(formats strfmt.Registry) error {
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

func (m *HashicorpCloudPacker20220411BucketLatestIteration) validateRevokeAt(formats strfmt.Registry) error {
	if swag.IsZero(m.RevokeAt) { // not required
		return nil
	}

	if err := validate.FormatOf("revoke_at", "body", "date-time", m.RevokeAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPacker20220411BucketLatestIteration) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer 20220411 bucket latest iteration based on the context it is used
func (m *HashicorpCloudPacker20220411BucketLatestIteration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuilds(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *HashicorpCloudPacker20220411BucketLatestIteration) contextValidateBuilds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Builds); i++ {

		if m.Builds[i] != nil {
			if err := m.Builds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("builds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("builds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudPacker20220411BucketLatestIteration) contextValidateRevocationInheritedFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.RevocationInheritedFrom != nil {
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

func (m *HashicorpCloudPacker20220411BucketLatestIteration) contextValidateRevocationType(ctx context.Context, formats strfmt.Registry) error {

	if m.RevocationType != nil {
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
func (m *HashicorpCloudPacker20220411BucketLatestIteration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411BucketLatestIteration) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411BucketLatestIteration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
