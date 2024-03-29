// Code generated by go-swagger; DO NOT EDIT.

package vproxy_client_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EventLoop event loop
//
// swagger:model EventLoop
type EventLoop struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this event loop
func (m *EventLoop) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this event loop based on context it is used
func (m *EventLoop) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventLoop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventLoop) UnmarshalBinary(b []byte) error {
	var res EventLoop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
