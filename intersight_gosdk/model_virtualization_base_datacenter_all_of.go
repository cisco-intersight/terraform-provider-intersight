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

// VirtualizationBaseDatacenterAllOf Definition of the list of properties defined in 'virtualization.BaseDatacenter', excluding properties defined in parent classes.
type VirtualizationBaseDatacenterAllOf struct {
	// Internally generated identity of this datacenter. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is a MOR (managed object reference).
	Identity *string `json:"Identity,omitempty"`
	// User provided name for the datacenter. Usually, this name is subject to manipulations by user. It is not the identity of the datacenter.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseDatacenterAllOf VirtualizationBaseDatacenterAllOf

// NewVirtualizationBaseDatacenterAllOf instantiates a new VirtualizationBaseDatacenterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseDatacenterAllOf() *VirtualizationBaseDatacenterAllOf {
	this := VirtualizationBaseDatacenterAllOf{}
	return &this
}

// NewVirtualizationBaseDatacenterAllOfWithDefaults instantiates a new VirtualizationBaseDatacenterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseDatacenterAllOfWithDefaults() *VirtualizationBaseDatacenterAllOf {
	this := VirtualizationBaseDatacenterAllOf{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationBaseDatacenterAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatacenterAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationBaseDatacenterAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationBaseDatacenterAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseDatacenterAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatacenterAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseDatacenterAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseDatacenterAllOf) SetName(v string) {
	o.Name = &v
}

func (o VirtualizationBaseDatacenterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseDatacenterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationBaseDatacenterAllOf := _VirtualizationBaseDatacenterAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationBaseDatacenterAllOf); err == nil {
		*o = VirtualizationBaseDatacenterAllOf(varVirtualizationBaseDatacenterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationBaseDatacenterAllOf struct {
	value *VirtualizationBaseDatacenterAllOf
	isSet bool
}

func (v NullableVirtualizationBaseDatacenterAllOf) Get() *VirtualizationBaseDatacenterAllOf {
	return v.value
}

func (v *NullableVirtualizationBaseDatacenterAllOf) Set(val *VirtualizationBaseDatacenterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseDatacenterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseDatacenterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseDatacenterAllOf(val *VirtualizationBaseDatacenterAllOf) *NullableVirtualizationBaseDatacenterAllOf {
	return &NullableVirtualizationBaseDatacenterAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationBaseDatacenterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseDatacenterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
