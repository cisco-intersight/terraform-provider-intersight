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

// VnicFlogiSettingsAllOf Definition of the list of properties defined in 'vnic.FlogiSettings', excluding properties defined in parent classes.
type VnicFlogiSettingsAllOf struct {
	// The number of times that the system tries to log in to the fabric after the first failure.
	Retries *int64 `json:"Retries,omitempty"`
	// The number of milliseconds that the system waits before it tries to log in again.
	Timeout              *int64 `json:"Timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFlogiSettingsAllOf VnicFlogiSettingsAllOf

// NewVnicFlogiSettingsAllOf instantiates a new VnicFlogiSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFlogiSettingsAllOf() *VnicFlogiSettingsAllOf {
	this := VnicFlogiSettingsAllOf{}
	return &this
}

// NewVnicFlogiSettingsAllOfWithDefaults instantiates a new VnicFlogiSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFlogiSettingsAllOfWithDefaults() *VnicFlogiSettingsAllOf {
	this := VnicFlogiSettingsAllOf{}
	return &this
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *VnicFlogiSettingsAllOf) GetRetries() int64 {
	if o == nil || o.Retries == nil {
		var ret int64
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFlogiSettingsAllOf) GetRetriesOk() (*int64, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *VnicFlogiSettingsAllOf) HasRetries() bool {
	if o != nil && o.Retries != nil {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int64 and assigns it to the Retries field.
func (o *VnicFlogiSettingsAllOf) SetRetries(v int64) {
	o.Retries = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *VnicFlogiSettingsAllOf) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFlogiSettingsAllOf) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *VnicFlogiSettingsAllOf) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *VnicFlogiSettingsAllOf) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o VnicFlogiSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Retries != nil {
		toSerialize["Retries"] = o.Retries
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFlogiSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicFlogiSettingsAllOf := _VnicFlogiSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varVnicFlogiSettingsAllOf); err == nil {
		*o = VnicFlogiSettingsAllOf(varVnicFlogiSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Retries")
		delete(additionalProperties, "Timeout")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicFlogiSettingsAllOf struct {
	value *VnicFlogiSettingsAllOf
	isSet bool
}

func (v NullableVnicFlogiSettingsAllOf) Get() *VnicFlogiSettingsAllOf {
	return v.value
}

func (v *NullableVnicFlogiSettingsAllOf) Set(val *VnicFlogiSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFlogiSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFlogiSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFlogiSettingsAllOf(val *VnicFlogiSettingsAllOf) *NullableVnicFlogiSettingsAllOf {
	return &NullableVnicFlogiSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicFlogiSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFlogiSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
