// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowMessage Workflow:Message
//
// Intermediate Task or Workflow message with severity.
//
// swagger:model workflowMessage
type WorkflowMessage struct {
	MoBaseComplexType

	WorkflowMessageAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowMessage) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 WorkflowMessageAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.WorkflowMessageAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowMessage) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.WorkflowMessageAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow message
func (m *WorkflowMessage) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with WorkflowMessageAO1P1
	if err := m.WorkflowMessageAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowMessage) UnmarshalBinary(b []byte) error {
	var res WorkflowMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowMessageAO1P1 workflow message a o1 p1
// swagger:model WorkflowMessageAO1P1
type WorkflowMessageAO1P1 struct {

	// An i18n message that can be translated in multiple languages to support internationalization.
	//
	Message string `json:"Message,omitempty"`

	// The severity of the Task or Workflow message warning/error/info etc.
	//
	// Enum: [Info Warning Debug Error]
	Severity *string `json:"Severity,omitempty"`

	// workflow message a o1 p1
	WorkflowMessageAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *WorkflowMessageAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// An i18n message that can be translated in multiple languages to support internationalization.
		//
		Message string `json:"Message,omitempty"`

		// The severity of the Task or Workflow message warning/error/info etc.
		//
		// Enum: [Info Warning Debug Error]
		Severity *string `json:"Severity,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv WorkflowMessageAO1P1

	rcv.Message = stage1.Message

	rcv.Severity = stage1.Severity

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Message")

	delete(stage2, "Severity")

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
		m.WorkflowMessageAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m WorkflowMessageAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// An i18n message that can be translated in multiple languages to support internationalization.
		//
		Message string `json:"Message,omitempty"`

		// The severity of the Task or Workflow message warning/error/info etc.
		//
		// Enum: [Info Warning Debug Error]
		Severity *string `json:"Severity,omitempty"`
	}

	stage1.Message = m.Message

	stage1.Severity = m.Severity

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.WorkflowMessageAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.WorkflowMessageAO1P1)
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

// Validate validates this workflow message a o1 p1
func (m *WorkflowMessageAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var workflowMessageAO1P1TypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Info","Warning","Debug","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowMessageAO1P1TypeSeverityPropEnum = append(workflowMessageAO1P1TypeSeverityPropEnum, v)
	}
}

const (

	// WorkflowMessageAO1P1SeverityInfo captures enum value "Info"
	WorkflowMessageAO1P1SeverityInfo string = "Info"

	// WorkflowMessageAO1P1SeverityWarning captures enum value "Warning"
	WorkflowMessageAO1P1SeverityWarning string = "Warning"

	// WorkflowMessageAO1P1SeverityDebug captures enum value "Debug"
	WorkflowMessageAO1P1SeverityDebug string = "Debug"

	// WorkflowMessageAO1P1SeverityError captures enum value "Error"
	WorkflowMessageAO1P1SeverityError string = "Error"
)

// prop value enum
func (m *WorkflowMessageAO1P1) validateSeverityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowMessageAO1P1TypeSeverityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowMessageAO1P1) validateSeverity(formats strfmt.Registry) error {

	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("Severity", "body", *m.Severity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowMessageAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowMessageAO1P1) UnmarshalBinary(b []byte) error {
	var res WorkflowMessageAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}