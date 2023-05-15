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

// IPType binds ipv4 or ipv6
//
// swagger:model IPType
type IPType string

func NewIPType(value IPType) *IPType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IPType.
func (m IPType) Pointer() *IPType {
	return &m
}

const (

	// IPTypeV4 captures enum value "v4"
	IPTypeV4 IPType = "v4"

	// IPTypeV6 captures enum value "v6"
	IPTypeV6 IPType = "v6"
)

// for schema
var ipTypeEnum []interface{}

func init() {
	var res []IPType
	if err := json.Unmarshal([]byte(`["v4","v6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipTypeEnum = append(ipTypeEnum, v)
	}
}

func (m IPType) validateIPTypeEnum(path, location string, value IPType) error {
	if err := validate.EnumCase(path, location, value, ipTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this IP type
func (m IPType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIPTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this IP type based on context it is used
func (m IPType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
