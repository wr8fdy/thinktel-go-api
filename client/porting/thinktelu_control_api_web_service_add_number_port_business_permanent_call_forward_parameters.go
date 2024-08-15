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

// NewThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams creates a new ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams() *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParamsWithTimeout creates a new ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParamsWithContext creates a new ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service add number port business permanent call forward operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams struct {

	/* OrderForm.

	   Business permanent call forward port-in order.
	*/
	OrderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderBusinessPermanentCallForward

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service add number port business permanent call forward params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams) WithDefaults() *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service add number port business permanent call forward params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service add number port business permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service add number port business permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service add number port business permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service add number port business permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service add number port business permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service add number port business permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderForm adds the orderForm to the thinktel u control Api web service add number port business permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams) WithOrderForm(orderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderBusinessPermanentCallForward) *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams {
	o.SetOrderForm(orderForm)
	return o
}

// SetOrderForm adds the orderForm to the thinktel u control Api web service add number port business permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams) SetOrderForm(orderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderBusinessPermanentCallForward) {
	o.OrderForm = orderForm
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessPermanentCallForwardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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