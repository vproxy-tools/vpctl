// Code generated by go-swagger; DO NOT EDIT.

package dns_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new dns server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dns server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddDNSServer(params *AddDNSServerParams, opts ...ClientOption) (*AddDNSServerNoContent, error)

	DescribeDNSServer(params *DescribeDNSServerParams, opts ...ClientOption) (*DescribeDNSServerOK, error)

	GetDNSServer(params *GetDNSServerParams, opts ...ClientOption) (*GetDNSServerOK, error)

	ListDNSServer(params *ListDNSServerParams, opts ...ClientOption) (*ListDNSServerOK, error)

	RemoveDNSServer(params *RemoveDNSServerParams, opts ...ClientOption) (*RemoveDNSServerNoContent, error)

	UpdateDNSServer(params *UpdateDNSServerParams, opts ...ClientOption) (*UpdateDNSServerNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddDNSServer adds dns server
*/
func (a *Client) AddDNSServer(params *AddDNSServerParams, opts ...ClientOption) (*AddDNSServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddDNSServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addDnsServer",
		Method:             "POST",
		PathPattern:        "/dns-server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddDNSServerReader{formats: a.formats},
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
	success, ok := result.(*AddDNSServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DescribeDNSServer gets detailed info of one dns server
*/
func (a *Client) DescribeDNSServer(params *DescribeDNSServerParams, opts ...ClientOption) (*DescribeDNSServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeDNSServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeDnsServer",
		Method:             "GET",
		PathPattern:        "/dns-server/{dns}/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeDNSServerReader{formats: a.formats},
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
	success, ok := result.(*DescribeDNSServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for describeDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDNSServer gets dns server
*/
func (a *Client) GetDNSServer(params *GetDNSServerParams, opts ...ClientOption) (*GetDNSServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDNSServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDnsServer",
		Method:             "GET",
		PathPattern:        "/dns-server/{dns}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDNSServerReader{formats: a.formats},
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
	success, ok := result.(*GetDNSServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListDNSServer retrieves dns server list
*/
func (a *Client) ListDNSServer(params *ListDNSServerParams, opts ...ClientOption) (*ListDNSServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDNSServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listDnsServer",
		Method:             "GET",
		PathPattern:        "/dns-server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListDNSServerReader{formats: a.formats},
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
	success, ok := result.(*ListDNSServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveDNSServer removes dns server
*/
func (a *Client) RemoveDNSServer(params *RemoveDNSServerParams, opts ...ClientOption) (*RemoveDNSServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveDNSServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeDnsServer",
		Method:             "DELETE",
		PathPattern:        "/dns-server/{dns}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveDNSServerReader{formats: a.formats},
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
	success, ok := result.(*RemoveDNSServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDNSServer updates dns server
*/
func (a *Client) UpdateDNSServer(params *UpdateDNSServerParams, opts ...ClientOption) (*UpdateDNSServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDNSServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDnsServer",
		Method:             "PUT",
		PathPattern:        "/dns-server/{dns}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDNSServerReader{formats: a.formats},
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
	success, ok := result.(*UpdateDNSServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
