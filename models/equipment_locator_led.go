// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EquipmentLocatorLed Equipment:Locator Led
//
// Locator Led of an Equipment like Rack, Disk etc.
//
// swagger:model equipmentLocatorLed
type EquipmentLocatorLed struct {
	InventoryBase

	// color
	// Read Only: true
	Color string `json:"Color,omitempty"`

	// A collection of references to the [compute.Blade](mo://compute.Blade) Managed Object.
	// When this managed object is deleted, the referenced [compute.Blade](mo://compute.Blade) MO unsets its reference to this deleted MO.
	// Read Only: true
	ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

	// A collection of references to the [compute.RackUnit](mo://compute.RackUnit) Managed Object.
	// When this managed object is deleted, the referenced [compute.RackUnit](mo://compute.RackUnit) MO unsets its reference to this deleted MO.
	// Read Only: true
	ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

	// oper state
	// Read Only: true
	OperState string `json:"OperState,omitempty"`

	// The Device to which this Managed Object is associated.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// A collection of references to the [storage.PhysicalDisk](mo://storage.PhysicalDisk) Managed Object.
	// When this managed object is deleted, the referenced [storage.PhysicalDisk](mo://storage.PhysicalDisk) MO unsets its reference to this deleted MO.
	// Read Only: true
	StoragePhysicalDisk *StoragePhysicalDiskRef `json:"StoragePhysicalDisk,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *EquipmentLocatorLed) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InventoryBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InventoryBase = aO0

	// AO1
	var dataAO1 struct {
		Color string `json:"Color,omitempty"`

		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

		OperState string `json:"OperState,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		StoragePhysicalDisk *StoragePhysicalDiskRef `json:"StoragePhysicalDisk,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Color = dataAO1.Color

	m.ComputeBlade = dataAO1.ComputeBlade

	m.ComputeRackUnit = dataAO1.ComputeRackUnit

	m.OperState = dataAO1.OperState

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.StoragePhysicalDisk = dataAO1.StoragePhysicalDisk

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m EquipmentLocatorLed) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.InventoryBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Color string `json:"Color,omitempty"`

		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

		OperState string `json:"OperState,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		StoragePhysicalDisk *StoragePhysicalDiskRef `json:"StoragePhysicalDisk,omitempty"`
	}

	dataAO1.Color = m.Color

	dataAO1.ComputeBlade = m.ComputeBlade

	dataAO1.ComputeRackUnit = m.ComputeRackUnit

	dataAO1.OperState = m.OperState

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.StoragePhysicalDisk = m.StoragePhysicalDisk

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this equipment locator led
func (m *EquipmentLocatorLed) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InventoryBase
	if err := m.InventoryBase.Validate(formats); err != nil {
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

	if err := m.validateStoragePhysicalDisk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EquipmentLocatorLed) validateComputeBlade(formats strfmt.Registry) error {

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

func (m *EquipmentLocatorLed) validateComputeRackUnit(formats strfmt.Registry) error {

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

func (m *EquipmentLocatorLed) validateRegisteredDevice(formats strfmt.Registry) error {

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

func (m *EquipmentLocatorLed) validateStoragePhysicalDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.StoragePhysicalDisk) { // not required
		return nil
	}

	if m.StoragePhysicalDisk != nil {
		if err := m.StoragePhysicalDisk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StoragePhysicalDisk")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EquipmentLocatorLed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EquipmentLocatorLed) UnmarshalBinary(b []byte) error {
	var res EquipmentLocatorLed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
