/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// WorkflowDecisionCaseAllOf Definition of the list of properties defined in 'workflow.DecisionCase', excluding properties defined in parent classes.
type WorkflowDecisionCaseAllOf struct {
	// Description of this decision case.
	Description *string `json:"Description,omitempty"`
	// The name of the next task (Task names unique within workflow) to run.  In a graph model, denotes an edge to another Task Node.
	NextTask *string `json:"NextTask,omitempty"`
	// Value for the decision case.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowDecisionCaseAllOf WorkflowDecisionCaseAllOf

// NewWorkflowDecisionCaseAllOf instantiates a new WorkflowDecisionCaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDecisionCaseAllOf() *WorkflowDecisionCaseAllOf {
	this := WorkflowDecisionCaseAllOf{}
	return &this
}

// NewWorkflowDecisionCaseAllOfWithDefaults instantiates a new WorkflowDecisionCaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDecisionCaseAllOfWithDefaults() *WorkflowDecisionCaseAllOf {
	this := WorkflowDecisionCaseAllOf{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowDecisionCaseAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionCaseAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowDecisionCaseAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowDecisionCaseAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetNextTask returns the NextTask field value if set, zero value otherwise.
func (o *WorkflowDecisionCaseAllOf) GetNextTask() string {
	if o == nil || o.NextTask == nil {
		var ret string
		return ret
	}
	return *o.NextTask
}

// GetNextTaskOk returns a tuple with the NextTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionCaseAllOf) GetNextTaskOk() (*string, bool) {
	if o == nil || o.NextTask == nil {
		return nil, false
	}
	return o.NextTask, true
}

// HasNextTask returns a boolean if a field has been set.
func (o *WorkflowDecisionCaseAllOf) HasNextTask() bool {
	if o != nil && o.NextTask != nil {
		return true
	}

	return false
}

// SetNextTask gets a reference to the given string and assigns it to the NextTask field.
func (o *WorkflowDecisionCaseAllOf) SetNextTask(v string) {
	o.NextTask = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WorkflowDecisionCaseAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionCaseAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkflowDecisionCaseAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WorkflowDecisionCaseAllOf) SetValue(v string) {
	o.Value = &v
}

func (o WorkflowDecisionCaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.NextTask != nil {
		toSerialize["NextTask"] = o.NextTask
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowDecisionCaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowDecisionCaseAllOf := _WorkflowDecisionCaseAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowDecisionCaseAllOf); err == nil {
		*o = WorkflowDecisionCaseAllOf(varWorkflowDecisionCaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Description")
		delete(additionalProperties, "NextTask")
		delete(additionalProperties, "Value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowDecisionCaseAllOf struct {
	value *WorkflowDecisionCaseAllOf
	isSet bool
}

func (v NullableWorkflowDecisionCaseAllOf) Get() *WorkflowDecisionCaseAllOf {
	return v.value
}

func (v *NullableWorkflowDecisionCaseAllOf) Set(val *WorkflowDecisionCaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDecisionCaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDecisionCaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDecisionCaseAllOf(val *WorkflowDecisionCaseAllOf) *NullableWorkflowDecisionCaseAllOf {
	return &NullableWorkflowDecisionCaseAllOf{value: val, isSet: true}
}

func (v NullableWorkflowDecisionCaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDecisionCaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
