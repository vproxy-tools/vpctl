// Code generated by go-swagger; DO NOT EDIT.

package dns_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// AddDNSServerReader is a Reader for the AddDNSServer structure.
type AddDNSServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDNSServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddDNSServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddDNSServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddDNSServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddDNSServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddDNSServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddDNSServerNoContent creates a AddDNSServerNoContent with default headers values
func NewAddDNSServerNoContent() *AddDNSServerNoContent {
	return &AddDNSServerNoContent{}
}

/*
AddDNSServerNoContent describes a response with status code 204, with default header values.

ok
*/
type AddDNSServerNoContent struct {
}

// IsSuccess returns true when this add Dns server no content response has a 2xx status code
func (o *AddDNSServerNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add Dns server no content response has a 3xx status code
func (o *AddDNSServerNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add Dns server no content response has a 4xx status code
func (o *AddDNSServerNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add Dns server no content response has a 5xx status code
func (o *AddDNSServerNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add Dns server no content response a status code equal to that given
func (o *AddDNSServerNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the add Dns server no content response
func (o *AddDNSServerNoContent) Code() int {
	return 204
}

func (o *AddDNSServerNoContent) Error() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerNoContent ", 204)
}

func (o *AddDNSServerNoContent) String() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerNoContent ", 204)
}

func (o *AddDNSServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddDNSServerBadRequest creates a AddDNSServerBadRequest with default headers values
func NewAddDNSServerBadRequest() *AddDNSServerBadRequest {
	return &AddDNSServerBadRequest{}
}

/*
AddDNSServerBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type AddDNSServerBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this add Dns server bad request response has a 2xx status code
func (o *AddDNSServerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add Dns server bad request response has a 3xx status code
func (o *AddDNSServerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add Dns server bad request response has a 4xx status code
func (o *AddDNSServerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add Dns server bad request response has a 5xx status code
func (o *AddDNSServerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add Dns server bad request response a status code equal to that given
func (o *AddDNSServerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add Dns server bad request response
func (o *AddDNSServerBadRequest) Code() int {
	return 400
}

func (o *AddDNSServerBadRequest) Error() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerBadRequest  %+v", 400, o.Payload)
}

func (o *AddDNSServerBadRequest) String() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerBadRequest  %+v", 400, o.Payload)
}

func (o *AddDNSServerBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *AddDNSServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDNSServerNotFound creates a AddDNSServerNotFound with default headers values
func NewAddDNSServerNotFound() *AddDNSServerNotFound {
	return &AddDNSServerNotFound{}
}

/*
AddDNSServerNotFound describes a response with status code 404, with default header values.

resource not found
*/
type AddDNSServerNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this add Dns server not found response has a 2xx status code
func (o *AddDNSServerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add Dns server not found response has a 3xx status code
func (o *AddDNSServerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add Dns server not found response has a 4xx status code
func (o *AddDNSServerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add Dns server not found response has a 5xx status code
func (o *AddDNSServerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add Dns server not found response a status code equal to that given
func (o *AddDNSServerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add Dns server not found response
func (o *AddDNSServerNotFound) Code() int {
	return 404
}

func (o *AddDNSServerNotFound) Error() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerNotFound  %+v", 404, o.Payload)
}

func (o *AddDNSServerNotFound) String() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerNotFound  %+v", 404, o.Payload)
}

func (o *AddDNSServerNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *AddDNSServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDNSServerConflict creates a AddDNSServerConflict with default headers values
func NewAddDNSServerConflict() *AddDNSServerConflict {
	return &AddDNSServerConflict{}
}

/*
AddDNSServerConflict describes a response with status code 409, with default header values.

conflict
*/
type AddDNSServerConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this add Dns server conflict response has a 2xx status code
func (o *AddDNSServerConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add Dns server conflict response has a 3xx status code
func (o *AddDNSServerConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add Dns server conflict response has a 4xx status code
func (o *AddDNSServerConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add Dns server conflict response has a 5xx status code
func (o *AddDNSServerConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add Dns server conflict response a status code equal to that given
func (o *AddDNSServerConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add Dns server conflict response
func (o *AddDNSServerConflict) Code() int {
	return 409
}

func (o *AddDNSServerConflict) Error() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerConflict  %+v", 409, o.Payload)
}

func (o *AddDNSServerConflict) String() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerConflict  %+v", 409, o.Payload)
}

func (o *AddDNSServerConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *AddDNSServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDNSServerInternalServerError creates a AddDNSServerInternalServerError with default headers values
func NewAddDNSServerInternalServerError() *AddDNSServerInternalServerError {
	return &AddDNSServerInternalServerError{}
}

/*
AddDNSServerInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type AddDNSServerInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this add Dns server internal server error response has a 2xx status code
func (o *AddDNSServerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add Dns server internal server error response has a 3xx status code
func (o *AddDNSServerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add Dns server internal server error response has a 4xx status code
func (o *AddDNSServerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add Dns server internal server error response has a 5xx status code
func (o *AddDNSServerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add Dns server internal server error response a status code equal to that given
func (o *AddDNSServerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add Dns server internal server error response
func (o *AddDNSServerInternalServerError) Code() int {
	return 500
}

func (o *AddDNSServerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerInternalServerError  %+v", 500, o.Payload)
}

func (o *AddDNSServerInternalServerError) String() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerInternalServerError  %+v", 500, o.Payload)
}

func (o *AddDNSServerInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *AddDNSServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
