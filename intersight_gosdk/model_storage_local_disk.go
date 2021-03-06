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

// StorageLocalDisk This models disks in the disk group.
type StorageLocalDisk struct {
	MoBaseComplexType
	// The slot number of the disk to be referenced. As this is a policy, this slot number may or may not be valid depending on the number of disks in the associated server.
	SlotNumber           *int64 `json:"SlotNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageLocalDisk StorageLocalDisk

// NewStorageLocalDisk instantiates a new StorageLocalDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageLocalDisk() *StorageLocalDisk {
	this := StorageLocalDisk{}
	return &this
}

// NewStorageLocalDiskWithDefaults instantiates a new StorageLocalDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageLocalDiskWithDefaults() *StorageLocalDisk {
	this := StorageLocalDisk{}
	return &this
}

// GetSlotNumber returns the SlotNumber field value if set, zero value otherwise.
func (o *StorageLocalDisk) GetSlotNumber() int64 {
	if o == nil || o.SlotNumber == nil {
		var ret int64
		return ret
	}
	return *o.SlotNumber
}

// GetSlotNumberOk returns a tuple with the SlotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageLocalDisk) GetSlotNumberOk() (*int64, bool) {
	if o == nil || o.SlotNumber == nil {
		return nil, false
	}
	return o.SlotNumber, true
}

// HasSlotNumber returns a boolean if a field has been set.
func (o *StorageLocalDisk) HasSlotNumber() bool {
	if o != nil && o.SlotNumber != nil {
		return true
	}

	return false
}

// SetSlotNumber gets a reference to the given int64 and assigns it to the SlotNumber field.
func (o *StorageLocalDisk) SetSlotNumber(v int64) {
	o.SlotNumber = &v
}

func (o StorageLocalDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.SlotNumber != nil {
		toSerialize["SlotNumber"] = o.SlotNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageLocalDisk) UnmarshalJSON(bytes []byte) (err error) {
	type StorageLocalDiskWithoutEmbeddedStruct struct {
		// The slot number of the disk to be referenced. As this is a policy, this slot number may or may not be valid depending on the number of disks in the associated server.
		SlotNumber *int64 `json:"SlotNumber,omitempty"`
	}

	varStorageLocalDiskWithoutEmbeddedStruct := StorageLocalDiskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageLocalDiskWithoutEmbeddedStruct)
	if err == nil {
		varStorageLocalDisk := _StorageLocalDisk{}
		varStorageLocalDisk.SlotNumber = varStorageLocalDiskWithoutEmbeddedStruct.SlotNumber
		*o = StorageLocalDisk(varStorageLocalDisk)
	} else {
		return err
	}

	varStorageLocalDisk := _StorageLocalDisk{}

	err = json.Unmarshal(bytes, &varStorageLocalDisk)
	if err == nil {
		o.MoBaseComplexType = varStorageLocalDisk.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "SlotNumber")

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

type NullableStorageLocalDisk struct {
	value *StorageLocalDisk
	isSet bool
}

func (v NullableStorageLocalDisk) Get() *StorageLocalDisk {
	return v.value
}

func (v *NullableStorageLocalDisk) Set(val *StorageLocalDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageLocalDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageLocalDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageLocalDisk(val *StorageLocalDisk) *NullableStorageLocalDisk {
	return &NullableStorageLocalDisk{value: val, isSet: true}
}

func (v NullableStorageLocalDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageLocalDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
