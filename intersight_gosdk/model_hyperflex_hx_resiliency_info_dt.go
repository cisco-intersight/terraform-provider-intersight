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

// HyperflexHxResiliencyInfoDt struct for HyperflexHxResiliencyInfoDt
type HyperflexHxResiliencyInfoDt struct {
	MoBaseComplexType
	DataReplicationFactor *string   `json:"DataReplicationFactor,omitempty"`
	HddFailuresTolerable  *int64    `json:"HddFailuresTolerable,omitempty"`
	Messages              *[]string `json:"Messages,omitempty"`
	NodeFailuresTolerable *int64    `json:"NodeFailuresTolerable,omitempty"`
	PolicyCompliance      *string   `json:"PolicyCompliance,omitempty"`
	ResiliencyState       *string   `json:"ResiliencyState,omitempty"`
	SsdFailuresTolerable  *int64    `json:"SsdFailuresTolerable,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _HyperflexHxResiliencyInfoDt HyperflexHxResiliencyInfoDt

// NewHyperflexHxResiliencyInfoDt instantiates a new HyperflexHxResiliencyInfoDt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxResiliencyInfoDt() *HyperflexHxResiliencyInfoDt {
	this := HyperflexHxResiliencyInfoDt{}
	var dataReplicationFactor string = "ONE_COPY"
	this.DataReplicationFactor = &dataReplicationFactor
	var policyCompliance string = "UNKNOWN"
	this.PolicyCompliance = &policyCompliance
	var resiliencyState string = "UNKNOWN"
	this.ResiliencyState = &resiliencyState
	return &this
}

// NewHyperflexHxResiliencyInfoDtWithDefaults instantiates a new HyperflexHxResiliencyInfoDt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxResiliencyInfoDtWithDefaults() *HyperflexHxResiliencyInfoDt {
	this := HyperflexHxResiliencyInfoDt{}
	var dataReplicationFactor string = "ONE_COPY"
	this.DataReplicationFactor = &dataReplicationFactor
	var policyCompliance string = "UNKNOWN"
	this.PolicyCompliance = &policyCompliance
	var resiliencyState string = "UNKNOWN"
	this.ResiliencyState = &resiliencyState
	return &this
}

// GetDataReplicationFactor returns the DataReplicationFactor field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDt) GetDataReplicationFactor() string {
	if o == nil || o.DataReplicationFactor == nil {
		var ret string
		return ret
	}
	return *o.DataReplicationFactor
}

// GetDataReplicationFactorOk returns a tuple with the DataReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDt) GetDataReplicationFactorOk() (*string, bool) {
	if o == nil || o.DataReplicationFactor == nil {
		return nil, false
	}
	return o.DataReplicationFactor, true
}

// HasDataReplicationFactor returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDt) HasDataReplicationFactor() bool {
	if o != nil && o.DataReplicationFactor != nil {
		return true
	}

	return false
}

// SetDataReplicationFactor gets a reference to the given string and assigns it to the DataReplicationFactor field.
func (o *HyperflexHxResiliencyInfoDt) SetDataReplicationFactor(v string) {
	o.DataReplicationFactor = &v
}

// GetHddFailuresTolerable returns the HddFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDt) GetHddFailuresTolerable() int64 {
	if o == nil || o.HddFailuresTolerable == nil {
		var ret int64
		return ret
	}
	return *o.HddFailuresTolerable
}

// GetHddFailuresTolerableOk returns a tuple with the HddFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDt) GetHddFailuresTolerableOk() (*int64, bool) {
	if o == nil || o.HddFailuresTolerable == nil {
		return nil, false
	}
	return o.HddFailuresTolerable, true
}

// HasHddFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDt) HasHddFailuresTolerable() bool {
	if o != nil && o.HddFailuresTolerable != nil {
		return true
	}

	return false
}

// SetHddFailuresTolerable gets a reference to the given int64 and assigns it to the HddFailuresTolerable field.
func (o *HyperflexHxResiliencyInfoDt) SetHddFailuresTolerable(v int64) {
	o.HddFailuresTolerable = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDt) GetMessages() []string {
	if o == nil || o.Messages == nil {
		var ret []string
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDt) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDt) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *HyperflexHxResiliencyInfoDt) SetMessages(v []string) {
	o.Messages = &v
}

// GetNodeFailuresTolerable returns the NodeFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDt) GetNodeFailuresTolerable() int64 {
	if o == nil || o.NodeFailuresTolerable == nil {
		var ret int64
		return ret
	}
	return *o.NodeFailuresTolerable
}

// GetNodeFailuresTolerableOk returns a tuple with the NodeFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDt) GetNodeFailuresTolerableOk() (*int64, bool) {
	if o == nil || o.NodeFailuresTolerable == nil {
		return nil, false
	}
	return o.NodeFailuresTolerable, true
}

// HasNodeFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDt) HasNodeFailuresTolerable() bool {
	if o != nil && o.NodeFailuresTolerable != nil {
		return true
	}

	return false
}

// SetNodeFailuresTolerable gets a reference to the given int64 and assigns it to the NodeFailuresTolerable field.
func (o *HyperflexHxResiliencyInfoDt) SetNodeFailuresTolerable(v int64) {
	o.NodeFailuresTolerable = &v
}

// GetPolicyCompliance returns the PolicyCompliance field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDt) GetPolicyCompliance() string {
	if o == nil || o.PolicyCompliance == nil {
		var ret string
		return ret
	}
	return *o.PolicyCompliance
}

// GetPolicyComplianceOk returns a tuple with the PolicyCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDt) GetPolicyComplianceOk() (*string, bool) {
	if o == nil || o.PolicyCompliance == nil {
		return nil, false
	}
	return o.PolicyCompliance, true
}

// HasPolicyCompliance returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDt) HasPolicyCompliance() bool {
	if o != nil && o.PolicyCompliance != nil {
		return true
	}

	return false
}

// SetPolicyCompliance gets a reference to the given string and assigns it to the PolicyCompliance field.
func (o *HyperflexHxResiliencyInfoDt) SetPolicyCompliance(v string) {
	o.PolicyCompliance = &v
}

// GetResiliencyState returns the ResiliencyState field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDt) GetResiliencyState() string {
	if o == nil || o.ResiliencyState == nil {
		var ret string
		return ret
	}
	return *o.ResiliencyState
}

// GetResiliencyStateOk returns a tuple with the ResiliencyState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDt) GetResiliencyStateOk() (*string, bool) {
	if o == nil || o.ResiliencyState == nil {
		return nil, false
	}
	return o.ResiliencyState, true
}

// HasResiliencyState returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDt) HasResiliencyState() bool {
	if o != nil && o.ResiliencyState != nil {
		return true
	}

	return false
}

// SetResiliencyState gets a reference to the given string and assigns it to the ResiliencyState field.
func (o *HyperflexHxResiliencyInfoDt) SetResiliencyState(v string) {
	o.ResiliencyState = &v
}

// GetSsdFailuresTolerable returns the SsdFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDt) GetSsdFailuresTolerable() int64 {
	if o == nil || o.SsdFailuresTolerable == nil {
		var ret int64
		return ret
	}
	return *o.SsdFailuresTolerable
}

// GetSsdFailuresTolerableOk returns a tuple with the SsdFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDt) GetSsdFailuresTolerableOk() (*int64, bool) {
	if o == nil || o.SsdFailuresTolerable == nil {
		return nil, false
	}
	return o.SsdFailuresTolerable, true
}

// HasSsdFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDt) HasSsdFailuresTolerable() bool {
	if o != nil && o.SsdFailuresTolerable != nil {
		return true
	}

	return false
}

// SetSsdFailuresTolerable gets a reference to the given int64 and assigns it to the SsdFailuresTolerable field.
func (o *HyperflexHxResiliencyInfoDt) SetSsdFailuresTolerable(v int64) {
	o.SsdFailuresTolerable = &v
}

func (o HyperflexHxResiliencyInfoDt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.DataReplicationFactor != nil {
		toSerialize["DataReplicationFactor"] = o.DataReplicationFactor
	}
	if o.HddFailuresTolerable != nil {
		toSerialize["HddFailuresTolerable"] = o.HddFailuresTolerable
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.NodeFailuresTolerable != nil {
		toSerialize["NodeFailuresTolerable"] = o.NodeFailuresTolerable
	}
	if o.PolicyCompliance != nil {
		toSerialize["PolicyCompliance"] = o.PolicyCompliance
	}
	if o.ResiliencyState != nil {
		toSerialize["ResiliencyState"] = o.ResiliencyState
	}
	if o.SsdFailuresTolerable != nil {
		toSerialize["SsdFailuresTolerable"] = o.SsdFailuresTolerable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxResiliencyInfoDt) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxResiliencyInfoDtWithoutEmbeddedStruct struct {
		DataReplicationFactor *string   `json:"DataReplicationFactor,omitempty"`
		HddFailuresTolerable  *int64    `json:"HddFailuresTolerable,omitempty"`
		Messages              *[]string `json:"Messages,omitempty"`
		NodeFailuresTolerable *int64    `json:"NodeFailuresTolerable,omitempty"`
		PolicyCompliance      *string   `json:"PolicyCompliance,omitempty"`
		ResiliencyState       *string   `json:"ResiliencyState,omitempty"`
		SsdFailuresTolerable  *int64    `json:"SsdFailuresTolerable,omitempty"`
	}

	varHyperflexHxResiliencyInfoDtWithoutEmbeddedStruct := HyperflexHxResiliencyInfoDtWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxResiliencyInfoDtWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxResiliencyInfoDt := _HyperflexHxResiliencyInfoDt{}
		varHyperflexHxResiliencyInfoDt.DataReplicationFactor = varHyperflexHxResiliencyInfoDtWithoutEmbeddedStruct.DataReplicationFactor
		varHyperflexHxResiliencyInfoDt.HddFailuresTolerable = varHyperflexHxResiliencyInfoDtWithoutEmbeddedStruct.HddFailuresTolerable
		varHyperflexHxResiliencyInfoDt.Messages = varHyperflexHxResiliencyInfoDtWithoutEmbeddedStruct.Messages
		varHyperflexHxResiliencyInfoDt.NodeFailuresTolerable = varHyperflexHxResiliencyInfoDtWithoutEmbeddedStruct.NodeFailuresTolerable
		varHyperflexHxResiliencyInfoDt.PolicyCompliance = varHyperflexHxResiliencyInfoDtWithoutEmbeddedStruct.PolicyCompliance
		varHyperflexHxResiliencyInfoDt.ResiliencyState = varHyperflexHxResiliencyInfoDtWithoutEmbeddedStruct.ResiliencyState
		varHyperflexHxResiliencyInfoDt.SsdFailuresTolerable = varHyperflexHxResiliencyInfoDtWithoutEmbeddedStruct.SsdFailuresTolerable
		*o = HyperflexHxResiliencyInfoDt(varHyperflexHxResiliencyInfoDt)
	} else {
		return err
	}

	varHyperflexHxResiliencyInfoDt := _HyperflexHxResiliencyInfoDt{}

	err = json.Unmarshal(bytes, &varHyperflexHxResiliencyInfoDt)
	if err == nil {
		o.MoBaseComplexType = varHyperflexHxResiliencyInfoDt.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DataReplicationFactor")
		delete(additionalProperties, "HddFailuresTolerable")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "NodeFailuresTolerable")
		delete(additionalProperties, "PolicyCompliance")
		delete(additionalProperties, "ResiliencyState")
		delete(additionalProperties, "SsdFailuresTolerable")

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

type NullableHyperflexHxResiliencyInfoDt struct {
	value *HyperflexHxResiliencyInfoDt
	isSet bool
}

func (v NullableHyperflexHxResiliencyInfoDt) Get() *HyperflexHxResiliencyInfoDt {
	return v.value
}

func (v *NullableHyperflexHxResiliencyInfoDt) Set(val *HyperflexHxResiliencyInfoDt) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxResiliencyInfoDt) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxResiliencyInfoDt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxResiliencyInfoDt(val *HyperflexHxResiliencyInfoDt) *NullableHyperflexHxResiliencyInfoDt {
	return &NullableHyperflexHxResiliencyInfoDt{value: val, isSet: true}
}

func (v NullableHyperflexHxResiliencyInfoDt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxResiliencyInfoDt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
