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

// HyperflexHxdpVersion A HyperFlex Data Platform version.
type HyperflexHxdpVersion struct {
	MoBaseMo
	// The HyperFlex Data Platform version.
	Version              *string                          `json:"Version,omitempty"`
	AppCatalog           *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxdpVersion HyperflexHxdpVersion

// NewHyperflexHxdpVersion instantiates a new HyperflexHxdpVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxdpVersion() *HyperflexHxdpVersion {
	this := HyperflexHxdpVersion{}
	return &this
}

// NewHyperflexHxdpVersionWithDefaults instantiates a new HyperflexHxdpVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxdpVersionWithDefaults() *HyperflexHxdpVersion {
	this := HyperflexHxdpVersion{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexHxdpVersion) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxdpVersion) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexHxdpVersion) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexHxdpVersion) SetVersion(v string) {
	o.Version = &v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise.
func (o *HyperflexHxdpVersion) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || o.AppCatalog == nil {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxdpVersion) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil || o.AppCatalog == nil {
		return nil, false
	}
	return o.AppCatalog, true
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HyperflexHxdpVersion) HasAppCatalog() bool {
	if o != nil && o.AppCatalog != nil {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given HyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HyperflexHxdpVersion) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog = &v
}

func (o HyperflexHxdpVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.AppCatalog != nil {
		toSerialize["AppCatalog"] = o.AppCatalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxdpVersion) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxdpVersionWithoutEmbeddedStruct struct {
		// The HyperFlex Data Platform version.
		Version    *string                          `json:"Version,omitempty"`
		AppCatalog *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	}

	varHyperflexHxdpVersionWithoutEmbeddedStruct := HyperflexHxdpVersionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxdpVersionWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxdpVersion := _HyperflexHxdpVersion{}
		varHyperflexHxdpVersion.Version = varHyperflexHxdpVersionWithoutEmbeddedStruct.Version
		varHyperflexHxdpVersion.AppCatalog = varHyperflexHxdpVersionWithoutEmbeddedStruct.AppCatalog
		*o = HyperflexHxdpVersion(varHyperflexHxdpVersion)
	} else {
		return err
	}

	varHyperflexHxdpVersion := _HyperflexHxdpVersion{}

	err = json.Unmarshal(bytes, &varHyperflexHxdpVersion)
	if err == nil {
		o.MoBaseMo = varHyperflexHxdpVersion.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Version")
		delete(additionalProperties, "AppCatalog")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHyperflexHxdpVersion struct {
	value *HyperflexHxdpVersion
	isSet bool
}

func (v NullableHyperflexHxdpVersion) Get() *HyperflexHxdpVersion {
	return v.value
}

func (v *NullableHyperflexHxdpVersion) Set(val *HyperflexHxdpVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxdpVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxdpVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxdpVersion(val *HyperflexHxdpVersion) *NullableHyperflexHxdpVersion {
	return &NullableHyperflexHxdpVersion{value: val, isSet: true}
}

func (v NullableHyperflexHxdpVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxdpVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
