// Code generated by go-swagger; DO NOT EDIT.

package operators

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

// NewListSupportedOperatorsParams creates a new ListSupportedOperatorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSupportedOperatorsParams() *ListSupportedOperatorsParams {
	return &ListSupportedOperatorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSupportedOperatorsParamsWithTimeout creates a new ListSupportedOperatorsParams object
// with the ability to set a timeout on a request.
func NewListSupportedOperatorsParamsWithTimeout(timeout time.Duration) *ListSupportedOperatorsParams {
	return &ListSupportedOperatorsParams{
		timeout: timeout,
	}
}

// NewListSupportedOperatorsParamsWithContext creates a new ListSupportedOperatorsParams object
// with the ability to set a context for a request.
func NewListSupportedOperatorsParamsWithContext(ctx context.Context) *ListSupportedOperatorsParams {
	return &ListSupportedOperatorsParams{
		Context: ctx,
	}
}

// NewListSupportedOperatorsParamsWithHTTPClient creates a new ListSupportedOperatorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSupportedOperatorsParamsWithHTTPClient(client *http.Client) *ListSupportedOperatorsParams {
	return &ListSupportedOperatorsParams{
		HTTPClient: client,
	}
}

/* ListSupportedOperatorsParams contains all the parameters to send to the API endpoint
   for the list supported operators operation.

   Typically these are written to a http.Request.
*/
type ListSupportedOperatorsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list supported operators params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSupportedOperatorsParams) WithDefaults() *ListSupportedOperatorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list supported operators params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSupportedOperatorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list supported operators params
func (o *ListSupportedOperatorsParams) WithTimeout(timeout time.Duration) *ListSupportedOperatorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list supported operators params
func (o *ListSupportedOperatorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list supported operators params
func (o *ListSupportedOperatorsParams) WithContext(ctx context.Context) *ListSupportedOperatorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list supported operators params
func (o *ListSupportedOperatorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list supported operators params
func (o *ListSupportedOperatorsParams) WithHTTPClient(client *http.Client) *ListSupportedOperatorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list supported operators params
func (o *ListSupportedOperatorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListSupportedOperatorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
