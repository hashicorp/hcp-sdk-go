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

// HashicorpCloudConsul20200413ClusterState State is the state of the Consul cluster. Note that this state
// represents the abstract Consul cluster itself, not necessarily whether
// Consul cluster is currently available or not.
//
//  - UNSET: UNSET is a sentinel zero value so that an uninitialized value can be
// detected.
//  - PENDING: PENDING is the state the cluster is in while it is waiting to be created.
//  - CREATING: CREATING is the state the cluster is in while it is being provisioned for
// the first time.
//  - RUNNING: RUNNING is the steady state while the cluster is running.
//  - FAILED: FAILED is a failure state in which the cluster is unavailable and may
// required an operator restore action to recover.
//  - UPDATING: UPDATING is the state the cluster is in while undergoing a version
// update.
//  - RESTORING: RESTORING is the state the cluster is in while restoring from a snapshot.
//  - DELETING: DELETING is the state the cluster is in while it is being de-provisioned.
//  - DELETED: DELETED is the state the cluster is in when it has been de-provisioned. At
// this point, the cluster is eligible for garbage collection.
//
// swagger:model hashicorp.cloud.consul_20200413.Cluster.State
type HashicorpCloudConsul20200413ClusterState string

const (

	// HashicorpCloudConsul20200413ClusterStateUNSET captures enum value "UNSET"
	HashicorpCloudConsul20200413ClusterStateUNSET HashicorpCloudConsul20200413ClusterState = "UNSET"

	// HashicorpCloudConsul20200413ClusterStatePENDING captures enum value "PENDING"
	HashicorpCloudConsul20200413ClusterStatePENDING HashicorpCloudConsul20200413ClusterState = "PENDING"

	// HashicorpCloudConsul20200413ClusterStateCREATING captures enum value "CREATING"
	HashicorpCloudConsul20200413ClusterStateCREATING HashicorpCloudConsul20200413ClusterState = "CREATING"

	// HashicorpCloudConsul20200413ClusterStateRUNNING captures enum value "RUNNING"
	HashicorpCloudConsul20200413ClusterStateRUNNING HashicorpCloudConsul20200413ClusterState = "RUNNING"

	// HashicorpCloudConsul20200413ClusterStateFAILED captures enum value "FAILED"
	HashicorpCloudConsul20200413ClusterStateFAILED HashicorpCloudConsul20200413ClusterState = "FAILED"

	// HashicorpCloudConsul20200413ClusterStateUPDATING captures enum value "UPDATING"
	HashicorpCloudConsul20200413ClusterStateUPDATING HashicorpCloudConsul20200413ClusterState = "UPDATING"

	// HashicorpCloudConsul20200413ClusterStateRESTORING captures enum value "RESTORING"
	HashicorpCloudConsul20200413ClusterStateRESTORING HashicorpCloudConsul20200413ClusterState = "RESTORING"

	// HashicorpCloudConsul20200413ClusterStateDELETING captures enum value "DELETING"
	HashicorpCloudConsul20200413ClusterStateDELETING HashicorpCloudConsul20200413ClusterState = "DELETING"

	// HashicorpCloudConsul20200413ClusterStateDELETED captures enum value "DELETED"
	HashicorpCloudConsul20200413ClusterStateDELETED HashicorpCloudConsul20200413ClusterState = "DELETED"
)

// for schema
var hashicorpCloudConsul20200413ClusterStateEnum []interface{}

func init() {
	var res []HashicorpCloudConsul20200413ClusterState
	if err := json.Unmarshal([]byte(`["UNSET","PENDING","CREATING","RUNNING","FAILED","UPDATING","RESTORING","DELETING","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudConsul20200413ClusterStateEnum = append(hashicorpCloudConsul20200413ClusterStateEnum, v)
	}
}

func (m HashicorpCloudConsul20200413ClusterState) validateHashicorpCloudConsul20200413ClusterStateEnum(path, location string, value HashicorpCloudConsul20200413ClusterState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudConsul20200413ClusterStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud consul 20200413 cluster state
func (m HashicorpCloudConsul20200413ClusterState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudConsul20200413ClusterStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
