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

// GetUpstreamReader is a Reader for the GetUpstream structure.
type GetUpstreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUpstreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUpstreamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUpstreamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUpstreamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetUpstreamConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUpstreamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUpstreamOK creates a GetUpstreamOK with default headers values
func NewGetUpstreamOK() *GetUpstreamOK {
	return &GetUpstreamOK{}
}

/*
GetUpstreamOK describes a response with status code 200, with default header values.

ok
*/
type GetUpstreamOK struct {
	Payload *vproxy_client_model.Upstream
}

// IsSuccess returns true when this get upstream o k response has a 2xx status code
func (o *GetUpstreamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get upstream o k response has a 3xx status code
func (o *GetUpstreamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upstream o k response has a 4xx status code
func (o *GetUpstreamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upstream o k response has a 5xx status code
func (o *GetUpstreamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get upstream o k response a status code equal to that given
func (o *GetUpstreamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get upstream o k response
func (o *GetUpstreamOK) Code() int {
	return 200
}

func (o *GetUpstreamOK) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}][%d] getUpstreamOK  %+v", 200, o.Payload)
}

func (o *GetUpstreamOK) String() string {
	return fmt.Sprintf("[GET /upstream/{ups}][%d] getUpstreamOK  %+v", 200, o.Payload)
}

func (o *GetUpstreamOK) GetPayload() *vproxy_client_model.Upstream {
	return o.Payload
}

func (o *GetUpstreamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Upstream)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUpstreamBadRequest creates a GetUpstreamBadRequest with default headers values
func NewGetUpstreamBadRequest() *GetUpstreamBadRequest {
	return &GetUpstreamBadRequest{}
}

/*
GetUpstreamBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type GetUpstreamBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this get upstream bad request response has a 2xx status code
func (o *GetUpstreamBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get upstream bad request response has a 3xx status code
func (o *GetUpstreamBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upstream bad request response has a 4xx status code
func (o *GetUpstreamBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get upstream bad request response has a 5xx status code
func (o *GetUpstreamBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get upstream bad request response a status code equal to that given
func (o *GetUpstreamBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get upstream bad request response
func (o *GetUpstreamBadRequest) Code() int {
	return 400
}

func (o *GetUpstreamBadRequest) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}][%d] getUpstreamBadRequest  %+v", 400, o.Payload)
}

func (o *GetUpstreamBadRequest) String() string {
	return fmt.Sprintf("[GET /upstream/{ups}][%d] getUpstreamBadRequest  %+v", 400, o.Payload)
}

func (o *GetUpstreamBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *GetUpstreamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUpstreamNotFound creates a GetUpstreamNotFound with default headers values
func NewGetUpstreamNotFound() *GetUpstreamNotFound {
	return &GetUpstreamNotFound{}
}

/*
GetUpstreamNotFound describes a response with status code 404, with default header values.

resource not found
*/
type GetUpstreamNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this get upstream not found response has a 2xx status code
func (o *GetUpstreamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get upstream not found response has a 3xx status code
func (o *GetUpstreamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upstream not found response has a 4xx status code
func (o *GetUpstreamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get upstream not found response has a 5xx status code
func (o *GetUpstreamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get upstream not found response a status code equal to that given
func (o *GetUpstreamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get upstream not found response
func (o *GetUpstreamNotFound) Code() int {
	return 404
}

func (o *GetUpstreamNotFound) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}][%d] getUpstreamNotFound  %+v", 404, o.Payload)
}

func (o *GetUpstreamNotFound) String() string {
	return fmt.Sprintf("[GET /upstream/{ups}][%d] getUpstreamNotFound  %+v", 404, o.Payload)
}

func (o *GetUpstreamNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *GetUpstreamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUpstreamConflict creates a GetUpstreamConflict with default headers values
func NewGetUpstreamConflict() *GetUpstreamConflict {
	return &GetUpstreamConflict{}
}

/*
GetUpstreamConflict describes a response with status code 409, with default header values.

conflict
*/
type GetUpstreamConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this get upstream conflict response has a 2xx status code
func (o *GetUpstreamConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get upstream conflict response has a 3xx status code
func (o *GetUpstreamConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upstream conflict response has a 4xx status code
func (o *GetUpstreamConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get upstream conflict response has a 5xx status code
func (o *GetUpstreamConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get upstream conflict response a status code equal to that given
func (o *GetUpstreamConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get upstream conflict response
func (o *GetUpstreamConflict) Code() int {
	return 409
}

func (o *GetUpstreamConflict) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}][%d] getUpstreamConflict  %+v", 409, o.Payload)
}

func (o *GetUpstreamConflict) String() string {
	return fmt.Sprintf("[GET /upstream/{ups}][%d] getUpstreamConflict  %+v", 409, o.Payload)
}

func (o *GetUpstreamConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *GetUpstreamConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUpstreamInternalServerError creates a GetUpstreamInternalServerError with default headers values
func NewGetUpstreamInternalServerError() *GetUpstreamInternalServerError {
	return &GetUpstreamInternalServerError{}
}

/*
GetUpstreamInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type GetUpstreamInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this get upstream internal server error response has a 2xx status code
func (o *GetUpstreamInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get upstream internal server error response has a 3xx status code
func (o *GetUpstreamInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upstream internal server error response has a 4xx status code
func (o *GetUpstreamInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upstream internal server error response has a 5xx status code
func (o *GetUpstreamInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get upstream internal server error response a status code equal to that given
func (o *GetUpstreamInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get upstream internal server error response
func (o *GetUpstreamInternalServerError) Code() int {
	return 500
}

func (o *GetUpstreamInternalServerError) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}][%d] getUpstreamInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUpstreamInternalServerError) String() string {
	return fmt.Sprintf("[GET /upstream/{ups}][%d] getUpstreamInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUpstreamInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *GetUpstreamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
