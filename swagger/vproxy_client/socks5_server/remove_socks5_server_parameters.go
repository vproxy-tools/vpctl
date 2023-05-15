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

// NewRemoveSocks5ServerParams creates a new RemoveSocks5ServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveSocks5ServerParams() *RemoveSocks5ServerParams {
	return &RemoveSocks5ServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveSocks5ServerParamsWithTimeout creates a new RemoveSocks5ServerParams object
// with the ability to set a timeout on a request.
func NewRemoveSocks5ServerParamsWithTimeout(timeout time.Duration) *RemoveSocks5ServerParams {
	return &RemoveSocks5ServerParams{
		timeout: timeout,
	}
}

// NewRemoveSocks5ServerParamsWithContext creates a new RemoveSocks5ServerParams object
// with the ability to set a context for a request.
func NewRemoveSocks5ServerParamsWithContext(ctx context.Context) *RemoveSocks5ServerParams {
	return &RemoveSocks5ServerParams{
		Context: ctx,
	}
}

// NewRemoveSocks5ServerParamsWithHTTPClient creates a new RemoveSocks5ServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveSocks5ServerParamsWithHTTPClient(client *http.Client) *RemoveSocks5ServerParams {
	return &RemoveSocks5ServerParams{
		HTTPClient: client,
	}
}

/*
RemoveSocks5ServerParams contains all the parameters to send to the API endpoint

	for the remove socks5 server operation.

	Typically these are written to a http.Request.
*/
type RemoveSocks5ServerParams struct {

	/* Socks5.

	   name of the socks5-server
	*/
	Socks5 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove socks5 server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveSocks5ServerParams) WithDefaults() *RemoveSocks5ServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove socks5 server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveSocks5ServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove socks5 server params
func (o *RemoveSocks5ServerParams) WithTimeout(timeout time.Duration) *RemoveSocks5ServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove socks5 server params
func (o *RemoveSocks5ServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove socks5 server params
func (o *RemoveSocks5ServerParams) WithContext(ctx context.Context) *RemoveSocks5ServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove socks5 server params
func (o *RemoveSocks5ServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove socks5 server params
func (o *RemoveSocks5ServerParams) WithHTTPClient(client *http.Client) *RemoveSocks5ServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove socks5 server params
func (o *RemoveSocks5ServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSocks5 adds the socks5 to the remove socks5 server params
func (o *RemoveSocks5ServerParams) WithSocks5(socks5 string) *RemoveSocks5ServerParams {
	o.SetSocks5(socks5)
	return o
}

// SetSocks5 adds the socks5 to the remove socks5 server params
func (o *RemoveSocks5ServerParams) SetSocks5(socks5 string) {
	o.Socks5 = socks5
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveSocks5ServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
