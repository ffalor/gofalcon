// Code generated by go-swagger; DO NOT EDIT.

package threatgraph

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

// NewQueriesEdgetypesGetParams creates a new QueriesEdgetypesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueriesEdgetypesGetParams() *QueriesEdgetypesGetParams {
	return &QueriesEdgetypesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueriesEdgetypesGetParamsWithTimeout creates a new QueriesEdgetypesGetParams object
// with the ability to set a timeout on a request.
func NewQueriesEdgetypesGetParamsWithTimeout(timeout time.Duration) *QueriesEdgetypesGetParams {
	return &QueriesEdgetypesGetParams{
		timeout: timeout,
	}
}

// NewQueriesEdgetypesGetParamsWithContext creates a new QueriesEdgetypesGetParams object
// with the ability to set a context for a request.
func NewQueriesEdgetypesGetParamsWithContext(ctx context.Context) *QueriesEdgetypesGetParams {
	return &QueriesEdgetypesGetParams{
		Context: ctx,
	}
}

// NewQueriesEdgetypesGetParamsWithHTTPClient creates a new QueriesEdgetypesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueriesEdgetypesGetParamsWithHTTPClient(client *http.Client) *QueriesEdgetypesGetParams {
	return &QueriesEdgetypesGetParams{
		HTTPClient: client,
	}
}

/*
QueriesEdgetypesGetParams contains all the parameters to send to the API endpoint

	for the queries edgetypes get operation.

	Typically these are written to a http.Request.
*/
type QueriesEdgetypesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the queries edgetypes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueriesEdgetypesGetParams) WithDefaults() *QueriesEdgetypesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the queries edgetypes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueriesEdgetypesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the queries edgetypes get params
func (o *QueriesEdgetypesGetParams) WithTimeout(timeout time.Duration) *QueriesEdgetypesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the queries edgetypes get params
func (o *QueriesEdgetypesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the queries edgetypes get params
func (o *QueriesEdgetypesGetParams) WithContext(ctx context.Context) *QueriesEdgetypesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the queries edgetypes get params
func (o *QueriesEdgetypesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the queries edgetypes get params
func (o *QueriesEdgetypesGetParams) WithHTTPClient(client *http.Client) *QueriesEdgetypesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the queries edgetypes get params
func (o *QueriesEdgetypesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *QueriesEdgetypesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}