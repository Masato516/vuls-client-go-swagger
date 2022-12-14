// Code generated by go-swagger; DO NOT EDIT.

package role

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

// NewRoleDeleteRoleParams creates a new RoleDeleteRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRoleDeleteRoleParams() *RoleDeleteRoleParams {
	return &RoleDeleteRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRoleDeleteRoleParamsWithTimeout creates a new RoleDeleteRoleParams object
// with the ability to set a timeout on a request.
func NewRoleDeleteRoleParamsWithTimeout(timeout time.Duration) *RoleDeleteRoleParams {
	return &RoleDeleteRoleParams{
		timeout: timeout,
	}
}

// NewRoleDeleteRoleParamsWithContext creates a new RoleDeleteRoleParams object
// with the ability to set a context for a request.
func NewRoleDeleteRoleParamsWithContext(ctx context.Context) *RoleDeleteRoleParams {
	return &RoleDeleteRoleParams{
		Context: ctx,
	}
}

// NewRoleDeleteRoleParamsWithHTTPClient creates a new RoleDeleteRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewRoleDeleteRoleParamsWithHTTPClient(client *http.Client) *RoleDeleteRoleParams {
	return &RoleDeleteRoleParams{
		HTTPClient: client,
	}
}

/*
RoleDeleteRoleParams contains all the parameters to send to the API endpoint

	for the role delete role operation.

	Typically these are written to a http.Request.
*/
type RoleDeleteRoleParams struct {

	/* Authorization.

	   api key auth
	*/
	Authorization *string

	/* RoleID.

	   Role ID

	   Format: int64
	*/
	RoleID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the role delete role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoleDeleteRoleParams) WithDefaults() *RoleDeleteRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the role delete role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoleDeleteRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the role delete role params
func (o *RoleDeleteRoleParams) WithTimeout(timeout time.Duration) *RoleDeleteRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the role delete role params
func (o *RoleDeleteRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the role delete role params
func (o *RoleDeleteRoleParams) WithContext(ctx context.Context) *RoleDeleteRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the role delete role params
func (o *RoleDeleteRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the role delete role params
func (o *RoleDeleteRoleParams) WithHTTPClient(client *http.Client) *RoleDeleteRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the role delete role params
func (o *RoleDeleteRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the role delete role params
func (o *RoleDeleteRoleParams) WithAuthorization(authorization *string) *RoleDeleteRoleParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the role delete role params
func (o *RoleDeleteRoleParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithRoleID adds the roleID to the role delete role params
func (o *RoleDeleteRoleParams) WithRoleID(roleID int64) *RoleDeleteRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the role delete role params
func (o *RoleDeleteRoleParams) SetRoleID(roleID int64) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *RoleDeleteRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param roleID
	if err := r.SetPathParam("roleID", swag.FormatInt64(o.RoleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
