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

// ServerGroupDetail server group detail
//
// swagger:model ServerGroupDetail
type ServerGroupDetail struct {

	// annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// consider the server unhealthy after $down times of failed health checks
	Down int64 `json:"down,omitempty"`

	// event loop group
	EventLoopGroup *EventLoopGroupDetail `json:"eventLoopGroup,omitempty"`

	// method
	Method *ServerGroupMethod `json:"method,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// health check period (ms) (interval between two hc)
	Period int64 `json:"period,omitempty"`

	// protocol
	Protocol *CheckProtocol `json:"protocol,omitempty"`

	// server list
	ServerList []*ServerDetail `json:"serverList"`

	// health check timeout (ms) (timeout before getting expected response)
	Timeout int64 `json:"timeout,omitempty"`

	// consider the server healthy after $up times of successful health checks
	Up int64 `json:"up,omitempty"`
}

// Validate validates this server group detail
func (m *ServerGroupDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventLoopGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerGroupDetail) validateEventLoopGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.EventLoopGroup) { // not required
		return nil
	}

	if m.EventLoopGroup != nil {
		if err := m.EventLoopGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventLoopGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eventLoopGroup")
			}
			return err
		}
	}

	return nil
}

func (m *ServerGroupDetail) validateMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.Method) { // not required
		return nil
	}

	if m.Method != nil {
		if err := m.Method.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

func (m *ServerGroupDetail) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	if m.Protocol != nil {
		if err := m.Protocol.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

func (m *ServerGroupDetail) validateServerList(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerList) { // not required
		return nil
	}

	for i := 0; i < len(m.ServerList); i++ {
		if swag.IsZero(m.ServerList[i]) { // not required
			continue
		}

		if m.ServerList[i] != nil {
			if err := m.ServerList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this server group detail based on the context it is used
func (m *ServerGroupDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventLoopGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServerList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerGroupDetail) contextValidateEventLoopGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.EventLoopGroup != nil {
		if err := m.EventLoopGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventLoopGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eventLoopGroup")
			}
			return err
		}
	}

	return nil
}

func (m *ServerGroupDetail) contextValidateMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.Method != nil {
		if err := m.Method.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

func (m *ServerGroupDetail) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if m.Protocol != nil {
		if err := m.Protocol.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

func (m *ServerGroupDetail) contextValidateServerList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServerList); i++ {

		if m.ServerList[i] != nil {
			if err := m.ServerList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerGroupDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerGroupDetail) UnmarshalBinary(b []byte) error {
	var res ServerGroupDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
