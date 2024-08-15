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

// NewThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams creates a new ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams() *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams {
	return &ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParamsWithTimeout creates a new ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams {
	return &ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParamsWithContext creates a new ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams {
	return &ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams {
	return &ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service base test resource not found exception operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service base test resource not found exception params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams) WithDefaults() *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service base test resource not found exception params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service base test resource not found exception params
func (o *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service base test resource not found exception params
func (o *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service base test resource not found exception params
func (o *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service base test resource not found exception params
func (o *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service base test resource not found exception params
func (o *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service base test resource not found exception params
func (o *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}