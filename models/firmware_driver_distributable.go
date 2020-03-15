// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FirmwareDriverDistributable Firmware:Driver Distributable
//
// A device driver image distributed by Cisco.
//
// swagger:model firmwareDriverDistributable
type FirmwareDriverDistributable struct {
	FirmwareBaseDistributable

	// The catalog where this image is present.
	Catalog *SoftwarerepositoryCatalogRef `json:"Catalog,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FirmwareDriverDistributable) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 FirmwareBaseDistributable
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.FirmwareBaseDistributable = aO0

	// AO1
	var dataAO1 struct {
		Catalog *SoftwarerepositoryCatalogRef `json:"Catalog,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Catalog = dataAO1.Catalog

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FirmwareDriverDistributable) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.FirmwareBaseDistributable)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Catalog *SoftwarerepositoryCatalogRef `json:"Catalog,omitempty"`
	}

	dataAO1.Catalog = m.Catalog

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this firmware driver distributable
func (m *FirmwareDriverDistributable) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with FirmwareBaseDistributable
	if err := m.FirmwareBaseDistributable.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCatalog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareDriverDistributable) validateCatalog(formats strfmt.Registry) error {

	if swag.IsZero(m.Catalog) { // not required
		return nil
	}

	if m.Catalog != nil {
		if err := m.Catalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Catalog")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareDriverDistributable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareDriverDistributable) UnmarshalBinary(b []byte) error {
	var res FirmwareDriverDistributable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
