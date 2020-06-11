// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StoragePureHostUtilization Storage:Pure Host Utilization
//
// Storage space utilization of Pure FlashArray host entity.
//
// swagger:model storagePureHostUtilization
type StoragePureHostUtilization struct {
	StorageStorageUtilization

	StoragePureHostUtilizationAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StoragePureHostUtilization) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageStorageUtilization
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageStorageUtilization = aO0

	// AO1
	var aO1 StoragePureHostUtilizationAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.StoragePureHostUtilizationAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StoragePureHostUtilization) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.StorageStorageUtilization)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.StoragePureHostUtilizationAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage pure host utilization
func (m *StoragePureHostUtilization) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageStorageUtilization
	if err := m.StorageStorageUtilization.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with StoragePureHostUtilizationAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StoragePureHostUtilization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePureHostUtilization) UnmarshalBinary(b []byte) error {
	var res StoragePureHostUtilization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StoragePureHostUtilizationAllOf1 storage pure host utilization all of1
//
// swagger:model StoragePureHostUtilizationAllOf1
type StoragePureHostUtilizationAllOf1 interface{}