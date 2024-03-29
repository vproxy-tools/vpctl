// Code generated by go-swagger; DO NOT EDIT.

package event_loop

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

// NewGetEventLoopParams creates a new GetEventLoopParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEventLoopParams() *GetEventLoopParams {
	return &GetEventLoopParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventLoopParamsWithTimeout creates a new GetEventLoopParams object
// with the ability to set a timeout on a request.
func NewGetEventLoopParamsWithTimeout(timeout time.Duration) *GetEventLoopParams {
	return &GetEventLoopParams{
		timeout: timeout,
	}
}

// NewGetEventLoopParamsWithContext creates a new GetEventLoopParams object
// with the ability to set a context for a request.
func NewGetEventLoopParamsWithContext(ctx context.Context) *GetEventLoopParams {
	return &GetEventLoopParams{
		Context: ctx,
	}
}

// NewGetEventLoopParamsWithHTTPClient creates a new GetEventLoopParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEventLoopParamsWithHTTPClient(client *http.Client) *GetEventLoopParams {
	return &GetEventLoopParams{
		HTTPClient: client,
	}
}

/*
GetEventLoopParams contains all the parameters to send to the API endpoint

	for the get event loop operation.

	Typically these are written to a http.Request.
*/
type GetEventLoopParams struct {

	/* El.

	   name of the event-loop
	*/
	El string

	/* Elg.

	   name of the event-loop-group
	*/
	Elg string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get event loop params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEventLoopParams) WithDefaults() *GetEventLoopParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get event loop params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEventLoopParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get event loop params
func (o *GetEventLoopParams) WithTimeout(timeout time.Duration) *GetEventLoopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get event loop params
func (o *GetEventLoopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get event loop params
func (o *GetEventLoopParams) WithContext(ctx context.Context) *GetEventLoopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get event loop params
func (o *GetEventLoopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get event loop params
func (o *GetEventLoopParams) WithHTTPClient(client *http.Client) *GetEventLoopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get event loop params
func (o *GetEventLoopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEl adds the el to the get event loop params
func (o *GetEventLoopParams) WithEl(el string) *GetEventLoopParams {
	o.SetEl(el)
	return o
}

// SetEl adds the el to the get event loop params
func (o *GetEventLoopParams) SetEl(el string) {
	o.El = el
}

// WithElg adds the elg to the get event loop params
func (o *GetEventLoopParams) WithElg(elg string) *GetEventLoopParams {
	o.SetElg(elg)
	return o
}

// SetElg adds the elg to the get event loop params
func (o *GetEventLoopParams) SetElg(elg string) {
	o.Elg = elg
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventLoopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param el
	if err := r.SetPathParam("el", o.El); err != nil {
		return err
	}

	// path param elg
	if err := r.SetPathParam("elg", o.Elg); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
