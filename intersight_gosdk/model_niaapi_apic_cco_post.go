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

// NiaapiApicCcoPost The post reporting a new release is available for APIC.
type NiaapiApicCcoPost struct {
	NiaapiNewReleasePost
	AdditionalProperties map[string]interface{}
}

type _NiaapiApicCcoPost NiaapiApicCcoPost

// NewNiaapiApicCcoPost instantiates a new NiaapiApicCcoPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiApicCcoPost() *NiaapiApicCcoPost {
	this := NiaapiApicCcoPost{}
	return &this
}

// NewNiaapiApicCcoPostWithDefaults instantiates a new NiaapiApicCcoPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiApicCcoPostWithDefaults() *NiaapiApicCcoPost {
	this := NiaapiApicCcoPost{}
	return &this
}

func (o NiaapiApicCcoPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNiaapiNewReleasePost, errNiaapiNewReleasePost := json.Marshal(o.NiaapiNewReleasePost)
	if errNiaapiNewReleasePost != nil {
		return []byte{}, errNiaapiNewReleasePost
	}
	errNiaapiNewReleasePost = json.Unmarshal([]byte(serializedNiaapiNewReleasePost), &toSerialize)
	if errNiaapiNewReleasePost != nil {
		return []byte{}, errNiaapiNewReleasePost
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiApicCcoPost) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiApicCcoPostWithoutEmbeddedStruct struct {
	}

	varNiaapiApicCcoPostWithoutEmbeddedStruct := NiaapiApicCcoPostWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiApicCcoPostWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiApicCcoPost := _NiaapiApicCcoPost{}
		*o = NiaapiApicCcoPost(varNiaapiApicCcoPost)
	} else {
		return err
	}

	varNiaapiApicCcoPost := _NiaapiApicCcoPost{}

	err = json.Unmarshal(bytes, &varNiaapiApicCcoPost)
	if err == nil {
		o.NiaapiNewReleasePost = varNiaapiApicCcoPost.NiaapiNewReleasePost
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectNiaapiNewReleasePost := reflect.ValueOf(o.NiaapiNewReleasePost)
		for i := 0; i < reflectNiaapiNewReleasePost.Type().NumField(); i++ {
			t := reflectNiaapiNewReleasePost.Type().Field(i)

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

type NullableNiaapiApicCcoPost struct {
	value *NiaapiApicCcoPost
	isSet bool
}

func (v NullableNiaapiApicCcoPost) Get() *NiaapiApicCcoPost {
	return v.value
}

func (v *NullableNiaapiApicCcoPost) Set(val *NiaapiApicCcoPost) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiApicCcoPost) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiApicCcoPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiApicCcoPost(val *NiaapiApicCcoPost) *NullableNiaapiApicCcoPost {
	return &NullableNiaapiApicCcoPost{value: val, isSet: true}
}

func (v NullableNiaapiApicCcoPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiApicCcoPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
