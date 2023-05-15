// Code generated by go-swagger; DO NOT EDIT.

package upstream

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// DescribeUpstreamReader is a Reader for the DescribeUpstream structure.
type DescribeUpstreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeUpstreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeUpstreamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDescribeUpstreamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDescribeUpstreamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDescribeUpstreamConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDescribeUpstreamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDescribeUpstreamOK creates a DescribeUpstreamOK with default headers values
func NewDescribeUpstreamOK() *DescribeUpstreamOK {
	return &DescribeUpstreamOK{}
}

/*
DescribeUpstreamOK describes a response with status code 200, with default header values.

ok
*/
type DescribeUpstreamOK struct {
	Payload *vproxy_client_model.UpstreamDetail
}

// IsSuccess returns true when this describe upstream o k response has a 2xx status code
func (o *DescribeUpstreamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe upstream o k response has a 3xx status code
func (o *DescribeUpstreamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe upstream o k response has a 4xx status code
func (o *DescribeUpstreamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe upstream o k response has a 5xx status code
func (o *DescribeUpstreamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe upstream o k response a status code equal to that given
func (o *DescribeUpstreamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe upstream o k response
func (o *DescribeUpstreamOK) Code() int {
	return 200
}

func (o *DescribeUpstreamOK) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamOK  %+v", 200, o.Payload)
}

func (o *DescribeUpstreamOK) String() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamOK  %+v", 200, o.Payload)
}

func (o *DescribeUpstreamOK) GetPayload() *vproxy_client_model.UpstreamDetail {
	return o.Payload
}

func (o *DescribeUpstreamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.UpstreamDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeUpstreamBadRequest creates a DescribeUpstreamBadRequest with default headers values
func NewDescribeUpstreamBadRequest() *DescribeUpstreamBadRequest {
	return &DescribeUpstreamBadRequest{}
}

/*
DescribeUpstreamBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type DescribeUpstreamBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this describe upstream bad request response has a 2xx status code
func (o *DescribeUpstreamBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe upstream bad request response has a 3xx status code
func (o *DescribeUpstreamBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe upstream bad request response has a 4xx status code
func (o *DescribeUpstreamBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe upstream bad request response has a 5xx status code
func (o *DescribeUpstreamBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this describe upstream bad request response a status code equal to that given
func (o *DescribeUpstreamBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the describe upstream bad request response
func (o *DescribeUpstreamBadRequest) Code() int {
	return 400
}

func (o *DescribeUpstreamBadRequest) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeUpstreamBadRequest) String() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeUpstreamBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *DescribeUpstreamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeUpstreamNotFound creates a DescribeUpstreamNotFound with default headers values
func NewDescribeUpstreamNotFound() *DescribeUpstreamNotFound {
	return &DescribeUpstreamNotFound{}
}

/*
DescribeUpstreamNotFound describes a response with status code 404, with default header values.

resource not found
*/
type DescribeUpstreamNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this describe upstream not found response has a 2xx status code
func (o *DescribeUpstreamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe upstream not found response has a 3xx status code
func (o *DescribeUpstreamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe upstream not found response has a 4xx status code
func (o *DescribeUpstreamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe upstream not found response has a 5xx status code
func (o *DescribeUpstreamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this describe upstream not found response a status code equal to that given
func (o *DescribeUpstreamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the describe upstream not found response
func (o *DescribeUpstreamNotFound) Code() int {
	return 404
}

func (o *DescribeUpstreamNotFound) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamNotFound  %+v", 404, o.Payload)
}

func (o *DescribeUpstreamNotFound) String() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamNotFound  %+v", 404, o.Payload)
}

func (o *DescribeUpstreamNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *DescribeUpstreamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeUpstreamConflict creates a DescribeUpstreamConflict with default headers values
func NewDescribeUpstreamConflict() *DescribeUpstreamConflict {
	return &DescribeUpstreamConflict{}
}

/*
DescribeUpstreamConflict describes a response with status code 409, with default header values.

conflict
*/
type DescribeUpstreamConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this describe upstream conflict response has a 2xx status code
func (o *DescribeUpstreamConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe upstream conflict response has a 3xx status code
func (o *DescribeUpstreamConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe upstream conflict response has a 4xx status code
func (o *DescribeUpstreamConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe upstream conflict response has a 5xx status code
func (o *DescribeUpstreamConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this describe upstream conflict response a status code equal to that given
func (o *DescribeUpstreamConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the describe upstream conflict response
func (o *DescribeUpstreamConflict) Code() int {
	return 409
}

func (o *DescribeUpstreamConflict) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamConflict  %+v", 409, o.Payload)
}

func (o *DescribeUpstreamConflict) String() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamConflict  %+v", 409, o.Payload)
}

func (o *DescribeUpstreamConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *DescribeUpstreamConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeUpstreamInternalServerError creates a DescribeUpstreamInternalServerError with default headers values
func NewDescribeUpstreamInternalServerError() *DescribeUpstreamInternalServerError {
	return &DescribeUpstreamInternalServerError{}
}

/*
DescribeUpstreamInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type DescribeUpstreamInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this describe upstream internal server error response has a 2xx status code
func (o *DescribeUpstreamInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe upstream internal server error response has a 3xx status code
func (o *DescribeUpstreamInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe upstream internal server error response has a 4xx status code
func (o *DescribeUpstreamInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe upstream internal server error response has a 5xx status code
func (o *DescribeUpstreamInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this describe upstream internal server error response a status code equal to that given
func (o *DescribeUpstreamInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the describe upstream internal server error response
func (o *DescribeUpstreamInternalServerError) Code() int {
	return 500
}

func (o *DescribeUpstreamInternalServerError) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamInternalServerError  %+v", 500, o.Payload)
}

func (o *DescribeUpstreamInternalServerError) String() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamInternalServerError  %+v", 500, o.Payload)
}

func (o *DescribeUpstreamInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *DescribeUpstreamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
