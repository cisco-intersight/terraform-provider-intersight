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

// FeedbackFeedbackPostAllOf Definition of the list of properties defined in 'feedback.FeedbackPost', excluding properties defined in parent classes.
type FeedbackFeedbackPostAllOf struct {
	FeedbackData         *FeedbackFeedbackData   `json:"FeedbackData,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeedbackFeedbackPostAllOf FeedbackFeedbackPostAllOf

// NewFeedbackFeedbackPostAllOf instantiates a new FeedbackFeedbackPostAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackFeedbackPostAllOf() *FeedbackFeedbackPostAllOf {
	this := FeedbackFeedbackPostAllOf{}
	return &this
}

// NewFeedbackFeedbackPostAllOfWithDefaults instantiates a new FeedbackFeedbackPostAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackFeedbackPostAllOfWithDefaults() *FeedbackFeedbackPostAllOf {
	this := FeedbackFeedbackPostAllOf{}
	return &this
}

// GetFeedbackData returns the FeedbackData field value if set, zero value otherwise.
func (o *FeedbackFeedbackPostAllOf) GetFeedbackData() FeedbackFeedbackData {
	if o == nil || o.FeedbackData == nil {
		var ret FeedbackFeedbackData
		return ret
	}
	return *o.FeedbackData
}

// GetFeedbackDataOk returns a tuple with the FeedbackData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackPostAllOf) GetFeedbackDataOk() (*FeedbackFeedbackData, bool) {
	if o == nil || o.FeedbackData == nil {
		return nil, false
	}
	return o.FeedbackData, true
}

// HasFeedbackData returns a boolean if a field has been set.
func (o *FeedbackFeedbackPostAllOf) HasFeedbackData() bool {
	if o != nil && o.FeedbackData != nil {
		return true
	}

	return false
}

// SetFeedbackData gets a reference to the given FeedbackFeedbackData and assigns it to the FeedbackData field.
func (o *FeedbackFeedbackPostAllOf) SetFeedbackData(v FeedbackFeedbackData) {
	o.FeedbackData = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *FeedbackFeedbackPostAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackPostAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *FeedbackFeedbackPostAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *FeedbackFeedbackPostAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o FeedbackFeedbackPostAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeedbackData != nil {
		toSerialize["FeedbackData"] = o.FeedbackData
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FeedbackFeedbackPostAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFeedbackFeedbackPostAllOf := _FeedbackFeedbackPostAllOf{}

	if err = json.Unmarshal(bytes, &varFeedbackFeedbackPostAllOf); err == nil {
		*o = FeedbackFeedbackPostAllOf(varFeedbackFeedbackPostAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "FeedbackData")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeedbackFeedbackPostAllOf struct {
	value *FeedbackFeedbackPostAllOf
	isSet bool
}

func (v NullableFeedbackFeedbackPostAllOf) Get() *FeedbackFeedbackPostAllOf {
	return v.value
}

func (v *NullableFeedbackFeedbackPostAllOf) Set(val *FeedbackFeedbackPostAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackFeedbackPostAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackFeedbackPostAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackFeedbackPostAllOf(val *FeedbackFeedbackPostAllOf) *NullableFeedbackFeedbackPostAllOf {
	return &NullableFeedbackFeedbackPostAllOf{value: val, isSet: true}
}

func (v NullableFeedbackFeedbackPostAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackFeedbackPostAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
