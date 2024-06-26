// Code generated by go-swagger; DO NOT EDIT.

package exposure_management

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

// NewBlobDownloadExternalAssetsParams creates a new BlobDownloadExternalAssetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBlobDownloadExternalAssetsParams() *BlobDownloadExternalAssetsParams {
	return &BlobDownloadExternalAssetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBlobDownloadExternalAssetsParamsWithTimeout creates a new BlobDownloadExternalAssetsParams object
// with the ability to set a timeout on a request.
func NewBlobDownloadExternalAssetsParamsWithTimeout(timeout time.Duration) *BlobDownloadExternalAssetsParams {
	return &BlobDownloadExternalAssetsParams{
		timeout: timeout,
	}
}

// NewBlobDownloadExternalAssetsParamsWithContext creates a new BlobDownloadExternalAssetsParams object
// with the ability to set a context for a request.
func NewBlobDownloadExternalAssetsParamsWithContext(ctx context.Context) *BlobDownloadExternalAssetsParams {
	return &BlobDownloadExternalAssetsParams{
		Context: ctx,
	}
}

// NewBlobDownloadExternalAssetsParamsWithHTTPClient creates a new BlobDownloadExternalAssetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewBlobDownloadExternalAssetsParamsWithHTTPClient(client *http.Client) *BlobDownloadExternalAssetsParams {
	return &BlobDownloadExternalAssetsParams{
		HTTPClient: client,
	}
}

/*
BlobDownloadExternalAssetsParams contains all the parameters to send to the API endpoint

	for the blob download external assets operation.

	Typically these are written to a http.Request.
*/
type BlobDownloadExternalAssetsParams struct {

	/* AssetID.

	   The Asset ID
	*/
	AssetID string

	/* Hash.

	   The File Hash
	*/
	Hash string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the blob download external assets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BlobDownloadExternalAssetsParams) WithDefaults() *BlobDownloadExternalAssetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the blob download external assets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BlobDownloadExternalAssetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the blob download external assets params
func (o *BlobDownloadExternalAssetsParams) WithTimeout(timeout time.Duration) *BlobDownloadExternalAssetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the blob download external assets params
func (o *BlobDownloadExternalAssetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the blob download external assets params
func (o *BlobDownloadExternalAssetsParams) WithContext(ctx context.Context) *BlobDownloadExternalAssetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the blob download external assets params
func (o *BlobDownloadExternalAssetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the blob download external assets params
func (o *BlobDownloadExternalAssetsParams) WithHTTPClient(client *http.Client) *BlobDownloadExternalAssetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the blob download external assets params
func (o *BlobDownloadExternalAssetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetID adds the assetID to the blob download external assets params
func (o *BlobDownloadExternalAssetsParams) WithAssetID(assetID string) *BlobDownloadExternalAssetsParams {
	o.SetAssetID(assetID)
	return o
}

// SetAssetID adds the assetId to the blob download external assets params
func (o *BlobDownloadExternalAssetsParams) SetAssetID(assetID string) {
	o.AssetID = assetID
}

// WithHash adds the hash to the blob download external assets params
func (o *BlobDownloadExternalAssetsParams) WithHash(hash string) *BlobDownloadExternalAssetsParams {
	o.SetHash(hash)
	return o
}

// SetHash adds the hash to the blob download external assets params
func (o *BlobDownloadExternalAssetsParams) SetHash(hash string) {
	o.Hash = hash
}

// WriteToRequest writes these params to a swagger request
func (o *BlobDownloadExternalAssetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param assetId
	qrAssetID := o.AssetID
	qAssetID := qrAssetID
	if qAssetID != "" {

		if err := r.SetQueryParam("assetId", qAssetID); err != nil {
			return err
		}
	}

	// query param hash
	qrHash := o.Hash
	qHash := qrHash
	if qHash != "" {

		if err := r.SetQueryParam("hash", qHash); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
