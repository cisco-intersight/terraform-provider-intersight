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

// IppoolIpLeaseAllOf Definition of the list of properties defined in 'ippool.IpLease', excluding properties defined in parent classes.
type IppoolIpLeaseAllOf struct {
	// IPv4 Address given as a lease to an external entity like server profiles.
	IpV4Address          *string                       `json:"IpV4Address,omitempty"`
	IpV4Config           *IppoolIpV4Config             `json:"IpV4Config,omitempty"`
	Pool                 *IppoolPoolRelationship       `json:"Pool,omitempty"`
	PoolMember           *IppoolPoolMemberRelationship `json:"PoolMember,omitempty"`
	Universe             *IppoolUniverseRelationship   `json:"Universe,omitempty"`
	Vrf                  *VrfVrfRelationship           `json:"Vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolIpLeaseAllOf IppoolIpLeaseAllOf

// NewIppoolIpLeaseAllOf instantiates a new IppoolIpLeaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolIpLeaseAllOf() *IppoolIpLeaseAllOf {
	this := IppoolIpLeaseAllOf{}
	return &this
}

// NewIppoolIpLeaseAllOfWithDefaults instantiates a new IppoolIpLeaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolIpLeaseAllOfWithDefaults() *IppoolIpLeaseAllOf {
	this := IppoolIpLeaseAllOf{}
	return &this
}

// GetIpV4Address returns the IpV4Address field value if set, zero value otherwise.
func (o *IppoolIpLeaseAllOf) GetIpV4Address() string {
	if o == nil || o.IpV4Address == nil {
		var ret string
		return ret
	}
	return *o.IpV4Address
}

// GetIpV4AddressOk returns a tuple with the IpV4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLeaseAllOf) GetIpV4AddressOk() (*string, bool) {
	if o == nil || o.IpV4Address == nil {
		return nil, false
	}
	return o.IpV4Address, true
}

// HasIpV4Address returns a boolean if a field has been set.
func (o *IppoolIpLeaseAllOf) HasIpV4Address() bool {
	if o != nil && o.IpV4Address != nil {
		return true
	}

	return false
}

// SetIpV4Address gets a reference to the given string and assigns it to the IpV4Address field.
func (o *IppoolIpLeaseAllOf) SetIpV4Address(v string) {
	o.IpV4Address = &v
}

// GetIpV4Config returns the IpV4Config field value if set, zero value otherwise.
func (o *IppoolIpLeaseAllOf) GetIpV4Config() IppoolIpV4Config {
	if o == nil || o.IpV4Config == nil {
		var ret IppoolIpV4Config
		return ret
	}
	return *o.IpV4Config
}

// GetIpV4ConfigOk returns a tuple with the IpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLeaseAllOf) GetIpV4ConfigOk() (*IppoolIpV4Config, bool) {
	if o == nil || o.IpV4Config == nil {
		return nil, false
	}
	return o.IpV4Config, true
}

// HasIpV4Config returns a boolean if a field has been set.
func (o *IppoolIpLeaseAllOf) HasIpV4Config() bool {
	if o != nil && o.IpV4Config != nil {
		return true
	}

	return false
}

// SetIpV4Config gets a reference to the given IppoolIpV4Config and assigns it to the IpV4Config field.
func (o *IppoolIpLeaseAllOf) SetIpV4Config(v IppoolIpV4Config) {
	o.IpV4Config = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IppoolIpLeaseAllOf) GetPool() IppoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLeaseAllOf) GetPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolIpLeaseAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IppoolPoolRelationship and assigns it to the Pool field.
func (o *IppoolIpLeaseAllOf) SetPool(v IppoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *IppoolIpLeaseAllOf) GetPoolMember() IppoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret IppoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLeaseAllOf) GetPoolMemberOk() (*IppoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *IppoolIpLeaseAllOf) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given IppoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *IppoolIpLeaseAllOf) SetPoolMember(v IppoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *IppoolIpLeaseAllOf) GetUniverse() IppoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret IppoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLeaseAllOf) GetUniverseOk() (*IppoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *IppoolIpLeaseAllOf) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given IppoolUniverseRelationship and assigns it to the Universe field.
func (o *IppoolIpLeaseAllOf) SetUniverse(v IppoolUniverseRelationship) {
	o.Universe = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *IppoolIpLeaseAllOf) GetVrf() VrfVrfRelationship {
	if o == nil || o.Vrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLeaseAllOf) GetVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.Vrf == nil {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *IppoolIpLeaseAllOf) HasVrf() bool {
	if o != nil && o.Vrf != nil {
		return true
	}

	return false
}

// SetVrf gets a reference to the given VrfVrfRelationship and assigns it to the Vrf field.
func (o *IppoolIpLeaseAllOf) SetVrf(v VrfVrfRelationship) {
	o.Vrf = &v
}

func (o IppoolIpLeaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpV4Address != nil {
		toSerialize["IpV4Address"] = o.IpV4Address
	}
	if o.IpV4Config != nil {
		toSerialize["IpV4Config"] = o.IpV4Config
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}
	if o.Vrf != nil {
		toSerialize["Vrf"] = o.Vrf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolIpLeaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIppoolIpLeaseAllOf := _IppoolIpLeaseAllOf{}

	if err = json.Unmarshal(bytes, &varIppoolIpLeaseAllOf); err == nil {
		*o = IppoolIpLeaseAllOf(varIppoolIpLeaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "IpV4Address")
		delete(additionalProperties, "IpV4Config")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")
		delete(additionalProperties, "Vrf")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIppoolIpLeaseAllOf struct {
	value *IppoolIpLeaseAllOf
	isSet bool
}

func (v NullableIppoolIpLeaseAllOf) Get() *IppoolIpLeaseAllOf {
	return v.value
}

func (v *NullableIppoolIpLeaseAllOf) Set(val *IppoolIpLeaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolIpLeaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolIpLeaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolIpLeaseAllOf(val *IppoolIpLeaseAllOf) *NullableIppoolIpLeaseAllOf {
	return &NullableIppoolIpLeaseAllOf{value: val, isSet: true}
}

func (v NullableIppoolIpLeaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolIpLeaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
