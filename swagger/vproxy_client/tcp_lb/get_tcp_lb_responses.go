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

// GetTCPLbReader is a Reader for the GetTCPLb structure.
type GetTCPLbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTCPLbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTCPLbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTCPLbBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTCPLbNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetTCPLbConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTCPLbInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTCPLbOK creates a GetTCPLbOK with default headers values
func NewGetTCPLbOK() *GetTCPLbOK {
	return &GetTCPLbOK{}
}

/*
GetTCPLbOK describes a response with status code 200, with default header values.

ok
*/
type GetTCPLbOK struct {
	Payload *vproxy_client_model.TCPLb
}

// IsSuccess returns true when this get Tcp lb o k response has a 2xx status code
func (o *GetTCPLbOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Tcp lb o k response has a 3xx status code
func (o *GetTCPLbOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Tcp lb o k response has a 4xx status code
func (o *GetTCPLbOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Tcp lb o k response has a 5xx status code
func (o *GetTCPLbOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Tcp lb o k response a status code equal to that given
func (o *GetTCPLbOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Tcp lb o k response
func (o *GetTCPLbOK) Code() int {
	return 200
}

func (o *GetTCPLbOK) Error() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}][%d] getTcpLbOK  %+v", 200, o.Payload)
}

func (o *GetTCPLbOK) String() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}][%d] getTcpLbOK  %+v", 200, o.Payload)
}

func (o *GetTCPLbOK) GetPayload() *vproxy_client_model.TCPLb {
	return o.Payload
}

func (o *GetTCPLbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.TCPLb)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTCPLbBadRequest creates a GetTCPLbBadRequest with default headers values
func NewGetTCPLbBadRequest() *GetTCPLbBadRequest {
	return &GetTCPLbBadRequest{}
}

/*
GetTCPLbBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type GetTCPLbBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this get Tcp lb bad request response has a 2xx status code
func (o *GetTCPLbBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Tcp lb bad request response has a 3xx status code
func (o *GetTCPLbBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Tcp lb bad request response has a 4xx status code
func (o *GetTCPLbBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Tcp lb bad request response has a 5xx status code
func (o *GetTCPLbBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Tcp lb bad request response a status code equal to that given
func (o *GetTCPLbBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get Tcp lb bad request response
func (o *GetTCPLbBadRequest) Code() int {
	return 400
}

func (o *GetTCPLbBadRequest) Error() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}][%d] getTcpLbBadRequest  %+v", 400, o.Payload)
}

func (o *GetTCPLbBadRequest) String() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}][%d] getTcpLbBadRequest  %+v", 400, o.Payload)
}

func (o *GetTCPLbBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *GetTCPLbBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTCPLbNotFound creates a GetTCPLbNotFound with default headers values
func NewGetTCPLbNotFound() *GetTCPLbNotFound {
	return &GetTCPLbNotFound{}
}

/*
GetTCPLbNotFound describes a response with status code 404, with default header values.

resource not found
*/
type GetTCPLbNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this get Tcp lb not found response has a 2xx status code
func (o *GetTCPLbNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Tcp lb not found response has a 3xx status code
func (o *GetTCPLbNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Tcp lb not found response has a 4xx status code
func (o *GetTCPLbNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Tcp lb not found response has a 5xx status code
func (o *GetTCPLbNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Tcp lb not found response a status code equal to that given
func (o *GetTCPLbNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Tcp lb not found response
func (o *GetTCPLbNotFound) Code() int {
	return 404
}

func (o *GetTCPLbNotFound) Error() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}][%d] getTcpLbNotFound  %+v", 404, o.Payload)
}

func (o *GetTCPLbNotFound) String() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}][%d] getTcpLbNotFound  %+v", 404, o.Payload)
}

func (o *GetTCPLbNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *GetTCPLbNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTCPLbConflict creates a GetTCPLbConflict with default headers values
func NewGetTCPLbConflict() *GetTCPLbConflict {
	return &GetTCPLbConflict{}
}

/*
GetTCPLbConflict describes a response with status code 409, with default header values.

conflict
*/
type GetTCPLbConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this get Tcp lb conflict response has a 2xx status code
func (o *GetTCPLbConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Tcp lb conflict response has a 3xx status code
func (o *GetTCPLbConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Tcp lb conflict response has a 4xx status code
func (o *GetTCPLbConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Tcp lb conflict response has a 5xx status code
func (o *GetTCPLbConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get Tcp lb conflict response a status code equal to that given
func (o *GetTCPLbConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get Tcp lb conflict response
func (o *GetTCPLbConflict) Code() int {
	return 409
}

func (o *GetTCPLbConflict) Error() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}][%d] getTcpLbConflict  %+v", 409, o.Payload)
}

func (o *GetTCPLbConflict) String() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}][%d] getTcpLbConflict  %+v", 409, o.Payload)
}

func (o *GetTCPLbConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *GetTCPLbConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTCPLbInternalServerError creates a GetTCPLbInternalServerError with default headers values
func NewGetTCPLbInternalServerError() *GetTCPLbInternalServerError {
	return &GetTCPLbInternalServerError{}
}

/*
GetTCPLbInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type GetTCPLbInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this get Tcp lb internal server error response has a 2xx status code
func (o *GetTCPLbInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Tcp lb internal server error response has a 3xx status code
func (o *GetTCPLbInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Tcp lb internal server error response has a 4xx status code
func (o *GetTCPLbInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Tcp lb internal server error response has a 5xx status code
func (o *GetTCPLbInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Tcp lb internal server error response a status code equal to that given
func (o *GetTCPLbInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get Tcp lb internal server error response
func (o *GetTCPLbInternalServerError) Code() int {
	return 500
}

func (o *GetTCPLbInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}][%d] getTcpLbInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTCPLbInternalServerError) String() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}][%d] getTcpLbInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTCPLbInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *GetTCPLbInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}