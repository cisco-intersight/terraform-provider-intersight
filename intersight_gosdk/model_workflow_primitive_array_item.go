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
	"reflect"
	"strings"
)

// WorkflowPrimitiveArrayItem PrimitiveArrayItem is used to create an array of primitive datatype. In order to create an array of strings, integers or float use PrimitiveArrayItem object type within the array.
type WorkflowPrimitiveArrayItem struct {
	WorkflowArrayItem
	Properties           *WorkflowPrimitiveDataProperty `json:"Properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowPrimitiveArrayItem WorkflowPrimitiveArrayItem

// NewWorkflowPrimitiveArrayItem instantiates a new WorkflowPrimitiveArrayItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPrimitiveArrayItem() *WorkflowPrimitiveArrayItem {
	this := WorkflowPrimitiveArrayItem{}
	return &this
}

// NewWorkflowPrimitiveArrayItemWithDefaults instantiates a new WorkflowPrimitiveArrayItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPrimitiveArrayItemWithDefaults() *WorkflowPrimitiveArrayItem {
	this := WorkflowPrimitiveArrayItem{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowPrimitiveArrayItem) GetProperties() WorkflowPrimitiveDataProperty {
	if o == nil || o.Properties == nil {
		var ret WorkflowPrimitiveDataProperty
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveArrayItem) GetPropertiesOk() (*WorkflowPrimitiveDataProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowPrimitiveArrayItem) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given WorkflowPrimitiveDataProperty and assigns it to the Properties field.
func (o *WorkflowPrimitiveArrayItem) SetProperties(v WorkflowPrimitiveDataProperty) {
	o.Properties = &v
}

func (o WorkflowPrimitiveArrayItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowArrayItem, errWorkflowArrayItem := json.Marshal(o.WorkflowArrayItem)
	if errWorkflowArrayItem != nil {
		return []byte{}, errWorkflowArrayItem
	}
	errWorkflowArrayItem = json.Unmarshal([]byte(serializedWorkflowArrayItem), &toSerialize)
	if errWorkflowArrayItem != nil {
		return []byte{}, errWorkflowArrayItem
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPrimitiveArrayItem) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowPrimitiveArrayItemWithoutEmbeddedStruct struct {
		Properties *WorkflowPrimitiveDataProperty `json:"Properties,omitempty"`
	}

	varWorkflowPrimitiveArrayItemWithoutEmbeddedStruct := WorkflowPrimitiveArrayItemWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowPrimitiveArrayItemWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowPrimitiveArrayItem := _WorkflowPrimitiveArrayItem{}
		varWorkflowPrimitiveArrayItem.Properties = varWorkflowPrimitiveArrayItemWithoutEmbeddedStruct.Properties
		*o = WorkflowPrimitiveArrayItem(varWorkflowPrimitiveArrayItem)
	} else {
		return err
	}

	varWorkflowPrimitiveArrayItem := _WorkflowPrimitiveArrayItem{}

	err = json.Unmarshal(bytes, &varWorkflowPrimitiveArrayItem)
	if err == nil {
		o.WorkflowArrayItem = varWorkflowPrimitiveArrayItem.WorkflowArrayItem
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Properties")

		// remove fields from embedded structs
		reflectWorkflowArrayItem := reflect.ValueOf(o.WorkflowArrayItem)
		for i := 0; i < reflectWorkflowArrayItem.Type().NumField(); i++ {
			t := reflectWorkflowArrayItem.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowPrimitiveArrayItem struct {
	value *WorkflowPrimitiveArrayItem
	isSet bool
}

func (v NullableWorkflowPrimitiveArrayItem) Get() *WorkflowPrimitiveArrayItem {
	return v.value
}

func (v *NullableWorkflowPrimitiveArrayItem) Set(val *WorkflowPrimitiveArrayItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPrimitiveArrayItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPrimitiveArrayItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPrimitiveArrayItem(val *WorkflowPrimitiveArrayItem) *NullableWorkflowPrimitiveArrayItem {
	return &NullableWorkflowPrimitiveArrayItem{value: val, isSet: true}
}

func (v NullableWorkflowPrimitiveArrayItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPrimitiveArrayItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
