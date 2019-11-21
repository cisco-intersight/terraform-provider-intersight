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

// ResourceGroup Resource:Group
//
// A group of REST resources, such as a group of compute.Blade MOs. A ResourceGroup can contain static members which are specified as a set of object references, or it can contain dynamic members, which are specified through OData query filters. A Resource can be part of multiple resource groups.
//
// swagger:model resourceGroup
type ResourceGroup struct {
	MoBaseMo

	// The name of this resource group.
	//
	Name string `json:"Name,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// The list of all resources contained in this resource group.
	//
	// Read Only: true
	Resources []*MoBaseMoRef `json:"Resources"`

	// A list of ODATA filters to select resources. The group selectors may be include URLs of individual resources, or OData query with filters that match multiple queries. The URLs must be relative (i.e. do not include the host) and start with "/api/v1/".
	//
	//
	Selectors []string `json:"Selectors"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ResourceGroup) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Name string `json:"Name,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Resources []*MoBaseMoRef `json:"Resources"`

		Selectors []string `json:"Selectors"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Name = dataAO1.Name

	m.Organization = dataAO1.Organization

	m.Resources = dataAO1.Resources

	m.Selectors = dataAO1.Selectors

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ResourceGroup) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Name string `json:"Name,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Resources []*MoBaseMoRef `json:"Resources"`

		Selectors []string `json:"Selectors"`
	}

	dataAO1.Name = m.Name

	dataAO1.Organization = m.Organization

	dataAO1.Resources = m.Resources

	dataAO1.Selectors = m.Selectors

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this resource group
func (m *ResourceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceGroup) validateOrganization(formats strfmt.Registry) error {

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

func (m *ResourceGroup) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceGroup) UnmarshalBinary(b []byte) error {
	var res ResourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}