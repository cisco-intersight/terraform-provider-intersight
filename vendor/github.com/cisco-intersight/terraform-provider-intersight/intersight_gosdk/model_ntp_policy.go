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

// NtpPolicy Policy to configure the NTP Servers.
type NtpPolicy struct {
	PolicyAbstractPolicy
	// State of NTP service on the endpoint.
	Enabled    *bool     `json:"Enabled,omitempty"`
	NtpServers *[]string `json:"NtpServers,omitempty"`
	// Timezone of services on the endpoint.
	Timezone         *string                               `json:"Timezone,omitempty"`
	ApplianceAccount *IamAccountRelationship               `json:"ApplianceAccount,omitempty"`
	Organization     *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NtpPolicy NtpPolicy

// NewNtpPolicy instantiates a new NtpPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtpPolicy() *NtpPolicy {
	this := NtpPolicy{}
	var timezone string = "Pacific/Niue"
	this.Timezone = &timezone
	return &this
}

// NewNtpPolicyWithDefaults instantiates a new NtpPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtpPolicyWithDefaults() *NtpPolicy {
	this := NtpPolicy{}
	var timezone string = "Pacific/Niue"
	this.Timezone = &timezone
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NtpPolicy) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NtpPolicy) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NtpPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *NtpPolicy) GetNtpServers() []string {
	if o == nil || o.NtpServers == nil {
		var ret []string
		return ret
	}
	return *o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpPolicy) GetNtpServersOk() (*[]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *NtpPolicy) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *NtpPolicy) SetNtpServers(v []string) {
	o.NtpServers = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *NtpPolicy) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpPolicy) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *NtpPolicy) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *NtpPolicy) SetTimezone(v string) {
	o.Timezone = &v
}

// GetApplianceAccount returns the ApplianceAccount field value if set, zero value otherwise.
func (o *NtpPolicy) GetApplianceAccount() IamAccountRelationship {
	if o == nil || o.ApplianceAccount == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.ApplianceAccount
}

// GetApplianceAccountOk returns a tuple with the ApplianceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpPolicy) GetApplianceAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.ApplianceAccount == nil {
		return nil, false
	}
	return o.ApplianceAccount, true
}

// HasApplianceAccount returns a boolean if a field has been set.
func (o *NtpPolicy) HasApplianceAccount() bool {
	if o != nil && o.ApplianceAccount != nil {
		return true
	}

	return false
}

// SetApplianceAccount gets a reference to the given IamAccountRelationship and assigns it to the ApplianceAccount field.
func (o *NtpPolicy) SetApplianceAccount(v IamAccountRelationship) {
	o.ApplianceAccount = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *NtpPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *NtpPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *NtpPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NtpPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NtpPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *NtpPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *NtpPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o NtpPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.NtpServers != nil {
		toSerialize["NtpServers"] = o.NtpServers
	}
	if o.Timezone != nil {
		toSerialize["Timezone"] = o.Timezone
	}
	if o.ApplianceAccount != nil {
		toSerialize["ApplianceAccount"] = o.ApplianceAccount
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

func (o *NtpPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type NtpPolicyWithoutEmbeddedStruct struct {
		// State of NTP service on the endpoint.
		Enabled    *bool     `json:"Enabled,omitempty"`
		NtpServers *[]string `json:"NtpServers,omitempty"`
		// Timezone of services on the endpoint.
		Timezone         *string                               `json:"Timezone,omitempty"`
		ApplianceAccount *IamAccountRelationship               `json:"ApplianceAccount,omitempty"`
		Organization     *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varNtpPolicyWithoutEmbeddedStruct := NtpPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNtpPolicyWithoutEmbeddedStruct)
	if err == nil {
		varNtpPolicy := _NtpPolicy{}
		varNtpPolicy.Enabled = varNtpPolicyWithoutEmbeddedStruct.Enabled
		varNtpPolicy.NtpServers = varNtpPolicyWithoutEmbeddedStruct.NtpServers
		varNtpPolicy.Timezone = varNtpPolicyWithoutEmbeddedStruct.Timezone
		varNtpPolicy.ApplianceAccount = varNtpPolicyWithoutEmbeddedStruct.ApplianceAccount
		varNtpPolicy.Organization = varNtpPolicyWithoutEmbeddedStruct.Organization
		varNtpPolicy.Profiles = varNtpPolicyWithoutEmbeddedStruct.Profiles
		*o = NtpPolicy(varNtpPolicy)
	} else {
		return err
	}

	varNtpPolicy := _NtpPolicy{}

	err = json.Unmarshal(bytes, &varNtpPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varNtpPolicy.PolicyAbstractPolicy
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "NtpServers")
		delete(additionalProperties, "Timezone")
		delete(additionalProperties, "ApplianceAccount")
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

type NullableNtpPolicy struct {
	value *NtpPolicy
	isSet bool
}

func (v NullableNtpPolicy) Get() *NtpPolicy {
	return v.value
}

func (v *NullableNtpPolicy) Set(val *NtpPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableNtpPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableNtpPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNtpPolicy(val *NtpPolicy) *NullableNtpPolicy {
	return &NullableNtpPolicy{value: val, isSet: true}
}

func (v NullableNtpPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNtpPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
