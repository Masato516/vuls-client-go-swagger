// Code generated by go-swagger; DO NOT EDIT.

package task

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

	"github.com/Masato516/vuls-client-go-swagger/models"
)

// NewTaskUpdateTaskIgnoreParams creates a new TaskUpdateTaskIgnoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTaskUpdateTaskIgnoreParams() *TaskUpdateTaskIgnoreParams {
	return &TaskUpdateTaskIgnoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTaskUpdateTaskIgnoreParamsWithTimeout creates a new TaskUpdateTaskIgnoreParams object
// with the ability to set a timeout on a request.
func NewTaskUpdateTaskIgnoreParamsWithTimeout(timeout time.Duration) *TaskUpdateTaskIgnoreParams {
	return &TaskUpdateTaskIgnoreParams{
		timeout: timeout,
	}
}

// NewTaskUpdateTaskIgnoreParamsWithContext creates a new TaskUpdateTaskIgnoreParams object
// with the ability to set a context for a request.
func NewTaskUpdateTaskIgnoreParamsWithContext(ctx context.Context) *TaskUpdateTaskIgnoreParams {
	return &TaskUpdateTaskIgnoreParams{
		Context: ctx,
	}
}

// NewTaskUpdateTaskIgnoreParamsWithHTTPClient creates a new TaskUpdateTaskIgnoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewTaskUpdateTaskIgnoreParamsWithHTTPClient(client *http.Client) *TaskUpdateTaskIgnoreParams {
	return &TaskUpdateTaskIgnoreParams{
		HTTPClient: client,
	}
}

/*
TaskUpdateTaskIgnoreParams contains all the parameters to send to the API endpoint

	for the task update task ignore operation.

	Typically these are written to a http.Request.
*/
type TaskUpdateTaskIgnoreParams struct {

	/* Authorization.

	   api key auth
	*/
	Authorization *string

	// UpdateTaskIgnoreRequestBody.
	UpdateTaskIgnoreRequestBody *models.TaskUpdateTaskIgnoreRequestBody

	/* TaskID.

	   Task ID

	   Format: int64
	*/
	TaskID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the task update task ignore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TaskUpdateTaskIgnoreParams) WithDefaults() *TaskUpdateTaskIgnoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the task update task ignore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TaskUpdateTaskIgnoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) WithTimeout(timeout time.Duration) *TaskUpdateTaskIgnoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) WithContext(ctx context.Context) *TaskUpdateTaskIgnoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) WithHTTPClient(client *http.Client) *TaskUpdateTaskIgnoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) WithAuthorization(authorization *string) *TaskUpdateTaskIgnoreParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithUpdateTaskIgnoreRequestBody adds the updateTaskIgnoreRequestBody to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) WithUpdateTaskIgnoreRequestBody(updateTaskIgnoreRequestBody *models.TaskUpdateTaskIgnoreRequestBody) *TaskUpdateTaskIgnoreParams {
	o.SetUpdateTaskIgnoreRequestBody(updateTaskIgnoreRequestBody)
	return o
}

// SetUpdateTaskIgnoreRequestBody adds the updateTaskIgnoreRequestBody to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) SetUpdateTaskIgnoreRequestBody(updateTaskIgnoreRequestBody *models.TaskUpdateTaskIgnoreRequestBody) {
	o.UpdateTaskIgnoreRequestBody = updateTaskIgnoreRequestBody
}

// WithTaskID adds the taskID to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) WithTaskID(taskID int64) *TaskUpdateTaskIgnoreParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the task update task ignore params
func (o *TaskUpdateTaskIgnoreParams) SetTaskID(taskID int64) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *TaskUpdateTaskIgnoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.UpdateTaskIgnoreRequestBody != nil {
		if err := r.SetBodyParam(o.UpdateTaskIgnoreRequestBody); err != nil {
			return err
		}
	}

	// path param taskID
	if err := r.SetPathParam("taskID", swag.FormatInt64(o.TaskID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
