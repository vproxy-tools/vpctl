// Code generated by go-swagger; DO NOT EDIT.

package dns_server

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

// NewAddDNSServerParams creates a new AddDNSServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDNSServerParams() *AddDNSServerParams {
	return &AddDNSServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDNSServerParamsWithTimeout creates a new AddDNSServerParams object
// with the ability to set a timeout on a request.
func NewAddDNSServerParamsWithTimeout(timeout time.Duration) *AddDNSServerParams {
	return &AddDNSServerParams{
		timeout: timeout,
	}
}

// NewAddDNSServerParamsWithContext creates a new AddDNSServerParams object
// with the ability to set a context for a request.
func NewAddDNSServerParamsWithContext(ctx context.Context) *AddDNSServerParams {
	return &AddDNSServerParams{
		Context: ctx,
	}
}

// NewAddDNSServerParamsWithHTTPClient creates a new AddDNSServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDNSServerParamsWithHTTPClient(client *http.Client) *AddDNSServerParams {
	return &AddDNSServerParams{
		HTTPClient: client,
	}
}

/*
AddDNSServerParams contains all the parameters to send to the API endpoint

	for the add Dns server operation.

	Typically these are written to a http.Request.
*/
type AddDNSServerParams struct {

	// Body.
	Body *vproxy_client_model.DNSServerCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add Dns server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDNSServerParams) WithDefaults() *AddDNSServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add Dns server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDNSServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add Dns server params
func (o *AddDNSServerParams) WithTimeout(timeout time.Duration) *AddDNSServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Dns server params
func (o *AddDNSServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Dns server params
func (o *AddDNSServerParams) WithContext(ctx context.Context) *AddDNSServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Dns server params
func (o *AddDNSServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Dns server params
func (o *AddDNSServerParams) WithHTTPClient(client *http.Client) *AddDNSServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Dns server params
func (o *AddDNSServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add Dns server params
func (o *AddDNSServerParams) WithBody(body *vproxy_client_model.DNSServerCreate) *AddDNSServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add Dns server params
func (o *AddDNSServerParams) SetBody(body *vproxy_client_model.DNSServerCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddDNSServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
