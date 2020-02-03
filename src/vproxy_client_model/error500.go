// Code generated by go-swagger; DO NOT EDIT.

package vproxy_client_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Error500 error500
// swagger:model Error500
type Error500 struct {

	// code
	Code int64 `json:"code,omitempty"`

	// err Id
	ErrID string `json:"errId,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this error500
func (m *Error500) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Error500) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error500) UnmarshalBinary(b []byte) error {
	var res Error500
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}