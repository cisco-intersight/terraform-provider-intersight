// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// I18nMessageParam I18n:Message Param
//
// A key, value pair used to substitue variables when translating a message in a particular locale for
// internationalization purpose.
//
// swagger:model i18nMessageParam
type I18nMessageParam struct {
	I18nMessageParamAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *I18nMessageParam) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 I18nMessageParamAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.I18nMessageParamAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m I18nMessageParam) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.I18nMessageParamAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this i18n message param
func (m *I18nMessageParam) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with I18nMessageParamAO0P0
	if err := m.I18nMessageParamAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *I18nMessageParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *I18nMessageParam) UnmarshalBinary(b []byte) error {
	var res I18nMessageParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// I18nMessageParamAO0P0 i18n message param a o0 p0
// swagger:model I18nMessageParamAO0P0
type I18nMessageParamAO0P0 struct {

	// The name of a variable which is referenced in a i18n text template.
	//
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// The value of a variable which is substituted in a i18n text template.
	//
	// Read Only: true
	Value string `json:"Value,omitempty"`

	// i18n message param a o0 p0
	I18nMessageParamAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *I18nMessageParamAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The name of a variable which is referenced in a i18n text template.
		//
		// Read Only: true
		Name string `json:"Name,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The value of a variable which is substituted in a i18n text template.
		//
		// Read Only: true
		Value string `json:"Value,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv I18nMessageParamAO0P0

	rcv.Name = stage1.Name

	rcv.ObjectType = stage1.ObjectType

	rcv.Value = stage1.Value

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Name")

	delete(stage2, "ObjectType")

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
		m.I18nMessageParamAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m I18nMessageParamAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The name of a variable which is referenced in a i18n text template.
		//
		// Read Only: true
		Name string `json:"Name,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The value of a variable which is substituted in a i18n text template.
		//
		// Read Only: true
		Value string `json:"Value,omitempty"`
	}

	stage1.Name = m.Name

	stage1.ObjectType = m.ObjectType

	stage1.Value = m.Value

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.I18nMessageParamAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.I18nMessageParamAO0P0)
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

// Validate validates this i18n message param a o0 p0
func (m *I18nMessageParamAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *I18nMessageParamAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *I18nMessageParamAO0P0) UnmarshalBinary(b []byte) error {
	var res I18nMessageParamAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}