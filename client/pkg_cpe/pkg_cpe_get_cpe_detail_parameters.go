// Code generated by go-swagger; DO NOT EDIT.

package pkg_cpe

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

// NewPkgCpeGetCpeDetailParams creates a new PkgCpeGetCpeDetailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPkgCpeGetCpeDetailParams() *PkgCpeGetCpeDetailParams {
	return &PkgCpeGetCpeDetailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPkgCpeGetCpeDetailParamsWithTimeout creates a new PkgCpeGetCpeDetailParams object
// with the ability to set a timeout on a request.
func NewPkgCpeGetCpeDetailParamsWithTimeout(timeout time.Duration) *PkgCpeGetCpeDetailParams {
	return &PkgCpeGetCpeDetailParams{
		timeout: timeout,
	}
}

// NewPkgCpeGetCpeDetailParamsWithContext creates a new PkgCpeGetCpeDetailParams object
// with the ability to set a context for a request.
func NewPkgCpeGetCpeDetailParamsWithContext(ctx context.Context) *PkgCpeGetCpeDetailParams {
	return &PkgCpeGetCpeDetailParams{
		Context: ctx,
	}
}

// NewPkgCpeGetCpeDetailParamsWithHTTPClient creates a new PkgCpeGetCpeDetailParams object
// with the ability to set a custom HTTPClient for a request.
func NewPkgCpeGetCpeDetailParamsWithHTTPClient(client *http.Client) *PkgCpeGetCpeDetailParams {
	return &PkgCpeGetCpeDetailParams{
		HTTPClient: client,
	}
}

/*
PkgCpeGetCpeDetailParams contains all the parameters to send to the API endpoint

	for the pkg cpe get cpe detail operation.

	Typically these are written to a http.Request.
*/
type PkgCpeGetCpeDetailParams struct {

	/* Authorization.

	   api key auth
	*/
	Authorization *string

	/* CpeID.

	   cpe ID

	   Format: int64
	*/
	CpeID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pkg cpe get cpe detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PkgCpeGetCpeDetailParams) WithDefaults() *PkgCpeGetCpeDetailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pkg cpe get cpe detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PkgCpeGetCpeDetailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pkg cpe get cpe detail params
func (o *PkgCpeGetCpeDetailParams) WithTimeout(timeout time.Duration) *PkgCpeGetCpeDetailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pkg cpe get cpe detail params
func (o *PkgCpeGetCpeDetailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pkg cpe get cpe detail params
func (o *PkgCpeGetCpeDetailParams) WithContext(ctx context.Context) *PkgCpeGetCpeDetailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pkg cpe get cpe detail params
func (o *PkgCpeGetCpeDetailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pkg cpe get cpe detail params
func (o *PkgCpeGetCpeDetailParams) WithHTTPClient(client *http.Client) *PkgCpeGetCpeDetailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pkg cpe get cpe detail params
func (o *PkgCpeGetCpeDetailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the pkg cpe get cpe detail params
func (o *PkgCpeGetCpeDetailParams) WithAuthorization(authorization *string) *PkgCpeGetCpeDetailParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the pkg cpe get cpe detail params
func (o *PkgCpeGetCpeDetailParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithCpeID adds the cpeID to the pkg cpe get cpe detail params
func (o *PkgCpeGetCpeDetailParams) WithCpeID(cpeID int64) *PkgCpeGetCpeDetailParams {
	o.SetCpeID(cpeID)
	return o
}

// SetCpeID adds the cpeId to the pkg cpe get cpe detail params
func (o *PkgCpeGetCpeDetailParams) SetCpeID(cpeID int64) {
	o.CpeID = cpeID
}

// WriteToRequest writes these params to a swagger request
func (o *PkgCpeGetCpeDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param cpeID
	if err := r.SetPathParam("cpeID", swag.FormatInt64(o.CpeID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
