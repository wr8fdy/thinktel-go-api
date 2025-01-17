// Code generated by go-swagger; DO NOT EDIT.

package sip_proxies

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

// NewThinkteluControlAPIWebServiceListSIPProxiesV2Params creates a new ThinkteluControlAPIWebServiceListSIPProxiesV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceListSIPProxiesV2Params() *ThinkteluControlAPIWebServiceListSIPProxiesV2Params {
	return &ThinkteluControlAPIWebServiceListSIPProxiesV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceListSIPProxiesV2ParamsWithTimeout creates a new ThinkteluControlAPIWebServiceListSIPProxiesV2Params object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceListSIPProxiesV2ParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListSIPProxiesV2Params {
	return &ThinkteluControlAPIWebServiceListSIPProxiesV2Params{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceListSIPProxiesV2ParamsWithContext creates a new ThinkteluControlAPIWebServiceListSIPProxiesV2Params object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceListSIPProxiesV2ParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListSIPProxiesV2Params {
	return &ThinkteluControlAPIWebServiceListSIPProxiesV2Params{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceListSIPProxiesV2ParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceListSIPProxiesV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceListSIPProxiesV2ParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListSIPProxiesV2Params {
	return &ThinkteluControlAPIWebServiceListSIPProxiesV2Params{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceListSIPProxiesV2Params contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service list Sip proxies v2 operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceListSIPProxiesV2Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service list Sip proxies v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListSIPProxiesV2Params) WithDefaults() *ThinkteluControlAPIWebServiceListSIPProxiesV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service list Sip proxies v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListSIPProxiesV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service list Sip proxies v2 params
func (o *ThinkteluControlAPIWebServiceListSIPProxiesV2Params) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListSIPProxiesV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service list Sip proxies v2 params
func (o *ThinkteluControlAPIWebServiceListSIPProxiesV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service list Sip proxies v2 params
func (o *ThinkteluControlAPIWebServiceListSIPProxiesV2Params) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListSIPProxiesV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service list Sip proxies v2 params
func (o *ThinkteluControlAPIWebServiceListSIPProxiesV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service list Sip proxies v2 params
func (o *ThinkteluControlAPIWebServiceListSIPProxiesV2Params) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListSIPProxiesV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service list Sip proxies v2 params
func (o *ThinkteluControlAPIWebServiceListSIPProxiesV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceListSIPProxiesV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
