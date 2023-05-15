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

// DescribeCertKeyReader is a Reader for the DescribeCertKey structure.
type DescribeCertKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeCertKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeCertKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDescribeCertKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDescribeCertKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDescribeCertKeyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDescribeCertKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDescribeCertKeyOK creates a DescribeCertKeyOK with default headers values
func NewDescribeCertKeyOK() *DescribeCertKeyOK {
	return &DescribeCertKeyOK{}
}

/*
DescribeCertKeyOK describes a response with status code 200, with default header values.

ok
*/
type DescribeCertKeyOK struct {
	Payload *vproxy_client_model.CertKeyDetail
}

// IsSuccess returns true when this describe cert key o k response has a 2xx status code
func (o *DescribeCertKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe cert key o k response has a 3xx status code
func (o *DescribeCertKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe cert key o k response has a 4xx status code
func (o *DescribeCertKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe cert key o k response has a 5xx status code
func (o *DescribeCertKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe cert key o k response a status code equal to that given
func (o *DescribeCertKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe cert key o k response
func (o *DescribeCertKeyOK) Code() int {
	return 200
}

func (o *DescribeCertKeyOK) Error() string {
	return fmt.Sprintf("[GET /cert-key/{ck}/detail][%d] describeCertKeyOK  %+v", 200, o.Payload)
}

func (o *DescribeCertKeyOK) String() string {
	return fmt.Sprintf("[GET /cert-key/{ck}/detail][%d] describeCertKeyOK  %+v", 200, o.Payload)
}

func (o *DescribeCertKeyOK) GetPayload() *vproxy_client_model.CertKeyDetail {
	return o.Payload
}

func (o *DescribeCertKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.CertKeyDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeCertKeyBadRequest creates a DescribeCertKeyBadRequest with default headers values
func NewDescribeCertKeyBadRequest() *DescribeCertKeyBadRequest {
	return &DescribeCertKeyBadRequest{}
}

/*
DescribeCertKeyBadRequest describes a response with status code 400, with default header values.

invalid input parameters
*/
type DescribeCertKeyBadRequest struct {
	Payload *vproxy_client_model.Error400
}

// IsSuccess returns true when this describe cert key bad request response has a 2xx status code
func (o *DescribeCertKeyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe cert key bad request response has a 3xx status code
func (o *DescribeCertKeyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe cert key bad request response has a 4xx status code
func (o *DescribeCertKeyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe cert key bad request response has a 5xx status code
func (o *DescribeCertKeyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this describe cert key bad request response a status code equal to that given
func (o *DescribeCertKeyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the describe cert key bad request response
func (o *DescribeCertKeyBadRequest) Code() int {
	return 400
}

func (o *DescribeCertKeyBadRequest) Error() string {
	return fmt.Sprintf("[GET /cert-key/{ck}/detail][%d] describeCertKeyBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeCertKeyBadRequest) String() string {
	return fmt.Sprintf("[GET /cert-key/{ck}/detail][%d] describeCertKeyBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeCertKeyBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *DescribeCertKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeCertKeyNotFound creates a DescribeCertKeyNotFound with default headers values
func NewDescribeCertKeyNotFound() *DescribeCertKeyNotFound {
	return &DescribeCertKeyNotFound{}
}

/*
DescribeCertKeyNotFound describes a response with status code 404, with default header values.

resource not found
*/
type DescribeCertKeyNotFound struct {
	Payload *vproxy_client_model.Error404
}

// IsSuccess returns true when this describe cert key not found response has a 2xx status code
func (o *DescribeCertKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe cert key not found response has a 3xx status code
func (o *DescribeCertKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe cert key not found response has a 4xx status code
func (o *DescribeCertKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe cert key not found response has a 5xx status code
func (o *DescribeCertKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this describe cert key not found response a status code equal to that given
func (o *DescribeCertKeyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the describe cert key not found response
func (o *DescribeCertKeyNotFound) Code() int {
	return 404
}

func (o *DescribeCertKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /cert-key/{ck}/detail][%d] describeCertKeyNotFound  %+v", 404, o.Payload)
}

func (o *DescribeCertKeyNotFound) String() string {
	return fmt.Sprintf("[GET /cert-key/{ck}/detail][%d] describeCertKeyNotFound  %+v", 404, o.Payload)
}

func (o *DescribeCertKeyNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *DescribeCertKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeCertKeyConflict creates a DescribeCertKeyConflict with default headers values
func NewDescribeCertKeyConflict() *DescribeCertKeyConflict {
	return &DescribeCertKeyConflict{}
}

/*
DescribeCertKeyConflict describes a response with status code 409, with default header values.

conflict
*/
type DescribeCertKeyConflict struct {
	Payload *vproxy_client_model.Error409
}

// IsSuccess returns true when this describe cert key conflict response has a 2xx status code
func (o *DescribeCertKeyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe cert key conflict response has a 3xx status code
func (o *DescribeCertKeyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe cert key conflict response has a 4xx status code
func (o *DescribeCertKeyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe cert key conflict response has a 5xx status code
func (o *DescribeCertKeyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this describe cert key conflict response a status code equal to that given
func (o *DescribeCertKeyConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the describe cert key conflict response
func (o *DescribeCertKeyConflict) Code() int {
	return 409
}

func (o *DescribeCertKeyConflict) Error() string {
	return fmt.Sprintf("[GET /cert-key/{ck}/detail][%d] describeCertKeyConflict  %+v", 409, o.Payload)
}

func (o *DescribeCertKeyConflict) String() string {
	return fmt.Sprintf("[GET /cert-key/{ck}/detail][%d] describeCertKeyConflict  %+v", 409, o.Payload)
}

func (o *DescribeCertKeyConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *DescribeCertKeyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeCertKeyInternalServerError creates a DescribeCertKeyInternalServerError with default headers values
func NewDescribeCertKeyInternalServerError() *DescribeCertKeyInternalServerError {
	return &DescribeCertKeyInternalServerError{}
}

/*
DescribeCertKeyInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type DescribeCertKeyInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

// IsSuccess returns true when this describe cert key internal server error response has a 2xx status code
func (o *DescribeCertKeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe cert key internal server error response has a 3xx status code
func (o *DescribeCertKeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe cert key internal server error response has a 4xx status code
func (o *DescribeCertKeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe cert key internal server error response has a 5xx status code
func (o *DescribeCertKeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this describe cert key internal server error response a status code equal to that given
func (o *DescribeCertKeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the describe cert key internal server error response
func (o *DescribeCertKeyInternalServerError) Code() int {
	return 500
}

func (o *DescribeCertKeyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cert-key/{ck}/detail][%d] describeCertKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *DescribeCertKeyInternalServerError) String() string {
	return fmt.Sprintf("[GET /cert-key/{ck}/detail][%d] describeCertKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *DescribeCertKeyInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *DescribeCertKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
