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
)

// I18nMessage I18n:Message
//
// An i18n message that can be translated in multiple languages to support internationalization.
// An i18n message includes a unique message identifier, a text format string, a list of message parameters that can be used
// to substitute parameters, and a translated string based on a user's locale.
//
// swagger:model i18nMessage
type I18nMessage struct {
	MoBaseComplexType

	I18nMessageAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *I18nMessage) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 I18nMessageAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.I18nMessageAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m I18nMessage) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.I18nMessageAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this i18n message
func (m *I18nMessage) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with I18nMessageAO1P1
	if err := m.I18nMessageAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *I18nMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *I18nMessage) UnmarshalBinary(b []byte) error {
	var res I18nMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// I18nMessageAO1P1 i18n message a o1 p1
//
// swagger:model I18nMessageAO1P1
type I18nMessageAO1P1 struct {

	// The default (en-US) localized message. Default localized message will be stored and directly retrieved when
	// the user's locale setting is en-US.
	// Read Only: true
	Message string `json:"Message,omitempty"`

	// The unique message identitifer used to lookup text templates in a multi-language message catalog.
	// Read Only: true
	MessageID string `json:"MessageId,omitempty"`

	// The list of message parameters.
	MessageParams []*I18nMessageParam `json:"MessageParams"`

	// i18n message a o1 p1
	I18nMessageAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *I18nMessageAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The default (en-US) localized message. Default localized message will be stored and directly retrieved when
		// the user's locale setting is en-US.
		// Read Only: true
		Message string `json:"Message,omitempty"`

		// The unique message identitifer used to lookup text templates in a multi-language message catalog.
		// Read Only: true
		MessageID string `json:"MessageId,omitempty"`

		// The list of message parameters.
		MessageParams []*I18nMessageParam `json:"MessageParams"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv I18nMessageAO1P1

	rcv.Message = stage1.Message
	rcv.MessageID = stage1.MessageID
	rcv.MessageParams = stage1.MessageParams
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Message")
	delete(stage2, "MessageId")
	delete(stage2, "MessageParams")
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
		m.I18nMessageAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m I18nMessageAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The default (en-US) localized message. Default localized message will be stored and directly retrieved when
		// the user's locale setting is en-US.
		// Read Only: true
		Message string `json:"Message,omitempty"`

		// The unique message identitifer used to lookup text templates in a multi-language message catalog.
		// Read Only: true
		MessageID string `json:"MessageId,omitempty"`

		// The list of message parameters.
		MessageParams []*I18nMessageParam `json:"MessageParams"`
	}

	stage1.Message = m.Message
	stage1.MessageID = m.MessageID
	stage1.MessageParams = m.MessageParams

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.I18nMessageAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.I18nMessageAO1P1)
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

// Validate validates this i18n message a o1 p1
func (m *I18nMessageAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessageParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *I18nMessageAO1P1) validateMessageParams(formats strfmt.Registry) error {

	if swag.IsZero(m.MessageParams) { // not required
		return nil
	}

	for i := 0; i < len(m.MessageParams); i++ {
		if swag.IsZero(m.MessageParams[i]) { // not required
			continue
		}

		if m.MessageParams[i] != nil {
			if err := m.MessageParams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("MessageParams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *I18nMessageAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *I18nMessageAO1P1) UnmarshalBinary(b []byte) error {
	var res I18nMessageAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
