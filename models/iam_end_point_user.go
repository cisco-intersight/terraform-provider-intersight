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

// IamEndPointUser User
//
// Endpoint User or Local User.
//
// swagger:model iamEndPointUser
type IamEndPointUser struct {
	MoBaseMo

	// A collection of references to the [iam.EndPointUserRole](mo://iam.EndPointUserRole) Managed Object.
	// When this managed object is deleted, the referenced [iam.EndPointUserRole](mo://iam.EndPointUserRole) MOs unset their reference to this deleted MO.
	EndPointUserRole []*IamEndPointUserRoleRef `json:"EndPointUserRole"`

	// Name of the user created on the endpoint.
	Name string `json:"Name,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamEndPointUser) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		EndPointUserRole []*IamEndPointUserRoleRef `json:"EndPointUserRole"`

		Name string `json:"Name,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EndPointUserRole = dataAO1.EndPointUserRole

	m.Name = dataAO1.Name

	m.Organization = dataAO1.Organization

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamEndPointUser) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		EndPointUserRole []*IamEndPointUserRoleRef `json:"EndPointUserRole"`

		Name string `json:"Name,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
	}

	dataAO1.EndPointUserRole = m.EndPointUserRole

	dataAO1.Name = m.Name

	dataAO1.Organization = m.Organization

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam end point user
func (m *IamEndPointUser) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPointUserRole(formats); err != nil {
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

func (m *IamEndPointUser) validateEndPointUserRole(formats strfmt.Registry) error {

	if swag.IsZero(m.EndPointUserRole) { // not required
		return nil
	}

	for i := 0; i < len(m.EndPointUserRole); i++ {
		if swag.IsZero(m.EndPointUserRole[i]) { // not required
			continue
		}

		if m.EndPointUserRole[i] != nil {
			if err := m.EndPointUserRole[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EndPointUserRole" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamEndPointUser) validateOrganization(formats strfmt.Registry) error {

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
func (m *IamEndPointUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamEndPointUser) UnmarshalBinary(b []byte) error {
	var res IamEndPointUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
