// Code generated by go-swagger; DO NOT EDIT.

package discover

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

// NewGetLoginsParams creates a new GetLoginsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLoginsParams() *GetLoginsParams {
	return &GetLoginsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoginsParamsWithTimeout creates a new GetLoginsParams object
// with the ability to set a timeout on a request.
func NewGetLoginsParamsWithTimeout(timeout time.Duration) *GetLoginsParams {
	return &GetLoginsParams{
		timeout: timeout,
	}
}

// NewGetLoginsParamsWithContext creates a new GetLoginsParams object
// with the ability to set a context for a request.
func NewGetLoginsParamsWithContext(ctx context.Context) *GetLoginsParams {
	return &GetLoginsParams{
		Context: ctx,
	}
}

// NewGetLoginsParamsWithHTTPClient creates a new GetLoginsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLoginsParamsWithHTTPClient(client *http.Client) *GetLoginsParams {
	return &GetLoginsParams{
		HTTPClient: client,
	}
}

/* GetLoginsParams contains all the parameters to send to the API endpoint
   for the get logins operation.

   Typically these are written to a http.Request.
*/
type GetLoginsParams struct {

	/* Ids.

	   One or more login IDs (max: 100). Find login IDs with GET `/discover/queries/logins/v1`
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get logins params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoginsParams) WithDefaults() *GetLoginsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get logins params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoginsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get logins params
func (o *GetLoginsParams) WithTimeout(timeout time.Duration) *GetLoginsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logins params
func (o *GetLoginsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logins params
func (o *GetLoginsParams) WithContext(ctx context.Context) *GetLoginsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logins params
func (o *GetLoginsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logins params
func (o *GetLoginsParams) WithHTTPClient(client *http.Client) *GetLoginsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logins params
func (o *GetLoginsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get logins params
func (o *GetLoginsParams) WithIds(ids []string) *GetLoginsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get logins params
func (o *GetLoginsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoginsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetLogins binds the parameter ids
func (o *GetLoginsParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}
