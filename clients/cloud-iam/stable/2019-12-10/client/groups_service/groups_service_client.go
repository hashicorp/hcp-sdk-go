// Code generated by go-swagger; DO NOT EDIT.

package groups_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new groups service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for groups service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateGroup(params *CreateGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGroupOK, error)

	DeleteGroup(params *DeleteGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteGroupOK, error)

	GetGroup(params *GetGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGroupOK, error)

	ListGroupMembers(params *ListGroupMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGroupMembersOK, error)

	ListGroups(params *ListGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGroupsOK, error)

	UpdateGroup(params *UpdateGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGroupOK, error)

	UpdateGroupMembers(params *UpdateGroupMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGroupMembersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateGroup creates group creates a group
*/
func (a *Client) CreateGroup(params *CreateGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateGroup",
		Method:             "POST",
		PathPattern:        "/iam/2019-12-10/iam/{parent_resource_name}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteGroup deletes group deletes a group
*/
func (a *Client) DeleteGroup(params *DeleteGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteGroup",
		Method:             "DELETE",
		PathPattern:        "/iam/2019-12-10/{resource_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetGroup gets group gets a group using the organization ID and group name
*/
func (a *Client) GetGroup(params *GetGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetGroup",
		Method:             "GET",
		PathPattern:        "/iam/2019-12-10/{resource_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGroupMembers lists group members retrieves a list of member principals
*/
func (a *Client) ListGroupMembers(params *ListGroupMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGroupMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGroupMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListGroupMembers",
		Method:             "GET",
		PathPattern:        "/iam/2019-12-10/{resource_name}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListGroupMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGroupMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGroupMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGroups lists groups lists the groups in an organization that match the optional filters
*/
func (a *Client) ListGroups(params *ListGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListGroups",
		Method:             "GET",
		PathPattern:        "/iam/2019-12-10/iam/{parent_resource_name}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateGroup updates group updates a group s modifiable field
*/
func (a *Client) UpdateGroup(params *UpdateGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateGroup",
		Method:             "PUT",
		PathPattern:        "/iam/2019-12-10/{resource_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateGroupMembers updates group members updates the principal memberships of a group
*/
func (a *Client) UpdateGroupMembers(params *UpdateGroupMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGroupMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGroupMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateGroupMembers",
		Method:             "PUT",
		PathPattern:        "/iam/2019-12-10/{resource_name}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGroupMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateGroupMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateGroupMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
