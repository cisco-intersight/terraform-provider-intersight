// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StoragePureDiskUtilization Storage:Pure Disk Utilization
//
// Storage space utilization of Pure FlashArray drive entity.
//
// swagger:model storagePureDiskUtilization
type StoragePureDiskUtilization struct {
	StorageBaseCapacity

	StoragePureDiskUtilizationAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StoragePureDiskUtilization) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageBaseCapacity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageBaseCapacity = aO0

	// AO1
	var aO1 StoragePureDiskUtilizationAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.StoragePureDiskUtilizationAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StoragePureDiskUtilization) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.StorageBaseCapacity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.StoragePureDiskUtilizationAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage pure disk utilization
func (m *StoragePureDiskUtilization) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageBaseCapacity
	if err := m.StorageBaseCapacity.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with StoragePureDiskUtilizationAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StoragePureDiskUtilization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePureDiskUtilization) UnmarshalBinary(b []byte) error {
	var res StoragePureDiskUtilization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StoragePureDiskUtilizationAllOf1 storage pure disk utilization all of1
//
// swagger:model StoragePureDiskUtilizationAllOf1
type StoragePureDiskUtilizationAllOf1 interface{}
