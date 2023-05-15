// Code generated by go-swagger; DO NOT EDIT.

package vproxy_client_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Socks5ServerDetail socks5 server detail
//
// swagger:model Socks5ServerDetail
type Socks5ServerDetail struct {

	// acceptor loop group
	AcceptorLoopGroup *EventLoopGroupDetail `json:"acceptorLoopGroup,omitempty"`

	// binding l4addr
	Address string `json:"address,omitempty"`

	// allow or disallow to proxy to non-backend endpoints
	AllowNonBackend bool `json:"allowNonBackend,omitempty"`

	// backend
	Backend *UpstreamDetail `json:"backend,omitempty"`

	// input buffer size
	InBufferSize int64 `json:"inBufferSize,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// output buffer size
	OutBufferSize int64 `json:"outBufferSize,omitempty"`

	// security group
	SecurityGroup *SecurityGroupDetail `json:"securityGroup,omitempty"`

	// worker loop group
	WorkerLoopGroup *EventLoopGroupDetail `json:"workerLoopGroup,omitempty"`
}

// Validate validates this socks5 server detail
func (m *Socks5ServerDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptorLoopGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkerLoopGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Socks5ServerDetail) validateAcceptorLoopGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.AcceptorLoopGroup) { // not required
		return nil
	}

	if m.AcceptorLoopGroup != nil {
		if err := m.AcceptorLoopGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptorLoopGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptorLoopGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Socks5ServerDetail) validateBackend(formats strfmt.Registry) error {
	if swag.IsZero(m.Backend) { // not required
		return nil
	}

	if m.Backend != nil {
		if err := m.Backend.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backend")
			}
			return err
		}
	}

	return nil
}

func (m *Socks5ServerDetail) validateSecurityGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroup) { // not required
		return nil
	}

	if m.SecurityGroup != nil {
		if err := m.SecurityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Socks5ServerDetail) validateWorkerLoopGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkerLoopGroup) { // not required
		return nil
	}

	if m.WorkerLoopGroup != nil {
		if err := m.WorkerLoopGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workerLoopGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workerLoopGroup")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this socks5 server detail based on the context it is used
func (m *Socks5ServerDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcceptorLoopGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackend(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkerLoopGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Socks5ServerDetail) contextValidateAcceptorLoopGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.AcceptorLoopGroup != nil {
		if err := m.AcceptorLoopGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptorLoopGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptorLoopGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Socks5ServerDetail) contextValidateBackend(ctx context.Context, formats strfmt.Registry) error {

	if m.Backend != nil {
		if err := m.Backend.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backend")
			}
			return err
		}
	}

	return nil
}

func (m *Socks5ServerDetail) contextValidateSecurityGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityGroup != nil {
		if err := m.SecurityGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Socks5ServerDetail) contextValidateWorkerLoopGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkerLoopGroup != nil {
		if err := m.WorkerLoopGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workerLoopGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workerLoopGroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Socks5ServerDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Socks5ServerDetail) UnmarshalBinary(b []byte) error {
	var res Socks5ServerDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
