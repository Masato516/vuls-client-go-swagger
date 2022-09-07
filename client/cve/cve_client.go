// Code generated by go-swagger; DO NOT EDIT.

package cve

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cve API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cve API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CveGetCveDetail(params *CveGetCveDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CveGetCveDetailOK, error)

	CveGetCveList(params *CveGetCveListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CveGetCveListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CveGetCveDetail gets cve detail cve

cve detail
*/
func (a *Client) CveGetCveDetail(params *CveGetCveDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CveGetCveDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCveGetCveDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cve#getCveDetail",
		Method:             "GET",
		PathPattern:        "/v1/cve/{cveID}",
		ProducesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CveGetCveDetailReader{formats: a.formats},
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
	success, ok := result.(*CveGetCveDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cve#getCveDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CveGetCveList gets cve list cve

cve list
*/
func (a *Client) CveGetCveList(params *CveGetCveListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CveGetCveListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCveGetCveListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cve#getCveList",
		Method:             "GET",
		PathPattern:        "/v1/cves",
		ProducesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/gob", "application/json", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CveGetCveListReader{formats: a.formats},
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
	success, ok := result.(*CveGetCveListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cve#getCveList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
