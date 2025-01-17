// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindConfigMemtableOffheapSpaceInMbParams creates a new FindConfigMemtableOffheapSpaceInMbParams object
// with the default values initialized.
func NewFindConfigMemtableOffheapSpaceInMbParams() *FindConfigMemtableOffheapSpaceInMbParams {

	return &FindConfigMemtableOffheapSpaceInMbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigMemtableOffheapSpaceInMbParamsWithTimeout creates a new FindConfigMemtableOffheapSpaceInMbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigMemtableOffheapSpaceInMbParamsWithTimeout(timeout time.Duration) *FindConfigMemtableOffheapSpaceInMbParams {

	return &FindConfigMemtableOffheapSpaceInMbParams{

		timeout: timeout,
	}
}

// NewFindConfigMemtableOffheapSpaceInMbParamsWithContext creates a new FindConfigMemtableOffheapSpaceInMbParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigMemtableOffheapSpaceInMbParamsWithContext(ctx context.Context) *FindConfigMemtableOffheapSpaceInMbParams {

	return &FindConfigMemtableOffheapSpaceInMbParams{

		Context: ctx,
	}
}

// NewFindConfigMemtableOffheapSpaceInMbParamsWithHTTPClient creates a new FindConfigMemtableOffheapSpaceInMbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigMemtableOffheapSpaceInMbParamsWithHTTPClient(client *http.Client) *FindConfigMemtableOffheapSpaceInMbParams {

	return &FindConfigMemtableOffheapSpaceInMbParams{
		HTTPClient: client,
	}
}

/*
FindConfigMemtableOffheapSpaceInMbParams contains all the parameters to send to the API endpoint
for the find config memtable offheap space in mb operation typically these are written to a http.Request
*/
type FindConfigMemtableOffheapSpaceInMbParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config memtable offheap space in mb params
func (o *FindConfigMemtableOffheapSpaceInMbParams) WithTimeout(timeout time.Duration) *FindConfigMemtableOffheapSpaceInMbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config memtable offheap space in mb params
func (o *FindConfigMemtableOffheapSpaceInMbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config memtable offheap space in mb params
func (o *FindConfigMemtableOffheapSpaceInMbParams) WithContext(ctx context.Context) *FindConfigMemtableOffheapSpaceInMbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config memtable offheap space in mb params
func (o *FindConfigMemtableOffheapSpaceInMbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config memtable offheap space in mb params
func (o *FindConfigMemtableOffheapSpaceInMbParams) WithHTTPClient(client *http.Client) *FindConfigMemtableOffheapSpaceInMbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config memtable offheap space in mb params
func (o *FindConfigMemtableOffheapSpaceInMbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigMemtableOffheapSpaceInMbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
