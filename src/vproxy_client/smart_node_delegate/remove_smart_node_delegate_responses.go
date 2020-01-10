// Code generated by go-swagger; DO NOT EDIT.

package smart_node_delegate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveSmartNodeDelegateReader is a Reader for the RemoveSmartNodeDelegate structure.
type RemoveSmartNodeDelegateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveSmartNodeDelegateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveSmartNodeDelegateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveSmartNodeDelegateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveSmartNodeDelegateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveSmartNodeDelegateNoContent creates a RemoveSmartNodeDelegateNoContent with default headers values
func NewRemoveSmartNodeDelegateNoContent() *RemoveSmartNodeDelegateNoContent {
	return &RemoveSmartNodeDelegateNoContent{}
}

/*RemoveSmartNodeDelegateNoContent handles this case with default header values.

ok
*/
type RemoveSmartNodeDelegateNoContent struct {
}

func (o *RemoveSmartNodeDelegateNoContent) Error() string {
	return fmt.Sprintf("[DELETE /smart-node-delegate/{snd}][%d] removeSmartNodeDelegateNoContent ", 204)
}

func (o *RemoveSmartNodeDelegateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveSmartNodeDelegateBadRequest creates a RemoveSmartNodeDelegateBadRequest with default headers values
func NewRemoveSmartNodeDelegateBadRequest() *RemoveSmartNodeDelegateBadRequest {
	return &RemoveSmartNodeDelegateBadRequest{}
}

/*RemoveSmartNodeDelegateBadRequest handles this case with default header values.

Invalid input
*/
type RemoveSmartNodeDelegateBadRequest struct {
}

func (o *RemoveSmartNodeDelegateBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /smart-node-delegate/{snd}][%d] removeSmartNodeDelegateBadRequest ", 400)
}

func (o *RemoveSmartNodeDelegateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveSmartNodeDelegateNotFound creates a RemoveSmartNodeDelegateNotFound with default headers values
func NewRemoveSmartNodeDelegateNotFound() *RemoveSmartNodeDelegateNotFound {
	return &RemoveSmartNodeDelegateNotFound{}
}

/*RemoveSmartNodeDelegateNotFound handles this case with default header values.

SmartNodeDelegate not found
*/
type RemoveSmartNodeDelegateNotFound struct {
}

func (o *RemoveSmartNodeDelegateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /smart-node-delegate/{snd}][%d] removeSmartNodeDelegateNotFound ", 404)
}

func (o *RemoveSmartNodeDelegateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
