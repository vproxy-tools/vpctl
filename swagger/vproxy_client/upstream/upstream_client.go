// Code generated by go-swagger; DO NOT EDIT.

package upstream

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new upstream API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for upstream API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddUpstream(params *AddUpstreamParams, opts ...ClientOption) (*AddUpstreamNoContent, error)

	DescribeUpstream(params *DescribeUpstreamParams, opts ...ClientOption) (*DescribeUpstreamOK, error)

	GetUpstream(params *GetUpstreamParams, opts ...ClientOption) (*GetUpstreamOK, error)

	ListUpstream(params *ListUpstreamParams, opts ...ClientOption) (*ListUpstreamOK, error)

	RemoveUpstream(params *RemoveUpstreamParams, opts ...ClientOption) (*RemoveUpstreamNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddUpstream adds upstream
*/
func (a *Client) AddUpstream(params *AddUpstreamParams, opts ...ClientOption) (*AddUpstreamNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUpstreamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addUpstream",
		Method:             "POST",
		PathPattern:        "/upstream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddUpstreamReader{formats: a.formats},
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
	success, ok := result.(*AddUpstreamNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addUpstream: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DescribeUpstream gets detailed info of one upstream
*/
func (a *Client) DescribeUpstream(params *DescribeUpstreamParams, opts ...ClientOption) (*DescribeUpstreamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeUpstreamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeUpstream",
		Method:             "GET",
		PathPattern:        "/upstream/{ups}/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeUpstreamReader{formats: a.formats},
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
	success, ok := result.(*DescribeUpstreamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for describeUpstream: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUpstream gets upstream
*/
func (a *Client) GetUpstream(params *GetUpstreamParams, opts ...ClientOption) (*GetUpstreamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUpstreamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUpstream",
		Method:             "GET",
		PathPattern:        "/upstream/{ups}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUpstreamReader{formats: a.formats},
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
	success, ok := result.(*GetUpstreamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUpstream: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListUpstream retrieves upstream list
*/
func (a *Client) ListUpstream(params *ListUpstreamParams, opts ...ClientOption) (*ListUpstreamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUpstreamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listUpstream",
		Method:             "GET",
		PathPattern:        "/upstream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListUpstreamReader{formats: a.formats},
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
	success, ok := result.(*ListUpstreamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUpstream: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveUpstream removes upstream
*/
func (a *Client) RemoveUpstream(params *RemoveUpstreamParams, opts ...ClientOption) (*RemoveUpstreamNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveUpstreamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeUpstream",
		Method:             "DELETE",
		PathPattern:        "/upstream/{ups}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveUpstreamReader{formats: a.formats},
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
	success, ok := result.(*RemoveUpstreamNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeUpstream: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
