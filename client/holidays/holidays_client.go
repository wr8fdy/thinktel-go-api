// Code generated by go-swagger; DO NOT EDIT.

package holidays

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new holidays API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new holidays API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new holidays API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for holidays API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeApplicationXML sets the Content-Type header to "application/xml".
func WithContentTypeApplicationXML(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/xml"}
}

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptApplicationXML sets the Accept header to "application/xml".
func WithAcceptApplicationXML(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/xml"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	ThinkteluControlAPIWebServiceListHolidays(params *ThinkteluControlAPIWebServiceListHolidaysParams, opts ...ClientOption) (*ThinktelUControlAPIWebServiceListHolidaysOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ThinkteluControlAPIWebServiceListHolidays lists upcoming distributel holidays
*/
func (a *Client) ThinkteluControlAPIWebServiceListHolidays(params *ThinkteluControlAPIWebServiceListHolidaysParams, opts ...ClientOption) (*ThinktelUControlAPIWebServiceListHolidaysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewThinkteluControlAPIWebServiceListHolidaysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Thinktel.uControl.Api.WebService.ListHolidays",
		Method:             "GET",
		PathPattern:        "/Holidays",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ThinkteluControlAPIWebServiceListHolidaysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ThinktelUControlAPIWebServiceListHolidaysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ThinkteluControlAPIWebServiceListHolidaysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
