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

// FabricQosClassAllOf Definition of the list of properties defined in 'fabric.QosClass', excluding properties defined in parent classes.
type FabricQosClassAllOf struct {
	// Administrative state for this QoS class.
	AdminState *string `json:"AdminState,omitempty"`
	// Percentage of bandwidth received by the traffic tagged with this QoS.
	BandwidthPercent *int64 `json:"BandwidthPercent,omitempty"`
	// Class of service received by the traffic tagged with this QoS.
	Cos *int64 `json:"Cos,omitempty"`
	// Maximum transmission unit (MTU) is the largest size packet or frame, that can be sent in a packet- or frame-based network such as the Internet. User can select from the following: 1. Any value between 1500 and 9216 2. 'Normal' (default) mapping to a value of 1500. 3. 'FC' mapping to a value of 2240.
	Mtu *int64 `json:"Mtu,omitempty"`
	// If enabled, this QoS class will be optimized to send multiple packets.
	MulticastOptimize *bool `json:"MulticastOptimize,omitempty"`
	// The 'name' of this QoS Class.
	Name *string `json:"Name,omitempty"`
	// If enabled, this QoS class will allow packet drops within an acceptable limit.
	PacketDrop *bool `json:"PacketDrop,omitempty"`
	// The weight of the QoS Class controls the distribution of bandwidth between QoS Classes, with the same priority after the Guarantees for the QoS Classes are reached.
	Weight               *int64 `json:"Weight,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricQosClassAllOf FabricQosClassAllOf

// NewFabricQosClassAllOf instantiates a new FabricQosClassAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricQosClassAllOf() *FabricQosClassAllOf {
	this := FabricQosClassAllOf{}
	var adminState string = "Disabled"
	this.AdminState = &adminState
	return &this
}

// NewFabricQosClassAllOfWithDefaults instantiates a new FabricQosClassAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricQosClassAllOfWithDefaults() *FabricQosClassAllOf {
	this := FabricQosClassAllOf{}
	var adminState string = "Disabled"
	this.AdminState = &adminState
	return &this
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *FabricQosClassAllOf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClassAllOf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *FabricQosClassAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *FabricQosClassAllOf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetBandwidthPercent returns the BandwidthPercent field value if set, zero value otherwise.
func (o *FabricQosClassAllOf) GetBandwidthPercent() int64 {
	if o == nil || o.BandwidthPercent == nil {
		var ret int64
		return ret
	}
	return *o.BandwidthPercent
}

// GetBandwidthPercentOk returns a tuple with the BandwidthPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClassAllOf) GetBandwidthPercentOk() (*int64, bool) {
	if o == nil || o.BandwidthPercent == nil {
		return nil, false
	}
	return o.BandwidthPercent, true
}

// HasBandwidthPercent returns a boolean if a field has been set.
func (o *FabricQosClassAllOf) HasBandwidthPercent() bool {
	if o != nil && o.BandwidthPercent != nil {
		return true
	}

	return false
}

// SetBandwidthPercent gets a reference to the given int64 and assigns it to the BandwidthPercent field.
func (o *FabricQosClassAllOf) SetBandwidthPercent(v int64) {
	o.BandwidthPercent = &v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *FabricQosClassAllOf) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClassAllOf) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *FabricQosClassAllOf) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *FabricQosClassAllOf) SetCos(v int64) {
	o.Cos = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *FabricQosClassAllOf) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClassAllOf) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *FabricQosClassAllOf) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *FabricQosClassAllOf) SetMtu(v int64) {
	o.Mtu = &v
}

// GetMulticastOptimize returns the MulticastOptimize field value if set, zero value otherwise.
func (o *FabricQosClassAllOf) GetMulticastOptimize() bool {
	if o == nil || o.MulticastOptimize == nil {
		var ret bool
		return ret
	}
	return *o.MulticastOptimize
}

// GetMulticastOptimizeOk returns a tuple with the MulticastOptimize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClassAllOf) GetMulticastOptimizeOk() (*bool, bool) {
	if o == nil || o.MulticastOptimize == nil {
		return nil, false
	}
	return o.MulticastOptimize, true
}

// HasMulticastOptimize returns a boolean if a field has been set.
func (o *FabricQosClassAllOf) HasMulticastOptimize() bool {
	if o != nil && o.MulticastOptimize != nil {
		return true
	}

	return false
}

// SetMulticastOptimize gets a reference to the given bool and assigns it to the MulticastOptimize field.
func (o *FabricQosClassAllOf) SetMulticastOptimize(v bool) {
	o.MulticastOptimize = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricQosClassAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClassAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricQosClassAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricQosClassAllOf) SetName(v string) {
	o.Name = &v
}

// GetPacketDrop returns the PacketDrop field value if set, zero value otherwise.
func (o *FabricQosClassAllOf) GetPacketDrop() bool {
	if o == nil || o.PacketDrop == nil {
		var ret bool
		return ret
	}
	return *o.PacketDrop
}

// GetPacketDropOk returns a tuple with the PacketDrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClassAllOf) GetPacketDropOk() (*bool, bool) {
	if o == nil || o.PacketDrop == nil {
		return nil, false
	}
	return o.PacketDrop, true
}

// HasPacketDrop returns a boolean if a field has been set.
func (o *FabricQosClassAllOf) HasPacketDrop() bool {
	if o != nil && o.PacketDrop != nil {
		return true
	}

	return false
}

// SetPacketDrop gets a reference to the given bool and assigns it to the PacketDrop field.
func (o *FabricQosClassAllOf) SetPacketDrop(v bool) {
	o.PacketDrop = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *FabricQosClassAllOf) GetWeight() int64 {
	if o == nil || o.Weight == nil {
		var ret int64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClassAllOf) GetWeightOk() (*int64, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *FabricQosClassAllOf) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int64 and assigns it to the Weight field.
func (o *FabricQosClassAllOf) SetWeight(v int64) {
	o.Weight = &v
}

func (o FabricQosClassAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.BandwidthPercent != nil {
		toSerialize["BandwidthPercent"] = o.BandwidthPercent
	}
	if o.Cos != nil {
		toSerialize["Cos"] = o.Cos
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.MulticastOptimize != nil {
		toSerialize["MulticastOptimize"] = o.MulticastOptimize
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PacketDrop != nil {
		toSerialize["PacketDrop"] = o.PacketDrop
	}
	if o.Weight != nil {
		toSerialize["Weight"] = o.Weight
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricQosClassAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricQosClassAllOf := _FabricQosClassAllOf{}

	if err = json.Unmarshal(bytes, &varFabricQosClassAllOf); err == nil {
		*o = FabricQosClassAllOf(varFabricQosClassAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "BandwidthPercent")
		delete(additionalProperties, "Cos")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "MulticastOptimize")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PacketDrop")
		delete(additionalProperties, "Weight")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricQosClassAllOf struct {
	value *FabricQosClassAllOf
	isSet bool
}

func (v NullableFabricQosClassAllOf) Get() *FabricQosClassAllOf {
	return v.value
}

func (v *NullableFabricQosClassAllOf) Set(val *FabricQosClassAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricQosClassAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricQosClassAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricQosClassAllOf(val *FabricQosClassAllOf) *NullableFabricQosClassAllOf {
	return &NullableFabricQosClassAllOf{value: val, isSet: true}
}

func (v NullableFabricQosClassAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricQosClassAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
