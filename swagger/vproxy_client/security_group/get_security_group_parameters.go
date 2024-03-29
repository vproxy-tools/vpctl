// Code generated by go-swagger; DO NOT EDIT.

package security_group

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

// NewGetSecurityGroupParams creates a new GetSecurityGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSecurityGroupParams() *GetSecurityGroupParams {
	return &GetSecurityGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecurityGroupParamsWithTimeout creates a new GetSecurityGroupParams object
// with the ability to set a timeout on a request.
func NewGetSecurityGroupParamsWithTimeout(timeout time.Duration) *GetSecurityGroupParams {
	return &GetSecurityGroupParams{
		timeout: timeout,
	}
}

// NewGetSecurityGroupParamsWithContext creates a new GetSecurityGroupParams object
// with the ability to set a context for a request.
func NewGetSecurityGroupParamsWithContext(ctx context.Context) *GetSecurityGroupParams {
	return &GetSecurityGroupParams{
		Context: ctx,
	}
}

// NewGetSecurityGroupParamsWithHTTPClient creates a new GetSecurityGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSecurityGroupParamsWithHTTPClient(client *http.Client) *GetSecurityGroupParams {
	return &GetSecurityGroupParams{
		HTTPClient: client,
	}
}

/*
GetSecurityGroupParams contains all the parameters to send to the API endpoint

	for the get security group operation.

	Typically these are written to a http.Request.
*/
type GetSecurityGroupParams struct {

	/* Secg.

	   name of the security-group
	*/
	Secg string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get security group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecurityGroupParams) WithDefaults() *GetSecurityGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get security group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecurityGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get security group params
func (o *GetSecurityGroupParams) WithTimeout(timeout time.Duration) *GetSecurityGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get security group params
func (o *GetSecurityGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get security group params
func (o *GetSecurityGroupParams) WithContext(ctx context.Context) *GetSecurityGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get security group params
func (o *GetSecurityGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get security group params
func (o *GetSecurityGroupParams) WithHTTPClient(client *http.Client) *GetSecurityGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get security group params
func (o *GetSecurityGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSecg adds the secg to the get security group params
func (o *GetSecurityGroupParams) WithSecg(secg string) *GetSecurityGroupParams {
	o.SetSecg(secg)
	return o
}

// SetSecg adds the secg to the get security group params
func (o *GetSecurityGroupParams) SetSecg(secg string) {
	o.Secg = secg
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecurityGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param secg
	if err := r.SetPathParam("secg", o.Secg); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
