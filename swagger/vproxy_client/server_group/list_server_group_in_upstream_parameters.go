// Code generated by go-swagger; DO NOT EDIT.

package server_group

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

// NewListServerGroupInUpstreamParams creates a new ListServerGroupInUpstreamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListServerGroupInUpstreamParams() *ListServerGroupInUpstreamParams {
	return &ListServerGroupInUpstreamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListServerGroupInUpstreamParamsWithTimeout creates a new ListServerGroupInUpstreamParams object
// with the ability to set a timeout on a request.
func NewListServerGroupInUpstreamParamsWithTimeout(timeout time.Duration) *ListServerGroupInUpstreamParams {
	return &ListServerGroupInUpstreamParams{
		timeout: timeout,
	}
}

// NewListServerGroupInUpstreamParamsWithContext creates a new ListServerGroupInUpstreamParams object
// with the ability to set a context for a request.
func NewListServerGroupInUpstreamParamsWithContext(ctx context.Context) *ListServerGroupInUpstreamParams {
	return &ListServerGroupInUpstreamParams{
		Context: ctx,
	}
}

// NewListServerGroupInUpstreamParamsWithHTTPClient creates a new ListServerGroupInUpstreamParams object
// with the ability to set a custom HTTPClient for a request.
func NewListServerGroupInUpstreamParamsWithHTTPClient(client *http.Client) *ListServerGroupInUpstreamParams {
	return &ListServerGroupInUpstreamParams{
		HTTPClient: client,
	}
}

/*
ListServerGroupInUpstreamParams contains all the parameters to send to the API endpoint

	for the list server group in upstream operation.

	Typically these are written to a http.Request.
*/
type ListServerGroupInUpstreamParams struct {

	/* Ups.

	   name of the upstream
	*/
	Ups string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list server group in upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServerGroupInUpstreamParams) WithDefaults() *ListServerGroupInUpstreamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list server group in upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServerGroupInUpstreamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list server group in upstream params
func (o *ListServerGroupInUpstreamParams) WithTimeout(timeout time.Duration) *ListServerGroupInUpstreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list server group in upstream params
func (o *ListServerGroupInUpstreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list server group in upstream params
func (o *ListServerGroupInUpstreamParams) WithContext(ctx context.Context) *ListServerGroupInUpstreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list server group in upstream params
func (o *ListServerGroupInUpstreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list server group in upstream params
func (o *ListServerGroupInUpstreamParams) WithHTTPClient(client *http.Client) *ListServerGroupInUpstreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list server group in upstream params
func (o *ListServerGroupInUpstreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUps adds the ups to the list server group in upstream params
func (o *ListServerGroupInUpstreamParams) WithUps(ups string) *ListServerGroupInUpstreamParams {
	o.SetUps(ups)
	return o
}

// SetUps adds the ups to the list server group in upstream params
func (o *ListServerGroupInUpstreamParams) SetUps(ups string) {
	o.Ups = ups
}

// WriteToRequest writes these params to a swagger request
func (o *ListServerGroupInUpstreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
