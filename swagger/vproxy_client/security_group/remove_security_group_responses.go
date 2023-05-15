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

// RemoveSecurityGroupReader is a Reader for the RemoveSecurityGroup structure.
type RemoveSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveSecurityGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveSecurityGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveSecurityGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveSecurityGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveSecurityGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveSecurityGroupNoContent creates a RemoveSecurityGroupNoContent with default headers values
func NewRemoveSecurityGroupNoContent() *RemoveSecurityGroupNoContent {
	return &RemoveSecurityGroupNoContent{}
}

/*
RemoveSecurityGroupNoContent describes a response with status code 204, with default header values.

ok
*/
type RemoveSecurityGroupNoContent struct {
}

// IsSuccess returns true when this remove security group no content response has a 2xx status code
func (o *RemoveSecurityGroupNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove security group no content response has a 3xx status code
func (o *RemoveSecurityGroupNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove security group no content response has a 4xx status code
func (o *RemoveSecurityGroupNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove security group no content response has a 5xx status code
func (o *RemoveSecurityGroupNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this remove security group no content response a status code equal to that given
func (o *RemoveSecurityGroupNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the remove security group no content response
func (o *RemoveSecurityGroupNoContent) Code() int {
	return 204
}

func (o *RemoveSecurityGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}][%d] removeSecurityGroupNoContent ", 204)
}

func (o *RemoveSecurityGroupNoContent) String() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}][%d] removeSecurityGroupNoContent ", 204)
}

func (o *RemoveSecurityGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveSecurityGroupBadRequest creates a RemoveSecurityGroupBadRequest with default headers values
func NewRemoveSecurityGroupBadRequest() *RemoveSecurityGroupBadRequest {
	return &RemoveSecurityGroupBadRequest{}
}

/*
RemoveSecurityGroupBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type RemoveSecurityGroupBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this remove security group bad request response has a 2xx status code
func (o *RemoveSecurityGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove security group bad request response has a 3xx status code
func (o *RemoveSecurityGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove security group bad request response has a 4xx status code
func (o *RemoveSecurityGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove security group bad request response has a 5xx status code
func (o *RemoveSecurityGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove security group bad request response a status code equal to that given
func (o *RemoveSecurityGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove security group bad request response
func (o *RemoveSecurityGroupBadRequest) Code() int {
	return 400
}

func (o *RemoveSecurityGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}][%d] removeSecurityGroupBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveSecurityGroupBadRequest) String() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}][%d] removeSecurityGroupBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveSecurityGroupBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *RemoveSecurityGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSecurityGroupNotFound creates a RemoveSecurityGroupNotFound with default headers values
func NewRemoveSecurityGroupNotFound() *RemoveSecurityGroupNotFound {
	return &RemoveSecurityGroupNotFound{}
}

/*
RemoveSecurityGroupNotFound describes a response with status code 404, with default header values.

resource not found
*/
type RemoveSecurityGroupNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this remove security group not found response has a 2xx status code
func (o *RemoveSecurityGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove security group not found response has a 3xx status code
func (o *RemoveSecurityGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove security group not found response has a 4xx status code
func (o *RemoveSecurityGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove security group not found response has a 5xx status code
func (o *RemoveSecurityGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove security group not found response a status code equal to that given
func (o *RemoveSecurityGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the remove security group not found response
func (o *RemoveSecurityGroupNotFound) Code() int {
	return 404
}

func (o *RemoveSecurityGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}][%d] removeSecurityGroupNotFound  %+v", 404, o.Payload)
}

func (o *RemoveSecurityGroupNotFound) String() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}][%d] removeSecurityGroupNotFound  %+v", 404, o.Payload)
}

func (o *RemoveSecurityGroupNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *RemoveSecurityGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSecurityGroupConflict creates a RemoveSecurityGroupConflict with default headers values
func NewRemoveSecurityGroupConflict() *RemoveSecurityGroupConflict {
	return &RemoveSecurityGroupConflict{}
}

/*
RemoveSecurityGroupConflict describes a response with status code 409, with default header values.

conflict
*/
type RemoveSecurityGroupConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this remove security group conflict response has a 2xx status code
func (o *RemoveSecurityGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove security group conflict response has a 3xx status code
func (o *RemoveSecurityGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove security group conflict response has a 4xx status code
func (o *RemoveSecurityGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove security group conflict response has a 5xx status code
func (o *RemoveSecurityGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this remove security group conflict response a status code equal to that given
func (o *RemoveSecurityGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the remove security group conflict response
func (o *RemoveSecurityGroupConflict) Code() int {
	return 409
}

func (o *RemoveSecurityGroupConflict) Error() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}][%d] removeSecurityGroupConflict  %+v", 409, o.Payload)
}

func (o *RemoveSecurityGroupConflict) String() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}][%d] removeSecurityGroupConflict  %+v", 409, o.Payload)
}

func (o *RemoveSecurityGroupConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *RemoveSecurityGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSecurityGroupInternalServerError creates a RemoveSecurityGroupInternalServerError with default headers values
func NewRemoveSecurityGroupInternalServerError() *RemoveSecurityGroupInternalServerError {
	return &RemoveSecurityGroupInternalServerError{}
}

/*
RemoveSecurityGroupInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type RemoveSecurityGroupInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this remove security group internal server error response has a 2xx status code
func (o *RemoveSecurityGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove security group internal server error response has a 3xx status code
func (o *RemoveSecurityGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove security group internal server error response has a 4xx status code
func (o *RemoveSecurityGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove security group internal server error response has a 5xx status code
func (o *RemoveSecurityGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove security group internal server error response a status code equal to that given
func (o *RemoveSecurityGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove security group internal server error response
func (o *RemoveSecurityGroupInternalServerError) Code() int {
	return 500
}

func (o *RemoveSecurityGroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}][%d] removeSecurityGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveSecurityGroupInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}][%d] removeSecurityGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveSecurityGroupInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *RemoveSecurityGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
