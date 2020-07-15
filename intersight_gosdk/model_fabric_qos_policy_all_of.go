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

// FabricQosPolicyAllOf Definition of the list of properties defined in 'fabric.QosPolicy', excluding properties defined in parent classes.
type FabricQosPolicyAllOf struct {
	Classes      *[]FabricQosClass                     `json:"Classes,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to fabricSwitchProfile resources.
	Profiles             []FabricSwitchProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricQosPolicyAllOf FabricQosPolicyAllOf

// NewFabricQosPolicyAllOf instantiates a new FabricQosPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricQosPolicyAllOf() *FabricQosPolicyAllOf {
	this := FabricQosPolicyAllOf{}
	return &this
}

// NewFabricQosPolicyAllOfWithDefaults instantiates a new FabricQosPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricQosPolicyAllOfWithDefaults() *FabricQosPolicyAllOf {
	this := FabricQosPolicyAllOf{}
	return &this
}

// GetClasses returns the Classes field value if set, zero value otherwise.
func (o *FabricQosPolicyAllOf) GetClasses() []FabricQosClass {
	if o == nil || o.Classes == nil {
		var ret []FabricQosClass
		return ret
	}
	return *o.Classes
}

// GetClassesOk returns a tuple with the Classes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosPolicyAllOf) GetClassesOk() (*[]FabricQosClass, bool) {
	if o == nil || o.Classes == nil {
		return nil, false
	}
	return o.Classes, true
}

// HasClasses returns a boolean if a field has been set.
func (o *FabricQosPolicyAllOf) HasClasses() bool {
	if o != nil && o.Classes != nil {
		return true
	}

	return false
}

// SetClasses gets a reference to the given []FabricQosClass and assigns it to the Classes field.
func (o *FabricQosPolicyAllOf) SetClasses(v []FabricQosClass) {
	o.Classes = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricQosPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricQosPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricQosPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricQosPolicyAllOf) GetProfiles() []FabricSwitchProfileRelationship {
	if o == nil {
		var ret []FabricSwitchProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricQosPolicyAllOf) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *FabricQosPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []FabricSwitchProfileRelationship and assigns it to the Profiles field.
func (o *FabricQosPolicyAllOf) SetProfiles(v []FabricSwitchProfileRelationship) {
	o.Profiles = v
}

func (o FabricQosPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Classes != nil {
		toSerialize["Classes"] = o.Classes
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

func (o *FabricQosPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricQosPolicyAllOf := _FabricQosPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varFabricQosPolicyAllOf); err == nil {
		*o = FabricQosPolicyAllOf(varFabricQosPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Classes")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricQosPolicyAllOf struct {
	value *FabricQosPolicyAllOf
	isSet bool
}

func (v NullableFabricQosPolicyAllOf) Get() *FabricQosPolicyAllOf {
	return v.value
}

func (v *NullableFabricQosPolicyAllOf) Set(val *FabricQosPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricQosPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricQosPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricQosPolicyAllOf(val *FabricQosPolicyAllOf) *NullableFabricQosPolicyAllOf {
	return &NullableFabricQosPolicyAllOf{value: val, isSet: true}
}

func (v NullableFabricQosPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricQosPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
