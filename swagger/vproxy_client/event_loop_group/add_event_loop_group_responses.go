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

// AddEventLoopGroupReader is a Reader for the AddEventLoopGroup structure.
type AddEventLoopGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddEventLoopGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddEventLoopGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddEventLoopGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddEventLoopGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddEventLoopGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddEventLoopGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddEventLoopGroupNoContent creates a AddEventLoopGroupNoContent with default headers values
func NewAddEventLoopGroupNoContent() *AddEventLoopGroupNoContent {
	return &AddEventLoopGroupNoContent{}
}

/*
AddEventLoopGroupNoContent describes a response with status code 204, with default header values.

ok
*/
type AddEventLoopGroupNoContent struct {
}

// IsSuccess returns true when this add event loop group no content response has a 2xx status code
func (o *AddEventLoopGroupNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add event loop group no content response has a 3xx status code
func (o *AddEventLoopGroupNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add event loop group no content response has a 4xx status code
func (o *AddEventLoopGroupNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add event loop group no content response has a 5xx status code
func (o *AddEventLoopGroupNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add event loop group no content response a status code equal to that given
func (o *AddEventLoopGroupNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the add event loop group no content response
func (o *AddEventLoopGroupNoContent) Code() int {
	return 204
}

func (o *AddEventLoopGroupNoContent) Error() string {
	return fmt.Sprintf("[POST /event-loop-group][%d] addEventLoopGroupNoContent ", 204)
}

func (o *AddEventLoopGroupNoContent) String() string {
	return fmt.Sprintf("[POST /event-loop-group][%d] addEventLoopGroupNoContent ", 204)
}

func (o *AddEventLoopGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddEventLoopGroupBadRequest creates a AddEventLoopGroupBadRequest with default headers values
func NewAddEventLoopGroupBadRequest() *AddEventLoopGroupBadRequest {
	return &AddEventLoopGroupBadRequest{}
}

/*
AddEventLoopGroupBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type AddEventLoopGroupBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this add event loop group bad request response has a 2xx status code
func (o *AddEventLoopGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add event loop group bad request response has a 3xx status code
func (o *AddEventLoopGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add event loop group bad request response has a 4xx status code
func (o *AddEventLoopGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add event loop group bad request response has a 5xx status code
func (o *AddEventLoopGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add event loop group bad request response a status code equal to that given
func (o *AddEventLoopGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add event loop group bad request response
func (o *AddEventLoopGroupBadRequest) Code() int {
	return 400
}

func (o *AddEventLoopGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /event-loop-group][%d] addEventLoopGroupBadRequest  %+v", 400, o.Payload)
}

func (o *AddEventLoopGroupBadRequest) String() string {
	return fmt.Sprintf("[POST /event-loop-group][%d] addEventLoopGroupBadRequest  %+v", 400, o.Payload)
}

func (o *AddEventLoopGroupBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *AddEventLoopGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEventLoopGroupNotFound creates a AddEventLoopGroupNotFound with default headers values
func NewAddEventLoopGroupNotFound() *AddEventLoopGroupNotFound {
	return &AddEventLoopGroupNotFound{}
}

/*
AddEventLoopGroupNotFound describes a response with status code 404, with default header values.

resource not found
*/
type AddEventLoopGroupNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this add event loop group not found response has a 2xx status code
func (o *AddEventLoopGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add event loop group not found response has a 3xx status code
func (o *AddEventLoopGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add event loop group not found response has a 4xx status code
func (o *AddEventLoopGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add event loop group not found response has a 5xx status code
func (o *AddEventLoopGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add event loop group not found response a status code equal to that given
func (o *AddEventLoopGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add event loop group not found response
func (o *AddEventLoopGroupNotFound) Code() int {
	return 404
}

func (o *AddEventLoopGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /event-loop-group][%d] addEventLoopGroupNotFound  %+v", 404, o.Payload)
}

func (o *AddEventLoopGroupNotFound) String() string {
	return fmt.Sprintf("[POST /event-loop-group][%d] addEventLoopGroupNotFound  %+v", 404, o.Payload)
}

func (o *AddEventLoopGroupNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *AddEventLoopGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEventLoopGroupConflict creates a AddEventLoopGroupConflict with default headers values
func NewAddEventLoopGroupConflict() *AddEventLoopGroupConflict {
	return &AddEventLoopGroupConflict{}
}

/*
AddEventLoopGroupConflict describes a response with status code 409, with default header values.

conflict
*/
type AddEventLoopGroupConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this add event loop group conflict response has a 2xx status code
func (o *AddEventLoopGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add event loop group conflict response has a 3xx status code
func (o *AddEventLoopGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add event loop group conflict response has a 4xx status code
func (o *AddEventLoopGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add event loop group conflict response has a 5xx status code
func (o *AddEventLoopGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add event loop group conflict response a status code equal to that given
func (o *AddEventLoopGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add event loop group conflict response
func (o *AddEventLoopGroupConflict) Code() int {
	return 409
}

func (o *AddEventLoopGroupConflict) Error() string {
	return fmt.Sprintf("[POST /event-loop-group][%d] addEventLoopGroupConflict  %+v", 409, o.Payload)
}

func (o *AddEventLoopGroupConflict) String() string {
	return fmt.Sprintf("[POST /event-loop-group][%d] addEventLoopGroupConflict  %+v", 409, o.Payload)
}

func (o *AddEventLoopGroupConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *AddEventLoopGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEventLoopGroupInternalServerError creates a AddEventLoopGroupInternalServerError with default headers values
func NewAddEventLoopGroupInternalServerError() *AddEventLoopGroupInternalServerError {
	return &AddEventLoopGroupInternalServerError{}
}

/*
AddEventLoopGroupInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type AddEventLoopGroupInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this add event loop group internal server error response has a 2xx status code
func (o *AddEventLoopGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add event loop group internal server error response has a 3xx status code
func (o *AddEventLoopGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add event loop group internal server error response has a 4xx status code
func (o *AddEventLoopGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add event loop group internal server error response has a 5xx status code
func (o *AddEventLoopGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add event loop group internal server error response a status code equal to that given
func (o *AddEventLoopGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add event loop group internal server error response
func (o *AddEventLoopGroupInternalServerError) Code() int {
	return 500
}

func (o *AddEventLoopGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /event-loop-group][%d] addEventLoopGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *AddEventLoopGroupInternalServerError) String() string {
	return fmt.Sprintf("[POST /event-loop-group][%d] addEventLoopGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *AddEventLoopGroupInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *AddEventLoopGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}