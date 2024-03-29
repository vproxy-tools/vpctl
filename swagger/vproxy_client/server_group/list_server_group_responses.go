// Code generated by go-swagger; DO NOT EDIT.

package server_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// ListServerGroupReader is a Reader for the ListServerGroup structure.
type ListServerGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServerGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServerGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListServerGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListServerGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListServerGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServerGroupOK creates a ListServerGroupOK with default headers values
func NewListServerGroupOK() *ListServerGroupOK {
	return &ListServerGroupOK{}
}

/*
ListServerGroupOK describes a response with status code 200, with default header values.

ok
*/
type ListServerGroupOK struct {
	Payload []*vproxy_client_model.ServerGroup
}

// IsSuccess returns true when this list server group o k response has a 2xx status code
func (o *ListServerGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list server group o k response has a 3xx status code
func (o *ListServerGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list server group o k response has a 4xx status code
func (o *ListServerGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list server group o k response has a 5xx status code
func (o *ListServerGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list server group o k response a status code equal to that given
func (o *ListServerGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list server group o k response
func (o *ListServerGroupOK) Code() int {
	return 200
}

func (o *ListServerGroupOK) Error() string {
	return fmt.Sprintf("[GET /server-group][%d] listServerGroupOK  %+v", 200, o.Payload)
}

func (o *ListServerGroupOK) String() string {
	return fmt.Sprintf("[GET /server-group][%d] listServerGroupOK  %+v", 200, o.Payload)
}

func (o *ListServerGroupOK) GetPayload() []*vproxy_client_model.ServerGroup {
	return o.Payload
}

func (o *ListServerGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServerGroupBadRequest creates a ListServerGroupBadRequest with default headers values
func NewListServerGroupBadRequest() *ListServerGroupBadRequest {
	return &ListServerGroupBadRequest{}
}

/*
ListServerGroupBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type ListServerGroupBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this list server group bad request response has a 2xx status code
func (o *ListServerGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list server group bad request response has a 3xx status code
func (o *ListServerGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list server group bad request response has a 4xx status code
func (o *ListServerGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list server group bad request response has a 5xx status code
func (o *ListServerGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list server group bad request response a status code equal to that given
func (o *ListServerGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list server group bad request response
func (o *ListServerGroupBadRequest) Code() int {
	return 400
}

func (o *ListServerGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /server-group][%d] listServerGroupBadRequest  %+v", 400, o.Payload)
}

func (o *ListServerGroupBadRequest) String() string {
	return fmt.Sprintf("[GET /server-group][%d] listServerGroupBadRequest  %+v", 400, o.Payload)
}

func (o *ListServerGroupBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *ListServerGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServerGroupConflict creates a ListServerGroupConflict with default headers values
func NewListServerGroupConflict() *ListServerGroupConflict {
	return &ListServerGroupConflict{}
}

/*
ListServerGroupConflict describes a response with status code 409, with default header values.

conflict
*/
type ListServerGroupConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this list server group conflict response has a 2xx status code
func (o *ListServerGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list server group conflict response has a 3xx status code
func (o *ListServerGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list server group conflict response has a 4xx status code
func (o *ListServerGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this list server group conflict response has a 5xx status code
func (o *ListServerGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this list server group conflict response a status code equal to that given
func (o *ListServerGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the list server group conflict response
func (o *ListServerGroupConflict) Code() int {
	return 409
}

func (o *ListServerGroupConflict) Error() string {
	return fmt.Sprintf("[GET /server-group][%d] listServerGroupConflict  %+v", 409, o.Payload)
}

func (o *ListServerGroupConflict) String() string {
	return fmt.Sprintf("[GET /server-group][%d] listServerGroupConflict  %+v", 409, o.Payload)
}

func (o *ListServerGroupConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *ListServerGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServerGroupInternalServerError creates a ListServerGroupInternalServerError with default headers values
func NewListServerGroupInternalServerError() *ListServerGroupInternalServerError {
	return &ListServerGroupInternalServerError{}
}

/*
ListServerGroupInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type ListServerGroupInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this list server group internal server error response has a 2xx status code
func (o *ListServerGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list server group internal server error response has a 3xx status code
func (o *ListServerGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list server group internal server error response has a 4xx status code
func (o *ListServerGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list server group internal server error response has a 5xx status code
func (o *ListServerGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list server group internal server error response a status code equal to that given
func (o *ListServerGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list server group internal server error response
func (o *ListServerGroupInternalServerError) Code() int {
	return 500
}

func (o *ListServerGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /server-group][%d] listServerGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ListServerGroupInternalServerError) String() string {
	return fmt.Sprintf("[GET /server-group][%d] listServerGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ListServerGroupInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *ListServerGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
