// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceconnectorPolicy Device Connector
//
// Policy to control configuration changes allowed from Cisco IMC.
//
// swagger:model deviceconnectorPolicy
type DeviceconnectorPolicy struct {
	PolicyAbstractPolicy

	// Enables configuration lockout on the endpoint.
	LockoutEnabled *bool `json:"LockoutEnabled,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// Relationship to the profile object.
	Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DeviceconnectorPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		LockoutEnabled *bool `json:"LockoutEnabled,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.LockoutEnabled = dataAO1.LockoutEnabled

	m.Organization = dataAO1.Organization

	m.Profiles = dataAO1.Profiles

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DeviceconnectorPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		LockoutEnabled *bool `json:"LockoutEnabled,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
	}

	dataAO1.LockoutEnabled = m.LockoutEnabled

	dataAO1.Organization = m.Organization

	dataAO1.Profiles = m.Profiles

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this deviceconnector policy
func (m *DeviceconnectorPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceconnectorPolicy) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceconnectorPolicy) validateProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Profiles) { // not required
		return nil
	}

	for i := 0; i < len(m.Profiles); i++ {
		if swag.IsZero(m.Profiles[i]) { // not required
			continue
		}

		if m.Profiles[i] != nil {
			if err := m.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceconnectorPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceconnectorPolicy) UnmarshalBinary(b []byte) error {
	var res DeviceconnectorPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
