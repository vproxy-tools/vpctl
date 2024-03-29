// Code generated by go-swagger; DO NOT EDIT.

package server_group

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

// NewGetServerGroupParams creates a new GetServerGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServerGroupParams() *GetServerGroupParams {
	return &GetServerGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerGroupParamsWithTimeout creates a new GetServerGroupParams object
// with the ability to set a timeout on a request.
func NewGetServerGroupParamsWithTimeout(timeout time.Duration) *GetServerGroupParams {
	return &GetServerGroupParams{
		timeout: timeout,
	}
}

// NewGetServerGroupParamsWithContext creates a new GetServerGroupParams object
// with the ability to set a context for a request.
func NewGetServerGroupParamsWithContext(ctx context.Context) *GetServerGroupParams {
	return &GetServerGroupParams{
		Context: ctx,
	}
}

// NewGetServerGroupParamsWithHTTPClient creates a new GetServerGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServerGroupParamsWithHTTPClient(client *http.Client) *GetServerGroupParams {
	return &GetServerGroupParams{
		HTTPClient: client,
	}
}

/*
GetServerGroupParams contains all the parameters to send to the API endpoint

	for the get server group operation.

	Typically these are written to a http.Request.
*/
type GetServerGroupParams struct {

	/* Sg.

	   name of the server-group
	*/
	Sg string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get server group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerGroupParams) WithDefaults() *GetServerGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get server group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get server group params
func (o *GetServerGroupParams) WithTimeout(timeout time.Duration) *GetServerGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server group params
func (o *GetServerGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server group params
func (o *GetServerGroupParams) WithContext(ctx context.Context) *GetServerGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server group params
func (o *GetServerGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server group params
func (o *GetServerGroupParams) WithHTTPClient(client *http.Client) *GetServerGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server group params
func (o *GetServerGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSg adds the sg to the get server group params
func (o *GetServerGroupParams) WithSg(sg string) *GetServerGroupParams {
	o.SetSg(sg)
	return o
}

// SetSg adds the sg to the get server group params
func (o *GetServerGroupParams) SetSg(sg string) {
	o.Sg = sg
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sg
	if err := r.SetPathParam("sg", o.Sg); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
