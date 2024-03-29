// Code generated by go-swagger; DO NOT EDIT.

package vproxy_client_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SecurityGroupProtocol security group protocol
//
// swagger:model SecurityGroupProtocol
type SecurityGroupProtocol string

func NewSecurityGroupProtocol(value SecurityGroupProtocol) *SecurityGroupProtocol {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SecurityGroupProtocol.
func (m SecurityGroupProtocol) Pointer() *SecurityGroupProtocol {
	return &m
}

const (

	// SecurityGroupProtocolTCP captures enum value "TCP"
	SecurityGroupProtocolTCP SecurityGroupProtocol = "TCP"

	// SecurityGroupProtocolUDP captures enum value "UDP"
	SecurityGroupProtocolUDP SecurityGroupProtocol = "UDP"
)

// for schema
var securityGroupProtocolEnum []interface{}

func init() {
	var res []SecurityGroupProtocol
	if err := json.Unmarshal([]byte(`["TCP","UDP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityGroupProtocolEnum = append(securityGroupProtocolEnum, v)
	}
}

func (m SecurityGroupProtocol) validateSecurityGroupProtocolEnum(path, location string, value SecurityGroupProtocol) error {
	if err := validate.EnumCase(path, location, value, securityGroupProtocolEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this security group protocol
func (m SecurityGroupProtocol) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSecurityGroupProtocolEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this security group protocol based on context it is used
func (m SecurityGroupProtocol) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
