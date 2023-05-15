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

// TCPLb Tcp lb
//
// swagger:model TcpLb
type TCPLb struct {

	// event loop group for accepting connections
	AcceptorLoopGroup string `json:"acceptorLoopGroup,omitempty"`

	// binding l4addr
	Address string `json:"address,omitempty"`

	// upstream reference for backend servers
	Backend string `json:"backend,omitempty"`

	// input buffer size
	InBufferSize int64 `json:"inBufferSize,omitempty"`

	// list of cert key
	ListOfCertKey []string `json:"listOfCertKey"`

	// name
	Name string `json:"name,omitempty"`

	// output buffer size
	OutBufferSize int64 `json:"outBufferSize,omitempty"`

	// protocol
	Protocol Protocol `json:"protocol,omitempty"`

	// security group reference for access control
	SecurityGroup string `json:"securityGroup,omitempty"`

	// event loop group for handling netflow
	WorkerLoopGroup string `json:"workerLoopGroup,omitempty"`
}

// Validate validates this Tcp lb
func (m *TCPLb) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TCPLb) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	if err := m.Protocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protocol")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("protocol")
		}
		return err
	}

	return nil
}

// ContextValidate validate this Tcp lb based on the context it is used
func (m *TCPLb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TCPLb) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Protocol.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protocol")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("protocol")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TCPLb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TCPLb) UnmarshalBinary(b []byte) error {
	var res TCPLb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
