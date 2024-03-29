// Code generated by go-swagger; DO NOT EDIT.

package socks5_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// ListSocks5ServerReader is a Reader for the ListSocks5Server structure.
type ListSocks5ServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSocks5ServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSocks5ServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListSocks5ServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListSocks5ServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListSocks5ServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSocks5ServerOK creates a ListSocks5ServerOK with default headers values
func NewListSocks5ServerOK() *ListSocks5ServerOK {
	return &ListSocks5ServerOK{}
}

/*
ListSocks5ServerOK describes a response with status code 200, with default header values.

ok
*/
type ListSocks5ServerOK struct {
	Payload []*vproxy_client_model.Socks5Server
}

// IsSuccess returns true when this list socks5 server o k response has a 2xx status code
func (o *ListSocks5ServerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list socks5 server o k response has a 3xx status code
func (o *ListSocks5ServerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list socks5 server o k response has a 4xx status code
func (o *ListSocks5ServerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list socks5 server o k response has a 5xx status code
func (o *ListSocks5ServerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list socks5 server o k response a status code equal to that given
func (o *ListSocks5ServerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list socks5 server o k response
func (o *ListSocks5ServerOK) Code() int {
	return 200
}

func (o *ListSocks5ServerOK) Error() string {
	return fmt.Sprintf("[GET /socks5-server][%d] listSocks5ServerOK  %+v", 200, o.Payload)
}

func (o *ListSocks5ServerOK) String() string {
	return fmt.Sprintf("[GET /socks5-server][%d] listSocks5ServerOK  %+v", 200, o.Payload)
}

func (o *ListSocks5ServerOK) GetPayload() []*vproxy_client_model.Socks5Server {
	return o.Payload
}

func (o *ListSocks5ServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSocks5ServerBadRequest creates a ListSocks5ServerBadRequest with default headers values
func NewListSocks5ServerBadRequest() *ListSocks5ServerBadRequest {
	return &ListSocks5ServerBadRequest{}
}

/*
ListSocks5ServerBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type ListSocks5ServerBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this list socks5 server bad request response has a 2xx status code
func (o *ListSocks5ServerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list socks5 server bad request response has a 3xx status code
func (o *ListSocks5ServerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list socks5 server bad request response has a 4xx status code
func (o *ListSocks5ServerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list socks5 server bad request response has a 5xx status code
func (o *ListSocks5ServerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list socks5 server bad request response a status code equal to that given
func (o *ListSocks5ServerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list socks5 server bad request response
func (o *ListSocks5ServerBadRequest) Code() int {
	return 400
}

func (o *ListSocks5ServerBadRequest) Error() string {
	return fmt.Sprintf("[GET /socks5-server][%d] listSocks5ServerBadRequest  %+v", 400, o.Payload)
}

func (o *ListSocks5ServerBadRequest) String() string {
	return fmt.Sprintf("[GET /socks5-server][%d] listSocks5ServerBadRequest  %+v", 400, o.Payload)
}

func (o *ListSocks5ServerBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *ListSocks5ServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSocks5ServerConflict creates a ListSocks5ServerConflict with default headers values
func NewListSocks5ServerConflict() *ListSocks5ServerConflict {
	return &ListSocks5ServerConflict{}
}

/*
ListSocks5ServerConflict describes a response with status code 409, with default header values.

conflict
*/
type ListSocks5ServerConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this list socks5 server conflict response has a 2xx status code
func (o *ListSocks5ServerConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list socks5 server conflict response has a 3xx status code
func (o *ListSocks5ServerConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list socks5 server conflict response has a 4xx status code
func (o *ListSocks5ServerConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this list socks5 server conflict response has a 5xx status code
func (o *ListSocks5ServerConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this list socks5 server conflict response a status code equal to that given
func (o *ListSocks5ServerConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the list socks5 server conflict response
func (o *ListSocks5ServerConflict) Code() int {
	return 409
}

func (o *ListSocks5ServerConflict) Error() string {
	return fmt.Sprintf("[GET /socks5-server][%d] listSocks5ServerConflict  %+v", 409, o.Payload)
}

func (o *ListSocks5ServerConflict) String() string {
	return fmt.Sprintf("[GET /socks5-server][%d] listSocks5ServerConflict  %+v", 409, o.Payload)
}

func (o *ListSocks5ServerConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *ListSocks5ServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSocks5ServerInternalServerError creates a ListSocks5ServerInternalServerError with default headers values
func NewListSocks5ServerInternalServerError() *ListSocks5ServerInternalServerError {
	return &ListSocks5ServerInternalServerError{}
}

/*
ListSocks5ServerInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type ListSocks5ServerInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this list socks5 server internal server error response has a 2xx status code
func (o *ListSocks5ServerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list socks5 server internal server error response has a 3xx status code
func (o *ListSocks5ServerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list socks5 server internal server error response has a 4xx status code
func (o *ListSocks5ServerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list socks5 server internal server error response has a 5xx status code
func (o *ListSocks5ServerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list socks5 server internal server error response a status code equal to that given
func (o *ListSocks5ServerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list socks5 server internal server error response
func (o *ListSocks5ServerInternalServerError) Code() int {
	return 500
}

func (o *ListSocks5ServerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /socks5-server][%d] listSocks5ServerInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSocks5ServerInternalServerError) String() string {
	return fmt.Sprintf("[GET /socks5-server][%d] listSocks5ServerInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSocks5ServerInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *ListSocks5ServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
