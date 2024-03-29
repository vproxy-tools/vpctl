// Code generated by go-swagger; DO NOT EDIT.

package socks5_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new socks5 server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for socks5 server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddSocks5Server(params *AddSocks5ServerParams, opts ...ClientOption) (*AddSocks5ServerNoContent, error)

	DescribeSocks5Server(params *DescribeSocks5ServerParams, opts ...ClientOption) (*DescribeSocks5ServerOK, error)

	GetSocks5Server(params *GetSocks5ServerParams, opts ...ClientOption) (*GetSocks5ServerOK, error)

	ListSocks5Server(params *ListSocks5ServerParams, opts ...ClientOption) (*ListSocks5ServerOK, error)

	RemoveSocks5Server(params *RemoveSocks5ServerParams, opts ...ClientOption) (*RemoveSocks5ServerNoContent, error)

	UpdateSocks5Server(params *UpdateSocks5ServerParams, opts ...ClientOption) (*UpdateSocks5ServerNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddSocks5Server adds socks5 server
*/
func (a *Client) AddSocks5Server(params *AddSocks5ServerParams, opts ...ClientOption) (*AddSocks5ServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSocks5ServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addSocks5Server",
		Method:             "POST",
		PathPattern:        "/socks5-server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddSocks5ServerReader{formats: a.formats},
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
	success, ok := result.(*AddSocks5ServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addSocks5Server: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DescribeSocks5Server gets detailed info of one socks5 server
*/
func (a *Client) DescribeSocks5Server(params *DescribeSocks5ServerParams, opts ...ClientOption) (*DescribeSocks5ServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeSocks5ServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeSocks5Server",
		Method:             "GET",
		PathPattern:        "/socks5-server/{socks5}/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeSocks5ServerReader{formats: a.formats},
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
	success, ok := result.(*DescribeSocks5ServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for describeSocks5Server: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSocks5Server gets socks5 server
*/
func (a *Client) GetSocks5Server(params *GetSocks5ServerParams, opts ...ClientOption) (*GetSocks5ServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSocks5ServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSocks5Server",
		Method:             "GET",
		PathPattern:        "/socks5-server/{socks5}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSocks5ServerReader{formats: a.formats},
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
	success, ok := result.(*GetSocks5ServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSocks5Server: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListSocks5Server retrieves socks5 server list
*/
func (a *Client) ListSocks5Server(params *ListSocks5ServerParams, opts ...ClientOption) (*ListSocks5ServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSocks5ServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSocks5Server",
		Method:             "GET",
		PathPattern:        "/socks5-server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSocks5ServerReader{formats: a.formats},
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
	success, ok := result.(*ListSocks5ServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listSocks5Server: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveSocks5Server removes socks5 server
*/
func (a *Client) RemoveSocks5Server(params *RemoveSocks5ServerParams, opts ...ClientOption) (*RemoveSocks5ServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveSocks5ServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeSocks5Server",
		Method:             "DELETE",
		PathPattern:        "/socks5-server/{socks5}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveSocks5ServerReader{formats: a.formats},
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
	success, ok := result.(*RemoveSocks5ServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeSocks5Server: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateSocks5Server updates socks5 server
*/
func (a *Client) UpdateSocks5Server(params *UpdateSocks5ServerParams, opts ...ClientOption) (*UpdateSocks5ServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSocks5ServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSocks5Server",
		Method:             "PUT",
		PathPattern:        "/socks5-server/{socks5}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSocks5ServerReader{formats: a.formats},
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
	success, ok := result.(*UpdateSocks5ServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSocks5Server: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
