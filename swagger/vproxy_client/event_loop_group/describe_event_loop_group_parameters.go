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

// NewDescribeEventLoopGroupParams creates a new DescribeEventLoopGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeEventLoopGroupParams() *DescribeEventLoopGroupParams {
	return &DescribeEventLoopGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeEventLoopGroupParamsWithTimeout creates a new DescribeEventLoopGroupParams object
// with the ability to set a timeout on a request.
func NewDescribeEventLoopGroupParamsWithTimeout(timeout time.Duration) *DescribeEventLoopGroupParams {
	return &DescribeEventLoopGroupParams{
		timeout: timeout,
	}
}

// NewDescribeEventLoopGroupParamsWithContext creates a new DescribeEventLoopGroupParams object
// with the ability to set a context for a request.
func NewDescribeEventLoopGroupParamsWithContext(ctx context.Context) *DescribeEventLoopGroupParams {
	return &DescribeEventLoopGroupParams{
		Context: ctx,
	}
}

// NewDescribeEventLoopGroupParamsWithHTTPClient creates a new DescribeEventLoopGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeEventLoopGroupParamsWithHTTPClient(client *http.Client) *DescribeEventLoopGroupParams {
	return &DescribeEventLoopGroupParams{
		HTTPClient: client,
	}
}

/*
DescribeEventLoopGroupParams contains all the parameters to send to the API endpoint

	for the describe event loop group operation.

	Typically these are written to a http.Request.
*/
type DescribeEventLoopGroupParams struct {

	/* Elg.

	   name of the event-loop-group
	*/
	Elg string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe event loop group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeEventLoopGroupParams) WithDefaults() *DescribeEventLoopGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe event loop group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeEventLoopGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe event loop group params
func (o *DescribeEventLoopGroupParams) WithTimeout(timeout time.Duration) *DescribeEventLoopGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe event loop group params
func (o *DescribeEventLoopGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe event loop group params
func (o *DescribeEventLoopGroupParams) WithContext(ctx context.Context) *DescribeEventLoopGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe event loop group params
func (o *DescribeEventLoopGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe event loop group params
func (o *DescribeEventLoopGroupParams) WithHTTPClient(client *http.Client) *DescribeEventLoopGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe event loop group params
func (o *DescribeEventLoopGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithElg adds the elg to the describe event loop group params
func (o *DescribeEventLoopGroupParams) WithElg(elg string) *DescribeEventLoopGroupParams {
	o.SetElg(elg)
	return o
}

// SetElg adds the elg to the describe event loop group params
func (o *DescribeEventLoopGroupParams) SetElg(elg string) {
	o.Elg = elg
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeEventLoopGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param elg
	if err := r.SetPathParam("elg", o.Elg); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
