// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EquipmentRackEnclosureSlot Equipment:Rack Enclosure Slot
//
// Rack Server Slot in a RackEnclosure.
//
// swagger:model equipmentRackEnclosureSlot
type EquipmentRackEnclosureSlot struct {
	EquipmentBase

	// A collection of references to the [equipment.RackEnclosure](mo://equipment.RackEnclosure) Managed Object.
	// When this managed object is deleted, the referenced [equipment.RackEnclosure](mo://equipment.RackEnclosure) MO unsets its reference to this deleted MO.
	// Read Only: true
	EquipmentRackEnclosure *EquipmentRackEnclosureRef `json:"EquipmentRackEnclosure,omitempty"`

	// rack Id
	// Read Only: true
	RackID int64 `json:"RackId,omitempty"`

	// rack unit
	// Read Only: true
	RackUnit *ComputeRackUnitRef `json:"RackUnit,omitempty"`

	// rack unit dn
	// Read Only: true
	RackUnitDn string `json:"RackUnitDn,omitempty"`

	// The Device to which this Managed Object is associated.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *EquipmentRackEnclosureSlot) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		EquipmentRackEnclosure *EquipmentRackEnclosureRef `json:"EquipmentRackEnclosure,omitempty"`

		RackID int64 `json:"RackId,omitempty"`

		RackUnit *ComputeRackUnitRef `json:"RackUnit,omitempty"`

		RackUnitDn string `json:"RackUnitDn,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EquipmentRackEnclosure = dataAO1.EquipmentRackEnclosure

	m.RackID = dataAO1.RackID

	m.RackUnit = dataAO1.RackUnit

	m.RackUnitDn = dataAO1.RackUnitDn

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m EquipmentRackEnclosureSlot) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		EquipmentRackEnclosure *EquipmentRackEnclosureRef `json:"EquipmentRackEnclosure,omitempty"`

		RackID int64 `json:"RackId,omitempty"`

		RackUnit *ComputeRackUnitRef `json:"RackUnit,omitempty"`

		RackUnitDn string `json:"RackUnitDn,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.EquipmentRackEnclosure = m.EquipmentRackEnclosure

	dataAO1.RackID = m.RackID

	dataAO1.RackUnit = m.RackUnit

	dataAO1.RackUnitDn = m.RackUnitDn

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this equipment rack enclosure slot
func (m *EquipmentRackEnclosureSlot) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEquipmentRackEnclosure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRackUnit(formats); err != nil {
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

func (m *EquipmentRackEnclosureSlot) validateEquipmentRackEnclosure(formats strfmt.Registry) error {

	if swag.IsZero(m.EquipmentRackEnclosure) { // not required
		return nil
	}

	if m.EquipmentRackEnclosure != nil {
		if err := m.EquipmentRackEnclosure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EquipmentRackEnclosure")
			}
			return err
		}
	}

	return nil
}

func (m *EquipmentRackEnclosureSlot) validateRackUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.RackUnit) { // not required
		return nil
	}

	if m.RackUnit != nil {
		if err := m.RackUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RackUnit")
			}
			return err
		}
	}

	return nil
}

func (m *EquipmentRackEnclosureSlot) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *EquipmentRackEnclosureSlot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EquipmentRackEnclosureSlot) UnmarshalBinary(b []byte) error {
	var res EquipmentRackEnclosureSlot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
