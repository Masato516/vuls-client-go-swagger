// Code generated by go-swagger; DO NOT EDIT.

package lockfile

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

// NewLockfileDeleteLockfileParams creates a new LockfileDeleteLockfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLockfileDeleteLockfileParams() *LockfileDeleteLockfileParams {
	return &LockfileDeleteLockfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLockfileDeleteLockfileParamsWithTimeout creates a new LockfileDeleteLockfileParams object
// with the ability to set a timeout on a request.
func NewLockfileDeleteLockfileParamsWithTimeout(timeout time.Duration) *LockfileDeleteLockfileParams {
	return &LockfileDeleteLockfileParams{
		timeout: timeout,
	}
}

// NewLockfileDeleteLockfileParamsWithContext creates a new LockfileDeleteLockfileParams object
// with the ability to set a context for a request.
func NewLockfileDeleteLockfileParamsWithContext(ctx context.Context) *LockfileDeleteLockfileParams {
	return &LockfileDeleteLockfileParams{
		Context: ctx,
	}
}

// NewLockfileDeleteLockfileParamsWithHTTPClient creates a new LockfileDeleteLockfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewLockfileDeleteLockfileParamsWithHTTPClient(client *http.Client) *LockfileDeleteLockfileParams {
	return &LockfileDeleteLockfileParams{
		HTTPClient: client,
	}
}

/*
LockfileDeleteLockfileParams contains all the parameters to send to the API endpoint

	for the lockfile delete lockfile operation.

	Typically these are written to a http.Request.
*/
type LockfileDeleteLockfileParams struct {

	/* Authorization.

	   api key auth
	*/
	Authorization *string

	/* LockfileID.

	   Lockfile ID

	   Format: int64
	*/
	LockfileID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lockfile delete lockfile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LockfileDeleteLockfileParams) WithDefaults() *LockfileDeleteLockfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lockfile delete lockfile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LockfileDeleteLockfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lockfile delete lockfile params
func (o *LockfileDeleteLockfileParams) WithTimeout(timeout time.Duration) *LockfileDeleteLockfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lockfile delete lockfile params
func (o *LockfileDeleteLockfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lockfile delete lockfile params
func (o *LockfileDeleteLockfileParams) WithContext(ctx context.Context) *LockfileDeleteLockfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lockfile delete lockfile params
func (o *LockfileDeleteLockfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lockfile delete lockfile params
func (o *LockfileDeleteLockfileParams) WithHTTPClient(client *http.Client) *LockfileDeleteLockfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lockfile delete lockfile params
func (o *LockfileDeleteLockfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the lockfile delete lockfile params
func (o *LockfileDeleteLockfileParams) WithAuthorization(authorization *string) *LockfileDeleteLockfileParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the lockfile delete lockfile params
func (o *LockfileDeleteLockfileParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithLockfileID adds the lockfileID to the lockfile delete lockfile params
func (o *LockfileDeleteLockfileParams) WithLockfileID(lockfileID int64) *LockfileDeleteLockfileParams {
	o.SetLockfileID(lockfileID)
	return o
}

// SetLockfileID adds the lockfileId to the lockfile delete lockfile params
func (o *LockfileDeleteLockfileParams) SetLockfileID(lockfileID int64) {
	o.LockfileID = lockfileID
}

// WriteToRequest writes these params to a swagger request
func (o *LockfileDeleteLockfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param lockfileID
	if err := r.SetPathParam("lockfileID", swag.FormatInt64(o.LockfileID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
