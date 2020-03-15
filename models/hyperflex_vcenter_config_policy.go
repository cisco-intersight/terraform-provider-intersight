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

// HyperflexVcenterConfigPolicy vCenter
//
// A policy specifying vCenter configuration.
//
// swagger:model hyperflexVcenterConfigPolicy
type HyperflexVcenterConfigPolicy struct {
	PolicyAbstractPolicy

	// List of cluster profiles using this policy.
	ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

	// The vCenter datacenter name.
	DataCenter string `json:"DataCenter,omitempty"`

	// The vCenter server FQDN or IP.
	Hostname string `json:"Hostname,omitempty"`

	// is password set
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// The password for authenticating with vCenter. Follow the corresponding password policy governed by vCenter.
	Password string `json:"Password,omitempty"`

	// Overrides the default vCenter Single Sign-On URL. Do not specify unless instructed by Cisco TAC.
	SsoURL string `json:"SsoUrl,omitempty"`

	// The vCenter username (e.g. administrator@vsphere.local).
	Username string `json:"Username,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexVcenterConfigPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

		DataCenter string `json:"DataCenter,omitempty"`

		Hostname string `json:"Hostname,omitempty"`

		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Password string `json:"Password,omitempty"`

		SsoURL string `json:"SsoUrl,omitempty"`

		Username string `json:"Username,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ClusterProfiles = dataAO1.ClusterProfiles

	m.DataCenter = dataAO1.DataCenter

	m.Hostname = dataAO1.Hostname

	m.IsPasswordSet = dataAO1.IsPasswordSet

	m.Organization = dataAO1.Organization

	m.Password = dataAO1.Password

	m.SsoURL = dataAO1.SsoURL

	m.Username = dataAO1.Username

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexVcenterConfigPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

		DataCenter string `json:"DataCenter,omitempty"`

		Hostname string `json:"Hostname,omitempty"`

		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Password string `json:"Password,omitempty"`

		SsoURL string `json:"SsoUrl,omitempty"`

		Username string `json:"Username,omitempty"`
	}

	dataAO1.ClusterProfiles = m.ClusterProfiles

	dataAO1.DataCenter = m.DataCenter

	dataAO1.Hostname = m.Hostname

	dataAO1.IsPasswordSet = m.IsPasswordSet

	dataAO1.Organization = m.Organization

	dataAO1.Password = m.Password

	dataAO1.SsoURL = m.SsoURL

	dataAO1.Username = m.Username

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex vcenter config policy
func (m *HyperflexVcenterConfigPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexVcenterConfigPolicy) validateClusterProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterProfiles) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterProfiles); i++ {
		if swag.IsZero(m.ClusterProfiles[i]) { // not required
			continue
		}

		if m.ClusterProfiles[i] != nil {
			if err := m.ClusterProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ClusterProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HyperflexVcenterConfigPolicy) validateOrganization(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *HyperflexVcenterConfigPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexVcenterConfigPolicy) UnmarshalBinary(b []byte) error {
	var res HyperflexVcenterConfigPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
