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

// Secrets20231128Integration secrets 20231128 integration
//
// swagger:model secrets_20231128Integration
type Secrets20231128Integration struct {

	// aws access keys
	AwsAccessKeys *Secrets20231128AwsAccessKeysResponse `json:"aws_access_keys,omitempty"`

	// aws federated workload identity
	AwsFederatedWorkloadIdentity *Secrets20231128AwsFederatedWorkloadIdentityResponse `json:"aws_federated_workload_identity,omitempty"`

	// azure client secret
	AzureClientSecret *Secrets20231128AzureClientSecretResponse `json:"azure_client_secret,omitempty"`

	// azure federated workload identity
	AzureFederatedWorkloadIdentity *Secrets20231128AzureFederatedWorkloadIdentityResponse `json:"azure_federated_workload_identity,omitempty"`

	// capabilities
	Capabilities []*Secrets20231128Capability `json:"capabilities"`

	// confluent static credentials
	ConfluentStaticCredentials *Secrets20231128ConfluentStaticCredentialsResponse `json:"confluent_static_credentials,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// gcp federated workload identity
	GcpFederatedWorkloadIdentity *Secrets20231128GcpFederatedWorkloadIdentityResponse `json:"gcp_federated_workload_identity,omitempty"`

	// gcp service account key
	GcpServiceAccountKey *Secrets20231128GcpServiceAccountKeyResponse `json:"gcp_service_account_key,omitempty"`

	// gitlab access token
	GitlabAccessToken Secrets20231128GitlabAccessTokenResponse `json:"gitlab_access_token,omitempty"`

	// mongo db atlas static credentials
	MongoDbAtlasStaticCredentials *Secrets20231128MongoDBAtlasStaticCredentialsResponse `json:"mongo_db_atlas_static_credentials,omitempty"`

	// mysql static credentials
	MysqlStaticCredentials *Secrets20231128MysqlStaticCredentialsResponse `json:"mysql_static_credentials,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// postgres static credentials
	PostgresStaticCredentials *Secrets20231128PostgresStaticCredentialsResponse `json:"postgres_static_credentials,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// resource id
	ResourceID string `json:"resource_id,omitempty"`

	// resource name
	ResourceName string `json:"resource_name,omitempty"`

	// twilio static credentials
	TwilioStaticCredentials *Secrets20231128TwilioStaticCredentialsResponse `json:"twilio_static_credentials,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// updated by id
	UpdatedByID string `json:"updated_by_id,omitempty"`

	// used by
	UsedBy map[string]Secrets20231128IntegrationUsage `json:"used_by,omitempty"`
}

