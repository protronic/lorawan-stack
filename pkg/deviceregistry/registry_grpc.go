// Copyright © 2017 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package deviceregistry

import (
	"github.com/TheThingsNetwork/ttn/pkg/component"
	"github.com/TheThingsNetwork/ttn/pkg/errors"
	"github.com/TheThingsNetwork/ttn/pkg/ttnpb"
	pbtypes "github.com/gogo/protobuf/types"
	"context"
)

// ErrDeviceNotFound represents the ErrDescriptor of the error returned
// when the device is not found.
var ErrDeviceNotFound = &errors.ErrDescriptor{
	MessageFormat: "Device not found",
	Type:          errors.NotFound,
	Code:          1,
}

// ErrTooManyDevices represents the ErrDescriptor of the error returned
// when there are too many devices associated with the identifiers specified.
var ErrTooManyDevices = &errors.ErrDescriptor{
	MessageFormat: "Too many devices are associated with identifiers specified",
	Type:          errors.Conflict,
	Code:          2,
}

func init() {
	ErrDeviceNotFound.Register()
	ErrTooManyDevices.Register()
}

// RegistryRPC implements the device registry gRPC service.
type RegistryRPC struct {
	Interface
	*component.Component

	checks struct {
		ListDevices  func(context.Context, *ttnpb.EndDeviceIdentifiers) error
		GetDevice    func(context.Context, *ttnpb.EndDeviceIdentifiers) error
		SetDevice    func(context.Context, *ttnpb.EndDevice) error
		DeleteDevice func(context.Context, *ttnpb.EndDeviceIdentifiers) error
	}
}

// Option represents RegistryRPC option
type Option func(*RegistryRPC)

// WithListDevicesCheck sets a check to ListDevices method of RegistryRPC instance.
// ListDevices first executes fn and if error is returned by it,
// returns error, otherwise execution advances as usual.
func WithListDevicesCheck(fn func(context.Context, *ttnpb.EndDeviceIdentifiers) error) Option {
	return func(r *RegistryRPC) { r.checks.ListDevices = fn }
}

// WithGetDeviceCheck sets a check to GetDevice method of RegistryRPC instance.
// GetDevice first executes fn and if error is returned by it,
// returns error, otherwise execution advances as usual.
func WithGetDeviceCheck(fn func(context.Context, *ttnpb.EndDeviceIdentifiers) error) Option {
	return func(r *RegistryRPC) { r.checks.GetDevice = fn }
}

// WithSetDeviceCheck sets a check to SetDevice method of RegistryRPC instance.
// SetDevice first executes fn and if error is returned by it,
// returns error, otherwise execution advances as usual.
func WithSetDeviceCheck(fn func(context.Context, *ttnpb.EndDevice) error) Option {
	return func(r *RegistryRPC) { r.checks.SetDevice = fn }
}

// WithDeleteDeviceCheck sets a check to DeleteDevice method of RegistryRPC instance.
// DeleteDevice first executes fn and if error is returned by it,
// returns error, otherwise execution advances as usual.
func WithDeleteDeviceCheck(fn func(context.Context, *ttnpb.EndDeviceIdentifiers) error) Option {
	return func(r *RegistryRPC) { r.checks.DeleteDevice = fn }
}

// NewRPC returns a new instance of RegistryRPC
func NewRPC(c *component.Component, r Interface, opts ...Option) *RegistryRPC {
	rpc := &RegistryRPC{
		Component: c,
		Interface: r,
	}
	for _, opt := range opts {
		opt(rpc)
	}
	return rpc
}

// ListDevices lists devices matching filter in underlying registry.
func (r *RegistryRPC) ListDevices(ctx context.Context, filter *ttnpb.EndDeviceIdentifiers) (*ttnpb.EndDevices, error) {
	if r.checks.ListDevices != nil {
		if err := r.checks.ListDevices(ctx, filter); err != nil {
			return nil, err
		}
	}

	devs, err := r.Interface.FindDeviceByIdentifiers(filter)
	if err != nil {
		return nil, err
	}
	if len(devs) == 0 {
		return nil, ErrDeviceNotFound.New(nil)
	}
	eds := make([]*ttnpb.EndDevice, len(devs))
	for i, dev := range devs {
		eds[i] = dev.EndDevice
	}
	return &ttnpb.EndDevices{eds}, nil
}

// GetDevice returns the device associated with id in underlying registry, if found.
func (r *RegistryRPC) GetDevice(ctx context.Context, id *ttnpb.EndDeviceIdentifiers) (*ttnpb.EndDevice, error) {
	if r.checks.GetDevice != nil {
		if err := r.checks.GetDevice(ctx, id); err != nil {
			return nil, err
		}
	}

	devs, err := r.Interface.FindDeviceByIdentifiers(id)
	if err != nil {
		return nil, err
	}
	switch len(devs) {
	case 0:
		return nil, ErrDeviceNotFound.New(nil)
	case 1:
		return devs[0].EndDevice, nil
	default:
		return nil, ErrTooManyDevices.New(nil)
	}
}

// SetDevice sets the device fields to match those of dev in underlying registry.
func (r *RegistryRPC) SetDevice(ctx context.Context, dev *ttnpb.EndDevice) (*pbtypes.Empty, error) {
	if r.checks.SetDevice != nil {
		if err := r.checks.SetDevice(ctx, dev); err != nil {
			return nil, err
		}
	}

	devs, err := r.Interface.FindDeviceByIdentifiers(&ttnpb.EndDeviceIdentifiers{DeviceID: dev.DeviceID})
	if err != nil {
		return nil, err
	}
	switch len(devs) {
	case 0:
		_, err := r.Interface.Create(dev)
		if err != nil {
			return nil, err
		}
		return &pbtypes.Empty{}, nil
	case 1:
		return &pbtypes.Empty{}, devs[0].Update()
	default:
		return nil, ErrTooManyDevices.New(nil)
	}
}

// DeleteDevice deletes the device associated with id from underlying registry.
func (r *RegistryRPC) DeleteDevice(ctx context.Context, id *ttnpb.EndDeviceIdentifiers) (*pbtypes.Empty, error) {
	if r.checks.DeleteDevice != nil {
		if err := r.checks.DeleteDevice(ctx, id); err != nil {
			return nil, err
		}
	}

	devs, err := r.Interface.FindDeviceByIdentifiers(id)
	if err != nil {
		return nil, err
	}
	switch len(devs) {
	case 0:
		return nil, ErrDeviceNotFound.New(nil)
	case 1:
		return &pbtypes.Empty{}, devs[0].Delete()
	default:
		return nil, ErrTooManyDevices.New(nil)
	}
}
