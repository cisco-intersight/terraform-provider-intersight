// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IpmioverlanPolicy IPMI Over LAN
//
// Intelligent Platform Management Interface Over LAN Policy.
//
// swagger:model ipmioverlanPolicy
type IpmioverlanPolicy struct {
	PolicyAbstractPolicy

	// State of the IPMI Over LAN service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`

	// The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters.
	EncryptionKey string `json:"EncryptionKey,omitempty"`

	// Indicates whether the value of the 'encryptionKey' property has been set.
	// Read Only: true
	IsEncryptionKeySet *bool `json:"IsEncryptionKeySet,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// The highest privilege level that can be assigned to an IPMI session on a server.
	// Enum: [admin user read-only]
	Privilege *string `json:"Privilege,omitempty"`

	// Relationship to the profile object.
	Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IpmioverlanPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		Enabled *bool `json:"Enabled,omitempty"`

		EncryptionKey string `json:"EncryptionKey,omitempty"`

		IsEncryptionKeySet *bool `json:"IsEncryptionKeySet,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Privilege *string `json:"Privilege,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Enabled = dataAO1.Enabled

	m.EncryptionKey = dataAO1.EncryptionKey

	m.IsEncryptionKeySet = dataAO1.IsEncryptionKeySet

	m.Organization = dataAO1.Organization

	m.Privilege = dataAO1.Privilege

	m.Profiles = dataAO1.Profiles

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IpmioverlanPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Enabled *bool `json:"Enabled,omitempty"`

		EncryptionKey string `json:"EncryptionKey,omitempty"`

		IsEncryptionKeySet *bool `json:"IsEncryptionKeySet,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Privilege *string `json:"Privilege,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
	}

	dataAO1.Enabled = m.Enabled

	dataAO1.EncryptionKey = m.EncryptionKey

	dataAO1.IsEncryptionKeySet = m.IsEncryptionKeySet

	dataAO1.Organization = m.Organization

	dataAO1.Privilege = m.Privilege

	dataAO1.Profiles = m.Profiles

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this ipmioverlan policy
func (m *IpmioverlanPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivilege(formats); err != nil {
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

func (m *IpmioverlanPolicy) validateOrganization(formats strfmt.Registry) error {

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

var ipmioverlanPolicyTypePrivilegePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["admin","user","read-only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipmioverlanPolicyTypePrivilegePropEnum = append(ipmioverlanPolicyTypePrivilegePropEnum, v)
	}
}

// property enum
func (m *IpmioverlanPolicy) validatePrivilegeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipmioverlanPolicyTypePrivilegePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IpmioverlanPolicy) validatePrivilege(formats strfmt.Registry) error {

	if swag.IsZero(m.Privilege) { // not required
		return nil
	}

	// value enum
	if err := m.validatePrivilegeEnum("Privilege", "body", *m.Privilege); err != nil {
		return err
	}

	return nil
}

func (m *IpmioverlanPolicy) validateProfiles(formats strfmt.Registry) error {

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
func (m *IpmioverlanPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpmioverlanPolicy) UnmarshalBinary(b []byte) error {
	var res IpmioverlanPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
