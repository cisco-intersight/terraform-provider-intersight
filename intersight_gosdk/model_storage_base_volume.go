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

// StorageBaseVolume Generic storage volume object. It is a provisioned storage identity which can be mapped to host to enable host access.
type StorageBaseVolume struct {
	MoBaseMo
	// Short description about the volume.
	Description *string `json:"Description,omitempty"`
	// NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor.
	NaaId *string `json:"NaaId,omitempty"`
	// Named entity of the volume.
	Name *string `json:"Name,omitempty"`
	// User provisioned volume size. It is the size exposed to host.
	Size                 *int64               `json:"Size,omitempty"`
	StorageUtilization   *StorageBaseCapacity `json:"StorageUtilization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseVolume StorageBaseVolume

// NewStorageBaseVolume instantiates a new StorageBaseVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseVolume() *StorageBaseVolume {
	this := StorageBaseVolume{}
	return &this
}

// NewStorageBaseVolumeWithDefaults instantiates a new StorageBaseVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseVolumeWithDefaults() *StorageBaseVolume {
	this := StorageBaseVolume{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageBaseVolume) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolume) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageBaseVolume) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageBaseVolume) SetDescription(v string) {
	o.Description = &v
}

// GetNaaId returns the NaaId field value if set, zero value otherwise.
func (o *StorageBaseVolume) GetNaaId() string {
	if o == nil || o.NaaId == nil {
		var ret string
		return ret
	}
	return *o.NaaId
}

// GetNaaIdOk returns a tuple with the NaaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolume) GetNaaIdOk() (*string, bool) {
	if o == nil || o.NaaId == nil {
		return nil, false
	}
	return o.NaaId, true
}

// HasNaaId returns a boolean if a field has been set.
func (o *StorageBaseVolume) HasNaaId() bool {
	if o != nil && o.NaaId != nil {
		return true
	}

	return false
}

// SetNaaId gets a reference to the given string and assigns it to the NaaId field.
func (o *StorageBaseVolume) SetNaaId(v string) {
	o.NaaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseVolume) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolume) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseVolume) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseVolume) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageBaseVolume) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolume) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageBaseVolume) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageBaseVolume) SetSize(v int64) {
	o.Size = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise.
func (o *StorageBaseVolume) GetStorageUtilization() StorageBaseCapacity {
	if o == nil || o.StorageUtilization == nil {
		var ret StorageBaseCapacity
		return ret
	}
	return *o.StorageUtilization
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolume) GetStorageUtilizationOk() (*StorageBaseCapacity, bool) {
	if o == nil || o.StorageUtilization == nil {
		return nil, false
	}
	return o.StorageUtilization, true
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageBaseVolume) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization != nil {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given StorageBaseCapacity and assigns it to the StorageUtilization field.
func (o *StorageBaseVolume) SetStorageUtilization(v StorageBaseCapacity) {
	o.StorageUtilization = &v
}

func (o StorageBaseVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.NaaId != nil {
		toSerialize["NaaId"] = o.NaaId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.StorageUtilization != nil {
		toSerialize["StorageUtilization"] = o.StorageUtilization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseVolume) UnmarshalJSON(bytes []byte) (err error) {
	type StorageBaseVolumeWithoutEmbeddedStruct struct {
		// Short description about the volume.
		Description *string `json:"Description,omitempty"`
		// NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor.
		NaaId *string `json:"NaaId,omitempty"`
		// Named entity of the volume.
		Name *string `json:"Name,omitempty"`
		// User provisioned volume size. It is the size exposed to host.
		Size               *int64               `json:"Size,omitempty"`
		StorageUtilization *StorageBaseCapacity `json:"StorageUtilization,omitempty"`
	}

	varStorageBaseVolumeWithoutEmbeddedStruct := StorageBaseVolumeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageBaseVolumeWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseVolume := _StorageBaseVolume{}
		varStorageBaseVolume.Description = varStorageBaseVolumeWithoutEmbeddedStruct.Description
		varStorageBaseVolume.NaaId = varStorageBaseVolumeWithoutEmbeddedStruct.NaaId
		varStorageBaseVolume.Name = varStorageBaseVolumeWithoutEmbeddedStruct.Name
		varStorageBaseVolume.Size = varStorageBaseVolumeWithoutEmbeddedStruct.Size
		varStorageBaseVolume.StorageUtilization = varStorageBaseVolumeWithoutEmbeddedStruct.StorageUtilization
		*o = StorageBaseVolume(varStorageBaseVolume)
	} else {
		return err
	}

	varStorageBaseVolume := _StorageBaseVolume{}

	err = json.Unmarshal(bytes, &varStorageBaseVolume)
	if err == nil {
		o.MoBaseMo = varStorageBaseVolume.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Description")
		delete(additionalProperties, "NaaId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "StorageUtilization")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableStorageBaseVolume struct {
	value *StorageBaseVolume
	isSet bool
}

func (v NullableStorageBaseVolume) Get() *StorageBaseVolume {
	return v.value
}

func (v *NullableStorageBaseVolume) Set(val *StorageBaseVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseVolume(val *StorageBaseVolume) *NullableStorageBaseVolume {
	return &NullableStorageBaseVolume{value: val, isSet: true}
}

func (v NullableStorageBaseVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
