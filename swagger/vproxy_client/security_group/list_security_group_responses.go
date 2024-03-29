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

// ListSecurityGroupReader is a Reader for the ListSecurityGroup structure.
type ListSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListSecurityGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListSecurityGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListSecurityGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSecurityGroupOK creates a ListSecurityGroupOK with default headers values
func NewListSecurityGroupOK() *ListSecurityGroupOK {
	return &ListSecurityGroupOK{}
}

/*
ListSecurityGroupOK describes a response with status code 200, with default header values.

ok
*/
type ListSecurityGroupOK struct {
	Payload []*vproxy_client_model.SecurityGroup
}

// IsSuccess returns true when this list security group o k response has a 2xx status code
func (o *ListSecurityGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list security group o k response has a 3xx status code
func (o *ListSecurityGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list security group o k response has a 4xx status code
func (o *ListSecurityGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list security group o k response has a 5xx status code
func (o *ListSecurityGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list security group o k response a status code equal to that given
func (o *ListSecurityGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list security group o k response
func (o *ListSecurityGroupOK) Code() int {
	return 200
}

func (o *ListSecurityGroupOK) Error() string {
	return fmt.Sprintf("[GET /security-group][%d] listSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *ListSecurityGroupOK) String() string {
	return fmt.Sprintf("[GET /security-group][%d] listSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *ListSecurityGroupOK) GetPayload() []*vproxy_client_model.SecurityGroup {
	return o.Payload
}

func (o *ListSecurityGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSecurityGroupBadRequest creates a ListSecurityGroupBadRequest with default headers values
func NewListSecurityGroupBadRequest() *ListSecurityGroupBadRequest {
	return &ListSecurityGroupBadRequest{}
}

/*
ListSecurityGroupBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type ListSecurityGroupBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this list security group bad request response has a 2xx status code
func (o *ListSecurityGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list security group bad request response has a 3xx status code
func (o *ListSecurityGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list security group bad request response has a 4xx status code
func (o *ListSecurityGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list security group bad request response has a 5xx status code
func (o *ListSecurityGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list security group bad request response a status code equal to that given
func (o *ListSecurityGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list security group bad request response
func (o *ListSecurityGroupBadRequest) Code() int {
	return 400
}

func (o *ListSecurityGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /security-group][%d] listSecurityGroupBadRequest  %+v", 400, o.Payload)
}

func (o *ListSecurityGroupBadRequest) String() string {
	return fmt.Sprintf("[GET /security-group][%d] listSecurityGroupBadRequest  %+v", 400, o.Payload)
}

func (o *ListSecurityGroupBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *ListSecurityGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSecurityGroupConflict creates a ListSecurityGroupConflict with default headers values
func NewListSecurityGroupConflict() *ListSecurityGroupConflict {
	return &ListSecurityGroupConflict{}
}

/*
ListSecurityGroupConflict describes a response with status code 409, with default header values.

conflict
*/
type ListSecurityGroupConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this list security group conflict response has a 2xx status code
func (o *ListSecurityGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list security group conflict response has a 3xx status code
func (o *ListSecurityGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list security group conflict response has a 4xx status code
func (o *ListSecurityGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this list security group conflict response has a 5xx status code
func (o *ListSecurityGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this list security group conflict response a status code equal to that given
func (o *ListSecurityGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the list security group conflict response
func (o *ListSecurityGroupConflict) Code() int {
	return 409
}

func (o *ListSecurityGroupConflict) Error() string {
	return fmt.Sprintf("[GET /security-group][%d] listSecurityGroupConflict  %+v", 409, o.Payload)
}

func (o *ListSecurityGroupConflict) String() string {
	return fmt.Sprintf("[GET /security-group][%d] listSecurityGroupConflict  %+v", 409, o.Payload)
}

func (o *ListSecurityGroupConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *ListSecurityGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSecurityGroupInternalServerError creates a ListSecurityGroupInternalServerError with default headers values
func NewListSecurityGroupInternalServerError() *ListSecurityGroupInternalServerError {
	return &ListSecurityGroupInternalServerError{}
}

/*
ListSecurityGroupInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type ListSecurityGroupInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this list security group internal server error response has a 2xx status code
func (o *ListSecurityGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list security group internal server error response has a 3xx status code
func (o *ListSecurityGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list security group internal server error response has a 4xx status code
func (o *ListSecurityGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list security group internal server error response has a 5xx status code
func (o *ListSecurityGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list security group internal server error response a status code equal to that given
func (o *ListSecurityGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list security group internal server error response
func (o *ListSecurityGroupInternalServerError) Code() int {
	return 500
}

func (o *ListSecurityGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /security-group][%d] listSecurityGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSecurityGroupInternalServerError) String() string {
	return fmt.Sprintf("[GET /security-group][%d] listSecurityGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSecurityGroupInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *ListSecurityGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
