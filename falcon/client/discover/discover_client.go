// Code generated by go-swagger; DO NOT EDIT.

package discover

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new discover API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for discover API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetHosts(params *GetHostsParams, opts ...ClientOption) (*GetHostsOK, error)

	QueryHosts(params *QueryHostsParams, opts ...ClientOption) (*QueryHostsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetHosts gets details on assets by providing one or more i ds
*/
func (a *Client) GetHosts(params *GetHostsParams, opts ...ClientOption) (*GetHostsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-hosts",
		Method:             "GET",
		PathPattern:        "/discover/entities/hosts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHostsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHostsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHostsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryHosts searches for assets in your environment by providing an f q l falcon query language filter and paging details returns a set of asset i ds which match the filter criteria
*/
func (a *Client) QueryHosts(params *QueryHostsParams, opts ...ClientOption) (*QueryHostsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryHostsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "query-hosts",
		Method:             "GET",
		PathPattern:        "/discover/queries/hosts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryHostsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryHostsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryHostsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}