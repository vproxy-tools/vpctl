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

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// NewAddUpstreamParams creates a new AddUpstreamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddUpstreamParams() *AddUpstreamParams {
	return &AddUpstreamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddUpstreamParamsWithTimeout creates a new AddUpstreamParams object
// with the ability to set a timeout on a request.
func NewAddUpstreamParamsWithTimeout(timeout time.Duration) *AddUpstreamParams {
	return &AddUpstreamParams{
		timeout: timeout,
	}
}

// NewAddUpstreamParamsWithContext creates a new AddUpstreamParams object
// with the ability to set a context for a request.
func NewAddUpstreamParamsWithContext(ctx context.Context) *AddUpstreamParams {
	return &AddUpstreamParams{
		Context: ctx,
	}
}

// NewAddUpstreamParamsWithHTTPClient creates a new AddUpstreamParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddUpstreamParamsWithHTTPClient(client *http.Client) *AddUpstreamParams {
	return &AddUpstreamParams{
		HTTPClient: client,
	}
}

/*
AddUpstreamParams contains all the parameters to send to the API endpoint

	for the add upstream operation.

	Typically these are written to a http.Request.
*/
type AddUpstreamParams struct {

	// Body.
	Body *vproxy_client_model.UpstreamCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddUpstreamParams) WithDefaults() *AddUpstreamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddUpstreamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add upstream params
func (o *AddUpstreamParams) WithTimeout(timeout time.Duration) *AddUpstreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add upstream params
func (o *AddUpstreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add upstream params
func (o *AddUpstreamParams) WithContext(ctx context.Context) *AddUpstreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add upstream params
func (o *AddUpstreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add upstream params
func (o *AddUpstreamParams) WithHTTPClient(client *http.Client) *AddUpstreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add upstream params
func (o *AddUpstreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add upstream params
func (o *AddUpstreamParams) WithBody(body *vproxy_client_model.UpstreamCreate) *AddUpstreamParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add upstream params
func (o *AddUpstreamParams) SetBody(body *vproxy_client_model.UpstreamCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddUpstreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}