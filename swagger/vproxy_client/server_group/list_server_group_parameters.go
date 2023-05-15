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

// NewListServerGroupParams creates a new ListServerGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListServerGroupParams() *ListServerGroupParams {
	return &ListServerGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListServerGroupParamsWithTimeout creates a new ListServerGroupParams object
// with the ability to set a timeout on a request.
func NewListServerGroupParamsWithTimeout(timeout time.Duration) *ListServerGroupParams {
	return &ListServerGroupParams{
		timeout: timeout,
	}
}

// NewListServerGroupParamsWithContext creates a new ListServerGroupParams object
// with the ability to set a context for a request.
func NewListServerGroupParamsWithContext(ctx context.Context) *ListServerGroupParams {
	return &ListServerGroupParams{
		Context: ctx,
	}
}

// NewListServerGroupParamsWithHTTPClient creates a new ListServerGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewListServerGroupParamsWithHTTPClient(client *http.Client) *ListServerGroupParams {
	return &ListServerGroupParams{
		HTTPClient: client,
	}
}

/*
ListServerGroupParams contains all the parameters to send to the API endpoint

	for the list server group operation.

	Typically these are written to a http.Request.
*/
type ListServerGroupParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list server group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServerGroupParams) WithDefaults() *ListServerGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list server group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServerGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list server group params
func (o *ListServerGroupParams) WithTimeout(timeout time.Duration) *ListServerGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list server group params
func (o *ListServerGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list server group params
func (o *ListServerGroupParams) WithContext(ctx context.Context) *ListServerGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list server group params
func (o *ListServerGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list server group params
func (o *ListServerGroupParams) WithHTTPClient(client *http.Client) *ListServerGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list server group params
func (o *ListServerGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListServerGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}