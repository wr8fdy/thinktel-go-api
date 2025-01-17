// Code generated by go-swagger; DO NOT EDIT.

package porting

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

// NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams creates a new ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams() *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams {
	return &ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoParamsWithTimeout creates a new ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams {
	return &ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoParamsWithContext creates a new ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams {
	return &ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams {
	return &ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service get port sys customer info operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service get port sys customer info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams) WithDefaults() *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service get port sys customer info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service get port sys customer info params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service get port sys customer info params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service get port sys customer info params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service get port sys customer info params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service get port sys customer info params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service get port sys customer info params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
