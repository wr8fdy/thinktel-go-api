// Code generated by go-swagger; DO NOT EDIT.

package rate_centers

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

// NewThinkteluControlAPIWebServiceGetRateCenterParams creates a new ThinkteluControlAPIWebServiceGetRateCenterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceGetRateCenterParams() *ThinkteluControlAPIWebServiceGetRateCenterParams {
	return &ThinkteluControlAPIWebServiceGetRateCenterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceGetRateCenterParamsWithTimeout creates a new ThinkteluControlAPIWebServiceGetRateCenterParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceGetRateCenterParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetRateCenterParams {
	return &ThinkteluControlAPIWebServiceGetRateCenterParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceGetRateCenterParamsWithContext creates a new ThinkteluControlAPIWebServiceGetRateCenterParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceGetRateCenterParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetRateCenterParams {
	return &ThinkteluControlAPIWebServiceGetRateCenterParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceGetRateCenterParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceGetRateCenterParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceGetRateCenterParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetRateCenterParams {
	return &ThinkteluControlAPIWebServiceGetRateCenterParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceGetRateCenterParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service get rate center operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceGetRateCenterParams struct {

	/* Name.

	   Rate center name.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service get rate center params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetRateCenterParams) WithDefaults() *ThinkteluControlAPIWebServiceGetRateCenterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service get rate center params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetRateCenterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service get rate center params
func (o *ThinkteluControlAPIWebServiceGetRateCenterParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetRateCenterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service get rate center params
func (o *ThinkteluControlAPIWebServiceGetRateCenterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service get rate center params
func (o *ThinkteluControlAPIWebServiceGetRateCenterParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetRateCenterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service get rate center params
func (o *ThinkteluControlAPIWebServiceGetRateCenterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service get rate center params
func (o *ThinkteluControlAPIWebServiceGetRateCenterParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetRateCenterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service get rate center params
func (o *ThinkteluControlAPIWebServiceGetRateCenterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the thinktel u control Api web service get rate center params
func (o *ThinkteluControlAPIWebServiceGetRateCenterParams) WithName(name string) *ThinkteluControlAPIWebServiceGetRateCenterParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the thinktel u control Api web service get rate center params
func (o *ThinkteluControlAPIWebServiceGetRateCenterParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceGetRateCenterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
