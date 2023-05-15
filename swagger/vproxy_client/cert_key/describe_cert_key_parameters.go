// Code generated by go-swagger; DO NOT EDIT.

package cert_key

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

// NewDescribeCertKeyParams creates a new DescribeCertKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeCertKeyParams() *DescribeCertKeyParams {
	return &DescribeCertKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeCertKeyParamsWithTimeout creates a new DescribeCertKeyParams object
// with the ability to set a timeout on a request.
func NewDescribeCertKeyParamsWithTimeout(timeout time.Duration) *DescribeCertKeyParams {
	return &DescribeCertKeyParams{
		timeout: timeout,
	}
}

// NewDescribeCertKeyParamsWithContext creates a new DescribeCertKeyParams object
// with the ability to set a context for a request.
func NewDescribeCertKeyParamsWithContext(ctx context.Context) *DescribeCertKeyParams {
	return &DescribeCertKeyParams{
		Context: ctx,
	}
}

// NewDescribeCertKeyParamsWithHTTPClient creates a new DescribeCertKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeCertKeyParamsWithHTTPClient(client *http.Client) *DescribeCertKeyParams {
	return &DescribeCertKeyParams{
		HTTPClient: client,
	}
}

/*
DescribeCertKeyParams contains all the parameters to send to the API endpoint

	for the describe cert key operation.

	Typically these are written to a http.Request.
*/
type DescribeCertKeyParams struct {

	/* Ck.

	   name of the cert-key
	*/
	Ck string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe cert key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeCertKeyParams) WithDefaults() *DescribeCertKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe cert key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeCertKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe cert key params
func (o *DescribeCertKeyParams) WithTimeout(timeout time.Duration) *DescribeCertKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe cert key params
func (o *DescribeCertKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe cert key params
func (o *DescribeCertKeyParams) WithContext(ctx context.Context) *DescribeCertKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe cert key params
func (o *DescribeCertKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe cert key params
func (o *DescribeCertKeyParams) WithHTTPClient(client *http.Client) *DescribeCertKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe cert key params
func (o *DescribeCertKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCk adds the ck to the describe cert key params
func (o *DescribeCertKeyParams) WithCk(ck string) *DescribeCertKeyParams {
	o.SetCk(ck)
	return o
}

// SetCk adds the ck to the describe cert key params
func (o *DescribeCertKeyParams) SetCk(ck string) {
	o.Ck = ck
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeCertKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ck
	if err := r.SetPathParam("ck", o.Ck); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}