// Code generated by go-swagger; DO NOT EDIT.

package vproxy_client_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecurityGroupDetail security group detail
//
// swagger:model SecurityGroupDetail
type SecurityGroupDetail struct {

	// default rule
	DefaultRule Rule `json:"defaultRule,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// rule list
	RuleList []*SecurityGroupRuleDetail `json:"ruleList"`
}

// Validate validates this security group detail
func (m *SecurityGroupDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupDetail) validateDefaultRule(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultRule) { // not required
		return nil
	}

	if err := m.DefaultRule.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("defaultRule")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("defaultRule")
		}
		return err
	}

	return nil
}

func (m *SecurityGroupDetail) validateRuleList(formats strfmt.Registry) error {
	if swag.IsZero(m.RuleList) { // not required
		return nil
	}

	for i := 0; i < len(m.RuleList); i++ {
		if swag.IsZero(m.RuleList[i]) { // not required
			continue
		}

		if m.RuleList[i] != nil {
			if err := m.RuleList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ruleList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ruleList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this security group detail based on the context it is used
func (m *SecurityGroupDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuleList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupDetail) contextValidateDefaultRule(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DefaultRule.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("defaultRule")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("defaultRule")
		}
		return err
	}

	return nil
}

func (m *SecurityGroupDetail) contextValidateRuleList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RuleList); i++ {

		if m.RuleList[i] != nil {
			if err := m.RuleList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ruleList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ruleList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityGroupDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityGroupDetail) UnmarshalBinary(b []byte) error {
	var res SecurityGroupDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
