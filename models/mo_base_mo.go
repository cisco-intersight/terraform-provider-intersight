// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoBaseMo Mo:Base Mo
//
// The base abstract class for all Cisco Intersight managed objects.
//
// swagger:model moBaseMo
type MoBaseMo struct {

	// The Account ID for this managed object.
	//
	// Read Only: true
	AccountMoid string `json:"AccountMoid,omitempty"`

	// The array containing the MO references of the ancestors in the object containment hierarchy.
	//
	// Read Only: true
	Ancestors []*MoBaseMoRef `json:"Ancestors"`

	// The time when this managed object was created.
	//
	// Read Only: true
	// Format: date-time
	CreateTime strfmt.DateTime `json:"CreateTime,omitempty"`

	// The DomainGroup ID for this managed object.
	//
	// Read Only: true
	DomainGroupMoid string `json:"DomainGroupMoid,omitempty"`

	// The time when this managed object was last modified.
	//
	// Read Only: true
	// Format: date-time
	ModTime strfmt.DateTime `json:"ModTime,omitempty"`

	// The unique identifier of this Managed Object instance.
	//
	Moid string `json:"Moid,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// The array of owners which represent effective ownership of this object.
	//
	// Read Only: true
	Owners []string `json:"Owners"`

	// The direct ancestor of this managed object in the containment hierarchy.
	//
	// Read Only: true
	Parent *MoBaseMoRef `json:"Parent,omitempty"`

	// Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.
	// Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.
	//
	// Read Only: true
	SharedScope string `json:"SharedScope,omitempty"`

	// The array of tags, which allow to add key, value meta-data to managed objects.
	//
	Tags []*MoTag `json:"Tags"`

	// The versioning info for this managed object.
	//
	// Read Only: true
	VersionContext *MoVersionContext `json:"VersionContext,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MoBaseMo) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		AccountMoid string `json:"AccountMoid,omitempty"`

		Ancestors []*MoBaseMoRef `json:"Ancestors"`

		CreateTime strfmt.DateTime `json:"CreateTime,omitempty"`

		DomainGroupMoid string `json:"DomainGroupMoid,omitempty"`

		ModTime strfmt.DateTime `json:"ModTime,omitempty"`

		Moid string `json:"Moid,omitempty"`

		ObjectType string `json:"ObjectType,omitempty"`

		Owners []string `json:"Owners"`

		Parent *MoBaseMoRef `json:"Parent,omitempty"`

		SharedScope string `json:"SharedScope,omitempty"`

		Tags []*MoTag `json:"Tags"`

		VersionContext *MoVersionContext `json:"VersionContext,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.AccountMoid = dataAO0.AccountMoid

	m.Ancestors = dataAO0.Ancestors

	m.CreateTime = dataAO0.CreateTime

	m.DomainGroupMoid = dataAO0.DomainGroupMoid

	m.ModTime = dataAO0.ModTime

	m.Moid = dataAO0.Moid

	m.ObjectType = dataAO0.ObjectType

	m.Owners = dataAO0.Owners

	m.Parent = dataAO0.Parent

	m.SharedScope = dataAO0.SharedScope

	m.Tags = dataAO0.Tags

	m.VersionContext = dataAO0.VersionContext

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MoBaseMo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		AccountMoid string `json:"AccountMoid,omitempty"`

		Ancestors []*MoBaseMoRef `json:"Ancestors"`

		CreateTime strfmt.DateTime `json:"CreateTime,omitempty"`

		DomainGroupMoid string `json:"DomainGroupMoid,omitempty"`

		ModTime strfmt.DateTime `json:"ModTime,omitempty"`

		Moid string `json:"Moid,omitempty"`

		ObjectType string `json:"ObjectType,omitempty"`

		Owners []string `json:"Owners"`

		Parent *MoBaseMoRef `json:"Parent,omitempty"`

		SharedScope string `json:"SharedScope,omitempty"`

		Tags []*MoTag `json:"Tags"`

		VersionContext *MoVersionContext `json:"VersionContext,omitempty"`
	}

	dataAO0.AccountMoid = m.AccountMoid

	dataAO0.Ancestors = m.Ancestors

	dataAO0.CreateTime = m.CreateTime

	dataAO0.DomainGroupMoid = m.DomainGroupMoid

	dataAO0.ModTime = m.ModTime

	dataAO0.Moid = m.Moid

	dataAO0.ObjectType = m.ObjectType

	dataAO0.Owners = m.Owners

	dataAO0.Parent = m.Parent

	dataAO0.SharedScope = m.SharedScope

	dataAO0.Tags = m.Tags

	dataAO0.VersionContext = m.VersionContext

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this mo base mo
func (m *MoBaseMo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAncestors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoBaseMo) validateAncestors(formats strfmt.Registry) error {

	if swag.IsZero(m.Ancestors) { // not required
		return nil
	}

	for i := 0; i < len(m.Ancestors); i++ {
		if swag.IsZero(m.Ancestors[i]) { // not required
			continue
		}

		if m.Ancestors[i] != nil {
			if err := m.Ancestors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Ancestors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoBaseMo) validateCreateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreateTime", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoBaseMo) validateModTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ModTime) { // not required
		return nil
	}

	if err := validate.FormatOf("ModTime", "body", "date-time", m.ModTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoBaseMo) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Parent")
			}
			return err
		}
	}

	return nil
}

func (m *MoBaseMo) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoBaseMo) validateVersionContext(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionContext) { // not required
		return nil
	}

	if m.VersionContext != nil {
		if err := m.VersionContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("VersionContext")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoBaseMo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoBaseMo) UnmarshalBinary(b []byte) error {
	var res MoBaseMo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}