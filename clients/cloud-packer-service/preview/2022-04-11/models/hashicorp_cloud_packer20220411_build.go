// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudPacker20220411Build hashicorp cloud packer 20220411 build
//
// swagger:model hashicorp.cloud.packer_20220411.Build
type HashicorpCloudPacker20220411Build struct {

	// The cloud provider that this build produced artifacts for.
	// For example, AWS, GCP, or Azure.
	CloudProvider string `json:"cloud_provider,omitempty"`

	// Internal Packer name for the builder or post-processor component used to
	// build this. For example, "amazon-ebs" or "azure-arm".
	ComponentType string `json:"component_type,omitempty"`

	// When the build was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Universally Unique Lexicographically Sortable Identifier (ULID) of the build.
	ID string `json:"id,omitempty"`

	// List of images associated with this build.
	Images []*HashicorpCloudPacker20220411Image `json:"images"`

	// Unique identifier of the iteration associated with this build; this was
	// created and set by the HCP Packer registry when the iteration was created.
	IterationID string `json:"iteration_id,omitempty"`

	// A key:value map for custom, user-settable metadata about your build.
	Labels map[string]string `json:"labels,omitempty"`

	// The UUID specific to this call to Packer build. If you use the manifest
	// post-processor, this UUID will match the UUID present there.
	PackerRunUUID string `json:"packer_run_uuid,omitempty"`

	// The parent image used as source for this build.
	Parent *HashicorpCloudPacker20220411ParentBuildStatus `json:"parent,omitempty"`

	// Unique identifier of the HCP Packer registry build used as the source
	// for this build. Used for tracking dependencies for build pipelines.
	SourceBuildUlid string `json:"source_build_ulid,omitempty"`

	// The ID or URL of the remote cloud source image. Used for tracking image
	// dependencies for build pipelines.
	SourceImageID string `json:"source_image_id,omitempty"`

	// Status of the build. The status can be RUNNING, DONE, CANCELLED, FAILED,
	// or UNSET.
	Status HashicorpCloudPacker20220411BuildStatus `json:"status,omitempty"`

	// When the build was most recently updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 build
func (m *HashicorpCloudPacker20220411Build) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *HashicorpCloudPacker20220411Build) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPacker20220411Build) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudPacker20220411Build) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPacker20220411Build) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *HashicorpCloudPacker20220411Build) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411Build) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411Build) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411Build
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
