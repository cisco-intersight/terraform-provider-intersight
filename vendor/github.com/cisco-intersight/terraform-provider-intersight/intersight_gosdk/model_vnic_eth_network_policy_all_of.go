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

// VnicEthNetworkPolicyAllOf Definition of the list of properties defined in 'vnic.EthNetworkPolicy', excluding properties defined in parent classes.
type VnicEthNetworkPolicyAllOf struct {
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.
	TargetPlatform       *string                               `json:"TargetPlatform,omitempty"`
	VlanSettings         *VnicVlanSettings                     `json:"VlanSettings,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthNetworkPolicyAllOf VnicEthNetworkPolicyAllOf

// NewVnicEthNetworkPolicyAllOf instantiates a new VnicEthNetworkPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthNetworkPolicyAllOf() *VnicEthNetworkPolicyAllOf {
	this := VnicEthNetworkPolicyAllOf{}
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewVnicEthNetworkPolicyAllOfWithDefaults instantiates a new VnicEthNetworkPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthNetworkPolicyAllOfWithDefaults() *VnicEthNetworkPolicyAllOf {
	this := VnicEthNetworkPolicyAllOf{}
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *VnicEthNetworkPolicyAllOf) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthNetworkPolicyAllOf) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicyAllOf) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *VnicEthNetworkPolicyAllOf) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetVlanSettings returns the VlanSettings field value if set, zero value otherwise.
func (o *VnicEthNetworkPolicyAllOf) GetVlanSettings() VnicVlanSettings {
	if o == nil || o.VlanSettings == nil {
		var ret VnicVlanSettings
		return ret
	}
	return *o.VlanSettings
}

// GetVlanSettingsOk returns a tuple with the VlanSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthNetworkPolicyAllOf) GetVlanSettingsOk() (*VnicVlanSettings, bool) {
	if o == nil || o.VlanSettings == nil {
		return nil, false
	}
	return o.VlanSettings, true
}

// HasVlanSettings returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicyAllOf) HasVlanSettings() bool {
	if o != nil && o.VlanSettings != nil {
		return true
	}

	return false
}

// SetVlanSettings gets a reference to the given VnicVlanSettings and assigns it to the VlanSettings field.
func (o *VnicEthNetworkPolicyAllOf) SetVlanSettings(v VnicVlanSettings) {
	o.VlanSettings = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicEthNetworkPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthNetworkPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicEthNetworkPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicEthNetworkPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.VlanSettings != nil {
		toSerialize["VlanSettings"] = o.VlanSettings
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicEthNetworkPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicEthNetworkPolicyAllOf := _VnicEthNetworkPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varVnicEthNetworkPolicyAllOf); err == nil {
		*o = VnicEthNetworkPolicyAllOf(varVnicEthNetworkPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "VlanSettings")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicEthNetworkPolicyAllOf struct {
	value *VnicEthNetworkPolicyAllOf
	isSet bool
}

func (v NullableVnicEthNetworkPolicyAllOf) Get() *VnicEthNetworkPolicyAllOf {
	return v.value
}

func (v *NullableVnicEthNetworkPolicyAllOf) Set(val *VnicEthNetworkPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthNetworkPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthNetworkPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthNetworkPolicyAllOf(val *VnicEthNetworkPolicyAllOf) *NullableVnicEthNetworkPolicyAllOf {
	return &NullableVnicEthNetworkPolicyAllOf{value: val, isSet: true}
}

func (v NullableVnicEthNetworkPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthNetworkPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
