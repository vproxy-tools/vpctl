// Code generated by go-swagger; DO NOT EDIT.

package server

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

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// NewAddServerParams creates a new AddServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddServerParams() *AddServerParams {
	return &AddServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddServerParamsWithTimeout creates a new AddServerParams object
// with the ability to set a timeout on a request.
func NewAddServerParamsWithTimeout(timeout time.Duration) *AddServerParams {
	return &AddServerParams{
		timeout: timeout,
	}
}

// NewAddServerParamsWithContext creates a new AddServerParams object
// with the ability to set a context for a request.
func NewAddServerParamsWithContext(ctx context.Context) *AddServerParams {
	return &AddServerParams{
		Context: ctx,
	}
}

// NewAddServerParamsWithHTTPClient creates a new AddServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddServerParamsWithHTTPClient(client *http.Client) *AddServerParams {
	return &AddServerParams{
		HTTPClient: client,
	}
}

/*
AddServerParams contains all the parameters to send to the API endpoint

	for the add server operation.

	Typically these are written to a http.Request.
*/
type AddServerParams struct {

	// Body.
	Body *vproxy_client_model.ServerCreate

	/* Sg.

	   name of the server-group
	*/
	Sg string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddServerParams) WithDefaults() *AddServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add server params
func (o *AddServerParams) WithTimeout(timeout time.Duration) *AddServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add server params
func (o *AddServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add server params
func (o *AddServerParams) WithContext(ctx context.Context) *AddServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add server params
func (o *AddServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add server params
func (o *AddServerParams) WithHTTPClient(client *http.Client) *AddServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add server params
func (o *AddServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add server params
func (o *AddServerParams) WithBody(body *vproxy_client_model.ServerCreate) *AddServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add server params
func (o *AddServerParams) SetBody(body *vproxy_client_model.ServerCreate) {
	o.Body = body
}

// WithSg adds the sg to the add server params
func (o *AddServerParams) WithSg(sg string) *AddServerParams {
	o.SetSg(sg)
	return o
}

// SetSg adds the sg to the add server params
func (o *AddServerParams) SetSg(sg string) {
	o.Sg = sg
}

// WriteToRequest writes these params to a swagger request
func (o *AddServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param sg
	if err := r.SetPathParam("sg", o.Sg); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
