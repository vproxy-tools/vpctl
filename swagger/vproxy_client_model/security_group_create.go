// Code generated by go-swagger; DO NOT EDIT.

package vproxy_client_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SecurityGroupCreate security group create
//
// swagger:model SecurityGroupCreate
type SecurityGroupCreate struct {

	// default rule
	// Required: true
	DefaultRule *Rule `json:"defaultRule"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this security group create
func (m *SecurityGroupCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupCreate) validateDefaultRule(formats strfmt.Registry) error {

	if err := validate.Required("defaultRule", "body", m.DefaultRule); err != nil {
		return err
	}

	if err := validate.Required("defaultRule", "body", m.DefaultRule); err != nil {
		return err
	}

	if m.DefaultRule != nil {
		if err := m.DefaultRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultRule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultRule")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityGroupCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this security group create based on the context it is used
func (m *SecurityGroupCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupCreate) contextValidateDefaultRule(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultRule != nil {
		if err := m.DefaultRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultRule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultRule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityGroupCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityGroupCreate) UnmarshalBinary(b []byte) error {
	var res SecurityGroupCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
