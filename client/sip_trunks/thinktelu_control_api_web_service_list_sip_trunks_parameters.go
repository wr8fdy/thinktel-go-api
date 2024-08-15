// Code generated by go-swagger; DO NOT EDIT.

package sip_trunks

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

// NewThinkteluControlAPIWebServiceListSIPTrunksParams creates a new ThinkteluControlAPIWebServiceListSIPTrunksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceListSIPTrunksParams() *ThinkteluControlAPIWebServiceListSIPTrunksParams {
	return &ThinkteluControlAPIWebServiceListSIPTrunksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceListSIPTrunksParamsWithTimeout creates a new ThinkteluControlAPIWebServiceListSIPTrunksParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceListSIPTrunksParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListSIPTrunksParams {
	return &ThinkteluControlAPIWebServiceListSIPTrunksParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceListSIPTrunksParamsWithContext creates a new ThinkteluControlAPIWebServiceListSIPTrunksParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceListSIPTrunksParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListSIPTrunksParams {
	return &ThinkteluControlAPIWebServiceListSIPTrunksParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceListSIPTrunksParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceListSIPTrunksParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceListSIPTrunksParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListSIPTrunksParams {
	return &ThinkteluControlAPIWebServiceListSIPTrunksParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceListSIPTrunksParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service list Sip trunks operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceListSIPTrunksParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service list Sip trunks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListSIPTrunksParams) WithDefaults() *ThinkteluControlAPIWebServiceListSIPTrunksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service list Sip trunks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListSIPTrunksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service list Sip trunks params
func (o *ThinkteluControlAPIWebServiceListSIPTrunksParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListSIPTrunksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service list Sip trunks params
func (o *ThinkteluControlAPIWebServiceListSIPTrunksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service list Sip trunks params
func (o *ThinkteluControlAPIWebServiceListSIPTrunksParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListSIPTrunksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service list Sip trunks params
func (o *ThinkteluControlAPIWebServiceListSIPTrunksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service list Sip trunks params
func (o *ThinkteluControlAPIWebServiceListSIPTrunksParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListSIPTrunksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service list Sip trunks params
func (o *ThinkteluControlAPIWebServiceListSIPTrunksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceListSIPTrunksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}