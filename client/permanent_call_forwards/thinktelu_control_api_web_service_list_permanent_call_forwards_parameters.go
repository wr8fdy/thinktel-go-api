// Code generated by go-swagger; DO NOT EDIT.

package permanent_call_forwards

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

// NewThinkteluControlAPIWebServiceListPermanentCallForwardsParams creates a new ThinkteluControlAPIWebServiceListPermanentCallForwardsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceListPermanentCallForwardsParams() *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams {
	return &ThinkteluControlAPIWebServiceListPermanentCallForwardsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceListPermanentCallForwardsParamsWithTimeout creates a new ThinkteluControlAPIWebServiceListPermanentCallForwardsParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceListPermanentCallForwardsParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams {
	return &ThinkteluControlAPIWebServiceListPermanentCallForwardsParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceListPermanentCallForwardsParamsWithContext creates a new ThinkteluControlAPIWebServiceListPermanentCallForwardsParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceListPermanentCallForwardsParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams {
	return &ThinkteluControlAPIWebServiceListPermanentCallForwardsParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceListPermanentCallForwardsParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceListPermanentCallForwardsParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceListPermanentCallForwardsParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams {
	return &ThinkteluControlAPIWebServiceListPermanentCallForwardsParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceListPermanentCallForwardsParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service list permanent call forwards operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceListPermanentCallForwardsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service list permanent call forwards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams) WithDefaults() *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service list permanent call forwards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service list permanent call forwards params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service list permanent call forwards params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service list permanent call forwards params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service list permanent call forwards params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service list permanent call forwards params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service list permanent call forwards params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}