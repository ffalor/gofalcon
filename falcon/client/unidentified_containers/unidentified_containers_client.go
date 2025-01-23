// Code generated by go-swagger; DO NOT EDIT.

package unidentified_containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new unidentified containers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for unidentified containers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Count(params *CountParams, opts ...ClientOption) (*CountOK, error)

	CountByDateRange(params *CountByDateRangeParams, opts ...ClientOption) (*CountByDateRangeOK, error)

	Search(params *SearchParams, opts ...ClientOption) (*SearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Count returns the total count of unidentified containers over a time period
*/
func (a *Client) Count(params *CountParams, opts ...ClientOption) (*CountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Count",
		Method:             "GET",
		PathPattern:        "/container-security/aggregates/unidentified-containers/count/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CountReader{formats: a.formats},
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
	success, ok := result.(*CountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Count: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CountByDateRange returns the count of unidentified containers over the last 7 days
*/
func (a *Client) CountByDateRange(params *CountByDateRangeParams, opts ...ClientOption) (*CountByDateRangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCountByDateRangeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CountByDateRange",
		Method:             "GET",
		PathPattern:        "/container-security/aggregates/unidentified-containers/count-by-date/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CountByDateRangeReader{formats: a.formats},
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
	success, ok := result.(*CountByDateRangeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CountByDateRange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Search maximums offset 10000 limit
*/
func (a *Client) Search(params *SearchParams, opts ...ClientOption) (*SearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Search",
		Method:             "GET",
		PathPattern:        "/container-security/combined/unidentified-containers/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchReader{formats: a.formats},
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
	success, ok := result.(*SearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
