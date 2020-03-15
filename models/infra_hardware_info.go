// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InfraHardwareInfo Infra:Hardware Info
//
// Information about the hardware platform (cpu, memory).
//
// swagger:model infraHardwareInfo
type InfraHardwareInfo struct {
	MoBaseComplexType

	InfraHardwareInfoAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InfraHardwareInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 InfraHardwareInfoAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.InfraHardwareInfoAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InfraHardwareInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.InfraHardwareInfoAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this infra hardware info
func (m *InfraHardwareInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with InfraHardwareInfoAO1P1
	if err := m.InfraHardwareInfoAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InfraHardwareInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraHardwareInfo) UnmarshalBinary(b []byte) error {
	var res InfraHardwareInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InfraHardwareInfoAO1P1 infra hardware info a o1 p1
//
// swagger:model InfraHardwareInfoAO1P1
type InfraHardwareInfoAO1P1 struct {

	// The number of cpu cores on this hardware platform.
	CPUCores int64 `json:"CpuCores,omitempty"`

	// Speed of cpu in MHz. Usually cpu speeds are reported for modern cpus in GHz but MHz makes it more precise.
	CPUSpeed int64 `json:"CpuSpeed,omitempty"`

	// The amount of memory allocated (bytes) to this hardware platform.
	MemorySize int64 `json:"MemorySize,omitempty"`

	// infra hardware info a o1 p1
	InfraHardwareInfoAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *InfraHardwareInfoAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The number of cpu cores on this hardware platform.
		CPUCores int64 `json:"CpuCores,omitempty"`

		// Speed of cpu in MHz. Usually cpu speeds are reported for modern cpus in GHz but MHz makes it more precise.
		CPUSpeed int64 `json:"CpuSpeed,omitempty"`

		// The amount of memory allocated (bytes) to this hardware platform.
		MemorySize int64 `json:"MemorySize,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv InfraHardwareInfoAO1P1

	rcv.CPUCores = stage1.CPUCores
	rcv.CPUSpeed = stage1.CPUSpeed
	rcv.MemorySize = stage1.MemorySize
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "CpuCores")
	delete(stage2, "CpuSpeed")
	delete(stage2, "MemorySize")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.InfraHardwareInfoAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m InfraHardwareInfoAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The number of cpu cores on this hardware platform.
		CPUCores int64 `json:"CpuCores,omitempty"`

		// Speed of cpu in MHz. Usually cpu speeds are reported for modern cpus in GHz but MHz makes it more precise.
		CPUSpeed int64 `json:"CpuSpeed,omitempty"`

		// The amount of memory allocated (bytes) to this hardware platform.
		MemorySize int64 `json:"MemorySize,omitempty"`
	}

	stage1.CPUCores = m.CPUCores
	stage1.CPUSpeed = m.CPUSpeed
	stage1.MemorySize = m.MemorySize

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.InfraHardwareInfoAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.InfraHardwareInfoAO1P1)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this infra hardware info a o1 p1
func (m *InfraHardwareInfoAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfraHardwareInfoAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraHardwareInfoAO1P1) UnmarshalBinary(b []byte) error {
	var res InfraHardwareInfoAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
