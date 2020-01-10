// Code generated by go-swagger; DO NOT EDIT.

package vproxy_client_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TCPLbUpdate Tcp lb update
// swagger:model TcpLbUpdate
type TCPLbUpdate struct {

	// in buffer size
	InBufferSize int64 `json:"inBufferSize,omitempty"`

	// out buffer size
	OutBufferSize int64 `json:"outBufferSize,omitempty"`

	// security group reference for access control
	SecurityGroup string `json:"securityGroup,omitempty"`
}

// Validate validates this Tcp lb update
func (m *TCPLbUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TCPLbUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TCPLbUpdate) UnmarshalBinary(b []byte) error {
	var res TCPLbUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
