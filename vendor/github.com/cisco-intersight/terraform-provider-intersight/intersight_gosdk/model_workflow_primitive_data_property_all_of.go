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

// WorkflowPrimitiveDataPropertyAllOf Definition of the list of properties defined in 'workflow.PrimitiveDataProperty', excluding properties defined in parent classes.
type WorkflowPrimitiveDataPropertyAllOf struct {
	Constraints       *WorkflowConstraints           `json:"Constraints,omitempty"`
	InventorySelector *[]WorkflowMoReferenceProperty `json:"InventorySelector,omitempty"`
	// Intersight supports secure properties as task input/output. The values of these properties are encrypted and stored in Intersight. This flag marks the property to be secure when it is set to true.
	Secure *bool `json:"Secure,omitempty"`
	// Specify the enum type for primitive data type.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowPrimitiveDataPropertyAllOf WorkflowPrimitiveDataPropertyAllOf

// NewWorkflowPrimitiveDataPropertyAllOf instantiates a new WorkflowPrimitiveDataPropertyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPrimitiveDataPropertyAllOf() *WorkflowPrimitiveDataPropertyAllOf {
	this := WorkflowPrimitiveDataPropertyAllOf{}
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// NewWorkflowPrimitiveDataPropertyAllOfWithDefaults instantiates a new WorkflowPrimitiveDataPropertyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPrimitiveDataPropertyAllOfWithDefaults() *WorkflowPrimitiveDataPropertyAllOf {
	this := WorkflowPrimitiveDataPropertyAllOf{}
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *WorkflowPrimitiveDataPropertyAllOf) GetConstraints() WorkflowConstraints {
	if o == nil || o.Constraints == nil {
		var ret WorkflowConstraints
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataPropertyAllOf) GetConstraintsOk() (*WorkflowConstraints, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *WorkflowPrimitiveDataPropertyAllOf) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given WorkflowConstraints and assigns it to the Constraints field.
func (o *WorkflowPrimitiveDataPropertyAllOf) SetConstraints(v WorkflowConstraints) {
	o.Constraints = &v
}

// GetInventorySelector returns the InventorySelector field value if set, zero value otherwise.
func (o *WorkflowPrimitiveDataPropertyAllOf) GetInventorySelector() []WorkflowMoReferenceProperty {
	if o == nil || o.InventorySelector == nil {
		var ret []WorkflowMoReferenceProperty
		return ret
	}
	return *o.InventorySelector
}

// GetInventorySelectorOk returns a tuple with the InventorySelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataPropertyAllOf) GetInventorySelectorOk() (*[]WorkflowMoReferenceProperty, bool) {
	if o == nil || o.InventorySelector == nil {
		return nil, false
	}
	return o.InventorySelector, true
}

// HasInventorySelector returns a boolean if a field has been set.
func (o *WorkflowPrimitiveDataPropertyAllOf) HasInventorySelector() bool {
	if o != nil && o.InventorySelector != nil {
		return true
	}

	return false
}

// SetInventorySelector gets a reference to the given []WorkflowMoReferenceProperty and assigns it to the InventorySelector field.
func (o *WorkflowPrimitiveDataPropertyAllOf) SetInventorySelector(v []WorkflowMoReferenceProperty) {
	o.InventorySelector = &v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *WorkflowPrimitiveDataPropertyAllOf) GetSecure() bool {
	if o == nil || o.Secure == nil {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataPropertyAllOf) GetSecureOk() (*bool, bool) {
	if o == nil || o.Secure == nil {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *WorkflowPrimitiveDataPropertyAllOf) HasSecure() bool {
	if o != nil && o.Secure != nil {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *WorkflowPrimitiveDataPropertyAllOf) SetSecure(v bool) {
	o.Secure = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowPrimitiveDataPropertyAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataPropertyAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowPrimitiveDataPropertyAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowPrimitiveDataPropertyAllOf) SetType(v string) {
	o.Type = &v
}

func (o WorkflowPrimitiveDataPropertyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Constraints != nil {
		toSerialize["Constraints"] = o.Constraints
	}
	if o.InventorySelector != nil {
		toSerialize["InventorySelector"] = o.InventorySelector
	}
	if o.Secure != nil {
		toSerialize["Secure"] = o.Secure
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPrimitiveDataPropertyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowPrimitiveDataPropertyAllOf := _WorkflowPrimitiveDataPropertyAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowPrimitiveDataPropertyAllOf); err == nil {
		*o = WorkflowPrimitiveDataPropertyAllOf(varWorkflowPrimitiveDataPropertyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Constraints")
		delete(additionalProperties, "InventorySelector")
		delete(additionalProperties, "Secure")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowPrimitiveDataPropertyAllOf struct {
	value *WorkflowPrimitiveDataPropertyAllOf
	isSet bool
}

func (v NullableWorkflowPrimitiveDataPropertyAllOf) Get() *WorkflowPrimitiveDataPropertyAllOf {
	return v.value
}

func (v *NullableWorkflowPrimitiveDataPropertyAllOf) Set(val *WorkflowPrimitiveDataPropertyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPrimitiveDataPropertyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPrimitiveDataPropertyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPrimitiveDataPropertyAllOf(val *WorkflowPrimitiveDataPropertyAllOf) *NullableWorkflowPrimitiveDataPropertyAllOf {
	return &NullableWorkflowPrimitiveDataPropertyAllOf{value: val, isSet: true}
}

func (v NullableWorkflowPrimitiveDataPropertyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPrimitiveDataPropertyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
