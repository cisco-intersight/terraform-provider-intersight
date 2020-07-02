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

// HyperflexHxZoneResiliencyInfoDtAllOf Definition of the list of properties defined in 'hyperflex.HxZoneResiliencyInfoDt', excluding properties defined in parent classes.
type HyperflexHxZoneResiliencyInfoDtAllOf struct {
	Name           *string                      `json:"Name,omitempty"`
	ResiliencyInfo *HyperflexHxResiliencyInfoDt `json:"ResiliencyInfo,omitempty"`
}

// NewHyperflexHxZoneResiliencyInfoDtAllOf instantiates a new HyperflexHxZoneResiliencyInfoDtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxZoneResiliencyInfoDtAllOf() *HyperflexHxZoneResiliencyInfoDtAllOf {
	this := HyperflexHxZoneResiliencyInfoDtAllOf{}
	return &this
}

// NewHyperflexHxZoneResiliencyInfoDtAllOfWithDefaults instantiates a new HyperflexHxZoneResiliencyInfoDtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxZoneResiliencyInfoDtAllOfWithDefaults() *HyperflexHxZoneResiliencyInfoDtAllOf {
	this := HyperflexHxZoneResiliencyInfoDtAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetName(v string) {
	o.Name = &v
}

// GetResiliencyInfo returns the ResiliencyInfo field value if set, zero value otherwise.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetResiliencyInfo() HyperflexHxResiliencyInfoDt {
	if o == nil || o.ResiliencyInfo == nil {
		var ret HyperflexHxResiliencyInfoDt
		return ret
	}
	return *o.ResiliencyInfo
}

// GetResiliencyInfoOk returns a tuple with the ResiliencyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetResiliencyInfoOk() (*HyperflexHxResiliencyInfoDt, bool) {
	if o == nil || o.ResiliencyInfo == nil {
		return nil, false
	}
	return o.ResiliencyInfo, true
}

// HasResiliencyInfo returns a boolean if a field has been set.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) HasResiliencyInfo() bool {
	if o != nil && o.ResiliencyInfo != nil {
		return true
	}

	return false
}

// SetResiliencyInfo gets a reference to the given HyperflexHxResiliencyInfoDt and assigns it to the ResiliencyInfo field.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetResiliencyInfo(v HyperflexHxResiliencyInfoDt) {
	o.ResiliencyInfo = &v
}

func (o HyperflexHxZoneResiliencyInfoDtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ResiliencyInfo != nil {
		toSerialize["ResiliencyInfo"] = o.ResiliencyInfo
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexHxZoneResiliencyInfoDtAllOf struct {
	value *HyperflexHxZoneResiliencyInfoDtAllOf
	isSet bool
}

func (v NullableHyperflexHxZoneResiliencyInfoDtAllOf) Get() *HyperflexHxZoneResiliencyInfoDtAllOf {
	return v.value
}

func (v *NullableHyperflexHxZoneResiliencyInfoDtAllOf) Set(val *HyperflexHxZoneResiliencyInfoDtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxZoneResiliencyInfoDtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxZoneResiliencyInfoDtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxZoneResiliencyInfoDtAllOf(val *HyperflexHxZoneResiliencyInfoDtAllOf) *NullableHyperflexHxZoneResiliencyInfoDtAllOf {
	return &NullableHyperflexHxZoneResiliencyInfoDtAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxZoneResiliencyInfoDtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxZoneResiliencyInfoDtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}