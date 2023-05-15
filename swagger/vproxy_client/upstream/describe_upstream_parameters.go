// Code generated by go-swagger; DO NOT EDIT.

package upstream

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

// NewDescribeUpstreamParams creates a new DescribeUpstreamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeUpstreamParams() *DescribeUpstreamParams {
	return &DescribeUpstreamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeUpstreamParamsWithTimeout creates a new DescribeUpstreamParams object
// with the ability to set a timeout on a request.
func NewDescribeUpstreamParamsWithTimeout(timeout time.Duration) *DescribeUpstreamParams {
	return &DescribeUpstreamParams{
		timeout: timeout,
	}
}

// NewDescribeUpstreamParamsWithContext creates a new DescribeUpstreamParams object
// with the ability to set a context for a request.
func NewDescribeUpstreamParamsWithContext(ctx context.Context) *DescribeUpstreamParams {
	return &DescribeUpstreamParams{
		Context: ctx,
	}
}

// NewDescribeUpstreamParamsWithHTTPClient creates a new DescribeUpstreamParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeUpstreamParamsWithHTTPClient(client *http.Client) *DescribeUpstreamParams {
	return &DescribeUpstreamParams{
		HTTPClient: client,
	}
}

/*
DescribeUpstreamParams contains all the parameters to send to the API endpoint

	for the describe upstream operation.

	Typically these are written to a http.Request.
*/
type DescribeUpstreamParams struct {

	/* Ups.

	   name of the upstream
	*/
	Ups string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeUpstreamParams) WithDefaults() *DescribeUpstreamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeUpstreamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe upstream params
func (o *DescribeUpstreamParams) WithTimeout(timeout time.Duration) *DescribeUpstreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe upstream params
func (o *DescribeUpstreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe upstream params
func (o *DescribeUpstreamParams) WithContext(ctx context.Context) *DescribeUpstreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe upstream params
func (o *DescribeUpstreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe upstream params
func (o *DescribeUpstreamParams) WithHTTPClient(client *http.Client) *DescribeUpstreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe upstream params
func (o *DescribeUpstreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUps adds the ups to the describe upstream params
func (o *DescribeUpstreamParams) WithUps(ups string) *DescribeUpstreamParams {
	o.SetUps(ups)
	return o
}

// SetUps adds the ups to the describe upstream params
func (o *DescribeUpstreamParams) SetUps(ups string) {
	o.Ups = ups
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeUpstreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ups
	if err := r.SetPathParam("ups", o.Ups); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
