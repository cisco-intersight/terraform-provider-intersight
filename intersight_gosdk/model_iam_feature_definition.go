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

// IamFeatureDefinition Feature parameters specified for the account.
type IamFeatureDefinition struct {
	MoBaseComplexType
	// The beta feature that will be enabled for specific account.
	Feature              *string `json:"Feature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamFeatureDefinition IamFeatureDefinition

// NewIamFeatureDefinition instantiates a new IamFeatureDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamFeatureDefinition() *IamFeatureDefinition {
	this := IamFeatureDefinition{}
	var feature string = "IWO"
	this.Feature = &feature
	return &this
}

// NewIamFeatureDefinitionWithDefaults instantiates a new IamFeatureDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamFeatureDefinitionWithDefaults() *IamFeatureDefinition {
	this := IamFeatureDefinition{}
	var feature string = "IWO"
	this.Feature = &feature
	return &this
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *IamFeatureDefinition) GetFeature() string {
	if o == nil || o.Feature == nil {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamFeatureDefinition) GetFeatureOk() (*string, bool) {
	if o == nil || o.Feature == nil {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *IamFeatureDefinition) HasFeature() bool {
	if o != nil && o.Feature != nil {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *IamFeatureDefinition) SetFeature(v string) {
	o.Feature = &v
}

func (o IamFeatureDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Feature != nil {
		toSerialize["Feature"] = o.Feature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamFeatureDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type IamFeatureDefinitionWithoutEmbeddedStruct struct {
		// The beta feature that will be enabled for specific account.
		Feature *string `json:"Feature,omitempty"`
	}

	varIamFeatureDefinitionWithoutEmbeddedStruct := IamFeatureDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamFeatureDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varIamFeatureDefinition := _IamFeatureDefinition{}
		varIamFeatureDefinition.Feature = varIamFeatureDefinitionWithoutEmbeddedStruct.Feature
		*o = IamFeatureDefinition(varIamFeatureDefinition)
	} else {
		return err
	}

	varIamFeatureDefinition := _IamFeatureDefinition{}

	err = json.Unmarshal(bytes, &varIamFeatureDefinition)
	if err == nil {
		o.MoBaseComplexType = varIamFeatureDefinition.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Feature")

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

type NullableIamFeatureDefinition struct {
	value *IamFeatureDefinition
	isSet bool
}

func (v NullableIamFeatureDefinition) Get() *IamFeatureDefinition {
	return v.value
}

func (v *NullableIamFeatureDefinition) Set(val *IamFeatureDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableIamFeatureDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableIamFeatureDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamFeatureDefinition(val *IamFeatureDefinition) *NullableIamFeatureDefinition {
	return &NullableIamFeatureDefinition{value: val, isSet: true}
}

func (v NullableIamFeatureDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamFeatureDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
