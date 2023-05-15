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

// RemoveUpstreamReader is a Reader for the RemoveUpstream structure.
type RemoveUpstreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveUpstreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveUpstreamNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveUpstreamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveUpstreamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveUpstreamConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveUpstreamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveUpstreamNoContent creates a RemoveUpstreamNoContent with default headers values
func NewRemoveUpstreamNoContent() *RemoveUpstreamNoContent {
	return &RemoveUpstreamNoContent{}
}

/*
RemoveUpstreamNoContent describes a response with status code 204, with default header values.

ok
*/
type RemoveUpstreamNoContent struct {
}

// IsSuccess returns true when this remove upstream no content response has a 2xx status code
func (o *RemoveUpstreamNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove upstream no content response has a 3xx status code
func (o *RemoveUpstreamNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove upstream no content response has a 4xx status code
func (o *RemoveUpstreamNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove upstream no content response has a 5xx status code
func (o *RemoveUpstreamNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this remove upstream no content response a status code equal to that given
func (o *RemoveUpstreamNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the remove upstream no content response
func (o *RemoveUpstreamNoContent) Code() int {
	return 204
}

func (o *RemoveUpstreamNoContent) Error() string {
	return fmt.Sprintf("[DELETE /upstream/{ups}][%d] removeUpstreamNoContent ", 204)
}

func (o *RemoveUpstreamNoContent) String() string {
	return fmt.Sprintf("[DELETE /upstream/{ups}][%d] removeUpstreamNoContent ", 204)
}

func (o *RemoveUpstreamNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveUpstreamBadRequest creates a RemoveUpstreamBadRequest with default headers values
func NewRemoveUpstreamBadRequest() *RemoveUpstreamBadRequest {
	return &RemoveUpstreamBadRequest{}
}

/*
RemoveUpstreamBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type RemoveUpstreamBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this remove upstream bad request response has a 2xx status code
func (o *RemoveUpstreamBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove upstream bad request response has a 3xx status code
func (o *RemoveUpstreamBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove upstream bad request response has a 4xx status code
func (o *RemoveUpstreamBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove upstream bad request response has a 5xx status code
func (o *RemoveUpstreamBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove upstream bad request response a status code equal to that given
func (o *RemoveUpstreamBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove upstream bad request response
func (o *RemoveUpstreamBadRequest) Code() int {
	return 400
}

func (o *RemoveUpstreamBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /upstream/{ups}][%d] removeUpstreamBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveUpstreamBadRequest) String() string {
	return fmt.Sprintf("[DELETE /upstream/{ups}][%d] removeUpstreamBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveUpstreamBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *RemoveUpstreamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUpstreamNotFound creates a RemoveUpstreamNotFound with default headers values
func NewRemoveUpstreamNotFound() *RemoveUpstreamNotFound {
	return &RemoveUpstreamNotFound{}
}

/*
RemoveUpstreamNotFound describes a response with status code 404, with default header values.

resource not found
*/
type RemoveUpstreamNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this remove upstream not found response has a 2xx status code
func (o *RemoveUpstreamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove upstream not found response has a 3xx status code
func (o *RemoveUpstreamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove upstream not found response has a 4xx status code
func (o *RemoveUpstreamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove upstream not found response has a 5xx status code
func (o *RemoveUpstreamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove upstream not found response a status code equal to that given
func (o *RemoveUpstreamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the remove upstream not found response
func (o *RemoveUpstreamNotFound) Code() int {
	return 404
}

func (o *RemoveUpstreamNotFound) Error() string {
	return fmt.Sprintf("[DELETE /upstream/{ups}][%d] removeUpstreamNotFound  %+v", 404, o.Payload)
}

func (o *RemoveUpstreamNotFound) String() string {
	return fmt.Sprintf("[DELETE /upstream/{ups}][%d] removeUpstreamNotFound  %+v", 404, o.Payload)
}

func (o *RemoveUpstreamNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *RemoveUpstreamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUpstreamConflict creates a RemoveUpstreamConflict with default headers values
func NewRemoveUpstreamConflict() *RemoveUpstreamConflict {
	return &RemoveUpstreamConflict{}
}

/*
RemoveUpstreamConflict describes a response with status code 409, with default header values.

conflict
*/
type RemoveUpstreamConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this remove upstream conflict response has a 2xx status code
func (o *RemoveUpstreamConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove upstream conflict response has a 3xx status code
func (o *RemoveUpstreamConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove upstream conflict response has a 4xx status code
func (o *RemoveUpstreamConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove upstream conflict response has a 5xx status code
func (o *RemoveUpstreamConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this remove upstream conflict response a status code equal to that given
func (o *RemoveUpstreamConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the remove upstream conflict response
func (o *RemoveUpstreamConflict) Code() int {
	return 409
}

func (o *RemoveUpstreamConflict) Error() string {
	return fmt.Sprintf("[DELETE /upstream/{ups}][%d] removeUpstreamConflict  %+v", 409, o.Payload)
}

func (o *RemoveUpstreamConflict) String() string {
	return fmt.Sprintf("[DELETE /upstream/{ups}][%d] removeUpstreamConflict  %+v", 409, o.Payload)
}

func (o *RemoveUpstreamConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *RemoveUpstreamConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUpstreamInternalServerError creates a RemoveUpstreamInternalServerError with default headers values
func NewRemoveUpstreamInternalServerError() *RemoveUpstreamInternalServerError {
	return &RemoveUpstreamInternalServerError{}
}

/*
RemoveUpstreamInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type RemoveUpstreamInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this remove upstream internal server error response has a 2xx status code
func (o *RemoveUpstreamInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove upstream internal server error response has a 3xx status code
func (o *RemoveUpstreamInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove upstream internal server error response has a 4xx status code
func (o *RemoveUpstreamInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove upstream internal server error response has a 5xx status code
func (o *RemoveUpstreamInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove upstream internal server error response a status code equal to that given
func (o *RemoveUpstreamInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove upstream internal server error response
func (o *RemoveUpstreamInternalServerError) Code() int {
	return 500
}

func (o *RemoveUpstreamInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /upstream/{ups}][%d] removeUpstreamInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveUpstreamInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /upstream/{ups}][%d] removeUpstreamInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveUpstreamInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *RemoveUpstreamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
