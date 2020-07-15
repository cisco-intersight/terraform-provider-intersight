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

// AssetCloudConnectionAllOf Definition of the list of properties defined in 'asset.CloudConnection', excluding properties defined in parent classes.
type AssetCloudConnectionAllOf struct {
	Credential           *AssetCredential `json:"Credential,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetCloudConnectionAllOf AssetCloudConnectionAllOf

// NewAssetCloudConnectionAllOf instantiates a new AssetCloudConnectionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetCloudConnectionAllOf() *AssetCloudConnectionAllOf {
	this := AssetCloudConnectionAllOf{}
	return &this
}

// NewAssetCloudConnectionAllOfWithDefaults instantiates a new AssetCloudConnectionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetCloudConnectionAllOfWithDefaults() *AssetCloudConnectionAllOf {
	this := AssetCloudConnectionAllOf{}
	return &this
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *AssetCloudConnectionAllOf) GetCredential() AssetCredential {
	if o == nil || o.Credential == nil {
		var ret AssetCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCloudConnectionAllOf) GetCredentialOk() (*AssetCredential, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *AssetCloudConnectionAllOf) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given AssetCredential and assigns it to the Credential field.
func (o *AssetCloudConnectionAllOf) SetCredential(v AssetCredential) {
	o.Credential = &v
}

func (o AssetCloudConnectionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credential != nil {
		toSerialize["Credential"] = o.Credential
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetCloudConnectionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetCloudConnectionAllOf := _AssetCloudConnectionAllOf{}

	if err = json.Unmarshal(bytes, &varAssetCloudConnectionAllOf); err == nil {
		*o = AssetCloudConnectionAllOf(varAssetCloudConnectionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Credential")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetCloudConnectionAllOf struct {
	value *AssetCloudConnectionAllOf
	isSet bool
}

func (v NullableAssetCloudConnectionAllOf) Get() *AssetCloudConnectionAllOf {
	return v.value
}

func (v *NullableAssetCloudConnectionAllOf) Set(val *AssetCloudConnectionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetCloudConnectionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetCloudConnectionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetCloudConnectionAllOf(val *AssetCloudConnectionAllOf) *NullableAssetCloudConnectionAllOf {
	return &NullableAssetCloudConnectionAllOf{value: val, isSet: true}
}

func (v NullableAssetCloudConnectionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetCloudConnectionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
