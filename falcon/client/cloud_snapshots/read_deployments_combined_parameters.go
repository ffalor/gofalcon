// Code generated by go-swagger; DO NOT EDIT.

package cloud_snapshots

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

// NewReadDeploymentsCombinedParams creates a new ReadDeploymentsCombinedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadDeploymentsCombinedParams() *ReadDeploymentsCombinedParams {
	return &ReadDeploymentsCombinedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadDeploymentsCombinedParamsWithTimeout creates a new ReadDeploymentsCombinedParams object
// with the ability to set a timeout on a request.
func NewReadDeploymentsCombinedParamsWithTimeout(timeout time.Duration) *ReadDeploymentsCombinedParams {
	return &ReadDeploymentsCombinedParams{
		timeout: timeout,
	}
}

// NewReadDeploymentsCombinedParamsWithContext creates a new ReadDeploymentsCombinedParams object
// with the ability to set a context for a request.
func NewReadDeploymentsCombinedParamsWithContext(ctx context.Context) *ReadDeploymentsCombinedParams {
	return &ReadDeploymentsCombinedParams{
		Context: ctx,
	}
}

// NewReadDeploymentsCombinedParamsWithHTTPClient creates a new ReadDeploymentsCombinedParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadDeploymentsCombinedParamsWithHTTPClient(client *http.Client) *ReadDeploymentsCombinedParams {
	return &ReadDeploymentsCombinedParams{
		HTTPClient: client,
	}
}

/*
ReadDeploymentsCombinedParams contains all the parameters to send to the API endpoint

	for the read deployments combined operation.

	Typically these are written to a http.Request.
*/
type ReadDeploymentsCombinedParams struct {

	/* Filter.

	   Search snapshot jobs using a query in Falcon Query Language (FQL). Supported filters:  account_id,asset_identifier,cloud_provider,region,status
	*/
	Filter *string

	/* Limit.

	   The upper-bound on the number of records to retrieve.
	*/
	Limit *int64

	/* Offset.

	   The offset from where to begin.
	*/
	Offset *int64

	/* Sort.

	   The fields to sort the records on. Supported columns:  [account_id asset_identifier cloud_provider instance_type last_updated_timestamp region status]
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read deployments combined params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadDeploymentsCombinedParams) WithDefaults() *ReadDeploymentsCombinedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read deployments combined params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadDeploymentsCombinedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) WithTimeout(timeout time.Duration) *ReadDeploymentsCombinedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) WithContext(ctx context.Context) *ReadDeploymentsCombinedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) WithHTTPClient(client *http.Client) *ReadDeploymentsCombinedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) WithFilter(filter *string) *ReadDeploymentsCombinedParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) WithLimit(limit *int64) *ReadDeploymentsCombinedParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) WithOffset(offset *int64) *ReadDeploymentsCombinedParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) WithSort(sort *string) *ReadDeploymentsCombinedParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the read deployments combined params
func (o *ReadDeploymentsCombinedParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ReadDeploymentsCombinedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
