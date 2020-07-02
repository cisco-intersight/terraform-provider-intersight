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

// CapabilitySectionAllOf Definition of the list of properties defined in 'capability.Section', excluding properties defined in parent classes.
type CapabilitySectionAllOf struct {
	// Administrative action to initialize/load the catalog section from a particular source.
	Action *string `json:"Action,omitempty"`
	// The catalog name reference.
	CatalogName *string `json:"CatalogName,omitempty"`
	// A unique name for the section inside a catalog.
	Name *string `json:"Name,omitempty"`
	// The configured source for this section of the catalog.
	Source *string `json:"Source,omitempty"`
	// Version of the section inside a catalog.
	Version *string `json:"Version,omitempty"`
}

// NewCapabilitySectionAllOf instantiates a new CapabilitySectionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySectionAllOf() *CapabilitySectionAllOf {
	this := CapabilitySectionAllOf{}
	var action string = "None"
	this.Action = &action
	var source string = "Local"
	this.Source = &source
	return &this
}

// NewCapabilitySectionAllOfWithDefaults instantiates a new CapabilitySectionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySectionAllOfWithDefaults() *CapabilitySectionAllOf {
	this := CapabilitySectionAllOf{}
	var action string = "None"
	this.Action = &action
	var source string = "Local"
	this.Source = &source
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *CapabilitySectionAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *CapabilitySectionAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *CapabilitySectionAllOf) SetAction(v string) {
	o.Action = &v
}

// GetCatalogName returns the CatalogName field value if set, zero value otherwise.
func (o *CapabilitySectionAllOf) GetCatalogName() string {
	if o == nil || o.CatalogName == nil {
		var ret string
		return ret
	}
	return *o.CatalogName
}

// GetCatalogNameOk returns a tuple with the CatalogName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetCatalogNameOk() (*string, bool) {
	if o == nil || o.CatalogName == nil {
		return nil, false
	}
	return o.CatalogName, true
}

// HasCatalogName returns a boolean if a field has been set.
func (o *CapabilitySectionAllOf) HasCatalogName() bool {
	if o != nil && o.CatalogName != nil {
		return true
	}

	return false
}

// SetCatalogName gets a reference to the given string and assigns it to the CatalogName field.
func (o *CapabilitySectionAllOf) SetCatalogName(v string) {
	o.CatalogName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CapabilitySectionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CapabilitySectionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CapabilitySectionAllOf) SetName(v string) {
	o.Name = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CapabilitySectionAllOf) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CapabilitySectionAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CapabilitySectionAllOf) SetSource(v string) {
	o.Source = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CapabilitySectionAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CapabilitySectionAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CapabilitySectionAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o CapabilitySectionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.CatalogName != nil {
		toSerialize["CatalogName"] = o.CatalogName
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableCapabilitySectionAllOf struct {
	value *CapabilitySectionAllOf
	isSet bool
}

func (v NullableCapabilitySectionAllOf) Get() *CapabilitySectionAllOf {
	return v.value
}

func (v *NullableCapabilitySectionAllOf) Set(val *CapabilitySectionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySectionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySectionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySectionAllOf(val *CapabilitySectionAllOf) *NullableCapabilitySectionAllOf {
	return &NullableCapabilitySectionAllOf{value: val, isSet: true}
}

func (v NullableCapabilitySectionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySectionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}