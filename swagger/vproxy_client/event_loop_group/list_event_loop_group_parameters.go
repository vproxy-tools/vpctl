// Code generated by go-swagger; DO NOT EDIT.

package event_loop_group

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

// NewListEventLoopGroupParams creates a new ListEventLoopGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListEventLoopGroupParams() *ListEventLoopGroupParams {
	return &ListEventLoopGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListEventLoopGroupParamsWithTimeout creates a new ListEventLoopGroupParams object
// with the ability to set a timeout on a request.
func NewListEventLoopGroupParamsWithTimeout(timeout time.Duration) *ListEventLoopGroupParams {
	return &ListEventLoopGroupParams{
		timeout: timeout,
	}
}

// NewListEventLoopGroupParamsWithContext creates a new ListEventLoopGroupParams object
// with the ability to set a context for a request.
func NewListEventLoopGroupParamsWithContext(ctx context.Context) *ListEventLoopGroupParams {
	return &ListEventLoopGroupParams{
		Context: ctx,
	}
}

// NewListEventLoopGroupParamsWithHTTPClient creates a new ListEventLoopGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewListEventLoopGroupParamsWithHTTPClient(client *http.Client) *ListEventLoopGroupParams {
	return &ListEventLoopGroupParams{
		HTTPClient: client,
	}
}

/*
ListEventLoopGroupParams contains all the parameters to send to the API endpoint

	for the list event loop group operation.

	Typically these are written to a http.Request.
*/
type ListEventLoopGroupParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list event loop group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEventLoopGroupParams) WithDefaults() *ListEventLoopGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list event loop group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEventLoopGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list event loop group params
func (o *ListEventLoopGroupParams) WithTimeout(timeout time.Duration) *ListEventLoopGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list event loop group params
func (o *ListEventLoopGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list event loop group params
func (o *ListEventLoopGroupParams) WithContext(ctx context.Context) *ListEventLoopGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list event loop group params
func (o *ListEventLoopGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list event loop group params
func (o *ListEventLoopGroupParams) WithHTTPClient(client *http.Client) *ListEventLoopGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list event loop group params
func (o *ListEventLoopGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListEventLoopGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
