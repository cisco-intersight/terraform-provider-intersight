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

// InventoryUemConnection UemConnection MO maintains the current Uem connection status for a given UEM endpoint.
type InventoryUemConnection struct {
	MoBaseMo
	// Connections status of Uem endpoint.
	ConnectionStatus *string          `json:"ConnectionStatus,omitempty"`
	EpInfo           *InventoryEpInfo `json:"EpInfo,omitempty"`
	// Type of Uem endpoint (BMC, CMC or VIC).
	EpType *string `json:"EpType,omitempty"`
	// The IP address of the Uem endpoint.
	Ip *string `json:"Ip,omitempty"`
	// The member identity of the UEM connection being inventoried.
	MemberIdentity *string `json:"MemberIdentity,omitempty"`
	// The model of the Uem endpoint.
	Model *string `json:"Model,omitempty"`
	// The serial number of the Uem endpoint.
	Serial *string `json:"Serial,omitempty"`
	// The MoId address of the Uem endpoint.
	TargetMoId *string `json:"TargetMoId,omitempty"`
	// The vendor of the Uem endpoint.
	Vendor               *string                              `json:"Vendor,omitempty"`
	ComputeBlade         *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeRackUnit      *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryUemConnection InventoryUemConnection

// NewInventoryUemConnection instantiates a new InventoryUemConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryUemConnection() *InventoryUemConnection {
	this := InventoryUemConnection{}
	var connectionStatus string = "Unknown"
	this.ConnectionStatus = &connectionStatus
	return &this
}

// NewInventoryUemConnectionWithDefaults instantiates a new InventoryUemConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryUemConnectionWithDefaults() *InventoryUemConnection {
	this := InventoryUemConnection{}
	var connectionStatus string = "Unknown"
	this.ConnectionStatus = &connectionStatus
	return &this
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *InventoryUemConnection) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetEpInfo returns the EpInfo field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetEpInfo() InventoryEpInfo {
	if o == nil || o.EpInfo == nil {
		var ret InventoryEpInfo
		return ret
	}
	return *o.EpInfo
}

// GetEpInfoOk returns a tuple with the EpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetEpInfoOk() (*InventoryEpInfo, bool) {
	if o == nil || o.EpInfo == nil {
		return nil, false
	}
	return o.EpInfo, true
}

// HasEpInfo returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasEpInfo() bool {
	if o != nil && o.EpInfo != nil {
		return true
	}

	return false
}

// SetEpInfo gets a reference to the given InventoryEpInfo and assigns it to the EpInfo field.
func (o *InventoryUemConnection) SetEpInfo(v InventoryEpInfo) {
	o.EpInfo = &v
}

// GetEpType returns the EpType field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetEpType() string {
	if o == nil || o.EpType == nil {
		var ret string
		return ret
	}
	return *o.EpType
}

// GetEpTypeOk returns a tuple with the EpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetEpTypeOk() (*string, bool) {
	if o == nil || o.EpType == nil {
		return nil, false
	}
	return o.EpType, true
}

// HasEpType returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasEpType() bool {
	if o != nil && o.EpType != nil {
		return true
	}

	return false
}

// SetEpType gets a reference to the given string and assigns it to the EpType field.
func (o *InventoryUemConnection) SetEpType(v string) {
	o.EpType = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InventoryUemConnection) SetIp(v string) {
	o.Ip = &v
}

// GetMemberIdentity returns the MemberIdentity field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetMemberIdentity() string {
	if o == nil || o.MemberIdentity == nil {
		var ret string
		return ret
	}
	return *o.MemberIdentity
}

// GetMemberIdentityOk returns a tuple with the MemberIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetMemberIdentityOk() (*string, bool) {
	if o == nil || o.MemberIdentity == nil {
		return nil, false
	}
	return o.MemberIdentity, true
}

// HasMemberIdentity returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasMemberIdentity() bool {
	if o != nil && o.MemberIdentity != nil {
		return true
	}

	return false
}

// SetMemberIdentity gets a reference to the given string and assigns it to the MemberIdentity field.
func (o *InventoryUemConnection) SetMemberIdentity(v string) {
	o.MemberIdentity = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InventoryUemConnection) SetModel(v string) {
	o.Model = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InventoryUemConnection) SetSerial(v string) {
	o.Serial = &v
}

// GetTargetMoId returns the TargetMoId field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetTargetMoId() string {
	if o == nil || o.TargetMoId == nil {
		var ret string
		return ret
	}
	return *o.TargetMoId
}

// GetTargetMoIdOk returns a tuple with the TargetMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetTargetMoIdOk() (*string, bool) {
	if o == nil || o.TargetMoId == nil {
		return nil, false
	}
	return o.TargetMoId, true
}

// HasTargetMoId returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasTargetMoId() bool {
	if o != nil && o.TargetMoId != nil {
		return true
	}

	return false
}

