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

// NewThinkteluControlAPIWebServiceListCountryNPAsParams creates a new ThinkteluControlAPIWebServiceListCountryNPAsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceListCountryNPAsParams() *ThinkteluControlAPIWebServiceListCountryNPAsParams {
	return &ThinkteluControlAPIWebServiceListCountryNPAsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceListCountryNPAsParamsWithTimeout creates a new ThinkteluControlAPIWebServiceListCountryNPAsParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceListCountryNPAsParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListCountryNPAsParams {
	return &ThinkteluControlAPIWebServiceListCountryNPAsParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceListCountryNPAsParamsWithContext creates a new ThinkteluControlAPIWebServiceListCountryNPAsParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceListCountryNPAsParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListCountryNPAsParams {
	return &ThinkteluControlAPIWebServiceListCountryNPAsParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceListCountryNPAsParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceListCountryNPAsParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceListCountryNPAsParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListCountryNPAsParams {
	return &ThinkteluControlAPIWebServiceListCountryNPAsParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceListCountryNPAsParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service list country n p as operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceListCountryNPAsParams struct {

	/* CountryAlpha2.

	   Country code (ISO-3166-1 Alpha2 Code).
	*/
	CountryAlpha2 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service list country n p as params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListCountryNPAsParams) WithDefaults() *ThinkteluControlAPIWebServiceListCountryNPAsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service list country n p as params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceListCountryNPAsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service list country n p as params
func (o *ThinkteluControlAPIWebServiceListCountryNPAsParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceListCountryNPAsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service list country n p as params
func (o *ThinkteluControlAPIWebServiceListCountryNPAsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service list country n p as params
func (o *ThinkteluControlAPIWebServiceListCountryNPAsParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceListCountryNPAsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service list country n p as params
func (o *ThinkteluControlAPIWebServiceListCountryNPAsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service list country n p as params
func (o *ThinkteluControlAPIWebServiceListCountryNPAsParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceListCountryNPAsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service list country n p as params
func (o *ThinkteluControlAPIWebServiceListCountryNPAsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCountryAlpha2 adds the countryAlpha2 to the thinktel u control Api web service list country n p as params
func (o *ThinkteluControlAPIWebServiceListCountryNPAsParams) WithCountryAlpha2(countryAlpha2 string) *ThinkteluControlAPIWebServiceListCountryNPAsParams {
	o.SetCountryAlpha2(countryAlpha2)
	return o
}

// SetCountryAlpha2 adds the countryAlpha2 to the thinktel u control Api web service list country n p as params
func (o *ThinkteluControlAPIWebServiceListCountryNPAsParams) SetCountryAlpha2(countryAlpha2 string) {
	o.CountryAlpha2 = countryAlpha2
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceListCountryNPAsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param country_alpha_2
	if err := r.SetPathParam("country_alpha_2", o.CountryAlpha2); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
