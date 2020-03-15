// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FirmwareRunningFirmware Firmware:Running Firmware
//
// Running Firmware on an endpoint.
//
// swagger:model firmwareRunningFirmware
type FirmwareRunningFirmware struct {
	InventoryBase

	// A collection of references to the [bios.Unit](mo://bios.Unit) Managed Object.
	// When this managed object is deleted, the referenced [bios.Unit](mo://bios.Unit) MO unsets its reference to this deleted MO.
	// Read Only: true
	BiosUnit *BiosUnitRef `json:"BiosUnit,omitempty"`

	// Kind of the firmware - boot-booloader/system/kernel.
	// Read Only: true
	Component string `json:"Component,omitempty"`

	// A collection of references to the [management.Controller](mo://management.Controller) Managed Object.
	// When this managed object is deleted, the referenced [management.Controller](mo://management.Controller) MO unsets its reference to this deleted MO.
	// Read Only: true
	ManagementController *ManagementControllerRef `json:"ManagementController,omitempty"`

	// network elements
	NetworkElements []*NetworkElementRef `json:"NetworkElements"`

	// Package version which the firmware belongs to.
	// Read Only: true
	PackageVersion string `json:"PackageVersion,omitempty"`

	// The Device to which this Managed Object is associated.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// A collection of references to the [storage.Controller](mo://storage.Controller) Managed Object.
	// When this managed object is deleted, the referenced [storage.Controller](mo://storage.Controller) MO unsets its reference to this deleted MO.
	// Read Only: true
	StorageController *StorageControllerRef `json:"StorageController,omitempty"`

	// A collection of references to the [storage.PhysicalDisk](mo://storage.PhysicalDisk) Managed Object.
	// When this managed object is deleted, the referenced [storage.PhysicalDisk](mo://storage.PhysicalDisk) MO unsets its reference to this deleted MO.
	// Read Only: true
	StoragePhysicalDisk *StoragePhysicalDiskRef `json:"StoragePhysicalDisk,omitempty"`

	// Type of the firmware.
	// Read Only: true
	Type string `json:"Type,omitempty"`

	// Version of the firmware.
	// Read Only: true
	Version string `json:"Version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FirmwareRunningFirmware) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InventoryBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InventoryBase = aO0

	// AO1
	var dataAO1 struct {
		BiosUnit *BiosUnitRef `json:"BiosUnit,omitempty"`

		Component string `json:"Component,omitempty"`

		ManagementController *ManagementControllerRef `json:"ManagementController,omitempty"`

		NetworkElements []*NetworkElementRef `json:"NetworkElements"`

		PackageVersion string `json:"PackageVersion,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		StorageController *StorageControllerRef `json:"StorageController,omitempty"`

		StoragePhysicalDisk *StoragePhysicalDiskRef `json:"StoragePhysicalDisk,omitempty"`

		Type string `json:"Type,omitempty"`

		Version string `json:"Version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BiosUnit = dataAO1.BiosUnit

	m.Component = dataAO1.Component

	m.ManagementController = dataAO1.ManagementController

	m.NetworkElements = dataAO1.NetworkElements

	m.PackageVersion = dataAO1.PackageVersion

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.StorageController = dataAO1.StorageController

	m.StoragePhysicalDisk = dataAO1.StoragePhysicalDisk

	m.Type = dataAO1.Type

	m.Version = dataAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FirmwareRunningFirmware) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.InventoryBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		BiosUnit *BiosUnitRef `json:"BiosUnit,omitempty"`

		Component string `json:"Component,omitempty"`

		ManagementController *ManagementControllerRef `json:"ManagementController,omitempty"`

		NetworkElements []*NetworkElementRef `json:"NetworkElements"`

		PackageVersion string `json:"PackageVersion,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		StorageController *StorageControllerRef `json:"StorageController,omitempty"`

		StoragePhysicalDisk *StoragePhysicalDiskRef `json:"StoragePhysicalDisk,omitempty"`

		Type string `json:"Type,omitempty"`

		Version string `json:"Version,omitempty"`
	}

	dataAO1.BiosUnit = m.BiosUnit

	dataAO1.Component = m.Component

	dataAO1.ManagementController = m.ManagementController

	dataAO1.NetworkElements = m.NetworkElements

	dataAO1.PackageVersion = m.PackageVersion

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.StorageController = m.StorageController

	dataAO1.StoragePhysicalDisk = m.StoragePhysicalDisk

	dataAO1.Type = m.Type

	dataAO1.Version = m.Version

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this firmware running firmware
func (m *FirmwareRunningFirmware) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InventoryBase
	if err := m.InventoryBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBiosUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementController(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkElements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageController(formats); err != nil {
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

func (m *FirmwareRunningFirmware) validateBiosUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.BiosUnit) { // not required
		return nil
	}

	if m.BiosUnit != nil {
		if err := m.BiosUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BiosUnit")
			}
			return err
		}
	}

	return nil
}

func (m *FirmwareRunningFirmware) validateManagementController(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagementController) { // not required
		return nil
	}

	if m.ManagementController != nil {
		if err := m.ManagementController.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ManagementController")
			}
			return err
		}
	}

	return nil
}

func (m *FirmwareRunningFirmware) validateNetworkElements(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkElements) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkElements); i++ {
		if swag.IsZero(m.NetworkElements[i]) { // not required
			continue
		}

		if m.NetworkElements[i] != nil {
			if err := m.NetworkElements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NetworkElements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FirmwareRunningFirmware) validateRegisteredDevice(formats strfmt.Registry) error {

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

func (m *FirmwareRunningFirmware) validateStorageController(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageController) { // not required
		return nil
	}

	if m.StorageController != nil {
		if err := m.StorageController.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StorageController")
			}
			return err
		}
	}

	return nil
}

func (m *FirmwareRunningFirmware) validateStoragePhysicalDisk(formats strfmt.Registry) error {

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
func (m *FirmwareRunningFirmware) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareRunningFirmware) UnmarshalBinary(b []byte) error {
	var res FirmwareRunningFirmware
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
