// Code generated by go-swagger; DO NOT EDIT.

package vproxy_client_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CertKey cert key
//
// swagger:model CertKey
type CertKey struct {

	// certs
	Certs []string `json:"certs"`

	// file path of the key
	Key string `json:"key,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this cert key
func (m *CertKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cert key based on context it is used
func (m *CertKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertKey) UnmarshalBinary(b []byte) error {
	var res CertKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}