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

// HyperflexFeatureLimitInternalAllOf Definition of the list of properties defined in 'hyperflex.FeatureLimitInternal', excluding properties defined in parent classes.
type HyperflexFeatureLimitInternalAllOf struct {
	FeatureLimitEntries  *[]HyperflexFeatureLimitEntry    `json:"FeatureLimitEntries,omitempty"`
	AppCatalog           *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexFeatureLimitInternalAllOf HyperflexFeatureLimitInternalAllOf

// NewHyperflexFeatureLimitInternalAllOf instantiates a new HyperflexFeatureLimitInternalAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexFeatureLimitInternalAllOf() *HyperflexFeatureLimitInternalAllOf {
	this := HyperflexFeatureLimitInternalAllOf{}
	return &this
}

// NewHyperflexFeatureLimitInternalAllOfWithDefaults instantiates a new HyperflexFeatureLimitInternalAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexFeatureLimitInternalAllOfWithDefaults() *HyperflexFeatureLimitInternalAllOf {
	this := HyperflexFeatureLimitInternalAllOf{}
	return &this
}

// GetFeatureLimitEntries returns the FeatureLimitEntries field value if set, zero value otherwise.
func (o *HyperflexFeatureLimitInternalAllOf) GetFeatureLimitEntries() []HyperflexFeatureLimitEntry {
	if o == nil || o.FeatureLimitEntries == nil {
		var ret []HyperflexFeatureLimitEntry
		return ret
	}
	return *o.FeatureLimitEntries
}

// GetFeatureLimitEntriesOk returns a tuple with the FeatureLimitEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexFeatureLimitInternalAllOf) GetFeatureLimitEntriesOk() (*[]HyperflexFeatureLimitEntry, bool) {
	if o == nil || o.FeatureLimitEntries == nil {
		return nil, false
	}
	return o.FeatureLimitEntries, true
}

// HasFeatureLimitEntries returns a boolean if a field has been set.
func (o *HyperflexFeatureLimitInternalAllOf) HasFeatureLimitEntries() bool {
	if o != nil && o.FeatureLimitEntries != nil {
		return true
	}

	return false
}

// SetFeatureLimitEntries gets a reference to the given []HyperflexFeatureLimitEntry and assigns it to the FeatureLimitEntries field.
func (o *HyperflexFeatureLimitInternalAllOf) SetFeatureLimitEntries(v []HyperflexFeatureLimitEntry) {
	o.FeatureLimitEntries = &v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise.
func (o *HyperflexFeatureLimitInternalAllOf) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || o.AppCatalog == nil {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexFeatureLimitInternalAllOf) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil || o.AppCatalog == nil {
		return nil, false
	}
	return o.AppCatalog, true
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HyperflexFeatureLimitInternalAllOf) HasAppCatalog() bool {
	if o != nil && o.AppCatalog != nil {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given HyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HyperflexFeatureLimitInternalAllOf) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog = &v
}

func (o HyperflexFeatureLimitInternalAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureLimitEntries != nil {
		toSerialize["FeatureLimitEntries"] = o.FeatureLimitEntries
	}
	if o.AppCatalog != nil {
		toSerialize["AppCatalog"] = o.AppCatalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexFeatureLimitInternalAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexFeatureLimitInternalAllOf := _HyperflexFeatureLimitInternalAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexFeatureLimitInternalAllOf); err == nil {
		*o = HyperflexFeatureLimitInternalAllOf(varHyperflexFeatureLimitInternalAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "FeatureLimitEntries")
		delete(additionalProperties, "AppCatalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexFeatureLimitInternalAllOf struct {
	value *HyperflexFeatureLimitInternalAllOf
	isSet bool
}

func (v NullableHyperflexFeatureLimitInternalAllOf) Get() *HyperflexFeatureLimitInternalAllOf {
	return v.value
}

func (v *NullableHyperflexFeatureLimitInternalAllOf) Set(val *HyperflexFeatureLimitInternalAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexFeatureLimitInternalAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexFeatureLimitInternalAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexFeatureLimitInternalAllOf(val *HyperflexFeatureLimitInternalAllOf) *NullableHyperflexFeatureLimitInternalAllOf {
	return &NullableHyperflexFeatureLimitInternalAllOf{value: val, isSet: true}
}

func (v NullableHyperflexFeatureLimitInternalAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexFeatureLimitInternalAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
