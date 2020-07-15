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

// ResourceSelector Complex type representing the permission to roles.
type ResourceSelector struct {
	MoBaseComplexType
	// ODATA filter to select resources. The group selector may include URLs of individual resource, or OData query with filters that match multiple queries. The URLs must be relative (i.e. do not include the host).
	Selector             *string `json:"Selector,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSelector ResourceSelector

// NewResourceSelector instantiates a new ResourceSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSelector() *ResourceSelector {
	this := ResourceSelector{}
	return &this
}

// NewResourceSelectorWithDefaults instantiates a new ResourceSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSelectorWithDefaults() *ResourceSelector {
	this := ResourceSelector{}
	return &this
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *ResourceSelector) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelector) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *ResourceSelector) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *ResourceSelector) SetSelector(v string) {
	o.Selector = &v
}

func (o ResourceSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Selector != nil {
		toSerialize["Selector"] = o.Selector
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSelector) UnmarshalJSON(bytes []byte) (err error) {
	type ResourceSelectorWithoutEmbeddedStruct struct {
		// ODATA filter to select resources. The group selector may include URLs of individual resource, or OData query with filters that match multiple queries. The URLs must be relative (i.e. do not include the host).
		Selector *string `json:"Selector,omitempty"`
	}

	varResourceSelectorWithoutEmbeddedStruct := ResourceSelectorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varResourceSelectorWithoutEmbeddedStruct)
	if err == nil {
		varResourceSelector := _ResourceSelector{}
		varResourceSelector.Selector = varResourceSelectorWithoutEmbeddedStruct.Selector
		*o = ResourceSelector(varResourceSelector)
	} else {
		return err
	}

	varResourceSelector := _ResourceSelector{}

	err = json.Unmarshal(bytes, &varResourceSelector)
	if err == nil {
		o.MoBaseComplexType = varResourceSelector.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Selector")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableResourceSelector struct {
	value *ResourceSelector
	isSet bool
}

func (v NullableResourceSelector) Get() *ResourceSelector {
	return v.value
}

func (v *NullableResourceSelector) Set(val *ResourceSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSelector(val *ResourceSelector) *NullableResourceSelector {
	return &NullableResourceSelector{value: val, isSet: true}
}

func (v NullableResourceSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
