// Code generated by go-swagger; DO NOT EDIT.

package deprecated

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

// NewThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams creates a new ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams() *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParamsWithTimeout creates a new ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParamsWithContext creates a new ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service add number port permanent call forward operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams struct {

	/* OrderForm.

	   Permanent call forward port-in order.
	*/
	OrderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForward

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service add number port permanent call forward params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams) WithDefaults() *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service add number port permanent call forward params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service add number port permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service add number port permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service add number port permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service add number port permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service add number port permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service add number port permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderForm adds the orderForm to the thinktel u control Api web service add number port permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams) WithOrderForm(orderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForward) *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams {
	o.SetOrderForm(orderForm)
	return o
}

// SetOrderForm adds the orderForm to the thinktel u control Api web service add number port permanent call forward params
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams) SetOrderForm(orderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForward) {
	o.OrderForm = orderForm
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
