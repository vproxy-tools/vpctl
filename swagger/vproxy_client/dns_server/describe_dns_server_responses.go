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

// DescribeDNSServerReader is a Reader for the DescribeDNSServer structure.
type DescribeDNSServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeDNSServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeDNSServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDescribeDNSServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDescribeDNSServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDescribeDNSServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDescribeDNSServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDescribeDNSServerOK creates a DescribeDNSServerOK with default headers values
func NewDescribeDNSServerOK() *DescribeDNSServerOK {
	return &DescribeDNSServerOK{}
}

/*
DescribeDNSServerOK describes a response with status code 200, with default header values.

ok
*/
type DescribeDNSServerOK struct {
	Payload *vproxy_client_model.DNSServerDetail
}

// IsSuccess returns true when this describe Dns server o k response has a 2xx status code
func (o *DescribeDNSServerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe Dns server o k response has a 3xx status code
func (o *DescribeDNSServerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe Dns server o k response has a 4xx status code
func (o *DescribeDNSServerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe Dns server o k response has a 5xx status code
func (o *DescribeDNSServerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe Dns server o k response a status code equal to that given
func (o *DescribeDNSServerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe Dns server o k response
func (o *DescribeDNSServerOK) Code() int {
	return 200
}

func (o *DescribeDNSServerOK) Error() string {
	return fmt.Sprintf("[GET /dns-server/{dns}/detail][%d] describeDnsServerOK  %+v", 200, o.Payload)
}

func (o *DescribeDNSServerOK) String() string {
	return fmt.Sprintf("[GET /dns-server/{dns}/detail][%d] describeDnsServerOK  %+v", 200, o.Payload)
}

func (o *DescribeDNSServerOK) GetPayload() *vproxy_client_model.DNSServerDetail {
	return o.Payload
}

func (o *DescribeDNSServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.DNSServerDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeDNSServerBadRequest creates a DescribeDNSServerBadRequest with default headers values
func NewDescribeDNSServerBadRequest() *DescribeDNSServerBadRequest {
	return &DescribeDNSServerBadRequest{}
}

/*
DescribeDNSServerBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type DescribeDNSServerBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this describe Dns server bad request response has a 2xx status code
func (o *DescribeDNSServerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe Dns server bad request response has a 3xx status code
func (o *DescribeDNSServerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe Dns server bad request response has a 4xx status code
func (o *DescribeDNSServerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe Dns server bad request response has a 5xx status code
func (o *DescribeDNSServerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this describe Dns server bad request response a status code equal to that given
func (o *DescribeDNSServerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the describe Dns server bad request response
func (o *DescribeDNSServerBadRequest) Code() int {
	return 400
}

func (o *DescribeDNSServerBadRequest) Error() string {
	return fmt.Sprintf("[GET /dns-server/{dns}/detail][%d] describeDnsServerBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeDNSServerBadRequest) String() string {
	return fmt.Sprintf("[GET /dns-server/{dns}/detail][%d] describeDnsServerBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeDNSServerBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *DescribeDNSServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeDNSServerNotFound creates a DescribeDNSServerNotFound with default headers values
func NewDescribeDNSServerNotFound() *DescribeDNSServerNotFound {
	return &DescribeDNSServerNotFound{}
}

/*
DescribeDNSServerNotFound describes a response with status code 404, with default header values.

resource not found
*/
type DescribeDNSServerNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this describe Dns server not found response has a 2xx status code
func (o *DescribeDNSServerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe Dns server not found response has a 3xx status code
func (o *DescribeDNSServerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe Dns server not found response has a 4xx status code
func (o *DescribeDNSServerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe Dns server not found response has a 5xx status code
func (o *DescribeDNSServerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this describe Dns server not found response a status code equal to that given
func (o *DescribeDNSServerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the describe Dns server not found response
func (o *DescribeDNSServerNotFound) Code() int {
	return 404
}

func (o *DescribeDNSServerNotFound) Error() string {
	return fmt.Sprintf("[GET /dns-server/{dns}/detail][%d] describeDnsServerNotFound  %+v", 404, o.Payload)
}

func (o *DescribeDNSServerNotFound) String() string {
	return fmt.Sprintf("[GET /dns-server/{dns}/detail][%d] describeDnsServerNotFound  %+v", 404, o.Payload)
}

func (o *DescribeDNSServerNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *DescribeDNSServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeDNSServerConflict creates a DescribeDNSServerConflict with default headers values
func NewDescribeDNSServerConflict() *DescribeDNSServerConflict {
	return &DescribeDNSServerConflict{}
}

/*
DescribeDNSServerConflict describes a response with status code 409, with default header values.

conflict
*/
type DescribeDNSServerConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this describe Dns server conflict response has a 2xx status code
func (o *DescribeDNSServerConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe Dns server conflict response has a 3xx status code
func (o *DescribeDNSServerConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe Dns server conflict response has a 4xx status code
func (o *DescribeDNSServerConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe Dns server conflict response has a 5xx status code
func (o *DescribeDNSServerConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this describe Dns server conflict response a status code equal to that given
func (o *DescribeDNSServerConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the describe Dns server conflict response
func (o *DescribeDNSServerConflict) Code() int {
	return 409
}

func (o *DescribeDNSServerConflict) Error() string {
	return fmt.Sprintf("[GET /dns-server/{dns}/detail][%d] describeDnsServerConflict  %+v", 409, o.Payload)
}

func (o *DescribeDNSServerConflict) String() string {
	return fmt.Sprintf("[GET /dns-server/{dns}/detail][%d] describeDnsServerConflict  %+v", 409, o.Payload)
}

func (o *DescribeDNSServerConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *DescribeDNSServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeDNSServerInternalServerError creates a DescribeDNSServerInternalServerError with default headers values
func NewDescribeDNSServerInternalServerError() *DescribeDNSServerInternalServerError {
	return &DescribeDNSServerInternalServerError{}
}

/*
DescribeDNSServerInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type DescribeDNSServerInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this describe Dns server internal server error response has a 2xx status code
func (o *DescribeDNSServerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe Dns server internal server error response has a 3xx status code
func (o *DescribeDNSServerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe Dns server internal server error response has a 4xx status code
func (o *DescribeDNSServerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe Dns server internal server error response has a 5xx status code
func (o *DescribeDNSServerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this describe Dns server internal server error response a status code equal to that given
func (o *DescribeDNSServerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the describe Dns server internal server error response
func (o *DescribeDNSServerInternalServerError) Code() int {
	return 500
}

func (o *DescribeDNSServerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dns-server/{dns}/detail][%d] describeDnsServerInternalServerError  %+v", 500, o.Payload)
}

func (o *DescribeDNSServerInternalServerError) String() string {
	return fmt.Sprintf("[GET /dns-server/{dns}/detail][%d] describeDnsServerInternalServerError  %+v", 500, o.Payload)
}

func (o *DescribeDNSServerInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *DescribeDNSServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}