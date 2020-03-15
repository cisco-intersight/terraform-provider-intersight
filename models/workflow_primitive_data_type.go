// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkflowPrimitiveDataType Workflow:Primitive Data Type
//
// This data type is used to represent primitives like string, floats and integers.
//
// swagger:model workflowPrimitiveDataType
type WorkflowPrimitiveDataType struct {
	WorkflowBaseDataType

	// Primitive data type properties.
	Properties *WorkflowPrimitiveDataProperty `json:"Properties,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowPrimitiveDataType) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowBaseDataType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowBaseDataType = aO0

	// AO1
	var dataAO1 struct {
		Properties *WorkflowPrimitiveDataProperty `json:"Properties,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Properties = dataAO1.Properties

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowPrimitiveDataType) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WorkflowBaseDataType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Properties *WorkflowPrimitiveDataProperty `json:"Properties,omitempty"`
	}

	dataAO1.Properties = m.Properties

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow primitive data type
func (m *WorkflowPrimitiveDataType) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowBaseDataType
	if err := m.WorkflowBaseDataType.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowPrimitiveDataType) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Properties")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowPrimitiveDataType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowPrimitiveDataType) UnmarshalBinary(b []byte) error {
	var res WorkflowPrimitiveDataType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
