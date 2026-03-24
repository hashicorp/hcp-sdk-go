// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package build_service

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2022-12-02/models"
)

// NewGetImageByBuildLabelsParams creates a new GetImageByBuildLabelsParams.
func NewGetImageByBuildLabelsParams() *GetImageByBuildLabelsParams {
	return &GetImageByBuildLabelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageByBuildLabelsParamsWithContext creates GetImageByBuildLabelsParams with context.
func NewGetImageByBuildLabelsParamsWithContext(ctx context.Context) *GetImageByBuildLabelsParams {
	return &GetImageByBuildLabelsParams{
		Context: ctx,
	}
}

// GetImageByBuildLabelsParams contains all parameters for GetImageByBuildLabels.
type GetImageByBuildLabelsParams struct {
	Body *models.GetImageByBuildLabelsRequestBody

	BucketName string
	ChannelName string
	LocationOrganizationID string
	LocationProjectID      string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithBody sets the body.
func (o *GetImageByBuildLabelsParams) WithBody(body *models.GetImageByBuildLabelsRequestBody) *GetImageByBuildLabelsParams {
	o.SetBody(body)
	return o
}

// SetBody sets the body.
func (o *GetImageByBuildLabelsParams) SetBody(body *models.GetImageByBuildLabelsRequestBody) {
	o.Body = body
}

// WithBucketName sets the bucket name.
func (o *GetImageByBuildLabelsParams) WithBucketName(v string) *GetImageByBuildLabelsParams {
	o.SetBucketName(v)
	return o
}

// SetBucketName sets the bucket name.
func (o *GetImageByBuildLabelsParams) SetBucketName(v string) {
	o.BucketName = v
}

// WithChannelName sets the channel name.
func (o *GetImageByBuildLabelsParams) WithChannelName(v string) *GetImageByBuildLabelsParams {
	o.SetChannelName(v)
	return o
}

// SetChannelName sets the channel name.
func (o *GetImageByBuildLabelsParams) SetChannelName(v string) {
	o.ChannelName = v
}

// WithLocationOrganizationID sets the organization ID.
func (o *GetImageByBuildLabelsParams) WithLocationOrganizationID(v string) *GetImageByBuildLabelsParams {
	o.SetLocationOrganizationID(v)
	return o
}

// SetLocationOrganizationID sets the organization ID.
func (o *GetImageByBuildLabelsParams) SetLocationOrganizationID(v string) {
	o.LocationOrganizationID = v
}

// WithLocationProjectID sets the project ID.
func (o *GetImageByBuildLabelsParams) WithLocationProjectID(v string) *GetImageByBuildLabelsParams {
	o.SetLocationProjectID(v)
	return o
}

// SetLocationProjectID sets the project ID.
func (o *GetImageByBuildLabelsParams) SetLocationProjectID(v string) {
	o.LocationProjectID = v
}

// WithTimeout sets the timeout.
func (o *GetImageByBuildLabelsParams) WithTimeout(timeout time.Duration) *GetImageByBuildLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout sets the timeout.
func (o *GetImageByBuildLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext sets the context.
func (o *GetImageByBuildLabelsParams) WithContext(ctx context.Context) *GetImageByBuildLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext sets the context.
func (o *GetImageByBuildLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes params to the request.
func (o *GetImageByBuildLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetPathParam("bucket_name", o.BucketName); err != nil {
		res = append(res, err)
	}
	if err := r.SetPathParam("channel_name", o.ChannelName); err != nil {
		res = append(res, err)
	}
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		res = append(res, err)
	}
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		res = append(res, err)
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
