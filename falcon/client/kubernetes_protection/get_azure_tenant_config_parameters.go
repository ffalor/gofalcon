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

// NewGetAzureTenantConfigParams creates a new GetAzureTenantConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAzureTenantConfigParams() *GetAzureTenantConfigParams {
	return &GetAzureTenantConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAzureTenantConfigParamsWithTimeout creates a new GetAzureTenantConfigParams object
// with the ability to set a timeout on a request.
func NewGetAzureTenantConfigParamsWithTimeout(timeout time.Duration) *GetAzureTenantConfigParams {
	return &GetAzureTenantConfigParams{
		timeout: timeout,
	}
}

// NewGetAzureTenantConfigParamsWithContext creates a new GetAzureTenantConfigParams object
// with the ability to set a context for a request.
func NewGetAzureTenantConfigParamsWithContext(ctx context.Context) *GetAzureTenantConfigParams {
	return &GetAzureTenantConfigParams{
		Context: ctx,
	}
}

// NewGetAzureTenantConfigParamsWithHTTPClient creates a new GetAzureTenantConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAzureTenantConfigParamsWithHTTPClient(client *http.Client) *GetAzureTenantConfigParams {
	return &GetAzureTenantConfigParams{
		HTTPClient: client,
	}
}

/*
GetAzureTenantConfigParams contains all the parameters to send to the API endpoint

	for the get azure tenant config operation.

	Typically these are written to a http.Request.
*/
type GetAzureTenantConfigParams struct {

	/* Ids.

	   Azure Tenant IDs
	*/
	Ids []string

	/* Limit.

	   Limit returned accounts
	*/
	Limit *int64

	/* Offset.

	   Offset returned accounts
	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get azure tenant config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAzureTenantConfigParams) WithDefaults() *GetAzureTenantConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get azure tenant config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAzureTenantConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get azure tenant config params
func (o *GetAzureTenantConfigParams) WithTimeout(timeout time.Duration) *GetAzureTenantConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get azure tenant config params
func (o *GetAzureTenantConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get azure tenant config params
func (o *GetAzureTenantConfigParams) WithContext(ctx context.Context) *GetAzureTenantConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get azure tenant config params
func (o *GetAzureTenantConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get azure tenant config params
func (o *GetAzureTenantConfigParams) WithHTTPClient(client *http.Client) *GetAzureTenantConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get azure tenant config params
func (o *GetAzureTenantConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get azure tenant config params
func (o *GetAzureTenantConfigParams) WithIds(ids []string) *GetAzureTenantConfigParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get azure tenant config params
func (o *GetAzureTenantConfigParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithLimit adds the limit to the get azure tenant config params
func (o *GetAzureTenantConfigParams) WithLimit(limit *int64) *GetAzureTenantConfigParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get azure tenant config params
func (o *GetAzureTenantConfigParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get azure tenant config params
func (o *GetAzureTenantConfigParams) WithOffset(offset *int64) *GetAzureTenantConfigParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get azure tenant config params
func (o *GetAzureTenantConfigParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetAzureTenantConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAzureTenantConfig binds the parameter ids
func (o *GetAzureTenantConfigParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "csv"
	idsIS := swag.JoinByFormat(idsIC, "csv")

	return idsIS
}