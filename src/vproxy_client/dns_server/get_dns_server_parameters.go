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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDNSServerParams creates a new GetDNSServerParams object
// with the default values initialized.
func NewGetDNSServerParams() *GetDNSServerParams {
	var ()
	return &GetDNSServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDNSServerParamsWithTimeout creates a new GetDNSServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDNSServerParamsWithTimeout(timeout time.Duration) *GetDNSServerParams {
	var ()
	return &GetDNSServerParams{

		timeout: timeout,
	}
}

// NewGetDNSServerParamsWithContext creates a new GetDNSServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDNSServerParamsWithContext(ctx context.Context) *GetDNSServerParams {
	var ()
	return &GetDNSServerParams{

		Context: ctx,
	}
}

// NewGetDNSServerParamsWithHTTPClient creates a new GetDNSServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDNSServerParamsWithHTTPClient(client *http.Client) *GetDNSServerParams {
	var ()
	return &GetDNSServerParams{
		HTTPClient: client,
	}
}

/*GetDNSServerParams contains all the parameters to send to the API endpoint
for the get Dns server operation typically these are written to a http.Request
*/
type GetDNSServerParams struct {

	/*DNS
	  name of the dns-server

	*/
	DNS string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Dns server params
func (o *GetDNSServerParams) WithTimeout(timeout time.Duration) *GetDNSServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Dns server params
func (o *GetDNSServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Dns server params
func (o *GetDNSServerParams) WithContext(ctx context.Context) *GetDNSServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Dns server params
func (o *GetDNSServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Dns server params
func (o *GetDNSServerParams) WithHTTPClient(client *http.Client) *GetDNSServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Dns server params
func (o *GetDNSServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDNS adds the dns to the get Dns server params
func (o *GetDNSServerParams) WithDNS(dns string) *GetDNSServerParams {
	o.SetDNS(dns)
	return o
}

// SetDNS adds the dns to the get Dns server params
func (o *GetDNSServerParams) SetDNS(dns string) {
	o.DNS = dns
}

// WriteToRequest writes these params to a swagger request
func (o *GetDNSServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
