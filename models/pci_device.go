// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PciDevice Pci:Device
//
// PCI device present in a server.
//
// swagger:model pciDevice
type PciDevice struct {
	EquipmentBase

	// A collection of references to the [compute.Blade](mo://compute.Blade) Managed Object.
	// When this managed object is deleted, the referenced [compute.Blade](mo://compute.Blade) MO unsets its reference to this deleted MO.
	// Read Only: true
	ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

	// A collection of references to the [compute.RackUnit](mo://compute.RackUnit) Managed Object.
	// When this managed object is deleted, the referenced [compute.RackUnit](mo://compute.RackUnit) MO unsets its reference to this deleted MO.
	// Read Only: true
	ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

	// It shows the running firmware version.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// It shows the product identifier.
	Pid string `json:"Pid,omitempty"`

	// The Device to which this Managed Object is associated.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// It show PCI slot id of the device.
	// Read Only: true
	SlotID string `json:"SlotId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PciDevice) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

		FirmwareVersion string `json:"FirmwareVersion,omitempty"`

		Pid string `json:"Pid,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		SlotID string `json:"SlotId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ComputeBlade = dataAO1.ComputeBlade

	m.ComputeRackUnit = dataAO1.ComputeRackUnit

	m.FirmwareVersion = dataAO1.FirmwareVersion

	m.Pid = dataAO1.Pid

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.SlotID = dataAO1.SlotID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PciDevice) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

		FirmwareVersion string `json:"FirmwareVersion,omitempty"`

		Pid string `json:"Pid,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		SlotID string `json:"SlotId,omitempty"`
	}

	dataAO1.ComputeBlade = m.ComputeBlade

	dataAO1.ComputeRackUnit = m.ComputeRackUnit

	dataAO1.FirmwareVersion = m.FirmwareVersion

	dataAO1.Pid = m.Pid

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.SlotID = m.SlotID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pci device
func (m *PciDevice) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeBlade(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeRackUnit(formats); err != nil {
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

func (m *PciDevice) validateComputeBlade(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeBlade) { // not required
		return nil
	}

	if m.ComputeBlade != nil {
		if err := m.ComputeBlade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeBlade")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) validateComputeRackUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeRackUnit) { // not required
		return nil
	}

	if m.ComputeRackUnit != nil {
		if err := m.ComputeRackUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeRackUnit")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *PciDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PciDevice) UnmarshalBinary(b []byte) error {
	var res PciDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