// Validate validates this secrets 20231128 integration
func (m *Secrets20231128Integration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsAccessKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsFederatedWorkloadIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureFederatedWorkloadIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfluentStaticCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcpFederatedWorkloadIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcpServiceAccountKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMongoDbAtlasStaticCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMysqlStaticCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostgresStaticCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTwilioStaticCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128Integration) validateAwsAccessKeys(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsAccessKeys) { // not required
		return nil
	}

	if m.AwsAccessKeys != nil {
		if err := m.AwsAccessKeys.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_access_keys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_access_keys")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) validateAwsFederatedWorkloadIdentity(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsFederatedWorkloadIdentity) { // not required
		return nil
	}

	if m.AwsFederatedWorkloadIdentity != nil {
		if err := m.AwsFederatedWorkloadIdentity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_federated_workload_identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_federated_workload_identity")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) validateAzureClientSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureClientSecret) { // not required
		return nil
	}

	if m.AzureClientSecret != nil {
		if err := m.AzureClientSecret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_client_secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_client_secret")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) validateAzureFederatedWorkloadIdentity(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureFederatedWorkloadIdentity) { // not required
		return nil
	}

	if m.AzureFederatedWorkloadIdentity != nil {
		if err := m.AzureFederatedWorkloadIdentity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_federated_workload_identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_federated_workload_identity")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) validateCapabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.Capabilities); i++ {
		if swag.IsZero(m.Capabilities[i]) { // not required
			continue
		}

		if m.Capabilities[i] != nil {
			if err := m.Capabilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("capabilities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("capabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Secrets20231128Integration) validateConfluentStaticCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfluentStaticCredentials) { // not required
		return nil
	}

	if m.ConfluentStaticCredentials != nil {
		if err := m.ConfluentStaticCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("confluent_static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("confluent_static_credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128Integration) validateGcpFederatedWorkloadIdentity(formats strfmt.Registry) error {
	if swag.IsZero(m.GcpFederatedWorkloadIdentity) { // not required
		return nil
	}

	if m.GcpFederatedWorkloadIdentity != nil {
		if err := m.GcpFederatedWorkloadIdentity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_federated_workload_identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_federated_workload_identity")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) validateGcpServiceAccountKey(formats strfmt.Registry) error {
	if swag.IsZero(m.GcpServiceAccountKey) { // not required
		return nil
	}

	if m.GcpServiceAccountKey != nil {
		if err := m.GcpServiceAccountKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_service_account_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_service_account_key")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) validateMongoDbAtlasStaticCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.MongoDbAtlasStaticCredentials) { // not required
		return nil
	}

	if m.MongoDbAtlasStaticCredentials != nil {
		if err := m.MongoDbAtlasStaticCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongo_db_atlas_static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongo_db_atlas_static_credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) validateMysqlStaticCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.MysqlStaticCredentials) { // not required
		return nil
	}

	if m.MysqlStaticCredentials != nil {
		if err := m.MysqlStaticCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mysql_static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mysql_static_credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) validatePostgresStaticCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.PostgresStaticCredentials) { // not required
		return nil
	}

	if m.PostgresStaticCredentials != nil {
		if err := m.PostgresStaticCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postgres_static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postgres_static_credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) validateTwilioStaticCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.TwilioStaticCredentials) { // not required
		return nil
	}

	if m.TwilioStaticCredentials != nil {
		if err := m.TwilioStaticCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("twilio_static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("twilio_static_credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128Integration) validateUsedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.UsedBy) { // not required
		return nil
	}

	for k := range m.UsedBy {

		if err := validate.Required("used_by"+"."+k, "body", m.UsedBy[k]); err != nil {
			return err
		}
		if val, ok := m.UsedBy[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("used_by" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("used_by" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this secrets 20231128 integration based on the context it is used
func (m *Secrets20231128Integration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsAccessKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsFederatedWorkloadIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureClientSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureFederatedWorkloadIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfluentStaticCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcpFederatedWorkloadIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcpServiceAccountKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMongoDbAtlasStaticCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMysqlStaticCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostgresStaticCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTwilioStaticCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128Integration) contextValidateAwsAccessKeys(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsAccessKeys != nil {

		if swag.IsZero(m.AwsAccessKeys) { // not required
			return nil
		}

		if err := m.AwsAccessKeys.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_access_keys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_access_keys")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) contextValidateAwsFederatedWorkloadIdentity(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsFederatedWorkloadIdentity != nil {

		if swag.IsZero(m.AwsFederatedWorkloadIdentity) { // not required
			return nil
		}

		if err := m.AwsFederatedWorkloadIdentity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_federated_workload_identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_federated_workload_identity")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) contextValidateAzureClientSecret(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureClientSecret != nil {

		if swag.IsZero(m.AzureClientSecret) { // not required
			return nil
		}

		if err := m.AzureClientSecret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_client_secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_client_secret")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) contextValidateAzureFederatedWorkloadIdentity(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureFederatedWorkloadIdentity != nil {

		if swag.IsZero(m.AzureFederatedWorkloadIdentity) { // not required
			return nil
		}

		if err := m.AzureFederatedWorkloadIdentity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_federated_workload_identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_federated_workload_identity")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Capabilities); i++ {

		if m.Capabilities[i] != nil {

			if swag.IsZero(m.Capabilities[i]) { // not required
				return nil
			}

			if err := m.Capabilities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("capabilities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("capabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Secrets20231128Integration) contextValidateConfluentStaticCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfluentStaticCredentials != nil {

		if swag.IsZero(m.ConfluentStaticCredentials) { // not required
			return nil
		}

		if err := m.ConfluentStaticCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("confluent_static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("confluent_static_credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) contextValidateGcpFederatedWorkloadIdentity(ctx context.Context, formats strfmt.Registry) error {

	if m.GcpFederatedWorkloadIdentity != nil {

		if swag.IsZero(m.GcpFederatedWorkloadIdentity) { // not required
			return nil
		}

		if err := m.GcpFederatedWorkloadIdentity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_federated_workload_identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_federated_workload_identity")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) contextValidateGcpServiceAccountKey(ctx context.Context, formats strfmt.Registry) error {

	if m.GcpServiceAccountKey != nil {

		if swag.IsZero(m.GcpServiceAccountKey) { // not required
			return nil
		}

		if err := m.GcpServiceAccountKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_service_account_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_service_account_key")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) contextValidateMongoDbAtlasStaticCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.MongoDbAtlasStaticCredentials != nil {

		if swag.IsZero(m.MongoDbAtlasStaticCredentials) { // not required
			return nil
		}

		if err := m.MongoDbAtlasStaticCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongo_db_atlas_static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongo_db_atlas_static_credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) contextValidateMysqlStaticCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.MysqlStaticCredentials != nil {

		if swag.IsZero(m.MysqlStaticCredentials) { // not required
			return nil
		}

		if err := m.MysqlStaticCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mysql_static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mysql_static_credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) contextValidatePostgresStaticCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.PostgresStaticCredentials != nil {

		if swag.IsZero(m.PostgresStaticCredentials) { // not required
			return nil
		}

		if err := m.PostgresStaticCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postgres_static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postgres_static_credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) contextValidateTwilioStaticCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.TwilioStaticCredentials != nil {

		if swag.IsZero(m.TwilioStaticCredentials) { // not required
			return nil
		}

		if err := m.TwilioStaticCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("twilio_static_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("twilio_static_credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128Integration) contextValidateUsedBy(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.UsedBy {

		if val, ok := m.UsedBy[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128Integration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128Integration) UnmarshalBinary(b []byte) error {
	var res Secrets20231128Integration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
