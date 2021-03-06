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

// CapabilitySwitchCapabilityDef Abstract class that defines properties common for all Switch specific 'defs'.
type CapabilitySwitchCapabilityDef struct {
	CapabilityCapability
	// Product Identifier for a Switch/Fabric-Interconnect.
	Pid *string `json:"Pid,omitempty"`
	// SKU information for Switch/Fabric-Interconnect.
	Sku *string `json:"Sku,omitempty"`
	// VID information for Switch/Fabric-Interconnect.
	Vid                  *string `json:"Vid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySwitchCapabilityDef CapabilitySwitchCapabilityDef

// NewCapabilitySwitchCapabilityDef instantiates a new CapabilitySwitchCapabilityDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchCapabilityDef() *CapabilitySwitchCapabilityDef {
	this := CapabilitySwitchCapabilityDef{}
	return &this
}

// NewCapabilitySwitchCapabilityDefWithDefaults instantiates a new CapabilitySwitchCapabilityDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchCapabilityDefWithDefaults() *CapabilitySwitchCapabilityDef {
	this := CapabilitySwitchCapabilityDef{}
	return &this
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *CapabilitySwitchCapabilityDef) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchCapabilityDef) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *CapabilitySwitchCapabilityDef) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *CapabilitySwitchCapabilityDef) SetPid(v string) {
	o.Pid = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *CapabilitySwitchCapabilityDef) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchCapabilityDef) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *CapabilitySwitchCapabilityDef) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *CapabilitySwitchCapabilityDef) SetSku(v string) {
	o.Sku = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *CapabilitySwitchCapabilityDef) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchCapabilityDef) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *CapabilitySwitchCapabilityDef) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *CapabilitySwitchCapabilityDef) SetVid(v string) {
	o.Vid = &v
}

func (o CapabilitySwitchCapabilityDef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.Sku != nil {
		toSerialize["Sku"] = o.Sku
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySwitchCapabilityDef) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilitySwitchCapabilityDefWithoutEmbeddedStruct struct {
		// Product Identifier for a Switch/Fabric-Interconnect.
		Pid *string `json:"Pid,omitempty"`
		// SKU information for Switch/Fabric-Interconnect.
		Sku *string `json:"Sku,omitempty"`
		// VID information for Switch/Fabric-Interconnect.
		Vid *string `json:"Vid,omitempty"`
	}

	varCapabilitySwitchCapabilityDefWithoutEmbeddedStruct := CapabilitySwitchCapabilityDefWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchCapabilityDefWithoutEmbeddedStruct)
	if err == nil {
		varCapabilitySwitchCapabilityDef := _CapabilitySwitchCapabilityDef{}
		varCapabilitySwitchCapabilityDef.Pid = varCapabilitySwitchCapabilityDefWithoutEmbeddedStruct.Pid
		varCapabilitySwitchCapabilityDef.Sku = varCapabilitySwitchCapabilityDefWithoutEmbeddedStruct.Sku
		varCapabilitySwitchCapabilityDef.Vid = varCapabilitySwitchCapabilityDefWithoutEmbeddedStruct.Vid
		*o = CapabilitySwitchCapabilityDef(varCapabilitySwitchCapabilityDef)
	} else {
		return err
	}

	varCapabilitySwitchCapabilityDef := _CapabilitySwitchCapabilityDef{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchCapabilityDef)
	if err == nil {
		o.CapabilityCapability = varCapabilitySwitchCapabilityDef.CapabilityCapability
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "Sku")
		delete(additionalProperties, "Vid")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableCapabilitySwitchCapabilityDef struct {
	value *CapabilitySwitchCapabilityDef
	isSet bool
}

func (v NullableCapabilitySwitchCapabilityDef) Get() *CapabilitySwitchCapabilityDef {
	return v.value
}

func (v *NullableCapabilitySwitchCapabilityDef) Set(val *CapabilitySwitchCapabilityDef) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchCapabilityDef) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchCapabilityDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchCapabilityDef(val *CapabilitySwitchCapabilityDef) *NullableCapabilitySwitchCapabilityDef {
	return &NullableCapabilitySwitchCapabilityDef{value: val, isSet: true}
}

func (v NullableCapabilitySwitchCapabilityDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchCapabilityDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
