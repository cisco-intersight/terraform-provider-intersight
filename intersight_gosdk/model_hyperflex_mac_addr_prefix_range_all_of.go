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

// HyperflexMacAddrPrefixRangeAllOf Definition of the list of properties defined in 'hyperflex.MacAddrPrefixRange', excluding properties defined in parent classes.
type HyperflexMacAddrPrefixRangeAllOf struct {
	// The end MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
	EndAddr *string `json:"EndAddr,omitempty"`
	// The start MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
	StartAddr            *string `json:"StartAddr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexMacAddrPrefixRangeAllOf HyperflexMacAddrPrefixRangeAllOf

// NewHyperflexMacAddrPrefixRangeAllOf instantiates a new HyperflexMacAddrPrefixRangeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexMacAddrPrefixRangeAllOf() *HyperflexMacAddrPrefixRangeAllOf {
	this := HyperflexMacAddrPrefixRangeAllOf{}
	return &this
}

// NewHyperflexMacAddrPrefixRangeAllOfWithDefaults instantiates a new HyperflexMacAddrPrefixRangeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexMacAddrPrefixRangeAllOfWithDefaults() *HyperflexMacAddrPrefixRangeAllOf {
	this := HyperflexMacAddrPrefixRangeAllOf{}
	return &this
}

// GetEndAddr returns the EndAddr field value if set, zero value otherwise.
func (o *HyperflexMacAddrPrefixRangeAllOf) GetEndAddr() string {
	if o == nil || o.EndAddr == nil {
		var ret string
		return ret
	}
	return *o.EndAddr
}

// GetEndAddrOk returns a tuple with the EndAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexMacAddrPrefixRangeAllOf) GetEndAddrOk() (*string, bool) {
	if o == nil || o.EndAddr == nil {
		return nil, false
	}
	return o.EndAddr, true
}

// HasEndAddr returns a boolean if a field has been set.
func (o *HyperflexMacAddrPrefixRangeAllOf) HasEndAddr() bool {
	if o != nil && o.EndAddr != nil {
		return true
	}

	return false
}

// SetEndAddr gets a reference to the given string and assigns it to the EndAddr field.
func (o *HyperflexMacAddrPrefixRangeAllOf) SetEndAddr(v string) {
	o.EndAddr = &v
}

// GetStartAddr returns the StartAddr field value if set, zero value otherwise.
func (o *HyperflexMacAddrPrefixRangeAllOf) GetStartAddr() string {
	if o == nil || o.StartAddr == nil {
		var ret string
		return ret
	}
	return *o.StartAddr
}

// GetStartAddrOk returns a tuple with the StartAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexMacAddrPrefixRangeAllOf) GetStartAddrOk() (*string, bool) {
	if o == nil || o.StartAddr == nil {
		return nil, false
	}
	return o.StartAddr, true
}

// HasStartAddr returns a boolean if a field has been set.
func (o *HyperflexMacAddrPrefixRangeAllOf) HasStartAddr() bool {
	if o != nil && o.StartAddr != nil {
		return true
	}

	return false
}

// SetStartAddr gets a reference to the given string and assigns it to the StartAddr field.
func (o *HyperflexMacAddrPrefixRangeAllOf) SetStartAddr(v string) {
	o.StartAddr = &v
}

func (o HyperflexMacAddrPrefixRangeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndAddr != nil {
		toSerialize["EndAddr"] = o.EndAddr
	}
	if o.StartAddr != nil {
		toSerialize["StartAddr"] = o.StartAddr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexMacAddrPrefixRangeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexMacAddrPrefixRangeAllOf := _HyperflexMacAddrPrefixRangeAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexMacAddrPrefixRangeAllOf); err == nil {
		*o = HyperflexMacAddrPrefixRangeAllOf(varHyperflexMacAddrPrefixRangeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "EndAddr")
		delete(additionalProperties, "StartAddr")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexMacAddrPrefixRangeAllOf struct {
	value *HyperflexMacAddrPrefixRangeAllOf
	isSet bool
}

func (v NullableHyperflexMacAddrPrefixRangeAllOf) Get() *HyperflexMacAddrPrefixRangeAllOf {
	return v.value
}

func (v *NullableHyperflexMacAddrPrefixRangeAllOf) Set(val *HyperflexMacAddrPrefixRangeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexMacAddrPrefixRangeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexMacAddrPrefixRangeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexMacAddrPrefixRangeAllOf(val *HyperflexMacAddrPrefixRangeAllOf) *NullableHyperflexMacAddrPrefixRangeAllOf {
	return &NullableHyperflexMacAddrPrefixRangeAllOf{value: val, isSet: true}
}

func (v NullableHyperflexMacAddrPrefixRangeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexMacAddrPrefixRangeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
