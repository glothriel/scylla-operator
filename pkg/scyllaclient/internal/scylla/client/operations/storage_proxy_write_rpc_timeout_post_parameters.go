// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewStorageProxyWriteRPCTimeoutPostParams creates a new StorageProxyWriteRPCTimeoutPostParams object
// with the default values initialized.
func NewStorageProxyWriteRPCTimeoutPostParams() *StorageProxyWriteRPCTimeoutPostParams {
	var ()
	return &StorageProxyWriteRPCTimeoutPostParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyWriteRPCTimeoutPostParamsWithTimeout creates a new StorageProxyWriteRPCTimeoutPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyWriteRPCTimeoutPostParamsWithTimeout(timeout time.Duration) *StorageProxyWriteRPCTimeoutPostParams {
	var ()
	return &StorageProxyWriteRPCTimeoutPostParams{

		requestTimeout: timeout,
	}
}

// NewStorageProxyWriteRPCTimeoutPostParamsWithContext creates a new StorageProxyWriteRPCTimeoutPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyWriteRPCTimeoutPostParamsWithContext(ctx context.Context) *StorageProxyWriteRPCTimeoutPostParams {
	var ()
	return &StorageProxyWriteRPCTimeoutPostParams{

		Context: ctx,
	}
}

// NewStorageProxyWriteRPCTimeoutPostParamsWithHTTPClient creates a new StorageProxyWriteRPCTimeoutPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyWriteRPCTimeoutPostParamsWithHTTPClient(client *http.Client) *StorageProxyWriteRPCTimeoutPostParams {
	var ()
	return &StorageProxyWriteRPCTimeoutPostParams{
		HTTPClient: client,
	}
}

/*
StorageProxyWriteRPCTimeoutPostParams contains all the parameters to send to the API endpoint
for the storage proxy write Rpc timeout post operation typically these are written to a http.Request
*/
type StorageProxyWriteRPCTimeoutPostParams struct {

	/*Timeout
	  timeout in seconds

	*/
	Timeout string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the storage proxy write Rpc timeout post params
func (o *StorageProxyWriteRPCTimeoutPostParams) WithRequestTimeout(timeout time.Duration) *StorageProxyWriteRPCTimeoutPostParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the storage proxy write Rpc timeout post params
func (o *StorageProxyWriteRPCTimeoutPostParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the storage proxy write Rpc timeout post params
func (o *StorageProxyWriteRPCTimeoutPostParams) WithContext(ctx context.Context) *StorageProxyWriteRPCTimeoutPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy write Rpc timeout post params
func (o *StorageProxyWriteRPCTimeoutPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy write Rpc timeout post params
func (o *StorageProxyWriteRPCTimeoutPostParams) WithHTTPClient(client *http.Client) *StorageProxyWriteRPCTimeoutPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy write Rpc timeout post params
func (o *StorageProxyWriteRPCTimeoutPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimeout adds the timeout to the storage proxy write Rpc timeout post params
func (o *StorageProxyWriteRPCTimeoutPostParams) WithTimeout(timeout string) *StorageProxyWriteRPCTimeoutPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy write Rpc timeout post params
func (o *StorageProxyWriteRPCTimeoutPostParams) SetTimeout(timeout string) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyWriteRPCTimeoutPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	// query param timeout
	qrTimeout := o.Timeout
	qTimeout := qrTimeout
	if qTimeout != "" {
		if err := r.SetQueryParam("timeout", qTimeout); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
