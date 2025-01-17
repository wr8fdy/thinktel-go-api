// Code generated by go-swagger; DO NOT EDIT.

package residential_lines

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

// NewThinkteluControlAPIRestServiceGetResidentialLinePlanParams creates a new ThinkteluControlAPIRestServiceGetResidentialLinePlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIRestServiceGetResidentialLinePlanParams() *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams {
	return &ThinkteluControlAPIRestServiceGetResidentialLinePlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIRestServiceGetResidentialLinePlanParamsWithTimeout creates a new ThinkteluControlAPIRestServiceGetResidentialLinePlanParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIRestServiceGetResidentialLinePlanParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams {
	return &ThinkteluControlAPIRestServiceGetResidentialLinePlanParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIRestServiceGetResidentialLinePlanParamsWithContext creates a new ThinkteluControlAPIRestServiceGetResidentialLinePlanParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIRestServiceGetResidentialLinePlanParamsWithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams {
	return &ThinkteluControlAPIRestServiceGetResidentialLinePlanParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIRestServiceGetResidentialLinePlanParamsWithHTTPClient creates a new ThinkteluControlAPIRestServiceGetResidentialLinePlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIRestServiceGetResidentialLinePlanParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams {
	return &ThinkteluControlAPIRestServiceGetResidentialLinePlanParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIRestServiceGetResidentialLinePlanParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api rest service get residential line plan operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIRestServiceGetResidentialLinePlanParams struct {

	/* ID.

	   Offering identifier.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api rest service get residential line plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams) WithDefaults() *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api rest service get residential line plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api rest service get residential line plan params
func (o *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api rest service get residential line plan params
func (o *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api rest service get residential line plan params
func (o *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams) WithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api rest service get residential line plan params
func (o *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api rest service get residential line plan params
func (o *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api rest service get residential line plan params
func (o *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the thinktel u control Api rest service get residential line plan params
func (o *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams) WithID(id string) *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the thinktel u control Api rest service get residential line plan params
func (o *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIRestServiceGetResidentialLinePlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ID
	if err := r.SetPathParam("ID", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
