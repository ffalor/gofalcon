// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// NewContainerEnrichmentParams creates a new ContainerEnrichmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerEnrichmentParams() *ContainerEnrichmentParams {
	return &ContainerEnrichmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerEnrichmentParamsWithTimeout creates a new ContainerEnrichmentParams object
// with the ability to set a timeout on a request.
func NewContainerEnrichmentParamsWithTimeout(timeout time.Duration) *ContainerEnrichmentParams {
	return &ContainerEnrichmentParams{
		timeout: timeout,
	}
}

// NewContainerEnrichmentParamsWithContext creates a new ContainerEnrichmentParams object
// with the ability to set a context for a request.
func NewContainerEnrichmentParamsWithContext(ctx context.Context) *ContainerEnrichmentParams {
	return &ContainerEnrichmentParams{
		Context: ctx,
	}
}

// NewContainerEnrichmentParamsWithHTTPClient creates a new ContainerEnrichmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerEnrichmentParamsWithHTTPClient(client *http.Client) *ContainerEnrichmentParams {
	return &ContainerEnrichmentParams{
		HTTPClient: client,
	}
}

/*
ContainerEnrichmentParams contains all the parameters to send to the API endpoint

	for the container enrichment operation.

	Typically these are written to a http.Request.
*/
type ContainerEnrichmentParams struct {

	/* ContainerID.

	   One or more container ids for which to retrieve enrichment info
	*/
	ContainerID []string

	/* Filter.

	     Supported filter fields:
	- `last_seen`
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container enrichment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerEnrichmentParams) WithDefaults() *ContainerEnrichmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container enrichment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerEnrichmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container enrichment params
func (o *ContainerEnrichmentParams) WithTimeout(timeout time.Duration) *ContainerEnrichmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container enrichment params
func (o *ContainerEnrichmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container enrichment params
func (o *ContainerEnrichmentParams) WithContext(ctx context.Context) *ContainerEnrichmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container enrichment params
func (o *ContainerEnrichmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container enrichment params
func (o *ContainerEnrichmentParams) WithHTTPClient(client *http.Client) *ContainerEnrichmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container enrichment params
func (o *ContainerEnrichmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContainerID adds the containerID to the container enrichment params
func (o *ContainerEnrichmentParams) WithContainerID(containerID []string) *ContainerEnrichmentParams {
	o.SetContainerID(containerID)
	return o
}

// SetContainerID adds the containerId to the container enrichment params
func (o *ContainerEnrichmentParams) SetContainerID(containerID []string) {
	o.ContainerID = containerID
}

// WithFilter adds the filter to the container enrichment params
func (o *ContainerEnrichmentParams) WithFilter(filter *string) *ContainerEnrichmentParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the container enrichment params
func (o *ContainerEnrichmentParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerEnrichmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContainerID != nil {

		// binding items for container_id
		joinedContainerID := o.bindParamContainerID(reg)

		// query array param container_id
		if err := r.SetQueryParam("container_id", joinedContainerID...); err != nil {
			return err
		}
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamContainerEnrichment binds the parameter container_id
func (o *ContainerEnrichmentParams) bindParamContainerID(formats strfmt.Registry) []string {
	containerIDIR := o.ContainerID

	var containerIDIC []string
	for _, containerIDIIR := range containerIDIR { // explode []string

		containerIDIIV := containerIDIIR // string as string
		containerIDIC = append(containerIDIC, containerIDIIV)
	}

	// items.CollectionFormat: "csv"
	containerIDIS := swag.JoinByFormat(containerIDIC, "csv")

	return containerIDIS
}
