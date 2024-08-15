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

// NewThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams creates a new ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams() *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortResidentialLineParamsWithTimeout creates a new ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceAddNumberPortResidentialLineParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortResidentialLineParamsWithContext creates a new ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceAddNumberPortResidentialLineParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortResidentialLineParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceAddNumberPortResidentialLineParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service add number port residential line operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams struct {

	/* OrderForm.

	   Residential line port-in order.
	*/
	OrderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderResidentialLine

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service add number port residential line params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams) WithDefaults() *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service add number port residential line params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service add number port residential line params
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service add number port residential line params
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service add number port residential line params
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service add number port residential line params
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service add number port residential line params
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service add number port residential line params
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderForm adds the orderForm to the thinktel u control Api web service add number port residential line params
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams) WithOrderForm(orderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderResidentialLine) *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams {
	o.SetOrderForm(orderForm)
	return o
}

// SetOrderForm adds the orderForm to the thinktel u control Api web service add number port residential line params
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams) SetOrderForm(orderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderResidentialLine) {
	o.OrderForm = orderForm
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
