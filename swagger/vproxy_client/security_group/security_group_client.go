// Code generated by go-swagger; DO NOT EDIT.

package security_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new security group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for security group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddSecurityGroup(params *AddSecurityGroupParams, opts ...ClientOption) (*AddSecurityGroupNoContent, error)

	DescribeSecurityGroup(params *DescribeSecurityGroupParams, opts ...ClientOption) (*DescribeSecurityGroupOK, error)

	GetSecurityGroup(params *GetSecurityGroupParams, opts ...ClientOption) (*GetSecurityGroupOK, error)

	ListSecurityGroup(params *ListSecurityGroupParams, opts ...ClientOption) (*ListSecurityGroupOK, error)

	RemoveSecurityGroup(params *RemoveSecurityGroupParams, opts ...ClientOption) (*RemoveSecurityGroupNoContent, error)

	UpdateSecurityGroup(params *UpdateSecurityGroupParams, opts ...ClientOption) (*UpdateSecurityGroupNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddSecurityGroup adds security group
*/
func (a *Client) AddSecurityGroup(params *AddSecurityGroupParams, opts ...ClientOption) (*AddSecurityGroupNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSecurityGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addSecurityGroup",
		Method:             "POST",
		PathPattern:        "/security-group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddSecurityGroupReader{formats: a.formats},
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
	success, ok := result.(*AddSecurityGroupNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addSecurityGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DescribeSecurityGroup gets detailed info of one security group
*/
func (a *Client) DescribeSecurityGroup(params *DescribeSecurityGroupParams, opts ...ClientOption) (*DescribeSecurityGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeSecurityGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "describeSecurityGroup",
		Method:             "GET",
		PathPattern:        "/security-group/{secg}/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeSecurityGroupReader{formats: a.formats},
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
	success, ok := result.(*DescribeSecurityGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for describeSecurityGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSecurityGroup gets security group
*/
func (a *Client) GetSecurityGroup(params *GetSecurityGroupParams, opts ...ClientOption) (*GetSecurityGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSecurityGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSecurityGroup",
		Method:             "GET",
		PathPattern:        "/security-group/{secg}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSecurityGroupReader{formats: a.formats},
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
	success, ok := result.(*GetSecurityGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSecurityGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListSecurityGroup retrieves security group list
*/
func (a *Client) ListSecurityGroup(params *ListSecurityGroupParams, opts ...ClientOption) (*ListSecurityGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSecurityGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSecurityGroup",
		Method:             "GET",
		PathPattern:        "/security-group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSecurityGroupReader{formats: a.formats},
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
	success, ok := result.(*ListSecurityGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listSecurityGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveSecurityGroup removes security group
*/
func (a *Client) RemoveSecurityGroup(params *RemoveSecurityGroupParams, opts ...ClientOption) (*RemoveSecurityGroupNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveSecurityGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeSecurityGroup",
		Method:             "DELETE",
		PathPattern:        "/security-group/{secg}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveSecurityGroupReader{formats: a.formats},
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
	success, ok := result.(*RemoveSecurityGroupNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeSecurityGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateSecurityGroup updates security group
*/
func (a *Client) UpdateSecurityGroup(params *UpdateSecurityGroupParams, opts ...ClientOption) (*UpdateSecurityGroupNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSecurityGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSecurityGroup",
		Method:             "PUT",
		PathPattern:        "/security-group/{secg}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSecurityGroupReader{formats: a.formats},
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
	success, ok := result.(*UpdateSecurityGroupNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSecurityGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
