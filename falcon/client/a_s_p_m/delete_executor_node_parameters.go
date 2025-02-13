// Code generated by go-swagger; DO NOT EDIT.

package a_s_p_m

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

// NewDeleteExecutorNodeParams creates a new DeleteExecutorNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteExecutorNodeParams() *DeleteExecutorNodeParams {
	return &DeleteExecutorNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteExecutorNodeParamsWithTimeout creates a new DeleteExecutorNodeParams object
// with the ability to set a timeout on a request.
func NewDeleteExecutorNodeParamsWithTimeout(timeout time.Duration) *DeleteExecutorNodeParams {
	return &DeleteExecutorNodeParams{
		timeout: timeout,
	}
}

// NewDeleteExecutorNodeParamsWithContext creates a new DeleteExecutorNodeParams object
// with the ability to set a context for a request.
func NewDeleteExecutorNodeParamsWithContext(ctx context.Context) *DeleteExecutorNodeParams {
	return &DeleteExecutorNodeParams{
		Context: ctx,
	}
}

// NewDeleteExecutorNodeParamsWithHTTPClient creates a new DeleteExecutorNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteExecutorNodeParamsWithHTTPClient(client *http.Client) *DeleteExecutorNodeParams {
	return &DeleteExecutorNodeParams{
		HTTPClient: client,
	}
}

/*
DeleteExecutorNodeParams contains all the parameters to send to the API endpoint

	for the delete executor node operation.

	Typically these are written to a http.Request.
*/
type DeleteExecutorNodeParams struct {

	// ID.
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete executor node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExecutorNodeParams) WithDefaults() *DeleteExecutorNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete executor node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExecutorNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete executor node params
func (o *DeleteExecutorNodeParams) WithTimeout(timeout time.Duration) *DeleteExecutorNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete executor node params
func (o *DeleteExecutorNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete executor node params
func (o *DeleteExecutorNodeParams) WithContext(ctx context.Context) *DeleteExecutorNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete executor node params
func (o *DeleteExecutorNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete executor node params
func (o *DeleteExecutorNodeParams) WithHTTPClient(client *http.Client) *DeleteExecutorNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete executor node params
func (o *DeleteExecutorNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete executor node params
func (o *DeleteExecutorNodeParams) WithID(id int64) *DeleteExecutorNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete executor node params
func (o *DeleteExecutorNodeParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteExecutorNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ID
	if err := r.SetPathParam("ID", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
