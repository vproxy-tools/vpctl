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

// NewGetUpstreamParams creates a new GetUpstreamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUpstreamParams() *GetUpstreamParams {
	return &GetUpstreamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUpstreamParamsWithTimeout creates a new GetUpstreamParams object
// with the ability to set a timeout on a request.
func NewGetUpstreamParamsWithTimeout(timeout time.Duration) *GetUpstreamParams {
	return &GetUpstreamParams{
		timeout: timeout,
	}
}

// NewGetUpstreamParamsWithContext creates a new GetUpstreamParams object
// with the ability to set a context for a request.
func NewGetUpstreamParamsWithContext(ctx context.Context) *GetUpstreamParams {
	return &GetUpstreamParams{
		Context: ctx,
	}
}

// NewGetUpstreamParamsWithHTTPClient creates a new GetUpstreamParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUpstreamParamsWithHTTPClient(client *http.Client) *GetUpstreamParams {
	return &GetUpstreamParams{
		HTTPClient: client,
	}
}

/*
GetUpstreamParams contains all the parameters to send to the API endpoint

	for the get upstream operation.

	Typically these are written to a http.Request.
*/
type GetUpstreamParams struct {

	/* Ups.

	   name of the upstream
	*/
	Ups string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpstreamParams) WithDefaults() *GetUpstreamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpstreamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get upstream params
func (o *GetUpstreamParams) WithTimeout(timeout time.Duration) *GetUpstreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get upstream params
func (o *GetUpstreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get upstream params
func (o *GetUpstreamParams) WithContext(ctx context.Context) *GetUpstreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get upstream params
func (o *GetUpstreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get upstream params
func (o *GetUpstreamParams) WithHTTPClient(client *http.Client) *GetUpstreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get upstream params
func (o *GetUpstreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUps adds the ups to the get upstream params
func (o *GetUpstreamParams) WithUps(ups string) *GetUpstreamParams {
	o.SetUps(ups)
	return o
}

// SetUps adds the ups to the get upstream params
func (o *GetUpstreamParams) SetUps(ups string) {
	o.Ups = ups
}

// WriteToRequest writes these params to a swagger request
func (o *GetUpstreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
