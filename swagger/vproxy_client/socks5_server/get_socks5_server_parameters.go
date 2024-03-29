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

// NewGetSocks5ServerParams creates a new GetSocks5ServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSocks5ServerParams() *GetSocks5ServerParams {
	return &GetSocks5ServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSocks5ServerParamsWithTimeout creates a new GetSocks5ServerParams object
// with the ability to set a timeout on a request.
func NewGetSocks5ServerParamsWithTimeout(timeout time.Duration) *GetSocks5ServerParams {
	return &GetSocks5ServerParams{
		timeout: timeout,
	}
}

// NewGetSocks5ServerParamsWithContext creates a new GetSocks5ServerParams object
// with the ability to set a context for a request.
func NewGetSocks5ServerParamsWithContext(ctx context.Context) *GetSocks5ServerParams {
	return &GetSocks5ServerParams{
		Context: ctx,
	}
}

// NewGetSocks5ServerParamsWithHTTPClient creates a new GetSocks5ServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSocks5ServerParamsWithHTTPClient(client *http.Client) *GetSocks5ServerParams {
	return &GetSocks5ServerParams{
		HTTPClient: client,
	}
}

/*
GetSocks5ServerParams contains all the parameters to send to the API endpoint

	for the get socks5 server operation.

	Typically these are written to a http.Request.
*/
type GetSocks5ServerParams struct {

	/* Socks5.

	   name of the socks5-server
	*/
	Socks5 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get socks5 server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSocks5ServerParams) WithDefaults() *GetSocks5ServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get socks5 server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSocks5ServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get socks5 server params
func (o *GetSocks5ServerParams) WithTimeout(timeout time.Duration) *GetSocks5ServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get socks5 server params
func (o *GetSocks5ServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get socks5 server params
func (o *GetSocks5ServerParams) WithContext(ctx context.Context) *GetSocks5ServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get socks5 server params
func (o *GetSocks5ServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get socks5 server params
func (o *GetSocks5ServerParams) WithHTTPClient(client *http.Client) *GetSocks5ServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get socks5 server params
func (o *GetSocks5ServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSocks5 adds the socks5 to the get socks5 server params
func (o *GetSocks5ServerParams) WithSocks5(socks5 string) *GetSocks5ServerParams {
	o.SetSocks5(socks5)
	return o
}

// SetSocks5 adds the socks5 to the get socks5 server params
func (o *GetSocks5ServerParams) SetSocks5(socks5 string) {
	o.Socks5 = socks5
}

// WriteToRequest writes these params to a swagger request
func (o *GetSocks5ServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
