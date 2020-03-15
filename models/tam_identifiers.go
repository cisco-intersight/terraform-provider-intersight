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

// TamIdentifiers Tam:Identifiers
//
// Identifiers are used to identify an affected object in an alert defition. please refer to https://Intersight.com/apidocs/introduction/#get-requests-and-query-capabilities for more details.
//
// swagger:model tamIdentifiers
type TamIdentifiers struct {
	MoBaseComplexType

	TamIdentifiersAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TamIdentifiers) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 TamIdentifiersAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TamIdentifiersAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TamIdentifiers) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TamIdentifiersAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tam identifiers
func (m *TamIdentifiers) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TamIdentifiersAO1P1
	if err := m.TamIdentifiersAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TamIdentifiers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamIdentifiers) UnmarshalBinary(b []byte) error {
	var res TamIdentifiers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TamIdentifiersAO1P1 tam identifiers a o1 p1
//
// swagger:model TamIdentifiersAO1P1
type TamIdentifiersAO1P1 struct {

	// Name of the filter paramter.
	Name string `json:"Name,omitempty"`

	// Value of the filter paramter.
	Value string `json:"Value,omitempty"`

	// tam identifiers a o1 p1
	TamIdentifiersAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *TamIdentifiersAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Name of the filter paramter.
		Name string `json:"Name,omitempty"`

		// Value of the filter paramter.
		Value string `json:"Value,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv TamIdentifiersAO1P1

	rcv.Name = stage1.Name
	rcv.Value = stage1.Value
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Name")
	delete(stage2, "Value")
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
		m.TamIdentifiersAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m TamIdentifiersAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Name of the filter paramter.
		Name string `json:"Name,omitempty"`

		// Value of the filter paramter.
		Value string `json:"Value,omitempty"`
	}

	stage1.Name = m.Name
	stage1.Value = m.Value

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.TamIdentifiersAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.TamIdentifiersAO1P1)
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

// Validate validates this tam identifiers a o1 p1
func (m *TamIdentifiersAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TamIdentifiersAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamIdentifiersAO1P1) UnmarshalBinary(b []byte) error {
	var res TamIdentifiersAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
