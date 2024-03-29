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

// UpstreamDetail upstream detail
//
// swagger:model UpstreamDetail
type UpstreamDetail struct {

	// name
	Name string `json:"name,omitempty"`

	// server group list
	ServerGroupList []*ServerGroupInUpstreamDetail `json:"serverGroupList"`
}

// Validate validates this upstream detail
func (m *UpstreamDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerGroupList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpstreamDetail) validateServerGroupList(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerGroupList) { // not required
		return nil
	}

	for i := 0; i < len(m.ServerGroupList); i++ {
		if swag.IsZero(m.ServerGroupList[i]) { // not required
			continue
		}

		if m.ServerGroupList[i] != nil {
			if err := m.ServerGroupList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverGroupList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverGroupList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this upstream detail based on the context it is used
func (m *UpstreamDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServerGroupList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpstreamDetail) contextValidateServerGroupList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServerGroupList); i++ {

		if m.ServerGroupList[i] != nil {
			if err := m.ServerGroupList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverGroupList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverGroupList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpstreamDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpstreamDetail) UnmarshalBinary(b []byte) error {
	var res UpstreamDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
