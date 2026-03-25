// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

// GetImageByBuildLabelsRequestBody is the request body for GetImageByBuildLabels (2022-12-02).
type GetImageByBuildLabelsRequestBody struct {
	Labels        map[string]string `json:"labels,omitempty"`
	CloudProvider string            `json:"cloud_provider,omitempty"`
	Region        string            `json:"region,omitempty"`
}

// GetImageByBuildLabelsResponse is the response for GetImageByBuildLabels (2022-12-02).
type GetImageByBuildLabelsResponse struct {
	Artifact *ExternalArtifact `json:"artifact,omitempty"`
}

// ExternalArtifact contains bucket, version, and build for a single image (2022-12-02).
type ExternalArtifact struct {
	Bucket  *ExternalArtifactBucket  `json:"bucket,omitempty"`
	Version *ExternalArtifactVersion `json:"version,omitempty"`
	Build   *Build                   `json:"build,omitempty"`
}

// ExternalArtifactBucket is the bucket reference in ExternalArtifact.
type ExternalArtifactBucket struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ExternalArtifactVersion is the version reference in ExternalArtifact.
type ExternalArtifactVersion struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	RevokeAt    string `json:"revoke_at,omitempty"`
}

// Build is the 2022-12-02 build shape (matches API JSON; compatible with 2023-01-01 for parsing).
type Build struct {
	ID            string            `json:"id,omitempty"`
	VersionID     string            `json:"version_id,omitempty"`
	ComponentType string            `json:"component_type,omitempty"`
	PackerRunUUID string            `json:"packer_run_uuid,omitempty"`
	Platform      string            `json:"platform,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Artifacts     []*Artifact       `json:"artifacts,omitempty"`
}

// Artifact is the 2022-12-02 artifact shape (matches API JSON).
type Artifact struct {
	ID                 string `json:"id,omitempty"`
	ExternalIdentifier string `json:"external_identifier,omitempty"`
	Region             string `json:"region,omitempty"`
}
