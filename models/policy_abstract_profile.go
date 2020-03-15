// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PolicyAbstractProfile Policy:Abstract Profile
//
// Abstract base type for all profiles.
//
// swagger:model policyAbstractProfile
type PolicyAbstractProfile struct {
	MoBaseMo

	// Description of the profile.
	Description string `json:"Description,omitempty"`

	// Name of the concrete profile.
	Name string `json:"Name,omitempty"`

	// The source profile template to apply to the profile instance. All configuration settings from the profile template will be applied to the profile instance.
	SrcTemplate *PolicyAbstractProfileRef `json:"SrcTemplate,omitempty"`

	// Defines the type of the profile. Accepted value is instance.
	// Enum: [instance]
	Type *string `json:"Type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PolicyAbstractProfile) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Description string `json:"Description,omitempty"`

		Name string `json:"Name,omitempty"`

		SrcTemplate *PolicyAbstractProfileRef `json:"SrcTemplate,omitempty"`

		Type *string `json:"Type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Description = dataAO1.Description

	m.Name = dataAO1.Name

	m.SrcTemplate = dataAO1.SrcTemplate

	m.Type = dataAO1.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PolicyAbstractProfile) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Description string `json:"Description,omitempty"`

		Name string `json:"Name,omitempty"`

		SrcTemplate *PolicyAbstractProfileRef `json:"SrcTemplate,omitempty"`

		Type *string `json:"Type,omitempty"`
	}

	dataAO1.Description = m.Description

	dataAO1.Name = m.Name

	dataAO1.SrcTemplate = m.SrcTemplate

	dataAO1.Type = m.Type

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this policy abstract profile
func (m *PolicyAbstractProfile) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyAbstractProfile) validateSrcTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.SrcTemplate) { // not required
		return nil
	}

	if m.SrcTemplate != nil {
		if err := m.SrcTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SrcTemplate")
			}
			return err
		}
	}

	return nil
}

var policyAbstractProfileTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["instance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyAbstractProfileTypeTypePropEnum = append(policyAbstractProfileTypeTypePropEnum, v)
	}
}

// property enum
func (m *PolicyAbstractProfile) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, policyAbstractProfileTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PolicyAbstractProfile) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyAbstractProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyAbstractProfile) UnmarshalBinary(b []byte) error {
	var res PolicyAbstractProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
