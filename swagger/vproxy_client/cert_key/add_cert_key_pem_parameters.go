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

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// NewAddCertKeyPemParams creates a new AddCertKeyPemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddCertKeyPemParams() *AddCertKeyPemParams {
	return &AddCertKeyPemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddCertKeyPemParamsWithTimeout creates a new AddCertKeyPemParams object
// with the ability to set a timeout on a request.
func NewAddCertKeyPemParamsWithTimeout(timeout time.Duration) *AddCertKeyPemParams {
	return &AddCertKeyPemParams{
		timeout: timeout,
	}
}

// NewAddCertKeyPemParamsWithContext creates a new AddCertKeyPemParams object
// with the ability to set a context for a request.
func NewAddCertKeyPemParamsWithContext(ctx context.Context) *AddCertKeyPemParams {
	return &AddCertKeyPemParams{
		Context: ctx,
	}
}

// NewAddCertKeyPemParamsWithHTTPClient creates a new AddCertKeyPemParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddCertKeyPemParamsWithHTTPClient(client *http.Client) *AddCertKeyPemParams {
	return &AddCertKeyPemParams{
		HTTPClient: client,
	}
}

/*
AddCertKeyPemParams contains all the parameters to send to the API endpoint

	for the add cert key pem operation.

	Typically these are written to a http.Request.
*/
type AddCertKeyPemParams struct {

	// Body.
	Body *vproxy_client_model.CertKeyCreatePem

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add cert key pem params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddCertKeyPemParams) WithDefaults() *AddCertKeyPemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add cert key pem params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddCertKeyPemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add cert key pem params
func (o *AddCertKeyPemParams) WithTimeout(timeout time.Duration) *AddCertKeyPemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add cert key pem params
func (o *AddCertKeyPemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add cert key pem params
func (o *AddCertKeyPemParams) WithContext(ctx context.Context) *AddCertKeyPemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add cert key pem params
func (o *AddCertKeyPemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add cert key pem params
func (o *AddCertKeyPemParams) WithHTTPClient(client *http.Client) *AddCertKeyPemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add cert key pem params
func (o *AddCertKeyPemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add cert key pem params
func (o *AddCertKeyPemParams) WithBody(body *vproxy_client_model.CertKeyCreatePem) *AddCertKeyPemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add cert key pem params
func (o *AddCertKeyPemParams) SetBody(body *vproxy_client_model.CertKeyCreatePem) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddCertKeyPemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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