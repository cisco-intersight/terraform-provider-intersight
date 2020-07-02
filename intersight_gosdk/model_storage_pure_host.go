/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-19T15:15:17Z.
 *
 * API version: 1.0.9-1903
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// StoragePureHost A host entity in PureStorage FlashArray. It is an abstraction used by PureStorage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.
type StoragePureHost struct {
	StorageBaseHost
	// Name of host group where the host belongs to. Empty if host is not part of any HostGroup.
	HostGroupName      *string                                 `json:"HostGroupName,omitempty"`
	StorageUtilization *StoragePureHostUtilization             `json:"StorageUtilization,omitempty"`
	Array              *StoragePureArrayRelationship           `json:"Array,omitempty"`
	HostGroup          *StoragePureHostGroupRelationship       `json:"HostGroup,omitempty"`
	ProtectionGroup    *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice   *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
}

// NewStoragePureHost instantiates a new StoragePureHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureHost() *StoragePureHost {
	this := StoragePureHost{}
	return &this
}

// NewStoragePureHostWithDefaults instantiates a new StoragePureHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureHostWithDefaults() *StoragePureHost {
	this := StoragePureHost{}
	return &this
}

// GetHostGroupName returns the HostGroupName field value if set, zero value otherwise.
func (o *StoragePureHost) GetHostGroupName() string {
	if o == nil || o.HostGroupName == nil {
		var ret string
		return ret
	}
	return *o.HostGroupName
}

// GetHostGroupNameOk returns a tuple with the HostGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHost) GetHostGroupNameOk() (*string, bool) {
	if o == nil || o.HostGroupName == nil {
		return nil, false
	}
	return o.HostGroupName, true
}

// HasHostGroupName returns a boolean if a field has been set.
func (o *StoragePureHost) HasHostGroupName() bool {
	if o != nil && o.HostGroupName != nil {
		return true
	}

	return false
}

// SetHostGroupName gets a reference to the given string and assigns it to the HostGroupName field.
func (o *StoragePureHost) SetHostGroupName(v string) {
	o.HostGroupName = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise.
func (o *StoragePureHost) GetStorageUtilization() StoragePureHostUtilization {
	if o == nil || o.StorageUtilization == nil {
		var ret StoragePureHostUtilization
		return ret
	}
	return *o.StorageUtilization
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHost) GetStorageUtilizationOk() (*StoragePureHostUtilization, bool) {
	if o == nil || o.StorageUtilization == nil {
		return nil, false
	}
	return o.StorageUtilization, true
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StoragePureHost) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization != nil {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given StoragePureHostUtilization and assigns it to the StorageUtilization field.
func (o *StoragePureHost) SetStorageUtilization(v StoragePureHostUtilization) {
	o.StorageUtilization = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureHost) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHost) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureHost) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureHost) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetHostGroup returns the HostGroup field value if set, zero value otherwise.
func (o *StoragePureHost) GetHostGroup() StoragePureHostGroupRelationship {
	if o == nil || o.HostGroup == nil {
		var ret StoragePureHostGroupRelationship
		return ret
	}
	return *o.HostGroup
}

// GetHostGroupOk returns a tuple with the HostGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHost) GetHostGroupOk() (*StoragePureHostGroupRelationship, bool) {
	if o == nil || o.HostGroup == nil {
		return nil, false
	}
	return o.HostGroup, true
}

// HasHostGroup returns a boolean if a field has been set.
func (o *StoragePureHost) HasHostGroup() bool {
	if o != nil && o.HostGroup != nil {
		return true
	}

	return false
}

// SetHostGroup gets a reference to the given StoragePureHostGroupRelationship and assigns it to the HostGroup field.
func (o *StoragePureHost) SetHostGroup(v StoragePureHostGroupRelationship) {
	o.HostGroup = &v
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise.
func (o *StoragePureHost) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || o.ProtectionGroup == nil {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHost) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil || o.ProtectionGroup == nil {
		return nil, false
	}
	return o.ProtectionGroup, true
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureHost) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup != nil {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given StoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureHost) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureHost) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHost) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureHost) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureHost) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseHost, errStorageBaseHost := json.Marshal(o.StorageBaseHost)
	if errStorageBaseHost != nil {
		return []byte{}, errStorageBaseHost
	}
	errStorageBaseHost = json.Unmarshal([]byte(serializedStorageBaseHost), &toSerialize)
	if errStorageBaseHost != nil {
		return []byte{}, errStorageBaseHost
	}
	if o.HostGroupName != nil {
		toSerialize["HostGroupName"] = o.HostGroupName
	}
	if o.StorageUtilization != nil {
		toSerialize["StorageUtilization"] = o.StorageUtilization
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.HostGroup != nil {
		toSerialize["HostGroup"] = o.HostGroup
	}
	if o.ProtectionGroup != nil {
		toSerialize["ProtectionGroup"] = o.ProtectionGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableStoragePureHost struct {
	value *StoragePureHost
	isSet bool
}

func (v NullableStoragePureHost) Get() *StoragePureHost {
	return v.value
}

func (v *NullableStoragePureHost) Set(val *StoragePureHost) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureHost) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureHost(val *StoragePureHost) *NullableStoragePureHost {
	return &NullableStoragePureHost{value: val, isSet: true}
}

func (v NullableStoragePureHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}