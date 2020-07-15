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

// PoolAbstractLeaseAllOf Definition of the list of properties defined in 'pool.AbstractLease', excluding properties defined in parent classes.
type PoolAbstractLeaseAllOf struct {
	// Moid of the entity/server profile that owns this ID.
	AssignedToMoid *string `json:"AssignedToMoid,omitempty"`
	// Type of the entity that owns this ID.
	AssignedToType       *string `json:"AssignedToType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractLeaseAllOf PoolAbstractLeaseAllOf

// NewPoolAbstractLeaseAllOf instantiates a new PoolAbstractLeaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractLeaseAllOf() *PoolAbstractLeaseAllOf {
	this := PoolAbstractLeaseAllOf{}
	return &this
}

// NewPoolAbstractLeaseAllOfWithDefaults instantiates a new PoolAbstractLeaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractLeaseAllOfWithDefaults() *PoolAbstractLeaseAllOf {
	this := PoolAbstractLeaseAllOf{}
	return &this
}

// GetAssignedToMoid returns the AssignedToMoid field value if set, zero value otherwise.
func (o *PoolAbstractLeaseAllOf) GetAssignedToMoid() string {
	if o == nil || o.AssignedToMoid == nil {
		var ret string
		return ret
	}
	return *o.AssignedToMoid
}

// GetAssignedToMoidOk returns a tuple with the AssignedToMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractLeaseAllOf) GetAssignedToMoidOk() (*string, bool) {
	if o == nil || o.AssignedToMoid == nil {
		return nil, false
	}
	return o.AssignedToMoid, true
}

// HasAssignedToMoid returns a boolean if a field has been set.
func (o *PoolAbstractLeaseAllOf) HasAssignedToMoid() bool {
	if o != nil && o.AssignedToMoid != nil {
		return true
	}

	return false
}

// SetAssignedToMoid gets a reference to the given string and assigns it to the AssignedToMoid field.
func (o *PoolAbstractLeaseAllOf) SetAssignedToMoid(v string) {
	o.AssignedToMoid = &v
}

// GetAssignedToType returns the AssignedToType field value if set, zero value otherwise.
func (o *PoolAbstractLeaseAllOf) GetAssignedToType() string {
	if o == nil || o.AssignedToType == nil {
		var ret string
		return ret
	}
	return *o.AssignedToType
}

// GetAssignedToTypeOk returns a tuple with the AssignedToType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractLeaseAllOf) GetAssignedToTypeOk() (*string, bool) {
	if o == nil || o.AssignedToType == nil {
		return nil, false
	}
	return o.AssignedToType, true
}

// HasAssignedToType returns a boolean if a field has been set.
func (o *PoolAbstractLeaseAllOf) HasAssignedToType() bool {
	if o != nil && o.AssignedToType != nil {
		return true
	}

	return false
}

// SetAssignedToType gets a reference to the given string and assigns it to the AssignedToType field.
func (o *PoolAbstractLeaseAllOf) SetAssignedToType(v string) {
	o.AssignedToType = &v
}

func (o PoolAbstractLeaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignedToMoid != nil {
		toSerialize["AssignedToMoid"] = o.AssignedToMoid
	}
	if o.AssignedToType != nil {
		toSerialize["AssignedToType"] = o.AssignedToType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PoolAbstractLeaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPoolAbstractLeaseAllOf := _PoolAbstractLeaseAllOf{}

	if err = json.Unmarshal(bytes, &varPoolAbstractLeaseAllOf); err == nil {
		*o = PoolAbstractLeaseAllOf(varPoolAbstractLeaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AssignedToMoid")
		delete(additionalProperties, "AssignedToType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePoolAbstractLeaseAllOf struct {
	value *PoolAbstractLeaseAllOf
	isSet bool
}

func (v NullablePoolAbstractLeaseAllOf) Get() *PoolAbstractLeaseAllOf {
	return v.value
}

func (v *NullablePoolAbstractLeaseAllOf) Set(val *PoolAbstractLeaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractLeaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractLeaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractLeaseAllOf(val *PoolAbstractLeaseAllOf) *NullablePoolAbstractLeaseAllOf {
	return &NullablePoolAbstractLeaseAllOf{value: val, isSet: true}
}

func (v NullablePoolAbstractLeaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractLeaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
