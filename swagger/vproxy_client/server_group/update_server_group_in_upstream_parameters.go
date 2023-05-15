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

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// NewUpdateServerGroupInUpstreamParams creates a new UpdateServerGroupInUpstreamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateServerGroupInUpstreamParams() *UpdateServerGroupInUpstreamParams {
	return &UpdateServerGroupInUpstreamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServerGroupInUpstreamParamsWithTimeout creates a new UpdateServerGroupInUpstreamParams object
// with the ability to set a timeout on a request.
func NewUpdateServerGroupInUpstreamParamsWithTimeout(timeout time.Duration) *UpdateServerGroupInUpstreamParams {
	return &UpdateServerGroupInUpstreamParams{
		timeout: timeout,
	}
}

// NewUpdateServerGroupInUpstreamParamsWithContext creates a new UpdateServerGroupInUpstreamParams object
// with the ability to set a context for a request.
func NewUpdateServerGroupInUpstreamParamsWithContext(ctx context.Context) *UpdateServerGroupInUpstreamParams {
	return &UpdateServerGroupInUpstreamParams{
		Context: ctx,
	}
}

// NewUpdateServerGroupInUpstreamParamsWithHTTPClient creates a new UpdateServerGroupInUpstreamParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateServerGroupInUpstreamParamsWithHTTPClient(client *http.Client) *UpdateServerGroupInUpstreamParams {
	return &UpdateServerGroupInUpstreamParams{
		HTTPClient: client,
	}
}

/*
UpdateServerGroupInUpstreamParams contains all the parameters to send to the API endpoint

	for the update server group in upstream operation.

	Typically these are written to a http.Request.
*/
type UpdateServerGroupInUpstreamParams struct {

	// Body.
	Body *vproxy_client_model.ServerGroupInUpstreamUpdate

	/* Sg.

	   name of the server-group
	*/
	Sg string

	/* Ups.

	   name of the upstream
	*/
	Ups string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update server group in upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServerGroupInUpstreamParams) WithDefaults() *UpdateServerGroupInUpstreamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update server group in upstream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServerGroupInUpstreamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) WithTimeout(timeout time.Duration) *UpdateServerGroupInUpstreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) WithContext(ctx context.Context) *UpdateServerGroupInUpstreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) WithHTTPClient(client *http.Client) *UpdateServerGroupInUpstreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) WithBody(body *vproxy_client_model.ServerGroupInUpstreamUpdate) *UpdateServerGroupInUpstreamParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) SetBody(body *vproxy_client_model.ServerGroupInUpstreamUpdate) {
	o.Body = body
}

// WithSg adds the sg to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) WithSg(sg string) *UpdateServerGroupInUpstreamParams {
	o.SetSg(sg)
	return o
}

// SetSg adds the sg to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) SetSg(sg string) {
	o.Sg = sg
}

// WithUps adds the ups to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) WithUps(ups string) *UpdateServerGroupInUpstreamParams {
	o.SetUps(ups)
	return o
}

// SetUps adds the ups to the update server group in upstream params
func (o *UpdateServerGroupInUpstreamParams) SetUps(ups string) {
	o.Ups = ups
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServerGroupInUpstreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param sg
	if err := r.SetPathParam("sg", o.Sg); err != nil {
		return err
	}

	// path param ups
	if err := r.SetPathParam("ups", o.Ups); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
