// Code generated by go-swagger; DO NOT EDIT.

package smart_group_delegate

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

// NewAddSmartGroupDelegateParams creates a new AddSmartGroupDelegateParams object
// with the default values initialized.
func NewAddSmartGroupDelegateParams() *AddSmartGroupDelegateParams {
	var ()
	return &AddSmartGroupDelegateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddSmartGroupDelegateParamsWithTimeout creates a new AddSmartGroupDelegateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddSmartGroupDelegateParamsWithTimeout(timeout time.Duration) *AddSmartGroupDelegateParams {
	var ()
	return &AddSmartGroupDelegateParams{

		timeout: timeout,
	}
}

// NewAddSmartGroupDelegateParamsWithContext creates a new AddSmartGroupDelegateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddSmartGroupDelegateParamsWithContext(ctx context.Context) *AddSmartGroupDelegateParams {
	var ()
	return &AddSmartGroupDelegateParams{

		Context: ctx,
	}
}

// NewAddSmartGroupDelegateParamsWithHTTPClient creates a new AddSmartGroupDelegateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddSmartGroupDelegateParamsWithHTTPClient(client *http.Client) *AddSmartGroupDelegateParams {
	var ()
	return &AddSmartGroupDelegateParams{
		HTTPClient: client,
	}
}

/*AddSmartGroupDelegateParams contains all the parameters to send to the API endpoint
for the add smart group delegate operation typically these are written to a http.Request
*/
type AddSmartGroupDelegateParams struct {

	/*Body*/
	Body *vproxy_client_model.SmartGroupDelegateCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add smart group delegate params
func (o *AddSmartGroupDelegateParams) WithTimeout(timeout time.Duration) *AddSmartGroupDelegateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add smart group delegate params
func (o *AddSmartGroupDelegateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add smart group delegate params
func (o *AddSmartGroupDelegateParams) WithContext(ctx context.Context) *AddSmartGroupDelegateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add smart group delegate params
func (o *AddSmartGroupDelegateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add smart group delegate params
func (o *AddSmartGroupDelegateParams) WithHTTPClient(client *http.Client) *AddSmartGroupDelegateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add smart group delegate params
func (o *AddSmartGroupDelegateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add smart group delegate params
func (o *AddSmartGroupDelegateParams) WithBody(body *vproxy_client_model.SmartGroupDelegateCreate) *AddSmartGroupDelegateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add smart group delegate params
func (o *AddSmartGroupDelegateParams) SetBody(body *vproxy_client_model.SmartGroupDelegateCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddSmartGroupDelegateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
