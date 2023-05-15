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

// NewDescribeTCPLbParams creates a new DescribeTCPLbParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeTCPLbParams() *DescribeTCPLbParams {
	return &DescribeTCPLbParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeTCPLbParamsWithTimeout creates a new DescribeTCPLbParams object
// with the ability to set a timeout on a request.
func NewDescribeTCPLbParamsWithTimeout(timeout time.Duration) *DescribeTCPLbParams {
	return &DescribeTCPLbParams{
		timeout: timeout,
	}
}

// NewDescribeTCPLbParamsWithContext creates a new DescribeTCPLbParams object
// with the ability to set a context for a request.
func NewDescribeTCPLbParamsWithContext(ctx context.Context) *DescribeTCPLbParams {
	return &DescribeTCPLbParams{
		Context: ctx,
	}
}

// NewDescribeTCPLbParamsWithHTTPClient creates a new DescribeTCPLbParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeTCPLbParamsWithHTTPClient(client *http.Client) *DescribeTCPLbParams {
	return &DescribeTCPLbParams{
		HTTPClient: client,
	}
}

/*
DescribeTCPLbParams contains all the parameters to send to the API endpoint

	for the describe Tcp lb operation.

	Typically these are written to a http.Request.
*/
type DescribeTCPLbParams struct {

	/* Tl.

	   name of the tcp-lb
	*/
	Tl string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe Tcp lb params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeTCPLbParams) WithDefaults() *DescribeTCPLbParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe Tcp lb params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeTCPLbParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe Tcp lb params
func (o *DescribeTCPLbParams) WithTimeout(timeout time.Duration) *DescribeTCPLbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe Tcp lb params
func (o *DescribeTCPLbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe Tcp lb params
func (o *DescribeTCPLbParams) WithContext(ctx context.Context) *DescribeTCPLbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe Tcp lb params
func (o *DescribeTCPLbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe Tcp lb params
func (o *DescribeTCPLbParams) WithHTTPClient(client *http.Client) *DescribeTCPLbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe Tcp lb params
func (o *DescribeTCPLbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTl adds the tl to the describe Tcp lb params
func (o *DescribeTCPLbParams) WithTl(tl string) *DescribeTCPLbParams {
	o.SetTl(tl)
	return o
}

// SetTl adds the tl to the describe Tcp lb params
func (o *DescribeTCPLbParams) SetTl(tl string) {
	o.Tl = tl
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeTCPLbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
