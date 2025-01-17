// Code generated by go-swagger; DO NOT EDIT.

package sip_bindings

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

// NewThinkteluControlAPIWebServiceListSIPBindingsParams creates a new ThinkteluControlAPIWebServiceListSIPBindingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceListSIPBindingsParams() *ThinkteluControlAPIWebServiceListSIPBindingsParams {
	return &ThinkteluControlAPIWebServiceListSIPBindingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceListSIPBindingsParamsWithTimeout creates a new ThinkteluControlAPIWebServiceListSIPBindingsParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceListSIPBindingsParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListSIPBindingsParams {
	return &ThinkteluControlAPIWebServiceListSIPBindingsParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceListSIPBindingsParamsWithContext creates a new ThinkteluControlAPIWebServiceListSIPBindingsParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceListSIPBindingsParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListSIPBindingsParams {
	return &ThinkteluControlAPIWebServiceListSIPBindingsParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceListSIPBindingsParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceListSIPBindingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceListSIPBindingsParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListSIPBindingsParams {
	return &ThinkteluControlAPIWebServiceListSIPBindingsParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceListSIPBindingsParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service list Sip bindings operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceListSIPBindingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service list Sip bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListSIPBindingsParams) WithDefaults() *ThinkteluControlAPIWebServiceListSIPBindingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service list Sip bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListSIPBindingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service list Sip bindings params
func (o *ThinkteluControlAPIWebServiceListSIPBindingsParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListSIPBindingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service list Sip bindings params
func (o *ThinkteluControlAPIWebServiceListSIPBindingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service list Sip bindings params
func (o *ThinkteluControlAPIWebServiceListSIPBindingsParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListSIPBindingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service list Sip bindings params
func (o *ThinkteluControlAPIWebServiceListSIPBindingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service list Sip bindings params
func (o *ThinkteluControlAPIWebServiceListSIPBindingsParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListSIPBindingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service list Sip bindings params
func (o *ThinkteluControlAPIWebServiceListSIPBindingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceListSIPBindingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
