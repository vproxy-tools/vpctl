// Code generated by go-swagger; DO NOT EDIT.

package cert_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

// GetCertKeyReader is a Reader for the GetCertKey structure.
type GetCertKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCertKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCertKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCertKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCertKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetCertKeyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCertKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCertKeyOK creates a GetCertKeyOK with default headers values
func NewGetCertKeyOK() *GetCertKeyOK {
	return &GetCertKeyOK{}
}

/*
GetCertKeyOK describes a response with status code 200, with default header values.

ok
*/
type GetCertKeyOK struct {
	Payload *vproxy_client_model.CertKey
}

// IsSuccess returns true when this get cert key o k response has a 2xx status code
func (o *GetCertKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cert key o k response has a 3xx status code
func (o *GetCertKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cert key o k response has a 4xx status code
func (o *GetCertKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cert key o k response has a 5xx status code
func (o *GetCertKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cert key o k response a status code equal to that given
func (o *GetCertKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cert key o k response
func (o *GetCertKeyOK) Code() int {
	return 200
}

func (o *GetCertKeyOK) Error() string {
	return fmt.Sprintf("[GET /cert-key/{ck}][%d] getCertKeyOK  %+v", 200, o.Payload)
}

func (o *GetCertKeyOK) String() string {
	return fmt.Sprintf("[GET /cert-key/{ck}][%d] getCertKeyOK  %+v", 200, o.Payload)
}

func (o *GetCertKeyOK) GetPayload() *vproxy_client_model.CertKey {
	return o.Payload
}

func (o *GetCertKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.CertKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertKeyBadRequest creates a GetCertKeyBadRequest with default headers values
func NewGetCertKeyBadRequest() *GetCertKeyBadRequest {
	return &GetCertKeyBadRequest{}
}

/*
GetCertKeyBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type GetCertKeyBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this get cert key bad request response has a 2xx status code
func (o *GetCertKeyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cert key bad request response has a 3xx status code
func (o *GetCertKeyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cert key bad request response has a 4xx status code
func (o *GetCertKeyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cert key bad request response has a 5xx status code
func (o *GetCertKeyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cert key bad request response a status code equal to that given
func (o *GetCertKeyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get cert key bad request response
func (o *GetCertKeyBadRequest) Code() int {
	return 400
}

func (o *GetCertKeyBadRequest) Error() string {
	return fmt.Sprintf("[GET /cert-key/{ck}][%d] getCertKeyBadRequest  %+v", 400, o.Payload)
}

func (o *GetCertKeyBadRequest) String() string {
	return fmt.Sprintf("[GET /cert-key/{ck}][%d] getCertKeyBadRequest  %+v", 400, o.Payload)
}

func (o *GetCertKeyBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *GetCertKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertKeyNotFound creates a GetCertKeyNotFound with default headers values
func NewGetCertKeyNotFound() *GetCertKeyNotFound {
	return &GetCertKeyNotFound{}
}

/*
GetCertKeyNotFound describes a response with status code 404, with default header values.

resource not found
*/
type GetCertKeyNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this get cert key not found response has a 2xx status code
func (o *GetCertKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cert key not found response has a 3xx status code
func (o *GetCertKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cert key not found response has a 4xx status code
func (o *GetCertKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cert key not found response has a 5xx status code
func (o *GetCertKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get cert key not found response a status code equal to that given
func (o *GetCertKeyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get cert key not found response
func (o *GetCertKeyNotFound) Code() int {
	return 404
}

func (o *GetCertKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /cert-key/{ck}][%d] getCertKeyNotFound  %+v", 404, o.Payload)
}

func (o *GetCertKeyNotFound) String() string {
	return fmt.Sprintf("[GET /cert-key/{ck}][%d] getCertKeyNotFound  %+v", 404, o.Payload)
}

func (o *GetCertKeyNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *GetCertKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertKeyConflict creates a GetCertKeyConflict with default headers values
func NewGetCertKeyConflict() *GetCertKeyConflict {
	return &GetCertKeyConflict{}
}

/*
GetCertKeyConflict describes a response with status code 409, with default header values.

conflict
*/
type GetCertKeyConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this get cert key conflict response has a 2xx status code
func (o *GetCertKeyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cert key conflict response has a 3xx status code
func (o *GetCertKeyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cert key conflict response has a 4xx status code
func (o *GetCertKeyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cert key conflict response has a 5xx status code
func (o *GetCertKeyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get cert key conflict response a status code equal to that given
func (o *GetCertKeyConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get cert key conflict response
func (o *GetCertKeyConflict) Code() int {
	return 409
}

func (o *GetCertKeyConflict) Error() string {
	return fmt.Sprintf("[GET /cert-key/{ck}][%d] getCertKeyConflict  %+v", 409, o.Payload)
}

func (o *GetCertKeyConflict) String() string {
	return fmt.Sprintf("[GET /cert-key/{ck}][%d] getCertKeyConflict  %+v", 409, o.Payload)
}

func (o *GetCertKeyConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *GetCertKeyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertKeyInternalServerError creates a GetCertKeyInternalServerError with default headers values
func NewGetCertKeyInternalServerError() *GetCertKeyInternalServerError {
	return &GetCertKeyInternalServerError{}
}

/*
GetCertKeyInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type GetCertKeyInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this get cert key internal server error response has a 2xx status code
func (o *GetCertKeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cert key internal server error response has a 3xx status code
func (o *GetCertKeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cert key internal server error response has a 4xx status code
func (o *GetCertKeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cert key internal server error response has a 5xx status code
func (o *GetCertKeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cert key internal server error response a status code equal to that given
func (o *GetCertKeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cert key internal server error response
func (o *GetCertKeyInternalServerError) Code() int {
	return 500
}

func (o *GetCertKeyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cert-key/{ck}][%d] getCertKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCertKeyInternalServerError) String() string {
	return fmt.Sprintf("[GET /cert-key/{ck}][%d] getCertKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCertKeyInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *GetCertKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}