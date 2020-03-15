// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NiaapiApicHweol Niaapi:Apic Hweol
//
// The hardware end of life notice for APIC.
//
// swagger:model niaapiApicHweol
type NiaapiApicHweol struct {
	NiaapiHardwareEol

	NiaapiApicHweolAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NiaapiApicHweol) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NiaapiHardwareEol
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NiaapiHardwareEol = aO0

	// AO1
	var aO1 NiaapiApicHweolAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.NiaapiApicHweolAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NiaapiApicHweol) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.NiaapiHardwareEol)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.NiaapiApicHweolAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this niaapi apic hweol
func (m *NiaapiApicHweol) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NiaapiHardwareEol
	if err := m.NiaapiHardwareEol.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with NiaapiApicHweolAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NiaapiApicHweol) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NiaapiApicHweol) UnmarshalBinary(b []byte) error {
	var res NiaapiApicHweol
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NiaapiApicHweolAllOf1 niaapi apic hweol all of1
//
// swagger:model NiaapiApicHweolAllOf1
type NiaapiApicHweolAllOf1 interface{}
