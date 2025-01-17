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

// NewColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams creates a new ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams() *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParamsWithContext creates a new ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics read latency estimated recent histogram by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics read latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics read latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics read latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics read latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics read latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics read latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics read latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams) WithName(name string) *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics read latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsReadLatencyEstimatedRecentHistogramByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
