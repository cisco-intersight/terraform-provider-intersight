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

// FabricConfigResultEntry This provides detailed information for the deploy and validation profile configuration results.
type FabricConfigResultEntry struct {
	PolicyAbstractConfigResultEntry
	ConfigResult         *FabricConfigResultRelationship `json:"ConfigResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricConfigResultEntry FabricConfigResultEntry

// NewFabricConfigResultEntry instantiates a new FabricConfigResultEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricConfigResultEntry() *FabricConfigResultEntry {
	this := FabricConfigResultEntry{}
	return &this
}

// NewFabricConfigResultEntryWithDefaults instantiates a new FabricConfigResultEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricConfigResultEntryWithDefaults() *FabricConfigResultEntry {
	this := FabricConfigResultEntry{}
	return &this
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *FabricConfigResultEntry) GetConfigResult() FabricConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret FabricConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricConfigResultEntry) GetConfigResultOk() (*FabricConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *FabricConfigResultEntry) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given FabricConfigResultRelationship and assigns it to the ConfigResult field.
func (o *FabricConfigResultEntry) SetConfigResult(v FabricConfigResultRelationship) {
	o.ConfigResult = &v
}

func (o FabricConfigResultEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigResultEntry, errPolicyAbstractConfigResultEntry := json.Marshal(o.PolicyAbstractConfigResultEntry)
	if errPolicyAbstractConfigResultEntry != nil {
		return []byte{}, errPolicyAbstractConfigResultEntry
	}
	errPolicyAbstractConfigResultEntry = json.Unmarshal([]byte(serializedPolicyAbstractConfigResultEntry), &toSerialize)
	if errPolicyAbstractConfigResultEntry != nil {
		return []byte{}, errPolicyAbstractConfigResultEntry
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricConfigResultEntry) UnmarshalJSON(bytes []byte) (err error) {
	type FabricConfigResultEntryWithoutEmbeddedStruct struct {
		ConfigResult *FabricConfigResultRelationship `json:"ConfigResult,omitempty"`
	}

	varFabricConfigResultEntryWithoutEmbeddedStruct := FabricConfigResultEntryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricConfigResultEntryWithoutEmbeddedStruct)
	if err == nil {
		varFabricConfigResultEntry := _FabricConfigResultEntry{}
		varFabricConfigResultEntry.ConfigResult = varFabricConfigResultEntryWithoutEmbeddedStruct.ConfigResult
		*o = FabricConfigResultEntry(varFabricConfigResultEntry)
	} else {
		return err
	}

	varFabricConfigResultEntry := _FabricConfigResultEntry{}

	err = json.Unmarshal(bytes, &varFabricConfigResultEntry)
	if err == nil {
		o.PolicyAbstractConfigResultEntry = varFabricConfigResultEntry.PolicyAbstractConfigResultEntry
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ConfigResult")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigResultEntry := reflect.ValueOf(o.PolicyAbstractConfigResultEntry)
		for i := 0; i < reflectPolicyAbstractConfigResultEntry.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigResultEntry.Type().Field(i)

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

type NullableFabricConfigResultEntry struct {
	value *FabricConfigResultEntry
	isSet bool
}

func (v NullableFabricConfigResultEntry) Get() *FabricConfigResultEntry {
	return v.value
}

func (v *NullableFabricConfigResultEntry) Set(val *FabricConfigResultEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricConfigResultEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricConfigResultEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricConfigResultEntry(val *FabricConfigResultEntry) *NullableFabricConfigResultEntry {
	return &NullableFabricConfigResultEntry{value: val, isSet: true}
}

func (v NullableFabricConfigResultEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricConfigResultEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
