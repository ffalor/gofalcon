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
)

// NewGroupContainersByManagedParams creates a new GroupContainersByManagedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupContainersByManagedParams() *GroupContainersByManagedParams {
	return &GroupContainersByManagedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupContainersByManagedParamsWithTimeout creates a new GroupContainersByManagedParams object
// with the ability to set a timeout on a request.
func NewGroupContainersByManagedParamsWithTimeout(timeout time.Duration) *GroupContainersByManagedParams {
	return &GroupContainersByManagedParams{
		timeout: timeout,
	}
}

// NewGroupContainersByManagedParamsWithContext creates a new GroupContainersByManagedParams object
// with the ability to set a context for a request.
func NewGroupContainersByManagedParamsWithContext(ctx context.Context) *GroupContainersByManagedParams {
	return &GroupContainersByManagedParams{
		Context: ctx,
	}
}

// NewGroupContainersByManagedParamsWithHTTPClient creates a new GroupContainersByManagedParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupContainersByManagedParamsWithHTTPClient(client *http.Client) *GroupContainersByManagedParams {
	return &GroupContainersByManagedParams{
		HTTPClient: client,
	}
}

/*
GroupContainersByManagedParams contains all the parameters to send to the API endpoint

	for the group containers by managed operation.

	Typically these are written to a http.Request.
*/
type GroupContainersByManagedParams struct {

	/* Filter.

	   Retrieve count of Kubernetes containers that match a query in Falcon Query Language (FQL). Supported filters:  agent_id,allow_privilege_escalation,cid,cloud_account_id,cloud_name,cloud_region,cluster_id,cluster_name,container_id,container_name,cve_id,detection_name,first_seen,image_detection_count,image_digest,image_has_been_assessed,image_id,image_registry,image_repository,image_tag,image_vulnerability_count,insecure_mount_source,insecure_mount_type,insecure_propagation_mode,interactive_mode,ipv4,ipv6,labels,last_seen,namespace,node_name,node_uid,pod_id,pod_name,port,privileged,root_write_access,run_as_root_group,run_as_root_user,running_status
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group containers by managed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupContainersByManagedParams) WithDefaults() *GroupContainersByManagedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group containers by managed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupContainersByManagedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group containers by managed params
func (o *GroupContainersByManagedParams) WithTimeout(timeout time.Duration) *GroupContainersByManagedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group containers by managed params
func (o *GroupContainersByManagedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group containers by managed params
func (o *GroupContainersByManagedParams) WithContext(ctx context.Context) *GroupContainersByManagedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group containers by managed params
func (o *GroupContainersByManagedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group containers by managed params
func (o *GroupContainersByManagedParams) WithHTTPClient(client *http.Client) *GroupContainersByManagedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group containers by managed params
func (o *GroupContainersByManagedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the group containers by managed params
func (o *GroupContainersByManagedParams) WithFilter(filter *string) *GroupContainersByManagedParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the group containers by managed params
func (o *GroupContainersByManagedParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *GroupContainersByManagedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}