// Code generated by go-swagger; DO NOT EDIT.

package security_group_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
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
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveSecurityGroupRuleNoContent creates a RemoveSecurityGroupRuleNoContent with default headers values
func NewRemoveSecurityGroupRuleNoContent() *RemoveSecurityGroupRuleNoContent {
	return &RemoveSecurityGroupRuleNoContent{}
}

/*RemoveSecurityGroupRuleNoContent handles this case with default header values.

ok
*/
type RemoveSecurityGroupRuleNoContent struct {
}

func (o *RemoveSecurityGroupRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /security-group/{secg}/security-group-rule/{secgr}][%d] removeSecurityGroupRuleNoContent ", 204)
}

func (o *RemoveSecurityGroupRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveSecurityGroupRuleBadRequest creates a RemoveSecurityGroupRuleBadRequest with default headers values
func NewRemoveSecurityGroupRuleBadRequest() *RemoveSecurityGroupRuleBadRequest {
	return &RemoveSecurityGroupRuleBadRequest{}
}

/*RemoveSecurityGroupRuleBadRequest handles this case with default header values.

invalid input parameters
*/
type RemoveSecurityGroupRuleBadRequest struct {
	Payload *vproxy_client_model.Error400
}

func (o *RemoveSecurityGroupRuleBadRequest) Error() string {
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

/*RemoveSecurityGroupRuleNotFound handles this case with default header values.

resource not found
*/
type RemoveSecurityGroupRuleNotFound struct {
	Payload *vproxy_client_model.Error404
}

func (o *RemoveSecurityGroupRuleNotFound) Error() string {
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

/*RemoveSecurityGroupRuleConflict handles this case with default header values.

conflict
*/
type RemoveSecurityGroupRuleConflict struct {
	Payload *vproxy_client_model.Error409
}

func (o *RemoveSecurityGroupRuleConflict) Error() string {
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

/*RemoveSecurityGroupRuleInternalServerError handles this case with default header values.

internal error
*/
type RemoveSecurityGroupRuleInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

func (o *RemoveSecurityGroupRuleInternalServerError) Error() string {
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
