// Code generated by go-swagger; DO NOT EDIT.

package version911

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

// NewThinkteluControlAPIWebServiceAddV911Params creates a new ThinkteluControlAPIWebServiceAddV911Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceAddV911Params() *ThinkteluControlAPIWebServiceAddV911Params {
	return &ThinkteluControlAPIWebServiceAddV911Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceAddV911ParamsWithTimeout creates a new ThinkteluControlAPIWebServiceAddV911Params object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceAddV911ParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddV911Params {
	return &ThinkteluControlAPIWebServiceAddV911Params{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceAddV911ParamsWithContext creates a new ThinkteluControlAPIWebServiceAddV911Params object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceAddV911ParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddV911Params {
	return &ThinkteluControlAPIWebServiceAddV911Params{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceAddV911ParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceAddV911Params object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceAddV911ParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddV911Params {
	return &ThinkteluControlAPIWebServiceAddV911Params{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceAddV911Params contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service add v911 operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceAddV911Params struct {

	/* V911.

	   V911 contact information record.
	*/
	V911 *models.ThinkteluControlAPIDataContractsV911

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service add v911 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddV911Params) WithDefaults() *ThinkteluControlAPIWebServiceAddV911Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service add v911 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddV911Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service add v911 params
func (o *ThinkteluControlAPIWebServiceAddV911Params) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddV911Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service add v911 params
func (o *ThinkteluControlAPIWebServiceAddV911Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service add v911 params
func (o *ThinkteluControlAPIWebServiceAddV911Params) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddV911Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service add v911 params
func (o *ThinkteluControlAPIWebServiceAddV911Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service add v911 params
func (o *ThinkteluControlAPIWebServiceAddV911Params) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddV911Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service add v911 params
func (o *ThinkteluControlAPIWebServiceAddV911Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV911 adds the v911 to the thinktel u control Api web service add v911 params
func (o *ThinkteluControlAPIWebServiceAddV911Params) WithV911(v911 *models.ThinkteluControlAPIDataContractsV911) *ThinkteluControlAPIWebServiceAddV911Params {
	o.SetV911(v911)
	return o
}

// SetV911 adds the v911 to the thinktel u control Api web service add v911 params
func (o *ThinkteluControlAPIWebServiceAddV911Params) SetV911(v911 *models.ThinkteluControlAPIDataContractsV911) {
	o.V911 = v911
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceAddV911Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V911 != nil {
		if err := r.SetBodyParam(o.V911); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
