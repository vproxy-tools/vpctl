// Code generated by go-swagger; DO NOT EDIT.

package socks5_server

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

// NewDescribeSocks5ServerParams creates a new DescribeSocks5ServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeSocks5ServerParams() *DescribeSocks5ServerParams {
	return &DescribeSocks5ServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeSocks5ServerParamsWithTimeout creates a new DescribeSocks5ServerParams object
// with the ability to set a timeout on a request.
func NewDescribeSocks5ServerParamsWithTimeout(timeout time.Duration) *DescribeSocks5ServerParams {
	return &DescribeSocks5ServerParams{
		timeout: timeout,
	}
}

// NewDescribeSocks5ServerParamsWithContext creates a new DescribeSocks5ServerParams object
// with the ability to set a context for a request.
func NewDescribeSocks5ServerParamsWithContext(ctx context.Context) *DescribeSocks5ServerParams {
	return &DescribeSocks5ServerParams{
		Context: ctx,
	}
}

// NewDescribeSocks5ServerParamsWithHTTPClient creates a new DescribeSocks5ServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeSocks5ServerParamsWithHTTPClient(client *http.Client) *DescribeSocks5ServerParams {
	return &DescribeSocks5ServerParams{
		HTTPClient: client,
	}
}

/*
DescribeSocks5ServerParams contains all the parameters to send to the API endpoint

	for the describe socks5 server operation.

	Typically these are written to a http.Request.
*/
type DescribeSocks5ServerParams struct {

	/* Socks5.

	   name of the socks5-server
	*/
	Socks5 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe socks5 server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeSocks5ServerParams) WithDefaults() *DescribeSocks5ServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe socks5 server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeSocks5ServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe socks5 server params
func (o *DescribeSocks5ServerParams) WithTimeout(timeout time.Duration) *DescribeSocks5ServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe socks5 server params
func (o *DescribeSocks5ServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe socks5 server params
func (o *DescribeSocks5ServerParams) WithContext(ctx context.Context) *DescribeSocks5ServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe socks5 server params
func (o *DescribeSocks5ServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe socks5 server params
func (o *DescribeSocks5ServerParams) WithHTTPClient(client *http.Client) *DescribeSocks5ServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe socks5 server params
func (o *DescribeSocks5ServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSocks5 adds the socks5 to the describe socks5 server params
func (o *DescribeSocks5ServerParams) WithSocks5(socks5 string) *DescribeSocks5ServerParams {
	o.SetSocks5(socks5)
	return o
}

// SetSocks5 adds the socks5 to the describe socks5 server params
func (o *DescribeSocks5ServerParams) SetSocks5(socks5 string) {
	o.Socks5 = socks5
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeSocks5ServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param socks5
	if err := r.SetPathParam("socks5", o.Socks5); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
