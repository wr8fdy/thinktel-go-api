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

	"github.com/wr8fdy/thinktel-go-api/models"
)

// NewThinkteluControlAPIWebServiceAddPermanentCallForwardParams creates a new ThinkteluControlAPIWebServiceAddPermanentCallForwardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceAddPermanentCallForwardParams() *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddPermanentCallForwardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceAddPermanentCallForwardParamsWithTimeout creates a new ThinkteluControlAPIWebServiceAddPermanentCallForwardParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceAddPermanentCallForwardParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddPermanentCallForwardParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceAddPermanentCallForwardParamsWithContext creates a new ThinkteluControlAPIWebServiceAddPermanentCallForwardParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceAddPermanentCallForwardParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddPermanentCallForwardParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceAddPermanentCallForwardParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceAddPermanentCallForwardParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceAddPermanentCallForwardParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddPermanentCallForwardParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceAddPermanentCallForwardParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service add permanent call forward operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceAddPermanentCallForwardParams struct {

	/* Order.

	   Permanent call forward order.
	*/
	Order *models.ThinkteluControlAPIDataContractsPermanentCallForwardOrder

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service add permanent call forward params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams) WithDefaults() *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service add permanent call forward params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service add permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service add permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service add permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service add permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service add permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service add permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrder adds the order to the thinktel u control Api web service add permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams) WithOrder(order *models.ThinkteluControlAPIDataContractsPermanentCallForwardOrder) *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the thinktel u control Api web service add permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams) SetOrder(order *models.ThinkteluControlAPIDataContractsPermanentCallForwardOrder) {
	o.Order = order
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceAddPermanentCallForwardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Order != nil {
		if err := r.SetBodyParam(o.Order); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
