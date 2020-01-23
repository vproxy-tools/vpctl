// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
)

// DescribeServerReader is a Reader for the DescribeServer structure.
type DescribeServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDescribeServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDescribeServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDescribeServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDescribeServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeServerOK creates a DescribeServerOK with default headers values
func NewDescribeServerOK() *DescribeServerOK {
	return &DescribeServerOK{}
}

/*DescribeServerOK handles this case with default header values.

ok
*/
type DescribeServerOK struct {
	Payload *vproxy_client_model.ServerDetail
}

func (o *DescribeServerOK) Error() string {
	return fmt.Sprintf("[GET /server-group/{sg}/server/{svr}/detail][%d] describeServerOK  %+v", 200, o.Payload)
}

func (o *DescribeServerOK) GetPayload() *vproxy_client_model.ServerDetail {
	return o.Payload
}

func (o *DescribeServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.ServerDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeServerBadRequest creates a DescribeServerBadRequest with default headers values
func NewDescribeServerBadRequest() *DescribeServerBadRequest {
	return &DescribeServerBadRequest{}
}

/*DescribeServerBadRequest handles this case with default header values.

invalid input parameters
*/
type DescribeServerBadRequest struct {
	Payload *vproxy_client_model.Error400
}

func (o *DescribeServerBadRequest) Error() string {
	return fmt.Sprintf("[GET /server-group/{sg}/server/{svr}/detail][%d] describeServerBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeServerBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *DescribeServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeServerNotFound creates a DescribeServerNotFound with default headers values
func NewDescribeServerNotFound() *DescribeServerNotFound {
	return &DescribeServerNotFound{}
}

/*DescribeServerNotFound handles this case with default header values.

resource not found
*/
type DescribeServerNotFound struct {
	Payload *vproxy_client_model.Error404
}

func (o *DescribeServerNotFound) Error() string {
	return fmt.Sprintf("[GET /server-group/{sg}/server/{svr}/detail][%d] describeServerNotFound  %+v", 404, o.Payload)
}

func (o *DescribeServerNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *DescribeServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeServerConflict creates a DescribeServerConflict with default headers values
func NewDescribeServerConflict() *DescribeServerConflict {
	return &DescribeServerConflict{}
}

/*DescribeServerConflict handles this case with default header values.

conflict
*/
type DescribeServerConflict struct {
	Payload *vproxy_client_model.Error409
}

func (o *DescribeServerConflict) Error() string {
	return fmt.Sprintf("[GET /server-group/{sg}/server/{svr}/detail][%d] describeServerConflict  %+v", 409, o.Payload)
}

func (o *DescribeServerConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *DescribeServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeServerInternalServerError creates a DescribeServerInternalServerError with default headers values
func NewDescribeServerInternalServerError() *DescribeServerInternalServerError {
	return &DescribeServerInternalServerError{}
}

/*DescribeServerInternalServerError handles this case with default header values.

internal error
*/
type DescribeServerInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

func (o *DescribeServerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /server-group/{sg}/server/{svr}/detail][%d] describeServerInternalServerError  %+v", 500, o.Payload)
}

func (o *DescribeServerInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *DescribeServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
