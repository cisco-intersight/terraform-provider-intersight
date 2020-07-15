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

// BiosBootMode The mode through which bios has booted.
type BiosBootMode struct {
	InventoryBase
	// The actual BIOS boot mode as reported by the platform BIOS.
	ActualBootMode       *string                              `json:"ActualBootMode,omitempty"`
	ComputeRackUnit      *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BiosBootMode BiosBootMode

// NewBiosBootMode instantiates a new BiosBootMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosBootMode() *BiosBootMode {
	this := BiosBootMode{}
	return &this
}

// NewBiosBootModeWithDefaults instantiates a new BiosBootMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosBootModeWithDefaults() *BiosBootMode {
	this := BiosBootMode{}
	return &this
}

// GetActualBootMode returns the ActualBootMode field value if set, zero value otherwise.
func (o *BiosBootMode) GetActualBootMode() string {
	if o == nil || o.ActualBootMode == nil {
		var ret string
		return ret
	}
	return *o.ActualBootMode
}

// GetActualBootModeOk returns a tuple with the ActualBootMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootMode) GetActualBootModeOk() (*string, bool) {
	if o == nil || o.ActualBootMode == nil {
		return nil, false
	}
	return o.ActualBootMode, true
}

// HasActualBootMode returns a boolean if a field has been set.
func (o *BiosBootMode) HasActualBootMode() bool {
	if o != nil && o.ActualBootMode != nil {
		return true
	}

	return false
}

// SetActualBootMode gets a reference to the given string and assigns it to the ActualBootMode field.
func (o *BiosBootMode) SetActualBootMode(v string) {
	o.ActualBootMode = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *BiosBootMode) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootMode) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *BiosBootMode) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *BiosBootMode) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *BiosBootMode) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootMode) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BiosBootMode) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BiosBootMode) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BiosBootMode) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootMode) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosBootMode) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosBootMode) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o BiosBootMode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if o.ActualBootMode != nil {
		toSerialize["ActualBootMode"] = o.ActualBootMode
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BiosBootMode) UnmarshalJSON(bytes []byte) (err error) {
	type BiosBootModeWithoutEmbeddedStruct struct {
		// The actual BIOS boot mode as reported by the platform BIOS.
		ActualBootMode      *string                              `json:"ActualBootMode,omitempty"`
		ComputeRackUnit     *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varBiosBootModeWithoutEmbeddedStruct := BiosBootModeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBiosBootModeWithoutEmbeddedStruct)
	if err == nil {
		varBiosBootMode := _BiosBootMode{}
		varBiosBootMode.ActualBootMode = varBiosBootModeWithoutEmbeddedStruct.ActualBootMode
		varBiosBootMode.ComputeRackUnit = varBiosBootModeWithoutEmbeddedStruct.ComputeRackUnit
		varBiosBootMode.InventoryDeviceInfo = varBiosBootModeWithoutEmbeddedStruct.InventoryDeviceInfo
		varBiosBootMode.RegisteredDevice = varBiosBootModeWithoutEmbeddedStruct.RegisteredDevice
		*o = BiosBootMode(varBiosBootMode)
	} else {
		return err
	}

	varBiosBootMode := _BiosBootMode{}

	err = json.Unmarshal(bytes, &varBiosBootMode)
	if err == nil {
		o.InventoryBase = varBiosBootMode.InventoryBase
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ActualBootMode")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableBiosBootMode struct {
	value *BiosBootMode
	isSet bool
}

func (v NullableBiosBootMode) Get() *BiosBootMode {
	return v.value
}

func (v *NullableBiosBootMode) Set(val *BiosBootMode) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosBootMode) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosBootMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosBootMode(val *BiosBootMode) *NullableBiosBootMode {
	return &NullableBiosBootMode{value: val, isSet: true}
}

func (v NullableBiosBootMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosBootMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
