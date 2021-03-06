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

// AssetServiceAllOf Definition of the list of properties defined in 'asset.Service', excluding properties defined in parent classes.
type AssetServiceAllOf struct {
	Options *AssetServiceOptions `json:"Options,omitempty"`
	// Status indicates if the respective Service can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target.
	Status *string `json:"Status,omitempty"`
	// When 'Status' is not Connected, statusErrorReason provides further details about why the device is not connected with Intersight.
	StatusErrorReason    *string `json:"StatusErrorReason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetServiceAllOf AssetServiceAllOf

// NewAssetServiceAllOf instantiates a new AssetServiceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetServiceAllOf() *AssetServiceAllOf {
	this := AssetServiceAllOf{}
	var status string = ""
	this.Status = &status
	return &this
}

// NewAssetServiceAllOfWithDefaults instantiates a new AssetServiceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetServiceAllOfWithDefaults() *AssetServiceAllOf {
	this := AssetServiceAllOf{}
	var status string = ""
	this.Status = &status
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AssetServiceAllOf) GetOptions() AssetServiceOptions {
	if o == nil || o.Options == nil {
		var ret AssetServiceOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetServiceAllOf) GetOptionsOk() (*AssetServiceOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AssetServiceAllOf) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AssetServiceOptions and assigns it to the Options field.
func (o *AssetServiceAllOf) SetOptions(v AssetServiceOptions) {
	o.Options = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetServiceAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetServiceAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetServiceAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssetServiceAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusErrorReason returns the StatusErrorReason field value if set, zero value otherwise.
func (o *AssetServiceAllOf) GetStatusErrorReason() string {
	if o == nil || o.StatusErrorReason == nil {
		var ret string
		return ret
	}
	return *o.StatusErrorReason
}

// GetStatusErrorReasonOk returns a tuple with the StatusErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetServiceAllOf) GetStatusErrorReasonOk() (*string, bool) {
	if o == nil || o.StatusErrorReason == nil {
		return nil, false
	}
	return o.StatusErrorReason, true
}

// HasStatusErrorReason returns a boolean if a field has been set.
func (o *AssetServiceAllOf) HasStatusErrorReason() bool {
	if o != nil && o.StatusErrorReason != nil {
		return true
	}

	return false
}

// SetStatusErrorReason gets a reference to the given string and assigns it to the StatusErrorReason field.
func (o *AssetServiceAllOf) SetStatusErrorReason(v string) {
	o.StatusErrorReason = &v
}

func (o AssetServiceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Options != nil {
		toSerialize["Options"] = o.Options
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusErrorReason != nil {
		toSerialize["StatusErrorReason"] = o.StatusErrorReason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetServiceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetServiceAllOf := _AssetServiceAllOf{}

	if err = json.Unmarshal(bytes, &varAssetServiceAllOf); err == nil {
		*o = AssetServiceAllOf(varAssetServiceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Options")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusErrorReason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetServiceAllOf struct {
	value *AssetServiceAllOf
	isSet bool
}

func (v NullableAssetServiceAllOf) Get() *AssetServiceAllOf {
	return v.value
}

func (v *NullableAssetServiceAllOf) Set(val *AssetServiceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetServiceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetServiceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetServiceAllOf(val *AssetServiceAllOf) *NullableAssetServiceAllOf {
	return &NullableAssetServiceAllOf{value: val, isSet: true}
}

func (v NullableAssetServiceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetServiceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
