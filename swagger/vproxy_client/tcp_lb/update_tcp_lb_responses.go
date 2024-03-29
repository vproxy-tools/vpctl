// Code generated by go-swagger; DO NOT EDIT.

package tcp_lb

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// UpdateTCPLbReader is a Reader for the UpdateTCPLb structure.
type UpdateTCPLbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTCPLbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateTCPLbNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTCPLbBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTCPLbNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateTCPLbConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTCPLbInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTCPLbNoContent creates a UpdateTCPLbNoContent with default headers values
func NewUpdateTCPLbNoContent() *UpdateTCPLbNoContent {
	return &UpdateTCPLbNoContent{}
}

/*
UpdateTCPLbNoContent describes a response with status code 204, with default header values.

ok
*/
type UpdateTCPLbNoContent struct {
}

// IsSuccess returns true when this update Tcp lb no content response has a 2xx status code
func (o *UpdateTCPLbNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update Tcp lb no content response has a 3xx status code
func (o *UpdateTCPLbNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Tcp lb no content response has a 4xx status code
func (o *UpdateTCPLbNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Tcp lb no content response has a 5xx status code
func (o *UpdateTCPLbNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update Tcp lb no content response a status code equal to that given
func (o *UpdateTCPLbNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update Tcp lb no content response
func (o *UpdateTCPLbNoContent) Code() int {
	return 204
}

func (o *UpdateTCPLbNoContent) Error() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbNoContent ", 204)
}

func (o *UpdateTCPLbNoContent) String() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbNoContent ", 204)
}

func (o *UpdateTCPLbNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTCPLbBadRequest creates a UpdateTCPLbBadRequest with default headers values
func NewUpdateTCPLbBadRequest() *UpdateTCPLbBadRequest {
	return &UpdateTCPLbBadRequest{}
}

/*
UpdateTCPLbBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type UpdateTCPLbBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this update Tcp lb bad request response has a 2xx status code
func (o *UpdateTCPLbBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Tcp lb bad request response has a 3xx status code
func (o *UpdateTCPLbBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Tcp lb bad request response has a 4xx status code
func (o *UpdateTCPLbBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Tcp lb bad request response has a 5xx status code
func (o *UpdateTCPLbBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update Tcp lb bad request response a status code equal to that given
func (o *UpdateTCPLbBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update Tcp lb bad request response
func (o *UpdateTCPLbBadRequest) Code() int {
	return 400
}

func (o *UpdateTCPLbBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTCPLbBadRequest) String() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTCPLbBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *UpdateTCPLbBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTCPLbNotFound creates a UpdateTCPLbNotFound with default headers values
func NewUpdateTCPLbNotFound() *UpdateTCPLbNotFound {
	return &UpdateTCPLbNotFound{}
}

/*
UpdateTCPLbNotFound describes a response with status code 404, with default header values.

resource not found
*/
type UpdateTCPLbNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this update Tcp lb not found response has a 2xx status code
func (o *UpdateTCPLbNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Tcp lb not found response has a 3xx status code
func (o *UpdateTCPLbNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Tcp lb not found response has a 4xx status code
func (o *UpdateTCPLbNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Tcp lb not found response has a 5xx status code
func (o *UpdateTCPLbNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update Tcp lb not found response a status code equal to that given
func (o *UpdateTCPLbNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update Tcp lb not found response
func (o *UpdateTCPLbNotFound) Code() int {
	return 404
}

func (o *UpdateTCPLbNotFound) Error() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTCPLbNotFound) String() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTCPLbNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *UpdateTCPLbNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTCPLbConflict creates a UpdateTCPLbConflict with default headers values
func NewUpdateTCPLbConflict() *UpdateTCPLbConflict {
	return &UpdateTCPLbConflict{}
}

/*
UpdateTCPLbConflict describes a response with status code 409, with default header values.

conflict
*/
type UpdateTCPLbConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this update Tcp lb conflict response has a 2xx status code
func (o *UpdateTCPLbConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Tcp lb conflict response has a 3xx status code
func (o *UpdateTCPLbConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Tcp lb conflict response has a 4xx status code
func (o *UpdateTCPLbConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Tcp lb conflict response has a 5xx status code
func (o *UpdateTCPLbConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update Tcp lb conflict response a status code equal to that given
func (o *UpdateTCPLbConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update Tcp lb conflict response
func (o *UpdateTCPLbConflict) Code() int {
	return 409
}

func (o *UpdateTCPLbConflict) Error() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbConflict  %+v", 409, o.Payload)
}

func (o *UpdateTCPLbConflict) String() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbConflict  %+v", 409, o.Payload)
}

func (o *UpdateTCPLbConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *UpdateTCPLbConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTCPLbInternalServerError creates a UpdateTCPLbInternalServerError with default headers values
func NewUpdateTCPLbInternalServerError() *UpdateTCPLbInternalServerError {
	return &UpdateTCPLbInternalServerError{}
}

/*
UpdateTCPLbInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type UpdateTCPLbInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this update Tcp lb internal server error response has a 2xx status code
func (o *UpdateTCPLbInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Tcp lb internal server error response has a 3xx status code
func (o *UpdateTCPLbInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Tcp lb internal server error response has a 4xx status code
func (o *UpdateTCPLbInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Tcp lb internal server error response has a 5xx status code
func (o *UpdateTCPLbInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update Tcp lb internal server error response a status code equal to that given
func (o *UpdateTCPLbInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update Tcp lb internal server error response
func (o *UpdateTCPLbInternalServerError) Code() int {
	return 500
}

func (o *UpdateTCPLbInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTCPLbInternalServerError) String() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTCPLbInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *UpdateTCPLbInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
