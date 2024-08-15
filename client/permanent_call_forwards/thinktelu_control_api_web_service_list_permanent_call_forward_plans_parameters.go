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

// NewThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams creates a new ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams() *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams {
	return &ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceListPermanentCallForwardPlansParamsWithTimeout creates a new ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceListPermanentCallForwardPlansParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams {
	return &ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceListPermanentCallForwardPlansParamsWithContext creates a new ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceListPermanentCallForwardPlansParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams {
	return &ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceListPermanentCallForwardPlansParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceListPermanentCallForwardPlansParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams {
	return &ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service list permanent call forward plans operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service list permanent call forward plans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams) WithDefaults() *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service list permanent call forward plans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service list permanent call forward plans params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service list permanent call forward plans params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service list permanent call forward plans params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service list permanent call forward plans params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service list permanent call forward plans params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service list permanent call forward plans params
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
