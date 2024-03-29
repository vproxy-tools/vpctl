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

// NewRemoveEventLoopParams creates a new RemoveEventLoopParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveEventLoopParams() *RemoveEventLoopParams {
	return &RemoveEventLoopParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveEventLoopParamsWithTimeout creates a new RemoveEventLoopParams object
// with the ability to set a timeout on a request.
func NewRemoveEventLoopParamsWithTimeout(timeout time.Duration) *RemoveEventLoopParams {
	return &RemoveEventLoopParams{
		timeout: timeout,
	}
}

// NewRemoveEventLoopParamsWithContext creates a new RemoveEventLoopParams object
// with the ability to set a context for a request.
func NewRemoveEventLoopParamsWithContext(ctx context.Context) *RemoveEventLoopParams {
	return &RemoveEventLoopParams{
		Context: ctx,
	}
}

// NewRemoveEventLoopParamsWithHTTPClient creates a new RemoveEventLoopParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveEventLoopParamsWithHTTPClient(client *http.Client) *RemoveEventLoopParams {
	return &RemoveEventLoopParams{
		HTTPClient: client,
	}
}

/*
RemoveEventLoopParams contains all the parameters to send to the API endpoint

	for the remove event loop operation.

	Typically these are written to a http.Request.
*/
type RemoveEventLoopParams struct {

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

// WithDefaults hydrates default values in the remove event loop params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveEventLoopParams) WithDefaults() *RemoveEventLoopParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove event loop params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveEventLoopParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove event loop params
func (o *RemoveEventLoopParams) WithTimeout(timeout time.Duration) *RemoveEventLoopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove event loop params
func (o *RemoveEventLoopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove event loop params
func (o *RemoveEventLoopParams) WithContext(ctx context.Context) *RemoveEventLoopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove event loop params
func (o *RemoveEventLoopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove event loop params
func (o *RemoveEventLoopParams) WithHTTPClient(client *http.Client) *RemoveEventLoopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove event loop params
func (o *RemoveEventLoopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEl adds the el to the remove event loop params
func (o *RemoveEventLoopParams) WithEl(el string) *RemoveEventLoopParams {
	o.SetEl(el)
	return o
}

// SetEl adds the el to the remove event loop params
func (o *RemoveEventLoopParams) SetEl(el string) {
	o.El = el
}

// WithElg adds the elg to the remove event loop params
func (o *RemoveEventLoopParams) WithElg(elg string) *RemoveEventLoopParams {
	o.SetElg(elg)
	return o
}

// SetElg adds the elg to the remove event loop params
func (o *RemoveEventLoopParams) SetElg(elg string) {
	o.Elg = elg
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveEventLoopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
