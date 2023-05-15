// Code generated by go-swagger; DO NOT EDIT.

package security_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// AddSecurityGroupReader is a Reader for the AddSecurityGroup structure.
type AddSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddSecurityGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddSecurityGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddSecurityGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddSecurityGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddSecurityGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddSecurityGroupNoContent creates a AddSecurityGroupNoContent with default headers values
func NewAddSecurityGroupNoContent() *AddSecurityGroupNoContent {
	return &AddSecurityGroupNoContent{}
}

/*
AddSecurityGroupNoContent describes a response with status code 204, with default header values.

ok
*/
type AddSecurityGroupNoContent struct {
}

// IsSuccess returns true when this add security group no content response has a 2xx status code
func (o *AddSecurityGroupNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add security group no content response has a 3xx status code
func (o *AddSecurityGroupNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add security group no content response has a 4xx status code
func (o *AddSecurityGroupNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add security group no content response has a 5xx status code
func (o *AddSecurityGroupNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add security group no content response a status code equal to that given
func (o *AddSecurityGroupNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the add security group no content response
func (o *AddSecurityGroupNoContent) Code() int {
	return 204
}

func (o *AddSecurityGroupNoContent) Error() string {
	return fmt.Sprintf("[POST /security-group][%d] addSecurityGroupNoContent ", 204)
}

func (o *AddSecurityGroupNoContent) String() string {
	return fmt.Sprintf("[POST /security-group][%d] addSecurityGroupNoContent ", 204)
}

func (o *AddSecurityGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSecurityGroupBadRequest creates a AddSecurityGroupBadRequest with default headers values
func NewAddSecurityGroupBadRequest() *AddSecurityGroupBadRequest {
	return &AddSecurityGroupBadRequest{}
}

/*
AddSecurityGroupBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type AddSecurityGroupBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this add security group bad request response has a 2xx status code
func (o *AddSecurityGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add security group bad request response has a 3xx status code
func (o *AddSecurityGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add security group bad request response has a 4xx status code
func (o *AddSecurityGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add security group bad request response has a 5xx status code
func (o *AddSecurityGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add security group bad request response a status code equal to that given
func (o *AddSecurityGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add security group bad request response
func (o *AddSecurityGroupBadRequest) Code() int {
	return 400
}

func (o *AddSecurityGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /security-group][%d] addSecurityGroupBadRequest  %+v", 400, o.Payload)
}

func (o *AddSecurityGroupBadRequest) String() string {
	return fmt.Sprintf("[POST /security-group][%d] addSecurityGroupBadRequest  %+v", 400, o.Payload)
}

func (o *AddSecurityGroupBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *AddSecurityGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSecurityGroupNotFound creates a AddSecurityGroupNotFound with default headers values
func NewAddSecurityGroupNotFound() *AddSecurityGroupNotFound {
	return &AddSecurityGroupNotFound{}
}

/*
AddSecurityGroupNotFound describes a response with status code 404, with default header values.

resource not found
*/
type AddSecurityGroupNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this add security group not found response has a 2xx status code
func (o *AddSecurityGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add security group not found response has a 3xx status code
func (o *AddSecurityGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add security group not found response has a 4xx status code
func (o *AddSecurityGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add security group not found response has a 5xx status code
func (o *AddSecurityGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add security group not found response a status code equal to that given
func (o *AddSecurityGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add security group not found response
func (o *AddSecurityGroupNotFound) Code() int {
	return 404
}

func (o *AddSecurityGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /security-group][%d] addSecurityGroupNotFound  %+v", 404, o.Payload)
}

func (o *AddSecurityGroupNotFound) String() string {
	return fmt.Sprintf("[POST /security-group][%d] addSecurityGroupNotFound  %+v", 404, o.Payload)
}

func (o *AddSecurityGroupNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *AddSecurityGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSecurityGroupConflict creates a AddSecurityGroupConflict with default headers values
func NewAddSecurityGroupConflict() *AddSecurityGroupConflict {
	return &AddSecurityGroupConflict{}
}

/*
AddSecurityGroupConflict describes a response with status code 409, with default header values.

conflict
*/
type AddSecurityGroupConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this add security group conflict response has a 2xx status code
func (o *AddSecurityGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add security group conflict response has a 3xx status code
func (o *AddSecurityGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add security group conflict response has a 4xx status code
func (o *AddSecurityGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add security group conflict response has a 5xx status code
func (o *AddSecurityGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add security group conflict response a status code equal to that given
func (o *AddSecurityGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add security group conflict response
func (o *AddSecurityGroupConflict) Code() int {
	return 409
}

func (o *AddSecurityGroupConflict) Error() string {
	return fmt.Sprintf("[POST /security-group][%d] addSecurityGroupConflict  %+v", 409, o.Payload)
}

func (o *AddSecurityGroupConflict) String() string {
	return fmt.Sprintf("[POST /security-group][%d] addSecurityGroupConflict  %+v", 409, o.Payload)
}

func (o *AddSecurityGroupConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *AddSecurityGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSecurityGroupInternalServerError creates a AddSecurityGroupInternalServerError with default headers values
func NewAddSecurityGroupInternalServerError() *AddSecurityGroupInternalServerError {
	return &AddSecurityGroupInternalServerError{}
}

/*
AddSecurityGroupInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type AddSecurityGroupInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this add security group internal server error response has a 2xx status code
func (o *AddSecurityGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add security group internal server error response has a 3xx status code
func (o *AddSecurityGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add security group internal server error response has a 4xx status code
func (o *AddSecurityGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add security group internal server error response has a 5xx status code
func (o *AddSecurityGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add security group internal server error response a status code equal to that given
func (o *AddSecurityGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add security group internal server error response
func (o *AddSecurityGroupInternalServerError) Code() int {
	return 500
}

func (o *AddSecurityGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /security-group][%d] addSecurityGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *AddSecurityGroupInternalServerError) String() string {
	return fmt.Sprintf("[POST /security-group][%d] addSecurityGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *AddSecurityGroupInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *AddSecurityGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
