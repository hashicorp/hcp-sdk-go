// Code generated by go-swagger; DO NOT EDIT.

package groups_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGroupsServiceUpdateGroupMembersParams creates a new GroupsServiceUpdateGroupMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupsServiceUpdateGroupMembersParams() *GroupsServiceUpdateGroupMembersParams {
	return &GroupsServiceUpdateGroupMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupsServiceUpdateGroupMembersParamsWithTimeout creates a new GroupsServiceUpdateGroupMembersParams object
// with the ability to set a timeout on a request.
func NewGroupsServiceUpdateGroupMembersParamsWithTimeout(timeout time.Duration) *GroupsServiceUpdateGroupMembersParams {
	return &GroupsServiceUpdateGroupMembersParams{
		timeout: timeout,
	}
}

// NewGroupsServiceUpdateGroupMembersParamsWithContext creates a new GroupsServiceUpdateGroupMembersParams object
// with the ability to set a context for a request.
func NewGroupsServiceUpdateGroupMembersParamsWithContext(ctx context.Context) *GroupsServiceUpdateGroupMembersParams {
	return &GroupsServiceUpdateGroupMembersParams{
		Context: ctx,
	}
}

// NewGroupsServiceUpdateGroupMembersParamsWithHTTPClient creates a new GroupsServiceUpdateGroupMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupsServiceUpdateGroupMembersParamsWithHTTPClient(client *http.Client) *GroupsServiceUpdateGroupMembersParams {
	return &GroupsServiceUpdateGroupMembersParams{
		HTTPClient: client,
	}
}

/*
GroupsServiceUpdateGroupMembersParams contains all the parameters to send to the API endpoint

	for the groups service update group members operation.

	Typically these are written to a http.Request.
*/
type GroupsServiceUpdateGroupMembersParams struct {

	// Body.
	Body GroupsServiceUpdateGroupMembersBody

	/* ResourceName.

	     resource_name is the resource name of the group.
	iam/organization/<org_id>/group/<group_id>.
	*/
	ResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the groups service update group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsServiceUpdateGroupMembersParams) WithDefaults() *GroupsServiceUpdateGroupMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the groups service update group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsServiceUpdateGroupMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the groups service update group members params
func (o *GroupsServiceUpdateGroupMembersParams) WithTimeout(timeout time.Duration) *GroupsServiceUpdateGroupMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the groups service update group members params
func (o *GroupsServiceUpdateGroupMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the groups service update group members params
func (o *GroupsServiceUpdateGroupMembersParams) WithContext(ctx context.Context) *GroupsServiceUpdateGroupMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the groups service update group members params
func (o *GroupsServiceUpdateGroupMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the groups service update group members params
func (o *GroupsServiceUpdateGroupMembersParams) WithHTTPClient(client *http.Client) *GroupsServiceUpdateGroupMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the groups service update group members params
func (o *GroupsServiceUpdateGroupMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the groups service update group members params
func (o *GroupsServiceUpdateGroupMembersParams) WithBody(body GroupsServiceUpdateGroupMembersBody) *GroupsServiceUpdateGroupMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the groups service update group members params
func (o *GroupsServiceUpdateGroupMembersParams) SetBody(body GroupsServiceUpdateGroupMembersBody) {
	o.Body = body
}

// WithResourceName adds the resourceName to the groups service update group members params
func (o *GroupsServiceUpdateGroupMembersParams) WithResourceName(resourceName string) *GroupsServiceUpdateGroupMembersParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the groups service update group members params
func (o *GroupsServiceUpdateGroupMembersParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *GroupsServiceUpdateGroupMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param resource_name
	if err := r.SetPathParam("resource_name", o.ResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}