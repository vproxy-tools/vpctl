// Code generated by go-swagger; DO NOT EDIT.

package tcp_lb

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
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
	case 404:
		result := NewGetTCPLbNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTCPLbOK creates a GetTCPLbOK with default headers values
func NewGetTCPLbOK() *GetTCPLbOK {
	return &GetTCPLbOK{}
}

/*GetTCPLbOK handles this case with default header values.

ok
*/
type GetTCPLbOK struct {
	Payload *vproxy_client_model.TCPLb
}

func (o *GetTCPLbOK) Error() string {
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

// NewGetTCPLbNotFound creates a GetTCPLbNotFound with default headers values
func NewGetTCPLbNotFound() *GetTCPLbNotFound {
	return &GetTCPLbNotFound{}
}

/*GetTCPLbNotFound handles this case with default header values.

not found
*/
type GetTCPLbNotFound struct {
}

func (o *GetTCPLbNotFound) Error() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}][%d] getTcpLbNotFound ", 404)
}

func (o *GetTCPLbNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
