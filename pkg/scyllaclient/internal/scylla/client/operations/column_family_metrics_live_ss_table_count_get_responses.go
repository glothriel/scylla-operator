// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// ColumnFamilyMetricsLiveSsTableCountGetReader is a Reader for the ColumnFamilyMetricsLiveSsTableCountGet structure.
type ColumnFamilyMetricsLiveSsTableCountGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsLiveSsTableCountGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsLiveSsTableCountGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsLiveSsTableCountGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsLiveSsTableCountGetOK creates a ColumnFamilyMetricsLiveSsTableCountGetOK with default headers values
func NewColumnFamilyMetricsLiveSsTableCountGetOK() *ColumnFamilyMetricsLiveSsTableCountGetOK {
	return &ColumnFamilyMetricsLiveSsTableCountGetOK{}
}

/*
ColumnFamilyMetricsLiveSsTableCountGetOK handles this case with default header values.

ColumnFamilyMetricsLiveSsTableCountGetOK column family metrics live ss table count get o k
*/
type ColumnFamilyMetricsLiveSsTableCountGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsLiveSsTableCountGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsLiveSsTableCountGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsLiveSsTableCountGetDefault creates a ColumnFamilyMetricsLiveSsTableCountGetDefault with default headers values
func NewColumnFamilyMetricsLiveSsTableCountGetDefault(code int) *ColumnFamilyMetricsLiveSsTableCountGetDefault {
	return &ColumnFamilyMetricsLiveSsTableCountGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsLiveSsTableCountGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsLiveSsTableCountGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics live ss table count get default response
func (o *ColumnFamilyMetricsLiveSsTableCountGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsLiveSsTableCountGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsLiveSsTableCountGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsLiveSsTableCountGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
