// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StoragePurePort Storage:Pure Port
//
// Port entity in Pure FlashArray.
//
// swagger:model storagePurePort
type StoragePurePort struct {
	StorageBasePhysicalPort

	// Storage array managed object.
	// Read Only: true
	Array *StoragePureArrayRef `json:"Array,omitempty"`

	// Parent storage array controller object.
	// Read Only: true
	Controller *StoragePureControllerRef `json:"Controller,omitempty"`

	// Name of the port to which this port has failed over.
	// Read Only: true
	Failover string `json:"Failover,omitempty"`

	// Ip address of iSCSI portal configured on the port.
	// Read Only: true
	Portal string `json:"Portal,omitempty"`

	// Device registration managed object that represents this storage array connection to Intersight.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StoragePurePort) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageBasePhysicalPort
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageBasePhysicalPort = aO0

	// AO1
	var dataAO1 struct {
		Array *StoragePureArrayRef `json:"Array,omitempty"`

		Controller *StoragePureControllerRef `json:"Controller,omitempty"`

		Failover string `json:"Failover,omitempty"`

		Portal string `json:"Portal,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Array = dataAO1.Array

	m.Controller = dataAO1.Controller

	m.Failover = dataAO1.Failover

	m.Portal = dataAO1.Portal

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StoragePurePort) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.StorageBasePhysicalPort)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Array *StoragePureArrayRef `json:"Array,omitempty"`

		Controller *StoragePureControllerRef `json:"Controller,omitempty"`

		Failover string `json:"Failover,omitempty"`

		Portal string `json:"Portal,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.Array = m.Array

	dataAO1.Controller = m.Controller

	dataAO1.Failover = m.Failover

	dataAO1.Portal = m.Portal

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage pure port
func (m *StoragePurePort) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageBasePhysicalPort
	if err := m.StorageBasePhysicalPort.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArray(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateController(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePurePort) validateArray(formats strfmt.Registry) error {

	if swag.IsZero(m.Array) { // not required
		return nil
	}

	if m.Array != nil {
		if err := m.Array.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Array")
			}
			return err
		}
	}

	return nil
}

func (m *StoragePurePort) validateController(formats strfmt.Registry) error {

	if swag.IsZero(m.Controller) { // not required
		return nil
	}

	if m.Controller != nil {
		if err := m.Controller.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Controller")
			}
			return err
		}
	}

	return nil
}

func (m *StoragePurePort) validateRegisteredDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.RegisteredDevice) { // not required
		return nil
	}

	if m.RegisteredDevice != nil {
		if err := m.RegisteredDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegisteredDevice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoragePurePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePurePort) UnmarshalBinary(b []byte) error {
	var res StoragePurePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