// SetTargetMoId gets a reference to the given string and assigns it to the TargetMoId field.
func (o *InventoryUemConnection) SetTargetMoId(v string) {
	o.TargetMoId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *InventoryUemConnection) SetVendor(v string) {
	o.Vendor = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *InventoryUemConnection) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *InventoryUemConnection) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *InventoryUemConnection) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemConnection) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *InventoryUemConnection) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *InventoryUemConnection) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o InventoryUemConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.EpInfo != nil {
		toSerialize["EpInfo"] = o.EpInfo
	}
	if o.EpType != nil {
		toSerialize["EpType"] = o.EpType
	}
	if o.Ip != nil {
		toSerialize["Ip"] = o.Ip
	}
	if o.MemberIdentity != nil {
		toSerialize["MemberIdentity"] = o.MemberIdentity
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.TargetMoId != nil {
		toSerialize["TargetMoId"] = o.TargetMoId
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryUemConnection) UnmarshalJSON(bytes []byte) (err error) {
	type InventoryUemConnectionWithoutEmbeddedStruct struct {
		// Connections status of Uem endpoint.
		ConnectionStatus *string          `json:"ConnectionStatus,omitempty"`
		EpInfo           *InventoryEpInfo `json:"EpInfo,omitempty"`
		// Type of Uem endpoint (BMC, CMC or VIC).
		EpType *string `json:"EpType,omitempty"`
		// The IP address of the Uem endpoint.
		Ip *string `json:"Ip,omitempty"`
		// The member identity of the UEM connection being inventoried.
		MemberIdentity *string `json:"MemberIdentity,omitempty"`
		// The model of the Uem endpoint.
		Model *string `json:"Model,omitempty"`
		// The serial number of the Uem endpoint.
		Serial *string `json:"Serial,omitempty"`
		// The MoId address of the Uem endpoint.
		TargetMoId *string `json:"TargetMoId,omitempty"`
		// The vendor of the Uem endpoint.
		Vendor           *string                              `json:"Vendor,omitempty"`
		ComputeBlade     *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
		ComputeRackUnit  *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varInventoryUemConnectionWithoutEmbeddedStruct := InventoryUemConnectionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInventoryUemConnectionWithoutEmbeddedStruct)
	if err == nil {
		varInventoryUemConnection := _InventoryUemConnection{}
		varInventoryUemConnection.ConnectionStatus = varInventoryUemConnectionWithoutEmbeddedStruct.ConnectionStatus
		varInventoryUemConnection.EpInfo = varInventoryUemConnectionWithoutEmbeddedStruct.EpInfo
		varInventoryUemConnection.EpType = varInventoryUemConnectionWithoutEmbeddedStruct.EpType
		varInventoryUemConnection.Ip = varInventoryUemConnectionWithoutEmbeddedStruct.Ip
		varInventoryUemConnection.MemberIdentity = varInventoryUemConnectionWithoutEmbeddedStruct.MemberIdentity
		varInventoryUemConnection.Model = varInventoryUemConnectionWithoutEmbeddedStruct.Model
		varInventoryUemConnection.Serial = varInventoryUemConnectionWithoutEmbeddedStruct.Serial
		varInventoryUemConnection.TargetMoId = varInventoryUemConnectionWithoutEmbeddedStruct.TargetMoId
		varInventoryUemConnection.Vendor = varInventoryUemConnectionWithoutEmbeddedStruct.Vendor
		varInventoryUemConnection.ComputeBlade = varInventoryUemConnectionWithoutEmbeddedStruct.ComputeBlade
		varInventoryUemConnection.ComputeRackUnit = varInventoryUemConnectionWithoutEmbeddedStruct.ComputeRackUnit
		varInventoryUemConnection.RegisteredDevice = varInventoryUemConnectionWithoutEmbeddedStruct.RegisteredDevice
		*o = InventoryUemConnection(varInventoryUemConnection)
	} else {
		return err
	}

	varInventoryUemConnection := _InventoryUemConnection{}

	err = json.Unmarshal(bytes, &varInventoryUemConnection)
	if err == nil {
		o.MoBaseMo = varInventoryUemConnection.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "EpInfo")
		delete(additionalProperties, "EpType")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "MemberIdentity")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "TargetMoId")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableInventoryUemConnection struct {
	value *InventoryUemConnection
	isSet bool
}

func (v NullableInventoryUemConnection) Get() *InventoryUemConnection {
	return v.value
}

func (v *NullableInventoryUemConnection) Set(val *InventoryUemConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryUemConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryUemConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryUemConnection(val *InventoryUemConnection) *NullableInventoryUemConnection {
	return &NullableInventoryUemConnection{value: val, isSet: true}
}

func (v NullableInventoryUemConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryUemConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
