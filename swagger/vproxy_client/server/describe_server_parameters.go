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
)

// NewDescribeServerParams creates a new DescribeServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeServerParams() *DescribeServerParams {
	return &DescribeServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeServerParamsWithTimeout creates a new DescribeServerParams object
// with the ability to set a timeout on a request.
func NewDescribeServerParamsWithTimeout(timeout time.Duration) *DescribeServerParams {
	return &DescribeServerParams{
		timeout: timeout,
	}
}

// NewDescribeServerParamsWithContext creates a new DescribeServerParams object
// with the ability to set a context for a request.
func NewDescribeServerParamsWithContext(ctx context.Context) *DescribeServerParams {
	return &DescribeServerParams{
		Context: ctx,
	}
}

// NewDescribeServerParamsWithHTTPClient creates a new DescribeServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeServerParamsWithHTTPClient(client *http.Client) *DescribeServerParams {
	return &DescribeServerParams{
		HTTPClient: client,
	}
}

/*
DescribeServerParams contains all the parameters to send to the API endpoint

	for the describe server operation.

	Typically these are written to a http.Request.
*/
type DescribeServerParams struct {

	/* Sg.

	   name of the server-group
	*/
	Sg string

	/* Svr.

	   name of the server
	*/
	Svr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeServerParams) WithDefaults() *DescribeServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe server params
func (o *DescribeServerParams) WithTimeout(timeout time.Duration) *DescribeServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe server params
func (o *DescribeServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe server params
func (o *DescribeServerParams) WithContext(ctx context.Context) *DescribeServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe server params
func (o *DescribeServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe server params
func (o *DescribeServerParams) WithHTTPClient(client *http.Client) *DescribeServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe server params
func (o *DescribeServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSg adds the sg to the describe server params
func (o *DescribeServerParams) WithSg(sg string) *DescribeServerParams {
	o.SetSg(sg)
	return o
}

// SetSg adds the sg to the describe server params
func (o *DescribeServerParams) SetSg(sg string) {
	o.Sg = sg
}

// WithSvr adds the svr to the describe server params
func (o *DescribeServerParams) WithSvr(svr string) *DescribeServerParams {
	o.SetSvr(svr)
	return o
}

// SetSvr adds the svr to the describe server params
func (o *DescribeServerParams) SetSvr(svr string) {
	o.Svr = svr
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sg
	if err := r.SetPathParam("sg", o.Sg); err != nil {
		return err
	}

	// path param svr
	if err := r.SetPathParam("svr", o.Svr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
