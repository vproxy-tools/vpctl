// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// AddServerReader is a Reader for the AddServer structure.
type AddServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddServerNoContent creates a AddServerNoContent with default headers values
func NewAddServerNoContent() *AddServerNoContent {
	return &AddServerNoContent{}
}

/*
AddServerNoContent describes a response with status code 204, with default header values.

ok
*/
type AddServerNoContent struct {
}

// IsSuccess returns true when this add server no content response has a 2xx status code
func (o *AddServerNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add server no content response has a 3xx status code
func (o *AddServerNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add server no content response has a 4xx status code
func (o *AddServerNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add server no content response has a 5xx status code
func (o *AddServerNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add server no content response a status code equal to that given
func (o *AddServerNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the add server no content response
func (o *AddServerNoContent) Code() int {
	return 204
}

func (o *AddServerNoContent) Error() string {
	return fmt.Sprintf("[POST /server-group/{sg}/server][%d] addServerNoContent ", 204)
}

func (o *AddServerNoContent) String() string {
	return fmt.Sprintf("[POST /server-group/{sg}/server][%d] addServerNoContent ", 204)
}

func (o *AddServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddServerBadRequest creates a AddServerBadRequest with default headers values
func NewAddServerBadRequest() *AddServerBadRequest {
	return &AddServerBadRequest{}
}

/*
AddServerBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type AddServerBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this add server bad request response has a 2xx status code
func (o *AddServerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add server bad request response has a 3xx status code
func (o *AddServerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add server bad request response has a 4xx status code
func (o *AddServerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add server bad request response has a 5xx status code
func (o *AddServerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add server bad request response a status code equal to that given
func (o *AddServerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add server bad request response
func (o *AddServerBadRequest) Code() int {
	return 400
}

func (o *AddServerBadRequest) Error() string {
	return fmt.Sprintf("[POST /server-group/{sg}/server][%d] addServerBadRequest  %+v", 400, o.Payload)
}

func (o *AddServerBadRequest) String() string {
	return fmt.Sprintf("[POST /server-group/{sg}/server][%d] addServerBadRequest  %+v", 400, o.Payload)
}

func (o *AddServerBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *AddServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddServerNotFound creates a AddServerNotFound with default headers values
func NewAddServerNotFound() *AddServerNotFound {
	return &AddServerNotFound{}
}

/*
AddServerNotFound describes a response with status code 404, with default header values.

resource not found
*/
type AddServerNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this add server not found response has a 2xx status code
func (o *AddServerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add server not found response has a 3xx status code
func (o *AddServerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add server not found response has a 4xx status code
func (o *AddServerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add server not found response has a 5xx status code
func (o *AddServerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add server not found response a status code equal to that given
func (o *AddServerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add server not found response
func (o *AddServerNotFound) Code() int {
	return 404
}

func (o *AddServerNotFound) Error() string {
	return fmt.Sprintf("[POST /server-group/{sg}/server][%d] addServerNotFound  %+v", 404, o.Payload)
}

func (o *AddServerNotFound) String() string {
	return fmt.Sprintf("[POST /server-group/{sg}/server][%d] addServerNotFound  %+v", 404, o.Payload)
}

func (o *AddServerNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *AddServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddServerConflict creates a AddServerConflict with default headers values
func NewAddServerConflict() *AddServerConflict {
	return &AddServerConflict{}
}

/*
AddServerConflict describes a response with status code 409, with default header values.

conflict
*/
type AddServerConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this add server conflict response has a 2xx status code
func (o *AddServerConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add server conflict response has a 3xx status code
func (o *AddServerConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add server conflict response has a 4xx status code
func (o *AddServerConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add server conflict response has a 5xx status code
func (o *AddServerConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add server conflict response a status code equal to that given
func (o *AddServerConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add server conflict response
func (o *AddServerConflict) Code() int {
	return 409
}

func (o *AddServerConflict) Error() string {
	return fmt.Sprintf("[POST /server-group/{sg}/server][%d] addServerConflict  %+v", 409, o.Payload)
}

func (o *AddServerConflict) String() string {
	return fmt.Sprintf("[POST /server-group/{sg}/server][%d] addServerConflict  %+v", 409, o.Payload)
}

func (o *AddServerConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *AddServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddServerInternalServerError creates a AddServerInternalServerError with default headers values
func NewAddServerInternalServerError() *AddServerInternalServerError {
	return &AddServerInternalServerError{}
}

/*
AddServerInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type AddServerInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this add server internal server error response has a 2xx status code
func (o *AddServerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add server internal server error response has a 3xx status code
func (o *AddServerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add server internal server error response has a 4xx status code
func (o *AddServerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add server internal server error response has a 5xx status code
func (o *AddServerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add server internal server error response a status code equal to that given
func (o *AddServerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add server internal server error response
func (o *AddServerInternalServerError) Code() int {
	return 500
}

func (o *AddServerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /server-group/{sg}/server][%d] addServerInternalServerError  %+v", 500, o.Payload)
}

func (o *AddServerInternalServerError) String() string {
	return fmt.Sprintf("[POST /server-group/{sg}/server][%d] addServerInternalServerError  %+v", 500, o.Payload)
}

func (o *AddServerInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *AddServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
