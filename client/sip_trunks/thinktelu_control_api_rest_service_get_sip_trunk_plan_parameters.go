// Code generated by go-swagger; DO NOT EDIT.

package sip_trunks

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

// NewThinkteluControlAPIRestServiceGetSIPTrunkPlanParams creates a new ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIRestServiceGetSIPTrunkPlanParams() *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams {
	return &ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIRestServiceGetSIPTrunkPlanParamsWithTimeout creates a new ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIRestServiceGetSIPTrunkPlanParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams {
	return &ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIRestServiceGetSIPTrunkPlanParamsWithContext creates a new ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIRestServiceGetSIPTrunkPlanParamsWithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams {
	return &ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIRestServiceGetSIPTrunkPlanParamsWithHTTPClient creates a new ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIRestServiceGetSIPTrunkPlanParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams {
	return &ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api rest service get Sip trunk plan operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams struct {

	/* ID.

	   Offering identifier.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api rest service get Sip trunk plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams) WithDefaults() *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api rest service get Sip trunk plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api rest service get Sip trunk plan params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api rest service get Sip trunk plan params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api rest service get Sip trunk plan params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams) WithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api rest service get Sip trunk plan params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api rest service get Sip trunk plan params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api rest service get Sip trunk plan params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the thinktel u control Api rest service get Sip trunk plan params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams) WithID(id string) *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the thinktel u control Api rest service get Sip trunk plan params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
