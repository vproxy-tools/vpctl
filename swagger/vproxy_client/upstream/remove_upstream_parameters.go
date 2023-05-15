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

// NewRemoveUpstreamParams creates a new RemoveUpstreamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveUpstreamParams() *RemoveUpstreamParams {
	return &RemoveUpstreamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveUpstreamParamsWithTimeout creates a new RemoveUpstreamParams object
// with the ability to set a timeout on a request.
func NewRemoveUpstreamParamsWithTimeout(timeout time.Duration) *RemoveUpstreamParams {
	return &RemoveUpstreamParams{
		timeout: timeout,
	}
}

// NewRemoveUpstreamParamsWithContext creates a new RemoveUpstreamParams object
// with the ability to set a context for a request.
func NewRemoveUpstreamParamsWithContext(ctx context.Context) *RemoveUpstreamParams {
	return &RemoveUpstreamParams{
		Context: ctx,
	}
}

// NewRemoveUpstreamParamsWithHTTPClient creates a new RemoveUpstreamParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveUpstreamParamsWithHTTPClient(client *http.Client) *RemoveUpstreamParams {
	return &RemoveUpstreamParams{
		HTTPClient: client,
	}
}

/*
RemoveUpstreamParams contains all the parameters to send to the API endpoint

	for the remove upstream operation.

	Typically these are written to a http.Request.
*/
type RemoveUpstreamParams struct {

	/* Ups.

	   name of the upstream
	*/
	Ups string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveUpstreamParams) WithDefaults() *RemoveUpstreamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveUpstreamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove upstream params
func (o *RemoveUpstreamParams) WithTimeout(timeout time.Duration) *RemoveUpstreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove upstream params
func (o *RemoveUpstreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove upstream params
func (o *RemoveUpstreamParams) WithContext(ctx context.Context) *RemoveUpstreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove upstream params
func (o *RemoveUpstreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove upstream params
func (o *RemoveUpstreamParams) WithHTTPClient(client *http.Client) *RemoveUpstreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove upstream params
func (o *RemoveUpstreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUps adds the ups to the remove upstream params
func (o *RemoveUpstreamParams) WithUps(ups string) *RemoveUpstreamParams {
	o.SetUps(ups)
	return o
}

// SetUps adds the ups to the remove upstream params
func (o *RemoveUpstreamParams) SetUps(ups string) {
	o.Ups = ups
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveUpstreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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