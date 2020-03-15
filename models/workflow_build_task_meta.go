// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkflowBuildTaskMeta Workflow:Build Task Meta
//
// Contains relationship for tasks within a workflow. It is used to dynamically generate a workflow.
//
// swagger:model workflowBuildTaskMeta
type WorkflowBuildTaskMeta struct {
	MoBaseMo

	// Name for the BuildTaskMeta instance used to created a dynamic workflow.
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// Microservice owner for the task in this workflow.
	// Read Only: true
	Src string `json:"Src,omitempty"`

	// Task list used to build the dynamic workflow.
	// Read Only: true
	TaskList interface{} `json:"TaskList,omitempty"`

	// The type of the task within this workflow.
	// Read Only: true
	TaskType string `json:"TaskType,omitempty"`

	// The type for the dynamic workflow.
	// Read Only: true
	WorkflowType string `json:"WorkflowType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowBuildTaskMeta) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Name string `json:"Name,omitempty"`

		Src string `json:"Src,omitempty"`

		TaskList interface{} `json:"TaskList,omitempty"`

		TaskType string `json:"TaskType,omitempty"`

		WorkflowType string `json:"WorkflowType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Name = dataAO1.Name

	m.Src = dataAO1.Src

	m.TaskList = dataAO1.TaskList

	m.TaskType = dataAO1.TaskType

	m.WorkflowType = dataAO1.WorkflowType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowBuildTaskMeta) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Name string `json:"Name,omitempty"`

		Src string `json:"Src,omitempty"`

		TaskList interface{} `json:"TaskList,omitempty"`

		TaskType string `json:"TaskType,omitempty"`

		WorkflowType string `json:"WorkflowType,omitempty"`
	}

	dataAO1.Name = m.Name

	dataAO1.Src = m.Src

	dataAO1.TaskList = m.TaskList

	dataAO1.TaskType = m.TaskType

	dataAO1.WorkflowType = m.WorkflowType

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow build task meta
func (m *WorkflowBuildTaskMeta) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowBuildTaskMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowBuildTaskMeta) UnmarshalBinary(b []byte) error {
	var res WorkflowBuildTaskMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
