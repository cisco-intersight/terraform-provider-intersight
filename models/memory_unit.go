// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MemoryUnit Memory:Unit
//
// This represents a memory DIMM on a server.
//
// swagger:model memoryUnit
type MemoryUnit struct {
	MemoryAbstractUnit

	// A collection of references to the [memory.Array](mo://memory.Array) Managed Object.
	// When this managed object is deleted, the referenced [memory.Array](mo://memory.Array) MO unsets its reference to this deleted MO.
	// Read Only: true
	MemoryArray *MemoryArrayRef `json:"MemoryArray,omitempty"`

	// This represents the ID of a regular DIMM on a server.
	// Read Only: true
	MemoryID int64 `json:"MemoryId,omitempty"`

	// The Device to which this Managed Object is associated.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MemoryUnit) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MemoryAbstractUnit
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MemoryAbstractUnit = aO0

	// AO1
	var dataAO1 struct {
		MemoryArray *MemoryArrayRef `json:"MemoryArray,omitempty"`

		MemoryID int64 `json:"MemoryId,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.MemoryArray = dataAO1.MemoryArray

	m.MemoryID = dataAO1.MemoryID

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MemoryUnit) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MemoryAbstractUnit)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		MemoryArray *MemoryArrayRef `json:"MemoryArray,omitempty"`

		MemoryID int64 `json:"MemoryId,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.MemoryArray = m.MemoryArray

	dataAO1.MemoryID = m.MemoryID

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this memory unit
func (m *MemoryUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MemoryAbstractUnit
	if err := m.MemoryAbstractUnit.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryArray(formats); err != nil {
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

func (m *MemoryUnit) validateMemoryArray(formats strfmt.Registry) error {

	if swag.IsZero(m.MemoryArray) { // not required
		return nil
	}

	if m.MemoryArray != nil {
		if err := m.MemoryArray.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MemoryArray")
			}
			return err
		}
	}

	return nil
}

func (m *MemoryUnit) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *MemoryUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemoryUnit) UnmarshalBinary(b []byte) error {
	var res MemoryUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
