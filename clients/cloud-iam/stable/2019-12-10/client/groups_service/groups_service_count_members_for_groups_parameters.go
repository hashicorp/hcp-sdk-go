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
	"github.com/go-openapi/swag"
)

// NewGroupsServiceCountMembersForGroupsParams creates a new GroupsServiceCountMembersForGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupsServiceCountMembersForGroupsParams() *GroupsServiceCountMembersForGroupsParams {
	return &GroupsServiceCountMembersForGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupsServiceCountMembersForGroupsParamsWithTimeout creates a new GroupsServiceCountMembersForGroupsParams object
// with the ability to set a timeout on a request.
func NewGroupsServiceCountMembersForGroupsParamsWithTimeout(timeout time.Duration) *GroupsServiceCountMembersForGroupsParams {
	return &GroupsServiceCountMembersForGroupsParams{
		timeout: timeout,
	}
}

// NewGroupsServiceCountMembersForGroupsParamsWithContext creates a new GroupsServiceCountMembersForGroupsParams object
// with the ability to set a context for a request.
func NewGroupsServiceCountMembersForGroupsParamsWithContext(ctx context.Context) *GroupsServiceCountMembersForGroupsParams {
	return &GroupsServiceCountMembersForGroupsParams{
		Context: ctx,
	}
}

// NewGroupsServiceCountMembersForGroupsParamsWithHTTPClient creates a new GroupsServiceCountMembersForGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupsServiceCountMembersForGroupsParamsWithHTTPClient(client *http.Client) *GroupsServiceCountMembersForGroupsParams {
	return &GroupsServiceCountMembersForGroupsParams{
		HTTPClient: client,
	}
}

/*
GroupsServiceCountMembersForGroupsParams contains all the parameters to send to the API endpoint

	for the groups service count members for groups operation.

	Typically these are written to a http.Request.
*/
type GroupsServiceCountMembersForGroupsParams struct {

	/* GroupResourceNames.

	   group_resource_names is a list of the group resource_names for which to count members.
	*/
	GroupResourceNames []string

	/* ParentResourceName.

	     parent_resource is the parent resource of the groups that the groups should belong.
	organization/<org_id>
	*/
	ParentResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the groups service count members for groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsServiceCountMembersForGroupsParams) WithDefaults() *GroupsServiceCountMembersForGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the groups service count members for groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsServiceCountMembersForGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the groups service count members for groups params
func (o *GroupsServiceCountMembersForGroupsParams) WithTimeout(timeout time.Duration) *GroupsServiceCountMembersForGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the groups service count members for groups params
func (o *GroupsServiceCountMembersForGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the groups service count members for groups params
func (o *GroupsServiceCountMembersForGroupsParams) WithContext(ctx context.Context) *GroupsServiceCountMembersForGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the groups service count members for groups params
func (o *GroupsServiceCountMembersForGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the groups service count members for groups params
func (o *GroupsServiceCountMembersForGroupsParams) WithHTTPClient(client *http.Client) *GroupsServiceCountMembersForGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the groups service count members for groups params
func (o *GroupsServiceCountMembersForGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupResourceNames adds the groupResourceNames to the groups service count members for groups params
func (o *GroupsServiceCountMembersForGroupsParams) WithGroupResourceNames(groupResourceNames []string) *GroupsServiceCountMembersForGroupsParams {
	o.SetGroupResourceNames(groupResourceNames)
	return o
}

// SetGroupResourceNames adds the groupResourceNames to the groups service count members for groups params
func (o *GroupsServiceCountMembersForGroupsParams) SetGroupResourceNames(groupResourceNames []string) {
	o.GroupResourceNames = groupResourceNames
}

// WithParentResourceName adds the parentResourceName to the groups service count members for groups params
func (o *GroupsServiceCountMembersForGroupsParams) WithParentResourceName(parentResourceName string) *GroupsServiceCountMembersForGroupsParams {
	o.SetParentResourceName(parentResourceName)
	return o
}

// SetParentResourceName adds the parentResourceName to the groups service count members for groups params
func (o *GroupsServiceCountMembersForGroupsParams) SetParentResourceName(parentResourceName string) {
	o.ParentResourceName = parentResourceName
}

// WriteToRequest writes these params to a swagger request
func (o *GroupsServiceCountMembersForGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GroupResourceNames != nil {

		// binding items for group_resource_names
		joinedGroupResourceNames := o.bindParamGroupResourceNames(reg)

		// query array param group_resource_names
		if err := r.SetQueryParam("group_resource_names", joinedGroupResourceNames...); err != nil {
			return err
		}
	}

	// path param parent_resource_name
	if err := r.SetPathParam("parent_resource_name", o.ParentResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGroupsServiceCountMembersForGroups binds the parameter group_resource_names
func (o *GroupsServiceCountMembersForGroupsParams) bindParamGroupResourceNames(formats strfmt.Registry) []string {
	groupResourceNamesIR := o.GroupResourceNames

	var groupResourceNamesIC []string
	for _, groupResourceNamesIIR := range groupResourceNamesIR { // explode []string

		groupResourceNamesIIV := groupResourceNamesIIR // string as string
		groupResourceNamesIC = append(groupResourceNamesIC, groupResourceNamesIIV)
	}

	// items.CollectionFormat: "multi"
	groupResourceNamesIS := swag.JoinByFormat(groupResourceNamesIC, "multi")

	return groupResourceNamesIS
}
