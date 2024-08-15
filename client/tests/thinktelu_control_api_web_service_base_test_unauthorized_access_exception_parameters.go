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

// NewThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams creates a new ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams() *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams {
	return &ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParamsWithTimeout creates a new ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams {
	return &ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParamsWithContext creates a new ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams {
	return &ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams {
	return &ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service base test unauthorized access exception operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service base test unauthorized access exception params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams) WithDefaults() *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service base test unauthorized access exception params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service base test unauthorized access exception params
func (o *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service base test unauthorized access exception params
func (o *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service base test unauthorized access exception params
func (o *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service base test unauthorized access exception params
func (o *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service base test unauthorized access exception params
func (o *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service base test unauthorized access exception params
func (o *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceBaseTestUnauthorizedAccessExceptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
