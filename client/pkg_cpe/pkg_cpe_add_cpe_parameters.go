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

	"github.com/Masato516/vuls-client-go-swagger/models"
)

// NewPkgCpeAddCpeParams creates a new PkgCpeAddCpeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPkgCpeAddCpeParams() *PkgCpeAddCpeParams {
	return &PkgCpeAddCpeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPkgCpeAddCpeParamsWithTimeout creates a new PkgCpeAddCpeParams object
// with the ability to set a timeout on a request.
func NewPkgCpeAddCpeParamsWithTimeout(timeout time.Duration) *PkgCpeAddCpeParams {
	return &PkgCpeAddCpeParams{
		timeout: timeout,
	}
}

// NewPkgCpeAddCpeParamsWithContext creates a new PkgCpeAddCpeParams object
// with the ability to set a context for a request.
func NewPkgCpeAddCpeParamsWithContext(ctx context.Context) *PkgCpeAddCpeParams {
	return &PkgCpeAddCpeParams{
		Context: ctx,
	}
}

// NewPkgCpeAddCpeParamsWithHTTPClient creates a new PkgCpeAddCpeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPkgCpeAddCpeParamsWithHTTPClient(client *http.Client) *PkgCpeAddCpeParams {
	return &PkgCpeAddCpeParams{
		HTTPClient: client,
	}
}

/*
PkgCpeAddCpeParams contains all the parameters to send to the API endpoint

	for the pkg cpe add cpe operation.

	Typically these are written to a http.Request.
*/
type PkgCpeAddCpeParams struct {

	// AddCpeRequestBody.
	AddCpeRequestBody *models.PkgCpeAddCpeRequestBody

	/* Authorization.

	   api key auth
	*/
	Authorization *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pkg cpe add cpe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PkgCpeAddCpeParams) WithDefaults() *PkgCpeAddCpeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pkg cpe add cpe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PkgCpeAddCpeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pkg cpe add cpe params
func (o *PkgCpeAddCpeParams) WithTimeout(timeout time.Duration) *PkgCpeAddCpeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pkg cpe add cpe params
func (o *PkgCpeAddCpeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pkg cpe add cpe params
func (o *PkgCpeAddCpeParams) WithContext(ctx context.Context) *PkgCpeAddCpeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pkg cpe add cpe params
func (o *PkgCpeAddCpeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pkg cpe add cpe params
func (o *PkgCpeAddCpeParams) WithHTTPClient(client *http.Client) *PkgCpeAddCpeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pkg cpe add cpe params
func (o *PkgCpeAddCpeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddCpeRequestBody adds the addCpeRequestBody to the pkg cpe add cpe params
func (o *PkgCpeAddCpeParams) WithAddCpeRequestBody(addCpeRequestBody *models.PkgCpeAddCpeRequestBody) *PkgCpeAddCpeParams {
	o.SetAddCpeRequestBody(addCpeRequestBody)
	return o
}

// SetAddCpeRequestBody adds the addCpeRequestBody to the pkg cpe add cpe params
func (o *PkgCpeAddCpeParams) SetAddCpeRequestBody(addCpeRequestBody *models.PkgCpeAddCpeRequestBody) {
	o.AddCpeRequestBody = addCpeRequestBody
}

// WithAuthorization adds the authorization to the pkg cpe add cpe params
func (o *PkgCpeAddCpeParams) WithAuthorization(authorization *string) *PkgCpeAddCpeParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the pkg cpe add cpe params
func (o *PkgCpeAddCpeParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *PkgCpeAddCpeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AddCpeRequestBody != nil {
		if err := r.SetBodyParam(o.AddCpeRequestBody); err != nil {
			return err
		}
	}

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}