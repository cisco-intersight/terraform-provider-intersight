// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AssetSudiInfo Asset:Sudi Info
//
// The SUDI is an X.509v3 certificate, which maintains the product identifier and serial number. The identity is implemented at manufacturing and chained to a publicly identifiable root certificate authority. It can be used as an unchangeable identity for configuration, security, auditing, and management. This strucure contains the SUDI information read from the device's Trust Anchor Module (TAM).
//
// swagger:model assetSudiInfo
type AssetSudiInfo struct {
	MoBaseComplexType

	AssetSudiInfoAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AssetSudiInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 AssetSudiInfoAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.AssetSudiInfoAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AssetSudiInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.AssetSudiInfoAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this asset sudi info
func (m *AssetSudiInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with AssetSudiInfoAO1P1
	if err := m.AssetSudiInfoAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AssetSudiInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetSudiInfo) UnmarshalBinary(b []byte) error {
	var res AssetSudiInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AssetSudiInfoAO1P1 asset sudi info a o1 p1
//
// swagger:model AssetSudiInfoAO1P1
type AssetSudiInfoAO1P1 struct {

	// The device model (PID) extracted from the X.509 SUDI Leaf Certificate.
	Pid string `json:"Pid,omitempty"`

	// The device SerialNumber extracted from the X.509 SUDI Leaf Certiicate.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The signature is obtained by taking the base64 encoding of the Serial Number + PID + Status, taking the SHA256 hash and then signing with the SUDI X.509 Leaf Certifiate.
	Signature string `json:"Signature,omitempty"`

	// The validation status of the device.
	// Enum: [DeviceStatusUnknown Verified CertificateValidationFailed UnsupportedFirmware UnsupportedHardware DeviceNotResponding]
	Status *string `json:"Status,omitempty"`

	// The X.509 SUDI Leaf Certificate from the Trust Anchor Module. The certificate is serialized in PEM format (Base64 encoded DER certificate).
	SudiCertificate *X509Certificate `json:"SudiCertificate,omitempty"`

	// asset sudi info a o1 p1
	AssetSudiInfoAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *AssetSudiInfoAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The device model (PID) extracted from the X.509 SUDI Leaf Certificate.
		Pid string `json:"Pid,omitempty"`

		// The device SerialNumber extracted from the X.509 SUDI Leaf Certiicate.
		SerialNumber string `json:"SerialNumber,omitempty"`

		// The signature is obtained by taking the base64 encoding of the Serial Number + PID + Status, taking the SHA256 hash and then signing with the SUDI X.509 Leaf Certifiate.
		Signature string `json:"Signature,omitempty"`

		// The validation status of the device.
		// Enum: [DeviceStatusUnknown Verified CertificateValidationFailed UnsupportedFirmware UnsupportedHardware DeviceNotResponding]
		Status *string `json:"Status,omitempty"`

		// The X.509 SUDI Leaf Certificate from the Trust Anchor Module. The certificate is serialized in PEM format (Base64 encoded DER certificate).
		SudiCertificate *X509Certificate `json:"SudiCertificate,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv AssetSudiInfoAO1P1

	rcv.Pid = stage1.Pid
	rcv.SerialNumber = stage1.SerialNumber
	rcv.Signature = stage1.Signature
	rcv.Status = stage1.Status
	rcv.SudiCertificate = stage1.SudiCertificate
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Pid")
	delete(stage2, "SerialNumber")
	delete(stage2, "Signature")
	delete(stage2, "Status")
	delete(stage2, "SudiCertificate")
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
		m.AssetSudiInfoAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m AssetSudiInfoAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The device model (PID) extracted from the X.509 SUDI Leaf Certificate.
		Pid string `json:"Pid,omitempty"`

		// The device SerialNumber extracted from the X.509 SUDI Leaf Certiicate.
		SerialNumber string `json:"SerialNumber,omitempty"`

		// The signature is obtained by taking the base64 encoding of the Serial Number + PID + Status, taking the SHA256 hash and then signing with the SUDI X.509 Leaf Certifiate.
		Signature string `json:"Signature,omitempty"`

		// The validation status of the device.
		// Enum: [DeviceStatusUnknown Verified CertificateValidationFailed UnsupportedFirmware UnsupportedHardware DeviceNotResponding]
		Status *string `json:"Status,omitempty"`

		// The X.509 SUDI Leaf Certificate from the Trust Anchor Module. The certificate is serialized in PEM format (Base64 encoded DER certificate).
		SudiCertificate *X509Certificate `json:"SudiCertificate,omitempty"`
	}

	stage1.Pid = m.Pid
	stage1.SerialNumber = m.SerialNumber
	stage1.Signature = m.Signature
	stage1.Status = m.Status
	stage1.SudiCertificate = m.SudiCertificate

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.AssetSudiInfoAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.AssetSudiInfoAO1P1)
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

// Validate validates this asset sudi info a o1 p1
func (m *AssetSudiInfoAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSudiCertificate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var assetSudiInfoAO1P1TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DeviceStatusUnknown","Verified","CertificateValidationFailed","UnsupportedFirmware","UnsupportedHardware","DeviceNotResponding"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetSudiInfoAO1P1TypeStatusPropEnum = append(assetSudiInfoAO1P1TypeStatusPropEnum, v)
	}
}

const (

	// AssetSudiInfoAO1P1StatusDeviceStatusUnknown captures enum value "DeviceStatusUnknown"
	AssetSudiInfoAO1P1StatusDeviceStatusUnknown string = "DeviceStatusUnknown"

	// AssetSudiInfoAO1P1StatusVerified captures enum value "Verified"
	AssetSudiInfoAO1P1StatusVerified string = "Verified"

	// AssetSudiInfoAO1P1StatusCertificateValidationFailed captures enum value "CertificateValidationFailed"
	AssetSudiInfoAO1P1StatusCertificateValidationFailed string = "CertificateValidationFailed"

	// AssetSudiInfoAO1P1StatusUnsupportedFirmware captures enum value "UnsupportedFirmware"
	AssetSudiInfoAO1P1StatusUnsupportedFirmware string = "UnsupportedFirmware"

	// AssetSudiInfoAO1P1StatusUnsupportedHardware captures enum value "UnsupportedHardware"
	AssetSudiInfoAO1P1StatusUnsupportedHardware string = "UnsupportedHardware"

	// AssetSudiInfoAO1P1StatusDeviceNotResponding captures enum value "DeviceNotResponding"
	AssetSudiInfoAO1P1StatusDeviceNotResponding string = "DeviceNotResponding"
)

// prop value enum
func (m *AssetSudiInfoAO1P1) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, assetSudiInfoAO1P1TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AssetSudiInfoAO1P1) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *AssetSudiInfoAO1P1) validateSudiCertificate(formats strfmt.Registry) error {

	if swag.IsZero(m.SudiCertificate) { // not required
		return nil
	}

	if m.SudiCertificate != nil {
		if err := m.SudiCertificate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SudiCertificate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetSudiInfoAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetSudiInfoAO1P1) UnmarshalBinary(b []byte) error {
	var res AssetSudiInfoAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
