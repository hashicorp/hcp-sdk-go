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

// Secrets20231128SyncIntegration secrets 20231128 sync integration
//
// swagger:model secrets_20231128SyncIntegration
type Secrets20231128SyncIntegration struct {

	// aws sm connection details
	AwsSmConnectionDetails *Secrets20231128AwsSmConnectionDetailsResponse `json:"aws_sm_connection_details,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by
	CreatedBy *Secrets20231128Principal `json:"created_by,omitempty"`

	// gh repo connection details
	GhRepoConnectionDetails *Secrets20231128GhRepoConnectionDetailsResponse `json:"gh_repo_connection_details,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization id
	OrganizationID string `json:"organization_id,omitempty"`

	// project id
	ProjectID string `json:"project_id,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy *Secrets20231128Principal `json:"updated_by,omitempty"`

	// vercel project connection details
	VercelProjectConnectionDetails *Secrets20231128VercelProjectConnectionDetailsResponse `json:"vercel_project_connection_details,omitempty"`
}

// Validate validates this secrets 20231128 sync integration
func (m *Secrets20231128SyncIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsSmConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGhRepoConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVercelProjectConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128SyncIntegration) validateAwsSmConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsSmConnectionDetails) { // not required
		return nil
	}

	if m.AwsSmConnectionDetails != nil {
		if err := m.AwsSmConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_sm_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_sm_connection_details")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128SyncIntegration) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128SyncIntegration) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128SyncIntegration) validateGhRepoConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.GhRepoConnectionDetails) { // not required
		return nil
	}

	if m.GhRepoConnectionDetails != nil {
		if err := m.GhRepoConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gh_repo_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gh_repo_connection_details")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128SyncIntegration) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128SyncIntegration) validateUpdatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedBy) { // not required
		return nil
	}

	if m.UpdatedBy != nil {
		if err := m.UpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128SyncIntegration) validateVercelProjectConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.VercelProjectConnectionDetails) { // not required
		return nil
	}

	if m.VercelProjectConnectionDetails != nil {
		if err := m.VercelProjectConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vercel_project_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vercel_project_connection_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secrets 20231128 sync integration based on the context it is used
func (m *Secrets20231128SyncIntegration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsSmConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGhRepoConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVercelProjectConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128SyncIntegration) contextValidateAwsSmConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsSmConnectionDetails != nil {

		if swag.IsZero(m.AwsSmConnectionDetails) { // not required
			return nil
		}

		if err := m.AwsSmConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_sm_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_sm_connection_details")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128SyncIntegration) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128SyncIntegration) contextValidateGhRepoConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.GhRepoConnectionDetails != nil {

		if swag.IsZero(m.GhRepoConnectionDetails) { // not required
			return nil
		}

		if err := m.GhRepoConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gh_repo_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gh_repo_connection_details")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128SyncIntegration) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdatedBy != nil {

		if swag.IsZero(m.UpdatedBy) { // not required
			return nil
		}

		if err := m.UpdatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128SyncIntegration) contextValidateVercelProjectConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.VercelProjectConnectionDetails != nil {

		if swag.IsZero(m.VercelProjectConnectionDetails) { // not required
			return nil
		}

		if err := m.VercelProjectConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vercel_project_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vercel_project_connection_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128SyncIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128SyncIntegration) UnmarshalBinary(b []byte) error {
	var res Secrets20231128SyncIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}