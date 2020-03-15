// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkflowDecisionTask Workflow:Decision Task
//
// A DecisionTask is a control task that executes a sequence of WorkflowTasks based off decision provided and evaluated by this task.
//
// swagger:model workflowDecisionTask
type WorkflowDecisionTask struct {
	WorkflowControlTask

	// The condition to evaluate for this decision task. The condition can be a workflow or task variable or an expression based on the input parameters. Example value for condition if its Workflow/task variable is -  "${task1.output.var1} or ${workflow.input.var2}" which evaluates to a value matching any of the decision case values. Example value for condition if its an expression is - "if ( $.element ! = null && $.element > 0 ) 'true'; else 'false'; " which evaluates to 'true' or 'false' and will match one of the decision case values. Here "element" is a variable in decisiontask's inputParameters JSON formatted map. You can also use javascript like functions indexOf, toUpperCase in the expression which will be evaluated by the expression evaluator.
	Condition string `json:"Condition,omitempty"`

	// A list of potential decision task flows based off a condition.
	DecisionCases []*WorkflowDecisionCase `json:"DecisionCases"`

	// The default next Task to execute if the decision cannot be evaluated to any of the DecisionCases.
	DefaultTask string `json:"DefaultTask,omitempty"`

	// JSON formatted map that defines the input given to the decision task. The inputs are used as variables in the condition property of decision task. The input variables can be static values like "hello" , "24", "true" OR previous task outputs/workflow inputs like "${task2.output.var1}}". The input variables are referrenced as $.inputVariableName in the condition property.
	InputParameters interface{} `json:"InputParameters,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowDecisionTask) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowControlTask
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowControlTask = aO0

	// AO1
	var dataAO1 struct {
		Condition string `json:"Condition,omitempty"`

		DecisionCases []*WorkflowDecisionCase `json:"DecisionCases"`

		DefaultTask string `json:"DefaultTask,omitempty"`

		InputParameters interface{} `json:"InputParameters,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Condition = dataAO1.Condition

	m.DecisionCases = dataAO1.DecisionCases

	m.DefaultTask = dataAO1.DefaultTask

	m.InputParameters = dataAO1.InputParameters

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowDecisionTask) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WorkflowControlTask)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Condition string `json:"Condition,omitempty"`

		DecisionCases []*WorkflowDecisionCase `json:"DecisionCases"`

		DefaultTask string `json:"DefaultTask,omitempty"`

		InputParameters interface{} `json:"InputParameters,omitempty"`
	}

	dataAO1.Condition = m.Condition

	dataAO1.DecisionCases = m.DecisionCases

	dataAO1.DefaultTask = m.DefaultTask

	dataAO1.InputParameters = m.InputParameters

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow decision task
func (m *WorkflowDecisionTask) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowControlTask
	if err := m.WorkflowControlTask.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDecisionCases(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowDecisionTask) validateDecisionCases(formats strfmt.Registry) error {

	if swag.IsZero(m.DecisionCases) { // not required
		return nil
	}

	for i := 0; i < len(m.DecisionCases); i++ {
		if swag.IsZero(m.DecisionCases[i]) { // not required
			continue
		}

		if m.DecisionCases[i] != nil {
			if err := m.DecisionCases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DecisionCases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowDecisionTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowDecisionTask) UnmarshalBinary(b []byte) error {
	var res WorkflowDecisionTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
