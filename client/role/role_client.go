// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new role API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for role API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	RoleDeleteRole(params *RoleDeleteRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RoleDeleteRoleOK, error)

	RoleGetRoleDetail(params *RoleGetRoleDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RoleGetRoleDetailOK, error)

	RoleGetRoleList(params *RoleGetRoleListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RoleGetRoleListOK, error)

	RoleUpdateRole(params *RoleUpdateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RoleUpdateRoleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
RoleDeleteRole deletes role role

role delete
*/
func (a *Client) RoleDeleteRole(params *RoleDeleteRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RoleDeleteRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoleDeleteRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "role#deleteRole",
		Method:             "DELETE",
		PathPattern:        "/v1/role/{roleID}",
		ProducesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RoleDeleteRoleReader{formats: a.formats},
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
	success, ok := result.(*RoleDeleteRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for role#deleteRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RoleGetRoleDetail gets role detail role

role detail
*/
func (a *Client) RoleGetRoleDetail(params *RoleGetRoleDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RoleGetRoleDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoleGetRoleDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "role#getRoleDetail",
		Method:             "GET",
		PathPattern:        "/v1/role/{roleID}",
		ProducesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RoleGetRoleDetailReader{formats: a.formats},
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
	success, ok := result.(*RoleGetRoleDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for role#getRoleDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RoleGetRoleList gets role list role

role list
*/
func (a *Client) RoleGetRoleList(params *RoleGetRoleListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RoleGetRoleListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoleGetRoleListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "role#getRoleList",
		Method:             "GET",
		PathPattern:        "/v1/roles",
		ProducesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RoleGetRoleListReader{formats: a.formats},
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
	success, ok := result.(*RoleGetRoleListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for role#getRoleList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RoleUpdateRole updates role role

update role
*/
func (a *Client) RoleUpdateRole(params *RoleUpdateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RoleUpdateRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoleUpdateRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "role#updateRole",
		Method:             "PUT",
		PathPattern:        "/v1/role/{roleID}",
		ProducesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RoleUpdateRoleReader{formats: a.formats},
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
	success, ok := result.(*RoleUpdateRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for role#updateRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
