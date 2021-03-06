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

// FabricEthNetworkPolicy A policy for all the Virtual LAN networks to be deployed on the Fabric Interconnect.
type FabricEthNetworkPolicy struct {
	PolicyAbstractPolicy
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to fabricSwitchProfile resources.
	Profiles             []FabricSwitchProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricEthNetworkPolicy FabricEthNetworkPolicy

// NewFabricEthNetworkPolicy instantiates a new FabricEthNetworkPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricEthNetworkPolicy() *FabricEthNetworkPolicy {
	this := FabricEthNetworkPolicy{}
	return &this
}

// NewFabricEthNetworkPolicyWithDefaults instantiates a new FabricEthNetworkPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricEthNetworkPolicyWithDefaults() *FabricEthNetworkPolicy {
	this := FabricEthNetworkPolicy{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricEthNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricEthNetworkPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricEthNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricEthNetworkPolicy) GetProfiles() []FabricSwitchProfileRelationship {
	if o == nil {
		var ret []FabricSwitchProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricEthNetworkPolicy) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *FabricEthNetworkPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []FabricSwitchProfileRelationship and assigns it to the Profiles field.
func (o *FabricEthNetworkPolicy) SetProfiles(v []FabricSwitchProfileRelationship) {
	o.Profiles = v
}

func (o FabricEthNetworkPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricEthNetworkPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type FabricEthNetworkPolicyWithoutEmbeddedStruct struct {
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to fabricSwitchProfile resources.
		Profiles []FabricSwitchProfileRelationship `json:"Profiles,omitempty"`
	}

	varFabricEthNetworkPolicyWithoutEmbeddedStruct := FabricEthNetworkPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricEthNetworkPolicyWithoutEmbeddedStruct)
	if err == nil {
		varFabricEthNetworkPolicy := _FabricEthNetworkPolicy{}
		varFabricEthNetworkPolicy.Organization = varFabricEthNetworkPolicyWithoutEmbeddedStruct.Organization
		varFabricEthNetworkPolicy.Profiles = varFabricEthNetworkPolicyWithoutEmbeddedStruct.Profiles
		*o = FabricEthNetworkPolicy(varFabricEthNetworkPolicy)
	} else {
		return err
	}

	varFabricEthNetworkPolicy := _FabricEthNetworkPolicy{}

	err = json.Unmarshal(bytes, &varFabricEthNetworkPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varFabricEthNetworkPolicy.PolicyAbstractPolicy
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableFabricEthNetworkPolicy struct {
	value *FabricEthNetworkPolicy
	isSet bool
}

func (v NullableFabricEthNetworkPolicy) Get() *FabricEthNetworkPolicy {
	return v.value
}

func (v *NullableFabricEthNetworkPolicy) Set(val *FabricEthNetworkPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricEthNetworkPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricEthNetworkPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricEthNetworkPolicy(val *FabricEthNetworkPolicy) *NullableFabricEthNetworkPolicy {
	return &NullableFabricEthNetworkPolicy{value: val, isSet: true}
}

func (v NullableFabricEthNetworkPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricEthNetworkPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
