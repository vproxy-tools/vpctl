// Code generated by go-swagger; DO NOT EDIT.

package security_group_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// AddSecurityGroupRuleReader is a Reader for the AddSecurityGroupRule structure.
type AddSecurityGroupRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSecurityGroupRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddSecurityGroupRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddSecurityGroupRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddSecurityGroupRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddSecurityGroupRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddSecurityGroupRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddSecurityGroupRuleNoContent creates a AddSecurityGroupRuleNoContent with default headers values
func NewAddSecurityGroupRuleNoContent() *AddSecurityGroupRuleNoContent {
	return &AddSecurityGroupRuleNoContent{}
}

/*
AddSecurityGroupRuleNoContent describes a response with status code 204, with default header values.

ok
*/
type AddSecurityGroupRuleNoContent struct {
}

// IsSuccess returns true when this add security group rule no content response has a 2xx status code
func (o *AddSecurityGroupRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add security group rule no content response has a 3xx status code
func (o *AddSecurityGroupRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add security group rule no content response has a 4xx status code
func (o *AddSecurityGroupRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add security group rule no content response has a 5xx status code
func (o *AddSecurityGroupRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add security group rule no content response a status code equal to that given
func (o *AddSecurityGroupRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the add security group rule no content response
func (o *AddSecurityGroupRuleNoContent) Code() int {
	return 204
}

func (o *AddSecurityGroupRuleNoContent) Error() string {
	return fmt.Sprintf("[POST /security-group/{secg}/security-group-rule][%d] addSecurityGroupRuleNoContent ", 204)
}

func (o *AddSecurityGroupRuleNoContent) String() string {
	return fmt.Sprintf("[POST /security-group/{secg}/security-group-rule][%d] addSecurityGroupRuleNoContent ", 204)
}

func (o *AddSecurityGroupRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSecurityGroupRuleBadRequest creates a AddSecurityGroupRuleBadRequest with default headers values
func NewAddSecurityGroupRuleBadRequest() *AddSecurityGroupRuleBadRequest {
	return &AddSecurityGroupRuleBadRequest{}
}

/*
AddSecurityGroupRuleBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type AddSecurityGroupRuleBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this add security group rule bad request response has a 2xx status code
func (o *AddSecurityGroupRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add security group rule bad request response has a 3xx status code
func (o *AddSecurityGroupRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add security group rule bad request response has a 4xx status code
func (o *AddSecurityGroupRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add security group rule bad request response has a 5xx status code
func (o *AddSecurityGroupRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add security group rule bad request response a status code equal to that given
func (o *AddSecurityGroupRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add security group rule bad request response
func (o *AddSecurityGroupRuleBadRequest) Code() int {
	return 400
}

func (o *AddSecurityGroupRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /security-group/{secg}/security-group-rule][%d] addSecurityGroupRuleBadRequest  %+v", 400, o.Payload)
}

func (o *AddSecurityGroupRuleBadRequest) String() string {
	return fmt.Sprintf("[POST /security-group/{secg}/security-group-rule][%d] addSecurityGroupRuleBadRequest  %+v", 400, o.Payload)
}

func (o *AddSecurityGroupRuleBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *AddSecurityGroupRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSecurityGroupRuleNotFound creates a AddSecurityGroupRuleNotFound with default headers values
func NewAddSecurityGroupRuleNotFound() *AddSecurityGroupRuleNotFound {
	return &AddSecurityGroupRuleNotFound{}
}

/*
AddSecurityGroupRuleNotFound describes a response with status code 404, with default header values.

resource not found
*/
type AddSecurityGroupRuleNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this add security group rule not found response has a 2xx status code
func (o *AddSecurityGroupRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add security group rule not found response has a 3xx status code
func (o *AddSecurityGroupRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add security group rule not found response has a 4xx status code
func (o *AddSecurityGroupRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add security group rule not found response has a 5xx status code
func (o *AddSecurityGroupRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add security group rule not found response a status code equal to that given
func (o *AddSecurityGroupRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add security group rule not found response
func (o *AddSecurityGroupRuleNotFound) Code() int {
	return 404
}

func (o *AddSecurityGroupRuleNotFound) Error() string {
	return fmt.Sprintf("[POST /security-group/{secg}/security-group-rule][%d] addSecurityGroupRuleNotFound  %+v", 404, o.Payload)
}

func (o *AddSecurityGroupRuleNotFound) String() string {
	return fmt.Sprintf("[POST /security-group/{secg}/security-group-rule][%d] addSecurityGroupRuleNotFound  %+v", 404, o.Payload)
}

func (o *AddSecurityGroupRuleNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *AddSecurityGroupRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSecurityGroupRuleConflict creates a AddSecurityGroupRuleConflict with default headers values
func NewAddSecurityGroupRuleConflict() *AddSecurityGroupRuleConflict {
	return &AddSecurityGroupRuleConflict{}
}

/*
AddSecurityGroupRuleConflict describes a response with status code 409, with default header values.

conflict
*/
type AddSecurityGroupRuleConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this add security group rule conflict response has a 2xx status code
func (o *AddSecurityGroupRuleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add security group rule conflict response has a 3xx status code
func (o *AddSecurityGroupRuleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add security group rule conflict response has a 4xx status code
func (o *AddSecurityGroupRuleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add security group rule conflict response has a 5xx status code
func (o *AddSecurityGroupRuleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add security group rule conflict response a status code equal to that given
func (o *AddSecurityGroupRuleConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add security group rule conflict response
func (o *AddSecurityGroupRuleConflict) Code() int {
	return 409
}

func (o *AddSecurityGroupRuleConflict) Error() string {
	return fmt.Sprintf("[POST /security-group/{secg}/security-group-rule][%d] addSecurityGroupRuleConflict  %+v", 409, o.Payload)
}

func (o *AddSecurityGroupRuleConflict) String() string {
	return fmt.Sprintf("[POST /security-group/{secg}/security-group-rule][%d] addSecurityGroupRuleConflict  %+v", 409, o.Payload)
}

func (o *AddSecurityGroupRuleConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *AddSecurityGroupRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSecurityGroupRuleInternalServerError creates a AddSecurityGroupRuleInternalServerError with default headers values
func NewAddSecurityGroupRuleInternalServerError() *AddSecurityGroupRuleInternalServerError {
	return &AddSecurityGroupRuleInternalServerError{}
}

/*
AddSecurityGroupRuleInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type AddSecurityGroupRuleInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this add security group rule internal server error response has a 2xx status code
func (o *AddSecurityGroupRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add security group rule internal server error response has a 3xx status code
func (o *AddSecurityGroupRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add security group rule internal server error response has a 4xx status code
func (o *AddSecurityGroupRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add security group rule internal server error response has a 5xx status code
func (o *AddSecurityGroupRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add security group rule internal server error response a status code equal to that given
func (o *AddSecurityGroupRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add security group rule internal server error response
func (o *AddSecurityGroupRuleInternalServerError) Code() int {
	return 500
}

func (o *AddSecurityGroupRuleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /security-group/{secg}/security-group-rule][%d] addSecurityGroupRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *AddSecurityGroupRuleInternalServerError) String() string {
	return fmt.Sprintf("[POST /security-group/{secg}/security-group-rule][%d] addSecurityGroupRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *AddSecurityGroupRuleInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *AddSecurityGroupRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
