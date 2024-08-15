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

	"github.com/wr8fdy/thinktel-go-api/models"
)

// NewThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params creates a new ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params() *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params {
	return &ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2ParamsWithTimeout creates a new ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2ParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params {
	return &ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2ParamsWithContext creates a new ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2ParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params {
	return &ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2ParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2ParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params {
	return &ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service add number port toll free DID v2 operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params struct {

	/* OrderForm.

	   Toll-Free SIP trunk DID port-in order.
	*/
	OrderForm *models.ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service add number port toll free DID v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params) WithDefaults() *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service add number port toll free DID v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service add number port toll free DID v2 params
func (o *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service add number port toll free DID v2 params
func (o *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service add number port toll free DID v2 params
func (o *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service add number port toll free DID v2 params
func (o *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service add number port toll free DID v2 params
func (o *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service add number port toll free DID v2 params
func (o *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderForm adds the orderForm to the thinktel u control Api web service add number port toll free DID v2 params
func (o *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params) WithOrderForm(orderForm *models.ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2) *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params {
	o.SetOrderForm(orderForm)
	return o
}

// SetOrderForm adds the orderForm to the thinktel u control Api web service add number port toll free DID v2 params
func (o *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params) SetOrderForm(orderForm *models.ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2) {
	o.OrderForm = orderForm
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceAddNumberPortTollFreeDIDV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.OrderForm != nil {
		if err := r.SetBodyParam(o.OrderForm); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
