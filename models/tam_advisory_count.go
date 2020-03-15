// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TamAdvisoryCount Tam:Advisory Count
//
// Total number of advisories currently affecting a given Account.
//
// swagger:model tamAdvisoryCount
type TamAdvisoryCount struct {
	MoBaseMo

	// The account associated with Advisory count.
	// Read Only: true
	Account *IamAccountRef `json:"Account,omitempty"`

	// Total number of advisories affecting the account.
	AdvisoryCount int64 `json:"AdvisoryCount,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TamAdvisoryCount) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		AdvisoryCount int64 `json:"AdvisoryCount,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.AdvisoryCount = dataAO1.AdvisoryCount

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TamAdvisoryCount) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		AdvisoryCount int64 `json:"AdvisoryCount,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.AdvisoryCount = m.AdvisoryCount

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tam advisory count
func (m *TamAdvisoryCount) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TamAdvisoryCount) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TamAdvisoryCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamAdvisoryCount) UnmarshalBinary(b []byte) error {
	var res TamAdvisoryCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
