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

// FirmwareDistributable An image distributed by Cisco.
type FirmwareDistributable struct {
	FirmwareBaseDistributable
	Catalog              *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareDistributable FirmwareDistributable

// NewFirmwareDistributable instantiates a new FirmwareDistributable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareDistributable() *FirmwareDistributable {
	this := FirmwareDistributable{}
	return &this
}

// NewFirmwareDistributableWithDefaults instantiates a new FirmwareDistributable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareDistributableWithDefaults() *FirmwareDistributable {
	this := FirmwareDistributable{}
	return &this
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *FirmwareDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *FirmwareDistributable) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *FirmwareDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o FirmwareDistributable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareBaseDistributable, errFirmwareBaseDistributable := json.Marshal(o.FirmwareBaseDistributable)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	errFirmwareBaseDistributable = json.Unmarshal([]byte(serializedFirmwareBaseDistributable), &toSerialize)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareDistributable) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareDistributableWithoutEmbeddedStruct struct {
		Catalog *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	}

	varFirmwareDistributableWithoutEmbeddedStruct := FirmwareDistributableWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareDistributableWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareDistributable := _FirmwareDistributable{}
		varFirmwareDistributable.Catalog = varFirmwareDistributableWithoutEmbeddedStruct.Catalog
		*o = FirmwareDistributable(varFirmwareDistributable)
	} else {
		return err
	}

	varFirmwareDistributable := _FirmwareDistributable{}

	err = json.Unmarshal(bytes, &varFirmwareDistributable)
	if err == nil {
		o.FirmwareBaseDistributable = varFirmwareDistributable.FirmwareBaseDistributable
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Catalog")

		// remove fields from embedded structs
		reflectFirmwareBaseDistributable := reflect.ValueOf(o.FirmwareBaseDistributable)
		for i := 0; i < reflectFirmwareBaseDistributable.Type().NumField(); i++ {
			t := reflectFirmwareBaseDistributable.Type().Field(i)

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

type NullableFirmwareDistributable struct {
	value *FirmwareDistributable
	isSet bool
}

func (v NullableFirmwareDistributable) Get() *FirmwareDistributable {
	return v.value
}

func (v *NullableFirmwareDistributable) Set(val *FirmwareDistributable) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareDistributable) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareDistributable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareDistributable(val *FirmwareDistributable) *NullableFirmwareDistributable {
	return &NullableFirmwareDistributable{value: val, isSet: true}
}

func (v NullableFirmwareDistributable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareDistributable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
