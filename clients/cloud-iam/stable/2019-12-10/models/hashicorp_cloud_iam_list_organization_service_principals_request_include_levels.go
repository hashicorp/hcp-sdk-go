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

// HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels IncludeLevels represents hierarchical levels below an organization under
// which service principals can be crated.
//
//   - UNSET: UNSET is the default value.
//
// UNSET will only include service principals at organization level.
//   - ALL: ALL will include all service principals on all hirarchical levels
//
// under the given organization.
//   - PROJECTS: PROJECTS includes service principals at project level
//
// in the list response.
//
// swagger:model hashicorp.cloud.iam.ListOrganizationServicePrincipalsRequest.IncludeLevels
type HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels string

func NewHashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels(value HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels) *HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels.
func (m HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels) Pointer() *HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels {
	return &m
}

const (

	// HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsUNSET captures enum value "UNSET"
	HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsUNSET HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels = "UNSET"

	// HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsALL captures enum value "ALL"
	HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsALL HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels = "ALL"

	// HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsPROJECTS captures enum value "PROJECTS"
	HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsPROJECTS HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels = "PROJECTS"
)

// for schema
var hashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsEnum []interface{}

func init() {
	var res []HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels
	if err := json.Unmarshal([]byte(`["UNSET","ALL","PROJECTS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsEnum = append(hashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsEnum, v)
	}
}

func (m HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels) validateHashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsEnum(path, location string, value HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud iam list organization service principals request include levels
func (m HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevelsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud iam list organization service principals request include levels based on context it is used
func (m HashicorpCloudIamListOrganizationServicePrincipalsRequestIncludeLevels) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
