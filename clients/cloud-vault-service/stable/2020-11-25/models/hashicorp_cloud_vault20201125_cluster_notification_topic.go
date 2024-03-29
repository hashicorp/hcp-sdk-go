// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpCloudVault20201125ClusterNotificationTopic Topic describes the event that triggered the notification
//
// swagger:model hashicorp.cloud.vault_20201125.Cluster.Notification.Topic
type HashicorpCloudVault20201125ClusterNotificationTopic string

func NewHashicorpCloudVault20201125ClusterNotificationTopic(value HashicorpCloudVault20201125ClusterNotificationTopic) *HashicorpCloudVault20201125ClusterNotificationTopic {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudVault20201125ClusterNotificationTopic.
func (m HashicorpCloudVault20201125ClusterNotificationTopic) Pointer() *HashicorpCloudVault20201125ClusterNotificationTopic {
	return &m
}

const (

	// HashicorpCloudVault20201125ClusterNotificationTopicCLUSTERNOTIFICATIONTOPICINVALID captures enum value "CLUSTER_NOTIFICATION_TOPIC_INVALID"
	HashicorpCloudVault20201125ClusterNotificationTopicCLUSTERNOTIFICATIONTOPICINVALID HashicorpCloudVault20201125ClusterNotificationTopic = "CLUSTER_NOTIFICATION_TOPIC_INVALID"

	// HashicorpCloudVault20201125ClusterNotificationTopicNEWMAJORVERSIONAVAILABLE captures enum value "NEW_MAJOR_VERSION_AVAILABLE"
	HashicorpCloudVault20201125ClusterNotificationTopicNEWMAJORVERSIONAVAILABLE HashicorpCloudVault20201125ClusterNotificationTopic = "NEW_MAJOR_VERSION_AVAILABLE"

	// HashicorpCloudVault20201125ClusterNotificationTopicMAJORVERSIONUPGRADEINPROGRESS captures enum value "MAJOR_VERSION_UPGRADE_IN_PROGRESS"
	HashicorpCloudVault20201125ClusterNotificationTopicMAJORVERSIONUPGRADEINPROGRESS HashicorpCloudVault20201125ClusterNotificationTopic = "MAJOR_VERSION_UPGRADE_IN_PROGRESS"

	// HashicorpCloudVault20201125ClusterNotificationTopicMAJORVERSIONUPGRADECOMPLETED captures enum value "MAJOR_VERSION_UPGRADE_COMPLETED"
	HashicorpCloudVault20201125ClusterNotificationTopicMAJORVERSIONUPGRADECOMPLETED HashicorpCloudVault20201125ClusterNotificationTopic = "MAJOR_VERSION_UPGRADE_COMPLETED"

	// HashicorpCloudVault20201125ClusterNotificationTopicMAJORVERSIONUPGRADEFAILED captures enum value "MAJOR_VERSION_UPGRADE_FAILED"
	HashicorpCloudVault20201125ClusterNotificationTopicMAJORVERSIONUPGRADEFAILED HashicorpCloudVault20201125ClusterNotificationTopic = "MAJOR_VERSION_UPGRADE_FAILED"

	// HashicorpCloudVault20201125ClusterNotificationTopicMAJORVERSIONUPGRADEMANUALREMINDER captures enum value "MAJOR_VERSION_UPGRADE_MANUAL_REMINDER"
	HashicorpCloudVault20201125ClusterNotificationTopicMAJORVERSIONUPGRADEMANUALREMINDER HashicorpCloudVault20201125ClusterNotificationTopic = "MAJOR_VERSION_UPGRADE_MANUAL_REMINDER"
)

// for schema
var hashicorpCloudVault20201125ClusterNotificationTopicEnum []interface{}

func init() {
	var res []HashicorpCloudVault20201125ClusterNotificationTopic
	if err := json.Unmarshal([]byte(`["CLUSTER_NOTIFICATION_TOPIC_INVALID","NEW_MAJOR_VERSION_AVAILABLE","MAJOR_VERSION_UPGRADE_IN_PROGRESS","MAJOR_VERSION_UPGRADE_COMPLETED","MAJOR_VERSION_UPGRADE_FAILED","MAJOR_VERSION_UPGRADE_MANUAL_REMINDER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVault20201125ClusterNotificationTopicEnum = append(hashicorpCloudVault20201125ClusterNotificationTopicEnum, v)
	}
}

func (m HashicorpCloudVault20201125ClusterNotificationTopic) validateHashicorpCloudVault20201125ClusterNotificationTopicEnum(path, location string, value HashicorpCloudVault20201125ClusterNotificationTopic) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVault20201125ClusterNotificationTopicEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vault 20201125 cluster notification topic
func (m HashicorpCloudVault20201125ClusterNotificationTopic) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVault20201125ClusterNotificationTopicEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20201125 cluster notification topic based on context it is used
func (m HashicorpCloudVault20201125ClusterNotificationTopic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
