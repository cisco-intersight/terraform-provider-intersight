// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VnicSanConnectivityPolicy SAN Connectivity
//
// SAN connectivity policy determines the network storage resources and the connections between the server and the SAN on the network. This policy enables you to configure vHBAs that the servers use to communicate with the storage network.
//
// swagger:model vnicSanConnectivityPolicy
type VnicSanConnectivityPolicy struct {
	PolicyAbstractPolicy

	// A collection of references to the [vnic.FcIf](mo://vnic.FcIf) Managed Object.
	//
	// When this managed object is deleted, the referenced [vnic.FcIf](mo://vnic.FcIf) MOs unset their reference to this deleted MO.
	//
	FcIfs []*VnicFcIfRef `json:"FcIfs"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// Relationship to the server profile.
	//
	Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicSanConnectivityPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		FcIfs []*VnicFcIfRef `json:"FcIfs"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.FcIfs = dataAO1.FcIfs

	m.Organization = dataAO1.Organization

	m.Profiles = dataAO1.Profiles

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicSanConnectivityPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		FcIfs []*VnicFcIfRef `json:"FcIfs"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
	}

	dataAO1.FcIfs = m.FcIfs

	dataAO1.Organization = m.Organization

	dataAO1.Profiles = m.Profiles

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic san connectivity policy
func (m *VnicSanConnectivityPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFcIfs(formats); err != nil {
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

func (m *VnicSanConnectivityPolicy) validateFcIfs(formats strfmt.Registry) error {

	if swag.IsZero(m.FcIfs) { // not required
		return nil
	}

	for i := 0; i < len(m.FcIfs); i++ {
		if swag.IsZero(m.FcIfs[i]) { // not required
			continue
		}

		if m.FcIfs[i] != nil {
			if err := m.FcIfs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FcIfs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VnicSanConnectivityPolicy) validateOrganization(formats strfmt.Registry) error {

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

func (m *VnicSanConnectivityPolicy) validateProfiles(formats strfmt.Registry) error {

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
func (m *VnicSanConnectivityPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicSanConnectivityPolicy) UnmarshalBinary(b []byte) error {
	var res VnicSanConnectivityPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}