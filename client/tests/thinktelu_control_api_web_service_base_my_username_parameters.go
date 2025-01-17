// Code generated by go-swagger; DO NOT EDIT.

package tests

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

// NewThinkteluControlAPIWebServiceBaseMyUsernameParams creates a new ThinkteluControlAPIWebServiceBaseMyUsernameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceBaseMyUsernameParams() *ThinkteluControlAPIWebServiceBaseMyUsernameParams {
	return &ThinkteluControlAPIWebServiceBaseMyUsernameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceBaseMyUsernameParamsWithTimeout creates a new ThinkteluControlAPIWebServiceBaseMyUsernameParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceBaseMyUsernameParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceBaseMyUsernameParams {
	return &ThinkteluControlAPIWebServiceBaseMyUsernameParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceBaseMyUsernameParamsWithContext creates a new ThinkteluControlAPIWebServiceBaseMyUsernameParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceBaseMyUsernameParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceBaseMyUsernameParams {
	return &ThinkteluControlAPIWebServiceBaseMyUsernameParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceBaseMyUsernameParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceBaseMyUsernameParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceBaseMyUsernameParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceBaseMyUsernameParams {
	return &ThinkteluControlAPIWebServiceBaseMyUsernameParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceBaseMyUsernameParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service base my username operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceBaseMyUsernameParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service base my username params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceBaseMyUsernameParams) WithDefaults() *ThinkteluControlAPIWebServiceBaseMyUsernameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service base my username params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceBaseMyUsernameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service base my username params
func (o *ThinkteluControlAPIWebServiceBaseMyUsernameParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceBaseMyUsernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service base my username params
func (o *ThinkteluControlAPIWebServiceBaseMyUsernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service base my username params
func (o *ThinkteluControlAPIWebServiceBaseMyUsernameParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceBaseMyUsernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service base my username params
func (o *ThinkteluControlAPIWebServiceBaseMyUsernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service base my username params
func (o *ThinkteluControlAPIWebServiceBaseMyUsernameParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceBaseMyUsernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service base my username params
func (o *ThinkteluControlAPIWebServiceBaseMyUsernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceBaseMyUsernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
