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

// NewThinkteluControlAPIWebServiceListRateCenterBlocksParams creates a new ThinkteluControlAPIWebServiceListRateCenterBlocksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceListRateCenterBlocksParams() *ThinkteluControlAPIWebServiceListRateCenterBlocksParams {
	return &ThinkteluControlAPIWebServiceListRateCenterBlocksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceListRateCenterBlocksParamsWithTimeout creates a new ThinkteluControlAPIWebServiceListRateCenterBlocksParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceListRateCenterBlocksParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListRateCenterBlocksParams {
	return &ThinkteluControlAPIWebServiceListRateCenterBlocksParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceListRateCenterBlocksParamsWithContext creates a new ThinkteluControlAPIWebServiceListRateCenterBlocksParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceListRateCenterBlocksParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListRateCenterBlocksParams {
	return &ThinkteluControlAPIWebServiceListRateCenterBlocksParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceListRateCenterBlocksParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceListRateCenterBlocksParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceListRateCenterBlocksParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListRateCenterBlocksParams {
	return &ThinkteluControlAPIWebServiceListRateCenterBlocksParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceListRateCenterBlocksParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service list rate center blocks operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceListRateCenterBlocksParams struct {

	/* Name.

	   Rate center name.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service list rate center blocks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListRateCenterBlocksParams) WithDefaults() *ThinkteluControlAPIWebServiceListRateCenterBlocksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service list rate center blocks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListRateCenterBlocksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service list rate center blocks params
func (o *ThinkteluControlAPIWebServiceListRateCenterBlocksParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListRateCenterBlocksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service list rate center blocks params
func (o *ThinkteluControlAPIWebServiceListRateCenterBlocksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service list rate center blocks params
func (o *ThinkteluControlAPIWebServiceListRateCenterBlocksParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListRateCenterBlocksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service list rate center blocks params
func (o *ThinkteluControlAPIWebServiceListRateCenterBlocksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service list rate center blocks params
func (o *ThinkteluControlAPIWebServiceListRateCenterBlocksParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListRateCenterBlocksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service list rate center blocks params
func (o *ThinkteluControlAPIWebServiceListRateCenterBlocksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the thinktel u control Api web service list rate center blocks params
func (o *ThinkteluControlAPIWebServiceListRateCenterBlocksParams) WithName(name string) *ThinkteluControlAPIWebServiceListRateCenterBlocksParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the thinktel u control Api web service list rate center blocks params
func (o *ThinkteluControlAPIWebServiceListRateCenterBlocksParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceListRateCenterBlocksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
