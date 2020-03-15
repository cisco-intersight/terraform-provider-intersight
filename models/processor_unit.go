// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProcessorUnit Processor:Unit
//
// The CPU present on a server.
//
// swagger:model processorUnit
type ProcessorUnit struct {
	EquipmentBase

	// architecture
	// Read Only: true
	Architecture string `json:"Architecture,omitempty"`

	// A collection of references to the [compute.Board](mo://compute.Board) Managed Object.
	// When this managed object is deleted, the referenced [compute.Board](mo://compute.Board) MO unsets its reference to this deleted MO.
	// Read Only: true
	ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

	// num cores
	// Read Only: true
	NumCores int64 `json:"NumCores,omitempty"`

	// num cores enabled
	// Read Only: true
	NumCoresEnabled string `json:"NumCoresEnabled,omitempty"`

	// num threads
	// Read Only: true
	NumThreads string `json:"NumThreads,omitempty"`

	// oper power state
	// Read Only: true
	OperPowerState string `json:"OperPowerState,omitempty"`

	// oper state
	// Read Only: true
	OperState string `json:"OperState,omitempty"`

	// operability
	// Read Only: true
	Operability string `json:"Operability,omitempty"`

	// presence
	// Read Only: true
	Presence string `json:"Presence,omitempty"`

	// processor Id
	// Read Only: true
	ProcessorID int64 `json:"ProcessorId,omitempty"`

	// The Device to which this Managed Object is associated.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// socket designation
	// Read Only: true
	SocketDesignation string `json:"SocketDesignation,omitempty"`

	// speed
	// Read Only: true
	Speed float32 `json:"Speed,omitempty"`

	// stepping
	// Read Only: true
	Stepping string `json:"Stepping,omitempty"`

	// thermal
	// Read Only: true
	Thermal string `json:"Thermal,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProcessorUnit) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		Architecture string `json:"Architecture,omitempty"`

		ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

		NumCores int64 `json:"NumCores,omitempty"`

		NumCoresEnabled string `json:"NumCoresEnabled,omitempty"`

		NumThreads string `json:"NumThreads,omitempty"`

		OperPowerState string `json:"OperPowerState,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Operability string `json:"Operability,omitempty"`

		Presence string `json:"Presence,omitempty"`

		ProcessorID int64 `json:"ProcessorId,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		SocketDesignation string `json:"SocketDesignation,omitempty"`

		Speed float32 `json:"Speed,omitempty"`

		Stepping string `json:"Stepping,omitempty"`

		Thermal string `json:"Thermal,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Architecture = dataAO1.Architecture

	m.ComputeBoard = dataAO1.ComputeBoard

	m.NumCores = dataAO1.NumCores

	m.NumCoresEnabled = dataAO1.NumCoresEnabled

	m.NumThreads = dataAO1.NumThreads

	m.OperPowerState = dataAO1.OperPowerState

	m.OperState = dataAO1.OperState

	m.Operability = dataAO1.Operability

	m.Presence = dataAO1.Presence

	m.ProcessorID = dataAO1.ProcessorID

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.SocketDesignation = dataAO1.SocketDesignation

	m.Speed = dataAO1.Speed

	m.Stepping = dataAO1.Stepping

	m.Thermal = dataAO1.Thermal

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProcessorUnit) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Architecture string `json:"Architecture,omitempty"`

		ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

		NumCores int64 `json:"NumCores,omitempty"`

		NumCoresEnabled string `json:"NumCoresEnabled,omitempty"`

		NumThreads string `json:"NumThreads,omitempty"`

		OperPowerState string `json:"OperPowerState,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Operability string `json:"Operability,omitempty"`

		Presence string `json:"Presence,omitempty"`

		ProcessorID int64 `json:"ProcessorId,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		SocketDesignation string `json:"SocketDesignation,omitempty"`

		Speed float32 `json:"Speed,omitempty"`

		Stepping string `json:"Stepping,omitempty"`

		Thermal string `json:"Thermal,omitempty"`
	}

	dataAO1.Architecture = m.Architecture

	dataAO1.ComputeBoard = m.ComputeBoard

	dataAO1.NumCores = m.NumCores

	dataAO1.NumCoresEnabled = m.NumCoresEnabled

	dataAO1.NumThreads = m.NumThreads

	dataAO1.OperPowerState = m.OperPowerState

	dataAO1.OperState = m.OperState

	dataAO1.Operability = m.Operability

	dataAO1.Presence = m.Presence

	dataAO1.ProcessorID = m.ProcessorID

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.SocketDesignation = m.SocketDesignation

	dataAO1.Speed = m.Speed

	dataAO1.Stepping = m.Stepping

	dataAO1.Thermal = m.Thermal

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this processor unit
func (m *ProcessorUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeBoard(formats); err != nil {
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

func (m *ProcessorUnit) validateComputeBoard(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeBoard) { // not required
		return nil
	}

	if m.ComputeBoard != nil {
		if err := m.ComputeBoard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeBoard")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessorUnit) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *ProcessorUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessorUnit) UnmarshalBinary(b []byte) error {
	var res ProcessorUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
