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

// CommIPV6Interface Comm:Ip V6 Interface
//
// The configuration data of a single IPv6 interface, including IP address, IPv6 prefix and default gateway.
//
// swagger:model commIpV6Interface
type CommIPV6Interface struct {
	MoBaseComplexType

	CommIPV6InterfaceAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CommIPV6Interface) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 CommIPV6InterfaceAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CommIPV6InterfaceAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CommIPV6Interface) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.CommIPV6InterfaceAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this comm Ip v6 interface
func (m *CommIPV6Interface) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with CommIPV6InterfaceAO1P1
	if err := m.CommIPV6InterfaceAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CommIPV6Interface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommIPV6Interface) UnmarshalBinary(b []byte) error {
	var res CommIPV6Interface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CommIPV6InterfaceAO1P1 comm IP v6 interface a o1 p1
//
// swagger:model CommIPV6InterfaceAO1P1
type CommIPV6InterfaceAO1P1 struct {

	// The IPv6 address of the default gateway.
	Gateway string `json:"Gateway,omitempty"`

	// The IPv6 Address, represented as eight groups of four hexadecimal digits, separated by colons.
	IPAddress string `json:"IpAddress,omitempty"`

	// The IPv6 Prefix, represented as eight groups of four hexadecimal digits, separated by colons.
	Prefix string `json:"Prefix,omitempty"`

	// comm IP v6 interface a o1 p1
	CommIPV6InterfaceAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *CommIPV6InterfaceAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The IPv6 address of the default gateway.
		Gateway string `json:"Gateway,omitempty"`

		// The IPv6 Address, represented as eight groups of four hexadecimal digits, separated by colons.
		IPAddress string `json:"IpAddress,omitempty"`

		// The IPv6 Prefix, represented as eight groups of four hexadecimal digits, separated by colons.
		Prefix string `json:"Prefix,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv CommIPV6InterfaceAO1P1

	rcv.Gateway = stage1.Gateway
	rcv.IPAddress = stage1.IPAddress
	rcv.Prefix = stage1.Prefix
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Gateway")
	delete(stage2, "IpAddress")
	delete(stage2, "Prefix")
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
		m.CommIPV6InterfaceAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m CommIPV6InterfaceAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The IPv6 address of the default gateway.
		Gateway string `json:"Gateway,omitempty"`

		// The IPv6 Address, represented as eight groups of four hexadecimal digits, separated by colons.
		IPAddress string `json:"IpAddress,omitempty"`

		// The IPv6 Prefix, represented as eight groups of four hexadecimal digits, separated by colons.
		Prefix string `json:"Prefix,omitempty"`
	}

	stage1.Gateway = m.Gateway
	stage1.IPAddress = m.IPAddress
	stage1.Prefix = m.Prefix

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.CommIPV6InterfaceAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.CommIPV6InterfaceAO1P1)
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

// Validate validates this comm IP v6 interface a o1 p1
func (m *CommIPV6InterfaceAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommIPV6InterfaceAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommIPV6InterfaceAO1P1) UnmarshalBinary(b []byte) error {
	var res CommIPV6InterfaceAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
