// Code generated by go-swagger; DO NOT EDIT.

package tcp_lb

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

	vproxy_client_model "vproxy_client_model"
)

// NewUpdateTCPLbParams creates a new UpdateTCPLbParams object
// with the default values initialized.
func NewUpdateTCPLbParams() *UpdateTCPLbParams {
	var ()
	return &UpdateTCPLbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTCPLbParamsWithTimeout creates a new UpdateTCPLbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTCPLbParamsWithTimeout(timeout time.Duration) *UpdateTCPLbParams {
	var ()
	return &UpdateTCPLbParams{

		timeout: timeout,
	}
}

// NewUpdateTCPLbParamsWithContext creates a new UpdateTCPLbParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTCPLbParamsWithContext(ctx context.Context) *UpdateTCPLbParams {
	var ()
	return &UpdateTCPLbParams{

		Context: ctx,
	}
}

// NewUpdateTCPLbParamsWithHTTPClient creates a new UpdateTCPLbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTCPLbParamsWithHTTPClient(client *http.Client) *UpdateTCPLbParams {
	var ()
	return &UpdateTCPLbParams{
		HTTPClient: client,
	}
}

/*UpdateTCPLbParams contains all the parameters to send to the API endpoint
for the update Tcp lb operation typically these are written to a http.Request
*/
type UpdateTCPLbParams struct {

	/*Body*/
	Body *vproxy_client_model.TCPLbUpdate
	/*Tl
	  name of the tcp-lb

	*/
	Tl string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update Tcp lb params
func (o *UpdateTCPLbParams) WithTimeout(timeout time.Duration) *UpdateTCPLbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Tcp lb params
func (o *UpdateTCPLbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Tcp lb params
func (o *UpdateTCPLbParams) WithContext(ctx context.Context) *UpdateTCPLbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Tcp lb params
func (o *UpdateTCPLbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Tcp lb params
func (o *UpdateTCPLbParams) WithHTTPClient(client *http.Client) *UpdateTCPLbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Tcp lb params
func (o *UpdateTCPLbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update Tcp lb params
func (o *UpdateTCPLbParams) WithBody(body *vproxy_client_model.TCPLbUpdate) *UpdateTCPLbParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update Tcp lb params
func (o *UpdateTCPLbParams) SetBody(body *vproxy_client_model.TCPLbUpdate) {
	o.Body = body
}

// WithTl adds the tl to the update Tcp lb params
func (o *UpdateTCPLbParams) WithTl(tl string) *UpdateTCPLbParams {
	o.SetTl(tl)
	return o
}

// SetTl adds the tl to the update Tcp lb params
func (o *UpdateTCPLbParams) SetTl(tl string) {
	o.Tl = tl
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTCPLbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param tl
	if err := r.SetPathParam("tl", o.Tl); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
