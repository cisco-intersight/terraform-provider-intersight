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

// IamRoleAllOf Definition of the list of properties defined in 'iam.Role', excluding properties defined in parent classes.
type IamRoleAllOf struct {
	// Informative description about each role.
	Description *string `json:"Description,omitempty"`
	// The name of the role which has to be granted to user.
	Name           *string                 `json:"Name,omitempty"`
	PrivilegeNames *[]string               `json:"PrivilegeNames,omitempty"`
	Account        *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to iamPrivilegeSet resources.
	PrivilegeSets []IamPrivilegeSetRelationship `json:"PrivilegeSets,omitempty"`
	System        *IamSystemRelationship        `json:"System,omitempty"`
}

// NewIamRoleAllOf instantiates a new IamRoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamRoleAllOf() *IamRoleAllOf {
	this := IamRoleAllOf{}
	return &this
}

// NewIamRoleAllOfWithDefaults instantiates a new IamRoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamRoleAllOfWithDefaults() *IamRoleAllOf {
	this := IamRoleAllOf{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamRoleAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamRoleAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamRoleAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamRoleAllOf) SetName(v string) {
	o.Name = &v
}

// GetPrivilegeNames returns the PrivilegeNames field value if set, zero value otherwise.
func (o *IamRoleAllOf) GetPrivilegeNames() []string {
	if o == nil || o.PrivilegeNames == nil {
		var ret []string
		return ret
	}
	return *o.PrivilegeNames
}

// GetPrivilegeNamesOk returns a tuple with the PrivilegeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleAllOf) GetPrivilegeNamesOk() (*[]string, bool) {
	if o == nil || o.PrivilegeNames == nil {
		return nil, false
	}
	return o.PrivilegeNames, true
}

// HasPrivilegeNames returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasPrivilegeNames() bool {
	if o != nil && o.PrivilegeNames != nil {
		return true
	}

	return false
}

// SetPrivilegeNames gets a reference to the given []string and assigns it to the PrivilegeNames field.
func (o *IamRoleAllOf) SetPrivilegeNames(v []string) {
	o.PrivilegeNames = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamRoleAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamRoleAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetPrivilegeSets returns the PrivilegeSets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamRoleAllOf) GetPrivilegeSets() []IamPrivilegeSetRelationship {
	if o == nil {
		var ret []IamPrivilegeSetRelationship
		return ret
	}
	return o.PrivilegeSets
}

// GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamRoleAllOf) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool) {
	if o == nil || o.PrivilegeSets == nil {
		return nil, false
	}
	return &o.PrivilegeSets, true
}

// HasPrivilegeSets returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasPrivilegeSets() bool {
	if o != nil && o.PrivilegeSets != nil {
		return true
	}

	return false
}

// SetPrivilegeSets gets a reference to the given []IamPrivilegeSetRelationship and assigns it to the PrivilegeSets field.
func (o *IamRoleAllOf) SetPrivilegeSets(v []IamPrivilegeSetRelationship) {
	o.PrivilegeSets = v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *IamRoleAllOf) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleAllOf) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *IamRoleAllOf) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

func (o IamRoleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PrivilegeNames != nil {
		toSerialize["PrivilegeNames"] = o.PrivilegeNames
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.PrivilegeSets != nil {
		toSerialize["PrivilegeSets"] = o.PrivilegeSets
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}
	return json.Marshal(toSerialize)
}

type NullableIamRoleAllOf struct {
	value *IamRoleAllOf
	isSet bool
}

func (v NullableIamRoleAllOf) Get() *IamRoleAllOf {
	return v.value
}

func (v *NullableIamRoleAllOf) Set(val *IamRoleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamRoleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamRoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamRoleAllOf(val *IamRoleAllOf) *NullableIamRoleAllOf {
	return &NullableIamRoleAllOf{value: val, isSet: true}
}

func (v NullableIamRoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamRoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}