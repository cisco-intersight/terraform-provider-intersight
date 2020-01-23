// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PkixSubjectAlternateName Pkix:Subject Alternate Name
//
// The additional hostnames to be protected by a given X.509 certificate.
//
// swagger:model pkixSubjectAlternateName
type PkixSubjectAlternateName struct {
	MoBaseComplexType

	PkixSubjectAlternateNameAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PkixSubjectAlternateName) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 PkixSubjectAlternateNameAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PkixSubjectAlternateNameAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PkixSubjectAlternateName) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PkixSubjectAlternateNameAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pkix subject alternate name
func (m *PkixSubjectAlternateName) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PkixSubjectAlternateNameAO1P1
	if err := m.PkixSubjectAlternateNameAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PkixSubjectAlternateName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkixSubjectAlternateName) UnmarshalBinary(b []byte) error {
	var res PkixSubjectAlternateName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PkixSubjectAlternateNameAO1P1 pkix subject alternate name a o1 p1
// swagger:model PkixSubjectAlternateNameAO1P1
type PkixSubjectAlternateNameAO1P1 struct {

	// Alternate DNS names for the host.
	//
	// Read Only: true
	DNSName []string `json:"DnsName"`

	// Alternate email addresses for the host.
	//
	// Read Only: true
	EmailAddress []string `json:"EmailAddress"`

	// Alternate IP addresses for the host.
	//
	// Read Only: true
	IPAddress []string `json:"IpAddress"`

	// Alternate URIs for the host.
	//
	// Read Only: true
	URI []string `json:"Uri"`

	// pkix subject alternate name a o1 p1
	PkixSubjectAlternateNameAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *PkixSubjectAlternateNameAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Alternate DNS names for the host.
		//
		// Read Only: true
		DNSName []string `json:"DnsName"`

		// Alternate email addresses for the host.
		//
		// Read Only: true
		EmailAddress []string `json:"EmailAddress"`

		// Alternate IP addresses for the host.
		//
		// Read Only: true
		IPAddress []string `json:"IpAddress"`

		// Alternate URIs for the host.
		//
		// Read Only: true
		URI []string `json:"Uri"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv PkixSubjectAlternateNameAO1P1

	rcv.DNSName = stage1.DNSName

	rcv.EmailAddress = stage1.EmailAddress

	rcv.IPAddress = stage1.IPAddress

	rcv.URI = stage1.URI

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "DnsName")

	delete(stage2, "EmailAddress")

	delete(stage2, "IpAddress")

	delete(stage2, "Uri")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.PkixSubjectAlternateNameAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m PkixSubjectAlternateNameAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Alternate DNS names for the host.
		//
		// Read Only: true
		DNSName []string `json:"DnsName"`

		// Alternate email addresses for the host.
		//
		// Read Only: true
		EmailAddress []string `json:"EmailAddress"`

		// Alternate IP addresses for the host.
		//
		// Read Only: true
		IPAddress []string `json:"IpAddress"`

		// Alternate URIs for the host.
		//
		// Read Only: true
		URI []string `json:"Uri"`
	}

	stage1.DNSName = m.DNSName

	stage1.EmailAddress = m.EmailAddress

	stage1.IPAddress = m.IPAddress

	stage1.URI = m.URI

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.PkixSubjectAlternateNameAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.PkixSubjectAlternateNameAO1P1)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this pkix subject alternate name a o1 p1
func (m *PkixSubjectAlternateNameAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PkixSubjectAlternateNameAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkixSubjectAlternateNameAO1P1) UnmarshalBinary(b []byte) error {
	var res PkixSubjectAlternateNameAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}