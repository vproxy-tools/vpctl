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

// RemoveSecurityGroupRuleReader is a Reader for the RemoveSecurityGroupRule structure.
type RemoveSecurityGroupRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveSecurityGroupRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveSecurityGroupRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveSecurityGroupRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveSecurityGroupRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveSecurityGroupRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveSecurityGroupRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveSecurityGroupRuleNoContent creates a RemoveSecurityGroupRuleNoContent with default headers values
func NewRemoveSecurityGroupRuleNoContent() *RemoveSecurityGroupRuleNoContent {
	return &RemoveSecurityGroupRuleNoContent{}
}

/*
RemoveSecurityGroupRuleNoContent describes a response with status code 204, with default header values.

ok
*/
type RemoveSecurityGroupRuleNoContent struct {
}

// IsSuccess returns true when this remove security group rule no content response has a 2xx status code
func (o *RemoveSecurityGroupRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove security group rule no content response has a 3xx status code
func (o *RemoveSecurityGroupRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove security group rule no content response has a 4xx status code
func (o *RemoveSecurityGroupRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove security group rule no content response has a 5xx status code
func (o *RemoveSecurityGroupRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this remove security group rule no content response a status code equal to that given
func (o *RemoveSecurityGroupRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the remove security group rule no content response
func (o *RemoveSecurityGroupRuleNoContent) Code() int {
	return 204
}

func (o *RemoveSecurityGroupRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}/security-group-rule/{secgr}][%d] removeSecurityGroupRuleNoContent ", 204)
}

func (o *RemoveSecurityGroupRuleNoContent) String() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}/security-group-rule/{secgr}][%d] removeSecurityGroupRuleNoContent ", 204)
}

func (o *RemoveSecurityGroupRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveSecurityGroupRuleBadRequest creates a RemoveSecurityGroupRuleBadRequest with default headers values
func NewRemoveSecurityGroupRuleBadRequest() *RemoveSecurityGroupRuleBadRequest {
	return &RemoveSecurityGroupRuleBadRequest{}
}

/*
RemoveSecurityGroupRuleBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type RemoveSecurityGroupRuleBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this remove security group rule bad request response has a 2xx status code
func (o *RemoveSecurityGroupRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove security group rule bad request response has a 3xx status code
func (o *RemoveSecurityGroupRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove security group rule bad request response has a 4xx status code
func (o *RemoveSecurityGroupRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove security group rule bad request response has a 5xx status code
func (o *RemoveSecurityGroupRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove security group rule bad request response a status code equal to that given
func (o *RemoveSecurityGroupRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove security group rule bad request response
func (o *RemoveSecurityGroupRuleBadRequest) Code() int {
	return 400
}

func (o *RemoveSecurityGroupRuleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}/security-group-rule/{secgr}][%d] removeSecurityGroupRuleBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveSecurityGroupRuleBadRequest) String() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}/security-group-rule/{secgr}][%d] removeSecurityGroupRuleBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveSecurityGroupRuleBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *RemoveSecurityGroupRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSecurityGroupRuleNotFound creates a RemoveSecurityGroupRuleNotFound with default headers values
func NewRemoveSecurityGroupRuleNotFound() *RemoveSecurityGroupRuleNotFound {
	return &RemoveSecurityGroupRuleNotFound{}
}

/*
RemoveSecurityGroupRuleNotFound describes a response with status code 404, with default header values.

resource not found
*/
type RemoveSecurityGroupRuleNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this remove security group rule not found response has a 2xx status code
func (o *RemoveSecurityGroupRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove security group rule not found response has a 3xx status code
func (o *RemoveSecurityGroupRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove security group rule not found response has a 4xx status code
func (o *RemoveSecurityGroupRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove security group rule not found response has a 5xx status code
func (o *RemoveSecurityGroupRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove security group rule not found response a status code equal to that given
func (o *RemoveSecurityGroupRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the remove security group rule not found response
func (o *RemoveSecurityGroupRuleNotFound) Code() int {
	return 404
}

func (o *RemoveSecurityGroupRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}/security-group-rule/{secgr}][%d] removeSecurityGroupRuleNotFound  %+v", 404, o.Payload)
}

func (o *RemoveSecurityGroupRuleNotFound) String() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}/security-group-rule/{secgr}][%d] removeSecurityGroupRuleNotFound  %+v", 404, o.Payload)
}

func (o *RemoveSecurityGroupRuleNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *RemoveSecurityGroupRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSecurityGroupRuleConflict creates a RemoveSecurityGroupRuleConflict with default headers values
func NewRemoveSecurityGroupRuleConflict() *RemoveSecurityGroupRuleConflict {
	return &RemoveSecurityGroupRuleConflict{}
}

/*
RemoveSecurityGroupRuleConflict describes a response with status code 409, with default header values.

conflict
*/
type RemoveSecurityGroupRuleConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this remove security group rule conflict response has a 2xx status code
func (o *RemoveSecurityGroupRuleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove security group rule conflict response has a 3xx status code
func (o *RemoveSecurityGroupRuleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove security group rule conflict response has a 4xx status code
func (o *RemoveSecurityGroupRuleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove security group rule conflict response has a 5xx status code
func (o *RemoveSecurityGroupRuleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this remove security group rule conflict response a status code equal to that given
func (o *RemoveSecurityGroupRuleConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the remove security group rule conflict response
func (o *RemoveSecurityGroupRuleConflict) Code() int {
	return 409
}

func (o *RemoveSecurityGroupRuleConflict) Error() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}/security-group-rule/{secgr}][%d] removeSecurityGroupRuleConflict  %+v", 409, o.Payload)
}

func (o *RemoveSecurityGroupRuleConflict) String() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}/security-group-rule/{secgr}][%d] removeSecurityGroupRuleConflict  %+v", 409, o.Payload)
}

func (o *RemoveSecurityGroupRuleConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *RemoveSecurityGroupRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSecurityGroupRuleInternalServerError creates a RemoveSecurityGroupRuleInternalServerError with default headers values
func NewRemoveSecurityGroupRuleInternalServerError() *RemoveSecurityGroupRuleInternalServerError {
	return &RemoveSecurityGroupRuleInternalServerError{}
}

/*
RemoveSecurityGroupRuleInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type RemoveSecurityGroupRuleInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this remove security group rule internal server error response has a 2xx status code
func (o *RemoveSecurityGroupRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove security group rule internal server error response has a 3xx status code
func (o *RemoveSecurityGroupRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove security group rule internal server error response has a 4xx status code
func (o *RemoveSecurityGroupRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove security group rule internal server error response has a 5xx status code
func (o *RemoveSecurityGroupRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove security group rule internal server error response a status code equal to that given
func (o *RemoveSecurityGroupRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove security group rule internal server error response
func (o *RemoveSecurityGroupRuleInternalServerError) Code() int {
	return 500
}

func (o *RemoveSecurityGroupRuleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}/security-group-rule/{secgr}][%d] removeSecurityGroupRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveSecurityGroupRuleInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}/security-group-rule/{secgr}][%d] removeSecurityGroupRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveSecurityGroupRuleInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *RemoveSecurityGroupRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
