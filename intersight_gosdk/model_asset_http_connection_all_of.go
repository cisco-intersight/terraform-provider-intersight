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

// AssetHttpConnectionAllOf Definition of the list of properties defined in 'asset.HttpConnection', excluding properties defined in parent classes.
type AssetHttpConnectionAllOf struct {
	Credential *AssetCredential `json:"Credential,omitempty"`
	// Indicates whether a connection to the target should be established using HTTPS.
	IsSecure *bool `json:"IsSecure,omitempty"`
	// The DNS hostname or IP Address, either IPv4 or IPv6, to be used to connect to the managed target.
	ManagementAddress *string `json:"ManagementAddress,omitempty"`
	// The port number to be used to to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection.
	Port                 *int64 `json:"Port,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetHttpConnectionAllOf AssetHttpConnectionAllOf

// NewAssetHttpConnectionAllOf instantiates a new AssetHttpConnectionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetHttpConnectionAllOf() *AssetHttpConnectionAllOf {
	this := AssetHttpConnectionAllOf{}
	return &this
}

// NewAssetHttpConnectionAllOfWithDefaults instantiates a new AssetHttpConnectionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetHttpConnectionAllOfWithDefaults() *AssetHttpConnectionAllOf {
	this := AssetHttpConnectionAllOf{}
	return &this
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *AssetHttpConnectionAllOf) GetCredential() AssetCredential {
	if o == nil || o.Credential == nil {
		var ret AssetCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetHttpConnectionAllOf) GetCredentialOk() (*AssetCredential, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *AssetHttpConnectionAllOf) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given AssetCredential and assigns it to the Credential field.
func (o *AssetHttpConnectionAllOf) SetCredential(v AssetCredential) {
	o.Credential = &v
}

// GetIsSecure returns the IsSecure field value if set, zero value otherwise.
func (o *AssetHttpConnectionAllOf) GetIsSecure() bool {
	if o == nil || o.IsSecure == nil {
		var ret bool
		return ret
	}
	return *o.IsSecure
}

// GetIsSecureOk returns a tuple with the IsSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetHttpConnectionAllOf) GetIsSecureOk() (*bool, bool) {
	if o == nil || o.IsSecure == nil {
		return nil, false
	}
	return o.IsSecure, true
}

// HasIsSecure returns a boolean if a field has been set.
func (o *AssetHttpConnectionAllOf) HasIsSecure() bool {
	if o != nil && o.IsSecure != nil {
		return true
	}

	return false
}

// SetIsSecure gets a reference to the given bool and assigns it to the IsSecure field.
func (o *AssetHttpConnectionAllOf) SetIsSecure(v bool) {
	o.IsSecure = &v
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *AssetHttpConnectionAllOf) GetManagementAddress() string {
	if o == nil || o.ManagementAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetHttpConnectionAllOf) GetManagementAddressOk() (*string, bool) {
	if o == nil || o.ManagementAddress == nil {
		return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *AssetHttpConnectionAllOf) HasManagementAddress() bool {
	if o != nil && o.ManagementAddress != nil {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *AssetHttpConnectionAllOf) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *AssetHttpConnectionAllOf) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetHttpConnectionAllOf) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *AssetHttpConnectionAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *AssetHttpConnectionAllOf) SetPort(v int64) {
	o.Port = &v
}

func (o AssetHttpConnectionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credential != nil {
		toSerialize["Credential"] = o.Credential
	}
	if o.IsSecure != nil {
		toSerialize["IsSecure"] = o.IsSecure
	}
	if o.ManagementAddress != nil {
		toSerialize["ManagementAddress"] = o.ManagementAddress
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetHttpConnectionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetHttpConnectionAllOf := _AssetHttpConnectionAllOf{}

	if err = json.Unmarshal(bytes, &varAssetHttpConnectionAllOf); err == nil {
		*o = AssetHttpConnectionAllOf(varAssetHttpConnectionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Credential")
		delete(additionalProperties, "IsSecure")
		delete(additionalProperties, "ManagementAddress")
		delete(additionalProperties, "Port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetHttpConnectionAllOf struct {
	value *AssetHttpConnectionAllOf
	isSet bool
}

func (v NullableAssetHttpConnectionAllOf) Get() *AssetHttpConnectionAllOf {
	return v.value
}

func (v *NullableAssetHttpConnectionAllOf) Set(val *AssetHttpConnectionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetHttpConnectionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetHttpConnectionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetHttpConnectionAllOf(val *AssetHttpConnectionAllOf) *NullableAssetHttpConnectionAllOf {
	return &NullableAssetHttpConnectionAllOf{value: val, isSet: true}
}

func (v NullableAssetHttpConnectionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetHttpConnectionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
