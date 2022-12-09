// Code generated by go-swagger; DO NOT EDIT.

package recon

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewAggregateNotificationsExposedDataRecordsV1Params creates a new AggregateNotificationsExposedDataRecordsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregateNotificationsExposedDataRecordsV1Params() *AggregateNotificationsExposedDataRecordsV1Params {
	return &AggregateNotificationsExposedDataRecordsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateNotificationsExposedDataRecordsV1ParamsWithTimeout creates a new AggregateNotificationsExposedDataRecordsV1Params object
// with the ability to set a timeout on a request.
func NewAggregateNotificationsExposedDataRecordsV1ParamsWithTimeout(timeout time.Duration) *AggregateNotificationsExposedDataRecordsV1Params {
	return &AggregateNotificationsExposedDataRecordsV1Params{
		timeout: timeout,
	}
}

// NewAggregateNotificationsExposedDataRecordsV1ParamsWithContext creates a new AggregateNotificationsExposedDataRecordsV1Params object
// with the ability to set a context for a request.
func NewAggregateNotificationsExposedDataRecordsV1ParamsWithContext(ctx context.Context) *AggregateNotificationsExposedDataRecordsV1Params {
	return &AggregateNotificationsExposedDataRecordsV1Params{
		Context: ctx,
	}
}

// NewAggregateNotificationsExposedDataRecordsV1ParamsWithHTTPClient creates a new AggregateNotificationsExposedDataRecordsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewAggregateNotificationsExposedDataRecordsV1ParamsWithHTTPClient(client *http.Client) *AggregateNotificationsExposedDataRecordsV1Params {
	return &AggregateNotificationsExposedDataRecordsV1Params{
		HTTPClient: client,
	}
}

/*
AggregateNotificationsExposedDataRecordsV1Params contains all the parameters to send to the API endpoint

	for the aggregate notifications exposed data records v1 operation.

	Typically these are written to a http.Request.
*/
type AggregateNotificationsExposedDataRecordsV1Params struct {

	// Body.
	Body []*models.MsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate notifications exposed data records v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateNotificationsExposedDataRecordsV1Params) WithDefaults() *AggregateNotificationsExposedDataRecordsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate notifications exposed data records v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateNotificationsExposedDataRecordsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aggregate notifications exposed data records v1 params
func (o *AggregateNotificationsExposedDataRecordsV1Params) WithTimeout(timeout time.Duration) *AggregateNotificationsExposedDataRecordsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate notifications exposed data records v1 params
func (o *AggregateNotificationsExposedDataRecordsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate notifications exposed data records v1 params
func (o *AggregateNotificationsExposedDataRecordsV1Params) WithContext(ctx context.Context) *AggregateNotificationsExposedDataRecordsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate notifications exposed data records v1 params
func (o *AggregateNotificationsExposedDataRecordsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate notifications exposed data records v1 params
func (o *AggregateNotificationsExposedDataRecordsV1Params) WithHTTPClient(client *http.Client) *AggregateNotificationsExposedDataRecordsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate notifications exposed data records v1 params
func (o *AggregateNotificationsExposedDataRecordsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the aggregate notifications exposed data records v1 params
func (o *AggregateNotificationsExposedDataRecordsV1Params) WithBody(body []*models.MsaAggregateQueryRequest) *AggregateNotificationsExposedDataRecordsV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the aggregate notifications exposed data records v1 params
func (o *AggregateNotificationsExposedDataRecordsV1Params) SetBody(body []*models.MsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateNotificationsExposedDataRecordsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}