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

// StorageSpanGroup This models a single span group of disks in a RAID group.
type StorageSpanGroup struct {
	MoBaseComplexType
	Disks                *[]StorageLocalDisk `json:"Disks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageSpanGroup StorageSpanGroup

// NewStorageSpanGroup instantiates a new StorageSpanGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSpanGroup() *StorageSpanGroup {
	this := StorageSpanGroup{}
	return &this
}

// NewStorageSpanGroupWithDefaults instantiates a new StorageSpanGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSpanGroupWithDefaults() *StorageSpanGroup {
	this := StorageSpanGroup{}
	return &this
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *StorageSpanGroup) GetDisks() []StorageLocalDisk {
	if o == nil || o.Disks == nil {
		var ret []StorageLocalDisk
		return ret
	}
	return *o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpanGroup) GetDisksOk() (*[]StorageLocalDisk, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *StorageSpanGroup) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []StorageLocalDisk and assigns it to the Disks field.
func (o *StorageSpanGroup) SetDisks(v []StorageLocalDisk) {
	o.Disks = &v
}

func (o StorageSpanGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Disks != nil {
		toSerialize["Disks"] = o.Disks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageSpanGroup) UnmarshalJSON(bytes []byte) (err error) {
	type StorageSpanGroupWithoutEmbeddedStruct struct {
		Disks *[]StorageLocalDisk `json:"Disks,omitempty"`
	}

	varStorageSpanGroupWithoutEmbeddedStruct := StorageSpanGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageSpanGroupWithoutEmbeddedStruct)
	if err == nil {
		varStorageSpanGroup := _StorageSpanGroup{}
		varStorageSpanGroup.Disks = varStorageSpanGroupWithoutEmbeddedStruct.Disks
		*o = StorageSpanGroup(varStorageSpanGroup)
	} else {
		return err
	}

	varStorageSpanGroup := _StorageSpanGroup{}

	err = json.Unmarshal(bytes, &varStorageSpanGroup)
	if err == nil {
		o.MoBaseComplexType = varStorageSpanGroup.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Disks")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableStorageSpanGroup struct {
	value *StorageSpanGroup
	isSet bool
}

func (v NullableStorageSpanGroup) Get() *StorageSpanGroup {
	return v.value
}

func (v *NullableStorageSpanGroup) Set(val *StorageSpanGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSpanGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSpanGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSpanGroup(val *StorageSpanGroup) *NullableStorageSpanGroup {
	return &NullableStorageSpanGroup{value: val, isSet: true}
}

func (v NullableStorageSpanGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSpanGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
