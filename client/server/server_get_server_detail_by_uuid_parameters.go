// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewServerGetServerDetailByUUIDParams creates a new ServerGetServerDetailByUUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServerGetServerDetailByUUIDParams() *ServerGetServerDetailByUUIDParams {
	return &ServerGetServerDetailByUUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServerGetServerDetailByUUIDParamsWithTimeout creates a new ServerGetServerDetailByUUIDParams object
// with the ability to set a timeout on a request.
func NewServerGetServerDetailByUUIDParamsWithTimeout(timeout time.Duration) *ServerGetServerDetailByUUIDParams {
	return &ServerGetServerDetailByUUIDParams{
		timeout: timeout,
	}
}

// NewServerGetServerDetailByUUIDParamsWithContext creates a new ServerGetServerDetailByUUIDParams object
// with the ability to set a context for a request.
func NewServerGetServerDetailByUUIDParamsWithContext(ctx context.Context) *ServerGetServerDetailByUUIDParams {
	return &ServerGetServerDetailByUUIDParams{
		Context: ctx,
	}
}

// NewServerGetServerDetailByUUIDParamsWithHTTPClient creates a new ServerGetServerDetailByUUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewServerGetServerDetailByUUIDParamsWithHTTPClient(client *http.Client) *ServerGetServerDetailByUUIDParams {
	return &ServerGetServerDetailByUUIDParams{
		HTTPClient: client,
	}
}

/*
ServerGetServerDetailByUUIDParams contains all the parameters to send to the API endpoint

	for the server get server detail by UUID operation.

	Typically these are written to a http.Request.
*/
type ServerGetServerDetailByUUIDParams struct {

	/* Authorization.

	   api key auth
	*/
	Authorization *string

	/* ServerUUID.

	   Server UUID
	*/
	ServerUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the server get server detail by UUID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerGetServerDetailByUUIDParams) WithDefaults() *ServerGetServerDetailByUUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the server get server detail by UUID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerGetServerDetailByUUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the server get server detail by UUID params
func (o *ServerGetServerDetailByUUIDParams) WithTimeout(timeout time.Duration) *ServerGetServerDetailByUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the server get server detail by UUID params
func (o *ServerGetServerDetailByUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the server get server detail by UUID params
func (o *ServerGetServerDetailByUUIDParams) WithContext(ctx context.Context) *ServerGetServerDetailByUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the server get server detail by UUID params
func (o *ServerGetServerDetailByUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the server get server detail by UUID params
func (o *ServerGetServerDetailByUUIDParams) WithHTTPClient(client *http.Client) *ServerGetServerDetailByUUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the server get server detail by UUID params
func (o *ServerGetServerDetailByUUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the server get server detail by UUID params
func (o *ServerGetServerDetailByUUIDParams) WithAuthorization(authorization *string) *ServerGetServerDetailByUUIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the server get server detail by UUID params
func (o *ServerGetServerDetailByUUIDParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithServerUUID adds the serverUUID to the server get server detail by UUID params
func (o *ServerGetServerDetailByUUIDParams) WithServerUUID(serverUUID string) *ServerGetServerDetailByUUIDParams {
	o.SetServerUUID(serverUUID)
	return o
}

// SetServerUUID adds the serverUuid to the server get server detail by UUID params
func (o *ServerGetServerDetailByUUIDParams) SetServerUUID(serverUUID string) {
	o.ServerUUID = serverUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ServerGetServerDetailByUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}
	}

	// path param serverUuid
	if err := r.SetPathParam("serverUuid", o.ServerUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
