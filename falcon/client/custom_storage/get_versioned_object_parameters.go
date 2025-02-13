// Code generated by go-swagger; DO NOT EDIT.

package custom_storage

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

// NewGetVersionedObjectParams creates a new GetVersionedObjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVersionedObjectParams() *GetVersionedObjectParams {
	return &GetVersionedObjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionedObjectParamsWithTimeout creates a new GetVersionedObjectParams object
// with the ability to set a timeout on a request.
func NewGetVersionedObjectParamsWithTimeout(timeout time.Duration) *GetVersionedObjectParams {
	return &GetVersionedObjectParams{
		timeout: timeout,
	}
}

// NewGetVersionedObjectParamsWithContext creates a new GetVersionedObjectParams object
// with the ability to set a context for a request.
func NewGetVersionedObjectParamsWithContext(ctx context.Context) *GetVersionedObjectParams {
	return &GetVersionedObjectParams{
		Context: ctx,
	}
}

// NewGetVersionedObjectParamsWithHTTPClient creates a new GetVersionedObjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVersionedObjectParamsWithHTTPClient(client *http.Client) *GetVersionedObjectParams {
	return &GetVersionedObjectParams{
		HTTPClient: client,
	}
}

/*
GetVersionedObjectParams contains all the parameters to send to the API endpoint

	for the get versioned object operation.

	Typically these are written to a http.Request.
*/
type GetVersionedObjectParams struct {

	/* CollectionName.

	   The name of the collection
	*/
	CollectionName string

	/* CollectionVersion.

	   The version of the collection
	*/
	CollectionVersion string

	/* ObjectKey.

	   The object key
	*/
	ObjectKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get versioned object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionedObjectParams) WithDefaults() *GetVersionedObjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get versioned object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionedObjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get versioned object params
func (o *GetVersionedObjectParams) WithTimeout(timeout time.Duration) *GetVersionedObjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get versioned object params
func (o *GetVersionedObjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get versioned object params
func (o *GetVersionedObjectParams) WithContext(ctx context.Context) *GetVersionedObjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get versioned object params
func (o *GetVersionedObjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get versioned object params
func (o *GetVersionedObjectParams) WithHTTPClient(client *http.Client) *GetVersionedObjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get versioned object params
func (o *GetVersionedObjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionName adds the collectionName to the get versioned object params
func (o *GetVersionedObjectParams) WithCollectionName(collectionName string) *GetVersionedObjectParams {
	o.SetCollectionName(collectionName)
	return o
}

// SetCollectionName adds the collectionName to the get versioned object params
func (o *GetVersionedObjectParams) SetCollectionName(collectionName string) {
	o.CollectionName = collectionName
}

// WithCollectionVersion adds the collectionVersion to the get versioned object params
func (o *GetVersionedObjectParams) WithCollectionVersion(collectionVersion string) *GetVersionedObjectParams {
	o.SetCollectionVersion(collectionVersion)
	return o
}

// SetCollectionVersion adds the collectionVersion to the get versioned object params
func (o *GetVersionedObjectParams) SetCollectionVersion(collectionVersion string) {
	o.CollectionVersion = collectionVersion
}

// WithObjectKey adds the objectKey to the get versioned object params
func (o *GetVersionedObjectParams) WithObjectKey(objectKey string) *GetVersionedObjectParams {
	o.SetObjectKey(objectKey)
	return o
}

// SetObjectKey adds the objectKey to the get versioned object params
func (o *GetVersionedObjectParams) SetObjectKey(objectKey string) {
	o.ObjectKey = objectKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionedObjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collection_name
	if err := r.SetPathParam("collection_name", o.CollectionName); err != nil {
		return err
	}

	// path param collection_version
	if err := r.SetPathParam("collection_version", o.CollectionVersion); err != nil {
		return err
	}

	// path param object_key
	if err := r.SetPathParam("object_key", o.ObjectKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
