// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpCloudConsul20210204ClusterConfigTier Tier represents the different tiers a Consul cluster can be provisioned as.
//
// swagger:model hashicorp.cloud.consul_20210204.ClusterConfig.Tier
type HashicorpCloudConsul20210204ClusterConfigTier string

const (

	// HashicorpCloudConsul20210204ClusterConfigTierUNSET captures enum value "UNSET"
	HashicorpCloudConsul20210204ClusterConfigTierUNSET HashicorpCloudConsul20210204ClusterConfigTier = "UNSET"

	// HashicorpCloudConsul20210204ClusterConfigTierDEVELOPMENT captures enum value "DEVELOPMENT"
	HashicorpCloudConsul20210204ClusterConfigTierDEVELOPMENT HashicorpCloudConsul20210204ClusterConfigTier = "DEVELOPMENT"

	// HashicorpCloudConsul20210204ClusterConfigTierSTANDARD captures enum value "STANDARD"
	HashicorpCloudConsul20210204ClusterConfigTierSTANDARD HashicorpCloudConsul20210204ClusterConfigTier = "STANDARD"

	// HashicorpCloudConsul20210204ClusterConfigTierPLUS captures enum value "PLUS"
	HashicorpCloudConsul20210204ClusterConfigTierPLUS HashicorpCloudConsul20210204ClusterConfigTier = "PLUS"
)

// for schema
var hashicorpCloudConsul20210204ClusterConfigTierEnum []interface{}

func init() {
	var res []HashicorpCloudConsul20210204ClusterConfigTier
	if err := json.Unmarshal([]byte(`["UNSET","DEVELOPMENT","STANDARD","PLUS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudConsul20210204ClusterConfigTierEnum = append(hashicorpCloudConsul20210204ClusterConfigTierEnum, v)
	}
}

func (m HashicorpCloudConsul20210204ClusterConfigTier) validateHashicorpCloudConsul20210204ClusterConfigTierEnum(path, location string, value HashicorpCloudConsul20210204ClusterConfigTier) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudConsul20210204ClusterConfigTierEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud consul 20210204 cluster config tier
func (m HashicorpCloudConsul20210204ClusterConfigTier) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudConsul20210204ClusterConfigTierEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
