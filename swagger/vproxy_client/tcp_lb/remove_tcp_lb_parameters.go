// Code generated by go-swagger; DO NOT EDIT.

package tcp_lb

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

// NewRemoveTCPLbParams creates a new RemoveTCPLbParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveTCPLbParams() *RemoveTCPLbParams {
	return &RemoveTCPLbParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveTCPLbParamsWithTimeout creates a new RemoveTCPLbParams object
// with the ability to set a timeout on a request.
func NewRemoveTCPLbParamsWithTimeout(timeout time.Duration) *RemoveTCPLbParams {
	return &RemoveTCPLbParams{
		timeout: timeout,
	}
}

// NewRemoveTCPLbParamsWithContext creates a new RemoveTCPLbParams object
// with the ability to set a context for a request.
func NewRemoveTCPLbParamsWithContext(ctx context.Context) *RemoveTCPLbParams {
	return &RemoveTCPLbParams{
		Context: ctx,
	}
}

// NewRemoveTCPLbParamsWithHTTPClient creates a new RemoveTCPLbParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveTCPLbParamsWithHTTPClient(client *http.Client) *RemoveTCPLbParams {
	return &RemoveTCPLbParams{
		HTTPClient: client,
	}
}

/*
RemoveTCPLbParams contains all the parameters to send to the API endpoint

	for the remove Tcp lb operation.

	Typically these are written to a http.Request.
*/
type RemoveTCPLbParams struct {

	/* Tl.

	   name of the tcp-lb
	*/
	Tl string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove Tcp lb params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTCPLbParams) WithDefaults() *RemoveTCPLbParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove Tcp lb params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTCPLbParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove Tcp lb params
func (o *RemoveTCPLbParams) WithTimeout(timeout time.Duration) *RemoveTCPLbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove Tcp lb params
func (o *RemoveTCPLbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove Tcp lb params
func (o *RemoveTCPLbParams) WithContext(ctx context.Context) *RemoveTCPLbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove Tcp lb params
func (o *RemoveTCPLbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove Tcp lb params
func (o *RemoveTCPLbParams) WithHTTPClient(client *http.Client) *RemoveTCPLbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove Tcp lb params
func (o *RemoveTCPLbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTl adds the tl to the remove Tcp lb params
func (o *RemoveTCPLbParams) WithTl(tl string) *RemoveTCPLbParams {
	o.SetTl(tl)
	return o
}

// SetTl adds the tl to the remove Tcp lb params
func (o *RemoveTCPLbParams) SetTl(tl string) {
	o.Tl = tl
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveTCPLbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tl
	if err := r.SetPathParam("tl", o.Tl); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
