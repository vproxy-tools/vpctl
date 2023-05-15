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
)

// NewDescribeDNSServerParams creates a new DescribeDNSServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeDNSServerParams() *DescribeDNSServerParams {
	return &DescribeDNSServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeDNSServerParamsWithTimeout creates a new DescribeDNSServerParams object
// with the ability to set a timeout on a request.
func NewDescribeDNSServerParamsWithTimeout(timeout time.Duration) *DescribeDNSServerParams {
	return &DescribeDNSServerParams{
		timeout: timeout,
	}
}

// NewDescribeDNSServerParamsWithContext creates a new DescribeDNSServerParams object
// with the ability to set a context for a request.
func NewDescribeDNSServerParamsWithContext(ctx context.Context) *DescribeDNSServerParams {
	return &DescribeDNSServerParams{
		Context: ctx,
	}
}

// NewDescribeDNSServerParamsWithHTTPClient creates a new DescribeDNSServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeDNSServerParamsWithHTTPClient(client *http.Client) *DescribeDNSServerParams {
	return &DescribeDNSServerParams{
		HTTPClient: client,
	}
}

/*
DescribeDNSServerParams contains all the parameters to send to the API endpoint

	for the describe Dns server operation.

	Typically these are written to a http.Request.
*/
type DescribeDNSServerParams struct {

	/* DNS.

	   name of the dns-server
	*/
	DNS string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe Dns server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeDNSServerParams) WithDefaults() *DescribeDNSServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe Dns server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeDNSServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe Dns server params
func (o *DescribeDNSServerParams) WithTimeout(timeout time.Duration) *DescribeDNSServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe Dns server params
func (o *DescribeDNSServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe Dns server params
func (o *DescribeDNSServerParams) WithContext(ctx context.Context) *DescribeDNSServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe Dns server params
func (o *DescribeDNSServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe Dns server params
func (o *DescribeDNSServerParams) WithHTTPClient(client *http.Client) *DescribeDNSServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe Dns server params
func (o *DescribeDNSServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDNS adds the dns to the describe Dns server params
func (o *DescribeDNSServerParams) WithDNS(dns string) *DescribeDNSServerParams {
	o.SetDNS(dns)
	return o
}

// SetDNS adds the dns to the describe Dns server params
func (o *DescribeDNSServerParams) SetDNS(dns string) {
	o.DNS = dns
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeDNSServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dns
	if err := r.SetPathParam("dns", o.DNS); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}