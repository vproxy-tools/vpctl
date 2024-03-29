// Code generated by go-swagger; DO NOT EDIT.

package event_loop_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// ListEventLoopGroupReader is a Reader for the ListEventLoopGroup structure.
type ListEventLoopGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEventLoopGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEventLoopGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListEventLoopGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListEventLoopGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListEventLoopGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListEventLoopGroupOK creates a ListEventLoopGroupOK with default headers values
func NewListEventLoopGroupOK() *ListEventLoopGroupOK {
	return &ListEventLoopGroupOK{}
}

/*
ListEventLoopGroupOK describes a response with status code 200, with default header values.

ok
*/
type ListEventLoopGroupOK struct {
	Payload []*vproxy_client_model.EventLoopGroup
}

// IsSuccess returns true when this list event loop group o k response has a 2xx status code
func (o *ListEventLoopGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list event loop group o k response has a 3xx status code
func (o *ListEventLoopGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list event loop group o k response has a 4xx status code
func (o *ListEventLoopGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list event loop group o k response has a 5xx status code
func (o *ListEventLoopGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list event loop group o k response a status code equal to that given
func (o *ListEventLoopGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list event loop group o k response
func (o *ListEventLoopGroupOK) Code() int {
	return 200
}

func (o *ListEventLoopGroupOK) Error() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupOK  %+v", 200, o.Payload)
}

func (o *ListEventLoopGroupOK) String() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupOK  %+v", 200, o.Payload)
}

func (o *ListEventLoopGroupOK) GetPayload() []*vproxy_client_model.EventLoopGroup {
	return o.Payload
}

func (o *ListEventLoopGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEventLoopGroupBadRequest creates a ListEventLoopGroupBadRequest with default headers values
func NewListEventLoopGroupBadRequest() *ListEventLoopGroupBadRequest {
	return &ListEventLoopGroupBadRequest{}
}

/*
ListEventLoopGroupBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type ListEventLoopGroupBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this list event loop group bad request response has a 2xx status code
func (o *ListEventLoopGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list event loop group bad request response has a 3xx status code
func (o *ListEventLoopGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list event loop group bad request response has a 4xx status code
func (o *ListEventLoopGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list event loop group bad request response has a 5xx status code
func (o *ListEventLoopGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list event loop group bad request response a status code equal to that given
func (o *ListEventLoopGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list event loop group bad request response
func (o *ListEventLoopGroupBadRequest) Code() int {
	return 400
}

func (o *ListEventLoopGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupBadRequest  %+v", 400, o.Payload)
}

func (o *ListEventLoopGroupBadRequest) String() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupBadRequest  %+v", 400, o.Payload)
}

func (o *ListEventLoopGroupBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *ListEventLoopGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEventLoopGroupConflict creates a ListEventLoopGroupConflict with default headers values
func NewListEventLoopGroupConflict() *ListEventLoopGroupConflict {
	return &ListEventLoopGroupConflict{}
}

/*
ListEventLoopGroupConflict describes a response with status code 409, with default header values.

conflict
*/
type ListEventLoopGroupConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this list event loop group conflict response has a 2xx status code
func (o *ListEventLoopGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list event loop group conflict response has a 3xx status code
func (o *ListEventLoopGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list event loop group conflict response has a 4xx status code
func (o *ListEventLoopGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this list event loop group conflict response has a 5xx status code
func (o *ListEventLoopGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this list event loop group conflict response a status code equal to that given
func (o *ListEventLoopGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the list event loop group conflict response
func (o *ListEventLoopGroupConflict) Code() int {
	return 409
}

func (o *ListEventLoopGroupConflict) Error() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupConflict  %+v", 409, o.Payload)
}

func (o *ListEventLoopGroupConflict) String() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupConflict  %+v", 409, o.Payload)
}

func (o *ListEventLoopGroupConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *ListEventLoopGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEventLoopGroupInternalServerError creates a ListEventLoopGroupInternalServerError with default headers values
func NewListEventLoopGroupInternalServerError() *ListEventLoopGroupInternalServerError {
	return &ListEventLoopGroupInternalServerError{}
}

/*
ListEventLoopGroupInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type ListEventLoopGroupInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this list event loop group internal server error response has a 2xx status code
func (o *ListEventLoopGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list event loop group internal server error response has a 3xx status code
func (o *ListEventLoopGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list event loop group internal server error response has a 4xx status code
func (o *ListEventLoopGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list event loop group internal server error response has a 5xx status code
func (o *ListEventLoopGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list event loop group internal server error response a status code equal to that given
func (o *ListEventLoopGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list event loop group internal server error response
func (o *ListEventLoopGroupInternalServerError) Code() int {
	return 500
}

func (o *ListEventLoopGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ListEventLoopGroupInternalServerError) String() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ListEventLoopGroupInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *ListEventLoopGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
