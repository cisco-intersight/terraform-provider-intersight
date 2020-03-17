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

// WorkflowWorkflowProperties Workflow:Workflow Properties
//
// Properties for a workflow definition.
//
// swagger:model workflowWorkflowProperties
type WorkflowWorkflowProperties struct {
	MoBaseComplexType

	WorkflowWorkflowPropertiesAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowWorkflowProperties) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 WorkflowWorkflowPropertiesAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.WorkflowWorkflowPropertiesAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowWorkflowProperties) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.WorkflowWorkflowPropertiesAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow workflow properties
func (m *WorkflowWorkflowProperties) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with WorkflowWorkflowPropertiesAO1P1
	if err := m.WorkflowWorkflowPropertiesAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowWorkflowProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowWorkflowProperties) UnmarshalBinary(b []byte) error {
	var res WorkflowWorkflowProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowWorkflowPropertiesAO1P1 workflow workflow properties a o1 p1
//
// swagger:model WorkflowWorkflowPropertiesAO1P1
type WorkflowWorkflowPropertiesAO1P1 struct {

	// When set to false the workflow is owned by the system and used for internal services. Such workflows cannot be directly used by external entities.
	ExternalMeta *bool `json:"ExternalMeta,omitempty"`

	// When true, this workflow can be retried if has not been modified for more than a period of 2 weeks.
	Retryable *bool `json:"Retryable,omitempty"`

	// Supported status of the definition.
	// Enum: [Supported Beta Deprecated]
	SupportStatus *string `json:"SupportStatus,omitempty"`

	// workflow workflow properties a o1 p1
	WorkflowWorkflowPropertiesAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *WorkflowWorkflowPropertiesAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// When set to false the workflow is owned by the system and used for internal services. Such workflows cannot be directly used by external entities.
		ExternalMeta *bool `json:"ExternalMeta,omitempty"`

		// When true, this workflow can be retried if has not been modified for more than a period of 2 weeks.
		Retryable *bool `json:"Retryable,omitempty"`

		// Supported status of the definition.
		// Enum: [Supported Beta Deprecated]
		SupportStatus *string `json:"SupportStatus,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv WorkflowWorkflowPropertiesAO1P1

	rcv.ExternalMeta = stage1.ExternalMeta
	rcv.Retryable = stage1.Retryable
	rcv.SupportStatus = stage1.SupportStatus
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ExternalMeta")
	delete(stage2, "Retryable")
	delete(stage2, "SupportStatus")
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
		m.WorkflowWorkflowPropertiesAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m WorkflowWorkflowPropertiesAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// When set to false the workflow is owned by the system and used for internal services. Such workflows cannot be directly used by external entities.
		ExternalMeta *bool `json:"ExternalMeta,omitempty"`

		// When true, this workflow can be retried if has not been modified for more than a period of 2 weeks.
		Retryable *bool `json:"Retryable,omitempty"`

		// Supported status of the definition.
		// Enum: [Supported Beta Deprecated]
		SupportStatus *string `json:"SupportStatus,omitempty"`
	}

	stage1.ExternalMeta = m.ExternalMeta
	stage1.Retryable = m.Retryable
	stage1.SupportStatus = m.SupportStatus

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.WorkflowWorkflowPropertiesAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.WorkflowWorkflowPropertiesAO1P1)
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

// Validate validates this workflow workflow properties a o1 p1
func (m *WorkflowWorkflowPropertiesAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSupportStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var workflowWorkflowPropertiesAO1P1TypeSupportStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Supported","Beta","Deprecated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowWorkflowPropertiesAO1P1TypeSupportStatusPropEnum = append(workflowWorkflowPropertiesAO1P1TypeSupportStatusPropEnum, v)
	}
}

const (

	// WorkflowWorkflowPropertiesAO1P1SupportStatusSupported captures enum value "Supported"
	WorkflowWorkflowPropertiesAO1P1SupportStatusSupported string = "Supported"

	// WorkflowWorkflowPropertiesAO1P1SupportStatusBeta captures enum value "Beta"
	WorkflowWorkflowPropertiesAO1P1SupportStatusBeta string = "Beta"

	// WorkflowWorkflowPropertiesAO1P1SupportStatusDeprecated captures enum value "Deprecated"
	WorkflowWorkflowPropertiesAO1P1SupportStatusDeprecated string = "Deprecated"
)

// prop value enum
func (m *WorkflowWorkflowPropertiesAO1P1) validateSupportStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowWorkflowPropertiesAO1P1TypeSupportStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowWorkflowPropertiesAO1P1) validateSupportStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateSupportStatusEnum("SupportStatus", "body", *m.SupportStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowWorkflowPropertiesAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowWorkflowPropertiesAO1P1) UnmarshalBinary(b []byte) error {
	var res WorkflowWorkflowPropertiesAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}